Helm Charts
---

## Current Status

Currently, either the [ShardingSphere Helm Charts](https://github.com/apache/shardingsphere-on-cloud/tree/main/charts/apache-shardingsphere-proxy-charts) and the [ShardingSphere Operator](https://github.com/apache/shardingsphere-on-cloud/tree/main/shardingsphere-operator) make up the latest version as (5.3.0) of ShardingSphere Proxy.

### Requirements

* [Kubernetes 1.18+]()
* [kubectl](https://kubernetes.io/docs/reference/kubectl/)
* [Helm 3.2.0+](https://helm.sh/)
* [StorageClass](https://kubernetes.io/docs/concepts/storage/storage-classes/) of [PV (Persistent Volumes)](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) can be dynamically applied for persistent data (Optional)


## ShardingSphere Helm Charts

ShardingSphere Helm Charts uses [Helm](https://helm.sh/) to provide guidance for the installation of a ShardingSphere-Proxy instance in a Kubernetes cluster.

### Quick Start

Please follow these [instructions](./apache-shardingsphere-proxy-charts/README.md) to deploy a ShardingSphere cluster with version 5.3.0.


## ShardingSphere-Cluster Operator Helm Charts

The ShardingSphere-Cluster Operator uses predefined CustomResourceDefinitions for describing an Apache ShardingSphere Deployment on Kubernetes.

### Quick Start 

Please follows [instructions](./apache-shardingsphere-cluster-operator-charts/README.md) to deploy a ShardingSphere cluster with version 5.3.0.

