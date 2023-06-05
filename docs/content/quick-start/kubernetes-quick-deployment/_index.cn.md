+++
pre = "<b>2.1 </b>"
title = "ShardingSphere Helm Charts 一键安装"
weight = 1
chapter = true
+++

## 操作步骤

### 在线安装

1. 添加 ShardingSphere-Proxy 到本地 Helm 仓库：

```shell
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
```

2. 安装 ShardingSphere-Proxy Charts：

```shell
helm install shardingsphere-proxy shardingsphere/apache-shardingsphere-proxy-charts 
```
