## 0.3.0

### New Feature

1. Introduce new CRD StorageNode，to be activated at `--feature-gates=StorageNode=true`
2. Introduce new CRD Chaos，to be activated at `--feature-gates=Chaos=true`
3. Add Grafana dashboard template for ShardingSphere Proxy Grafana

### Enhancement

1. Update AgentConfig of ComputeNode
2. Update the display of execute-latency in the Grafana panel
3. Add some checks about zookeeper before install proxy with Helm 
4. Add demo for sharding algorithm based on Wasm
5. Support Agent in Proxy CHarts
6. StorageNode suppors CloudNativePG cluster
7. Optimize ComputeNode conditions in Pending phase
8. Adjust default legents in Grafana template 
9. Support operator docker image pushed to Apache Dockerhub repo
10. Add golangci-lint action to Github Action 
11. Upgrade Zookeeper config to support arm arch in Proxy Charts

### Bug fix

1. Fix Arm image build in Github Action 
2. Fix Agent ConfigMap charts name template
3. Fix Status reconciliation error which cause unexpected loop

### Change Log

1. [MILESTONE](https://github.com/apache/shardingsphere-on-cloud/milestone/7)


## 0.2.0

### New Feature

1. Introduce new CRD ComputeNode，to be activated at `--feature-gates=ComputeNode=true`
2. Introduce subresource scale for ComputeNode to support kubectl scale command

### Enhancement

1. Add annotation rolling update support for CRD ShardingSphereProxy 
2. Write NodePort back to ComputeNode
3. Introducing Ginkgo test for ComputeNode

### Bug fix

1. Fix readyNodes and conditions error for ShardingSphereProxy Status in some cases
2. Fix NPE in non-MySQL configurations

### Change Log

1. [MILESTONE](https://github.com/apache/shardingsphere-on-cloud/milestone/6)

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
