## 0.1.2

### New Feature

1. Provide Grafana dashboard template for ShardingSphere-Operator metrics.
1. Updating and renaming renaming apache-shardingsphere-operator-cluster-charts to apache-shardingsphere-operator-charts, which supports deploying operator and cluster at the same time.

### Enhancement

1. Support monitoring metrics for ShardingSphere-Operator.
1. Support more parameters about health check, CloudWatch and alerts when deploying ShardingSphere clusters with Terraform and AWS CloudFormation.
1. Support HuaweiCloud with Terraform.

### Change Log

1. [MILESTONE](https://github.com/apache/shardingsphere-on-cloud/milestone/5)


## 0.1.1

### New Feature

1. Support using Terraform to deploy a ShardingSphere Proxy Cluster in  Multiple AZs

### Enhancement

1. Support using CloudFormation to deploy a ShardingSphere Proxy Cluster in  Multiple AZs
1. Add nameOverride parameter for helm charts, allowing users to override resource names

### Change Log

1. [MILESTONE](https://github.com/apache/shardingsphere-on-cloud/milestone/4)


## 0.1.0

### New Feature

1. Supports the use of proxy to configure and describe the ShardingSphere Proxy cluster
1. Support native ShardingSphere Proxy server.yaml configuration
1. Support automatic creation of HPA based on CPU metrics
1. Support automatic download of MySQL driver
1. Support deploy ShardingSphere Operator
1. Support deploy ShardingSphere Proxy cluster 
1. Support deploy Zookeeper by Bitnami
1. Support automatic configuration of the address of the governance center
1. Include ShardingSphere-Proxy charts into the on-cloud repository
1. Support deploying ShardingSphere-Proxy cluster on aws using cloudformation

### Change Log

1. [MILESTONE](https://github.com/apache/shardingsphere-on-cloud/milestone/3)
