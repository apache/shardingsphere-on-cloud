# Apache ShardingSphere Operator Charts
The Apache ShardingSphere-Operator is used to quickly install Apache ShardingSphere-Proxy Cluster. 
This Chart will install ShardingSphere-Proxy Cluster using the CRD provided by ShardingSphere-Operator.

## Install
Use the following command to install:
```shell
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts 
```

## Uninstall 
Use the following command to uninstall:
```shell
helm unstall [RELEASE_NAME]
```

## Manual

For further instructions, please check out the [Apache ShardingSphere on Cloud official documentations](https://shardingsphere.apache.org/oncloud/current/en/overview/).
