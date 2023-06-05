+++
pre = "<b>2.1 </b>"
title = "ShardingSphere Helm Charts User Manual"
weight = 1
chapter = true
+++

## Procedure

### Online Installation

1. Add ShardingSphere-Proxy to the local Helm warehouse:

```shell
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
```

2. Install ShardingSphere-Proxy Charts:

```shell
helm install shardingsphere-proxy shardingsphere/apache-shardingsphere-proxy-charts 
```
