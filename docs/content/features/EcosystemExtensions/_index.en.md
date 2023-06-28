+++
pre = "<b>3.4 </b>"
title = "Ecosystem"
weight = 4
chapter = true
+++


## WebAssembly (Wasm) Extensions

WebAssembly (abbreviated as Wasm) has now expanded its application beyond web browsers, despite its initial intention of improving JavaScript performance on webpages. 

With WebAssembly System Interface (WASI), Wasm can now run in various scenarios including trusted computing and edge computing. The majority of popular programming languages are compatible with Wasm, while ShardingSphere plugins (SPIs) currently only support the Java ecosystem. Introducing Wasm into ShardingSphere, can significantly enhance ShardingShphere's pluggable ecosystem with better flexibility, and attract more developers to the community.

### Using Wasm for Custom Sharding

Apache ShardingSphere currently uses Service Provider Interface (SPIs) to expand its pluggable architecture. For more information, please refer to the [ShardingSphere Developer Manual](https://shardingsphere.apache.org/document/current/en/dev-manual/). 

We have implemented a custom sharding demo using Wasm for sharding scenarios. The demo below shows the custom sharding logic when `sharding_count` is `3`:

1. Extract the sharding SPI logic from Apache ShardingSphere, for example, the auto-create sharding algorithm `MOD` from the [document](https://shardingsphere.apache.org/document/current/en/dev-manual/sharding/). Organize it into a separate [directory](https://github.com/apache/shardingsphere-on-cloud/tree/main/wasm/wasm-sharding-java/src/main/java/org/apache/shardingsphere):

```shell
├── pom.xml
├── src
│   └── main
│       └── java
│           └── org
│               └── apache
│                   └── shardingsphere
│                       ├── infra 
│                       ├── sharding 
```

2. Add [demo.java](https://github.com/apache/shardingsphere-on-cloud/blob/main/wasm/wasm-sharding-java/src/main/java/org/apache/shardingsphere/demo.java) to the above directory. Instantiate `StandardShardingAlgorithm` using `WasmShardingAlgorithm` provided by Wasm for sharding. Run the custom sharding logic and view the output. 

```java
// ...
        StandardShardingAlgorithm<?> shardingAlgorithm = new WasmShardingAlgorithm();
// ...
```

3. Write [custom sharding logic](https://github.com/apache/shardingsphere-on-cloud/tree/main/wasm/wasm-sharding-java/wasm-sharding) in Rust, and compile to Wasm module.

```rust
#[link(wasm_import_module = "sharding")]
extern "C" {
    fn poll_table(addr: i64, len: i32) -> i32;
}

// The value of sharding_count must be consistent with the value of the AvaliableTargetNames
const SHARDING_COUNT: u8 = 3;

#[no_mangle]
pub unsafe extern "C" fn do_work() -> i64 {
// ...
    let sharding =  column_value % SHARDING_COUNT;
// ...
    std::ptr::copy_nonoverlapping(table_name.as_mut_ptr() as *const _, buf.as_mut_ptr().add(len as usize), table_name.len());
    buf_ptr
}
```

4. Create [WasmShardingAlgorithm.java](https://github.com/apache/shardingsphere-on-cloud/blob/main/wasm/wasm-sharding-java/src/main/java/org/apache/shardingsphere/sharding/WasmShardingAlgorithm.java) under `src/main/java/org/apache/shardingsphere/sharding/`, to communicate with the custom sharding logic in Wasm:

```java
//...
public final class WasmShardingAlgorithm implements StandardShardingAlgorithm<Comparable<?>> {
// ...
    private static final String WASM_PATH = "./wasm-sharding/target/wasm32-wasi/debug/wasm_sharding.wasm";
    private String wasmDoSharding(final Collection<String> availableTargetNames, final PreciseShardingValue<Comparable<?>> shardingValue) {
// ...
    }

    @Override
    public String getType() {
        return "WASM";
    }
}

```

### Extend Custom Sharding Expressions with Wasm

ShardingSphere only supports Groovy for defining sharding rules within the Java ecosystem. With Wasm, you can now define sharding logic using your preferred language. `WASM-sharding-js` demonstrates how to define the CRC32MOD sharding algorithm using JavaScript. 

To make sharding easier, Wasm allows you to use your familiar language, which makes the extension of sharding algorithms even more effortless. [wasm-sharding-js](https://github.com/apache/shardingsphere-on-cloud/tree/main/wasm/wasm-sharding-js) provides an example of how to compile the sharding algorithms in JavaScript into Wasm extensions. 

The directory structure is as follows:
```shell
├── Cargo.lock
├── Cargo.toml
├── README.md
├── build.rs
├── lib
│   └── binding.rs
├── package-lock.json
├── package.json
├── sharding
│   ├── config.js
│   ├── crc32.js
│   ├── sharding.js
│   └── strgen.js
└── src
```
In the file `sharding/config.js`, two sharding resources are defined: `t_order_00${0..2}` and `ms_ds00${crc32(field_id)}`. For `t_order_00${0..2}`, it's expected to generate three sharded tables: `t_order_000`, `t_order_001`, and `t_order_002` after parsing. For `ms_ds00${crc32(field_id)}`, we expect the `field_id` to be hashed with `crc32` before sharding: 

```javascript
export let cc = "t_order_00${0..2}"
export let cc_crc32 = "ms_ds00${crc32(field_id)}"
```

Furthermore, the `pisa_crc32` function declared in the file `sharding/sharding.js` shows the parsing of the above two expressions using JavaScript:

```javascript
//...
function pisa_crc32(str, mod) {
    let c2 = crc32_str(str)
    let m = c2 % mod
    return m < 256 ? 0 : m < 512 ? 1: m<768 ? 2 : 3
}
//...
```
Thanks to Wasm, not only you can enhance the functionality of ShardingSphere, but also extend their technical capabilities to a wider range of stacks.
