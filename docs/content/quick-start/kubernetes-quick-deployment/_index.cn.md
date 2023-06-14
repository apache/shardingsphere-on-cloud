+++
pre = "<b>2.2 </b>"
title = "ShardingSphere Helm Charts 一键安装"
weight = 2
chapter = true
+++

## 操作说明

在 Kubernetes 上使用 Helm Charts 安装，可以使用 Charts 仓库在线安装，或者下载源码后本地编译安装。 

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

### 源码安装

1. 同步 Git 仓库

```shell
git clone git@github.com:apache/shardingsphere-on-cloud.git
```

2. 可以使用如下命令进行默认配置安装：

```shell
cd charts/apache-shardingsphere-proxy-charts/charts/governance
helm dependency build 
cd ../..
helm dependency build 
cd ..
helm install shardingsphere-proxy apache-shardingsphere-proxy-charts
```
