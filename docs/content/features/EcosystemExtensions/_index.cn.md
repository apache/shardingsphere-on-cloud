+++
pre = "<b>3.4 </b>"
title = "生态扩展"
weight = 4
chapter = true
+++

## WebAssembly 扩展

WebAssembly 起初是为了解决 JavaScript 在浏览器上的一些性能瓶颈，目前已经不局限于浏览器, WASI 规范使得其可以运行在更多的场景中，例如可信计算、边缘计算等。目前主流的语言基本都支持编译到 WASM 。

ShardingSphere 的插件（SPI）只支持 java 生态，把 WASM 引入到 ShardingSphere 中可以进一步开放生态，吸引更多的社区开发者。例如，社区针对自定义分片场景，创建了简单的 demo， 增加了 WasmShardingAlgorithm.java 模块，用于加载和运行用户提供的 wasm 自定义分片逻辑。

此外由于 ShardingSphere 定义分片规则只支持 Java 生态中的 Groovy 语法，通过 WASM 用户可以使用自己熟悉的语言生态来定义分片逻辑。WASM-sharding-js 则演示了如何通过 JavaScript 定义 CRC32MOD 分片算法。
