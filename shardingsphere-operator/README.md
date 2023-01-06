Apache ShardingSphere Operator
---
The ShardingSphere Operator uses predefined CustomResourceDefinitions for describing an Apache ShardingSphere Deployment on Kubernetes.

### Requirements

With the help of ShardingSphere-Operator you can create a ShardingSphere-Proxy cluster including the ZooKeeper cluster in minutes.

For installation of SharingSphere-Operator, you will need a Kubernetes cluster. No matter it is a managed Kubernetes service like AWS EKS or self-hosted Kubernetes, or just a mini-kube, you can easily install ShardingSphere-Operator with respective [ShardingSphere Operator Helm Charts](https://github.com/apache/shardingsphere-on-cloud/tree/main/charts/apache-shardingsphere-operator-charts), and apply the manifests in [ShardingSphere Operator Cluster Helm Charts](https://github.com/apache/shardingsphere-on-cloud/tree/main/charts/apache-shardingsphere-operator-cluster-charts) describing the expected Apache ShardingSphere deployment. **Kubernetes 1.18+ is recommended**. 

### Internal Architecture

![img.png](../docs/images/ss-operatorIA.png)

### Current Status

Minimum Viable Product

### Quick Start 

Please follows [instructions](./docs/shardingsphere-operator.md) to deploy a ShardingSphere cluster with version 5.2.0.

### Features

* Supports the use of proxy to configure and describe the shardingsphere-proxy cluster. For detailed configuration, see the deployment documentation.
* Support native shardingsphere proxy server.yaml configuration. For specific support items, please refer to the documentation
* Support automatic creation of HPA based on CPU metrics.
* Support automatic download of MySQL driver.

### Installation

#### Helm

* ShardingSphere-Operator chart
    * Support deploy shardingsphere operator.
* ShardingSphere-Cluster chart
    * Support deploy shardingsphere proxy cluster.
    * Support deploy Zookeeper by bitnami.
    * Support automatic configuration of the address of the governance center.
    * Use github pages to host repositories and support helm repo add to add repositories.

### Development Requirements

To build ShardingSphere Operator from scratch you will need to install the following tools:

* [Git](https://git-scm.com/)
* [Golang 1.17](https://golang.org/dl/)
* [make](https://www.gnu.org/savannah-checkouts/gnu/make/make.html)
* [Kubernetes 1.20+](https://github.com/kubernetes/kubernetes)
* [Kubebuilder 3.4.1+](https://github.com/kubernetes-sigs/kubebuilder)



