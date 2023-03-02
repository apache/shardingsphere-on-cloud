This is an experimental demo project using WebAssembly for crc32 sharding in Apache ShardingSphere

### How to use.
#### 1. Please confirm that WasmEdgeRuntime is installed in the environment.

### 2. build
```
cargo build --target wasm32-wasi --release
```

### 3. run sharding.js
```
wasmedge --dir .:. target/wasm32-wasi/release/wasmedge_quickjs.wasm sharding/sharding.js WasmEdge Runtime
```

### THANKS
Thanks wasmedge-quickjs for providing a reference