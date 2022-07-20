# shardingsphere-on-cloud
Contains various cloud solutions about ShardingSphere
## ShardingSphere Operator
Use the operator to create a ShardingSphere-Proxy cluster including the ZooKeeper cluster in the Kubernetes .
### Internal Architecture
![img.png](./doc/images/ss-operatorIA.png)
### Status
minimum viable product
### how to install

[quick start](./doc/deploy.md)
## Features
* Supports the use of proxy to configure and describe the shardingsphere-proxy cluster. For detailed configuration, see the deployment documentation
* Support native shardingsphere proxy server.yaml configuration. For specific support items, please refer to the documentation
* Support automatic creation of HPA based on cpu metrics
* Support automatic download of MySQL driver

## Deploy
### Helm
* ShardingSphere-Operator chart
    * Support deploy shardingsphere operator
* ShardingSphere-Cluster chart
    * Support deploy shardingsphere proxy cluster
    * Support deploy Zookeeper by bitnami
    * Support automatic configuration of the address of the governance center
    * Use github pages  to host repositories and support helm repo add to add repositories 
## Prerequisites
- Kubernetes 1.18+
## Deploy Prerequisites

To build ShardingSphere Operator from scratch you will need to install the following tools:

* Git
* [Golang 1.17](https://golang.org/dl/)
* make
* Kubernetes 1.20+ 
* Kubebuilder 3.4.1+
## Contributing
We welcome all kinds of contributions from the open-source community, individuals and partners.
[Contributing](./CONTRIBUTING.md)
## Community

[Slack channel](https://join.slack.com/t/apacheshardingsphere/shared_invite/zt-sbdde7ie-SjDqo9~I4rYcR18bq0SYTg). Best for: instant communications and online meetings, sharing your applications.

![wechat](./doc/images/ssvx.png)
## License
[Apache License 2.0](https://github.com/SphereEx/shardingsphere-on-cloud/blob/master/LICENSE)



