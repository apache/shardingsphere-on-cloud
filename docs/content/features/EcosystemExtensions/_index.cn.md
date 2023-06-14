+++
pre = "<b>3.4 </b>"
title = "生态扩展"
weight = 4
chapter = true
+++

## 概述

## WebAssembly 扩展

WebAssembly 起初是为了解决 JavaScript 在浏览器上的一些性能瓶颈，目前已经不局限于浏览器, Wasi 规范使得其可以运行在更多的场景中，例如可信计算、边缘计算等。目前主流的语言基本都支持编译到 Wasm。ShardingSphere 的插件（SPI）只支持 java 生态，把 Wasm 引入到 ShardingSphere 中可以进一步开放扩展，可以更加丰富其可插拔生态，吸引更多的社区开发者。

### 利用 Wasm 实现自定义分片算法

目前 Apache ShardingSphere 的可插拔架构使用 SPI 方式进行扩展，详见：[ShardingSphere 开发者手册](https://shardingsphere.apache.org/document/current/cn/dev-manual/)。

针对自定义分片场景，使用 WebAssembly 实现了自定义的分片 Demo。Demo 中演示了 `sharding_count` 为`3`的自定义分片逻辑，实现步骤如下：

1. 从 Apache ShardingSphere 中抽取数据分片 SPI 的相关逻辑，比如[文档](https://shardingsphere.apache.org/document/current/cn/dev-manual/sharding/)中提到的`MOD`自动分片算法，将其整理到单独的[目录](https://github.com/apache/shardingsphere-on-cloud/tree/main/wasm/wasm-sharding-java/src/main/java/org/apache/shardingsphere)中：

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

2. 在上述目录下增加 [demo.java](https://github.com/apache/shardingsphere-on-cloud/blob/main/wasm/wasm-sharding-java/src/main/java/org/apache/shardingsphere/demo.java) ，用 Wasm 提供的分片算法`WasmShardingAlgorithm`实例化`StandardShardingAlgorithm`， 运行自定义的分片逻辑并输出结果。

```java
// ...
        StandardShardingAlgorithm<?> shardingAlgorithm = new WasmShardingAlgorithm();
// ...

```

3. 使用 Rust 编写[自定义分片逻辑](https://github.com/apache/shardingsphere-on-cloud/tree/main/wasm/wasm-sharding-java/wasm-sharding)，并编译为 Wasm 制品，

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

4. 在`src/main/java/org/apache/shardingsphere/sharding/`下创建 [WasmShardingAlgorithm.java](https://github.com/apache/shardingsphere-on-cloud/blob/main/wasm/wasm-sharding-java/src/main/java/org/apache/shardingsphere/sharding/WasmShardingAlgorithm.java)， 以和 Wasm 中的自定义分片逻辑通信并获得结果：

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

### 利用 Wasm 实现自定义分片表达式

由于 ShardingSphere 定义分片规则只支持 Java 生态中的 Groovy 语法，通过 WASM 用户可以使用自己熟悉的语言生态来定义分片逻辑。WASM-sharding-js 则演示了如何通过 JavaScript 定义 CRC32MOD 分片算法。
