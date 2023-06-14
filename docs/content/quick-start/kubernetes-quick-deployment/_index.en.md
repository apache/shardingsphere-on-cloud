+++
pre = "<b>2.2 </b>"
title = "ShardingSphere Helm Charts Quick Start"
weight = 2
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
