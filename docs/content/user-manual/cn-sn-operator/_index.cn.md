+++
pre = "<b>4.2 </b>"
title = "ShardingSphere Operator 用户手册"
weight = 2
chapter = true
+++

## 概述 

ShardingSphere Operator 是 Kubernetes Operator 模型的实践。它将 ShardingSphere Proxy 的维护经验变成了可执行的程序，并借助 Kubernetes 的声明式和“调谐”特性进行实现。

ShardingSphere Operator 将从计算节点、存储节点甚至混沌故障都抽象为 Kubernetes 自定义资源对象 CRD，由用户者负责编写相应的 CRD 配置，由 Operator 负责执行和保障期望的状态。

如果想要安装试用，请阅读“安装 Operator”小节，如果想了解 CRD 的配置，请阅读“CRD 介绍”小节。

## 安装 Operator

Operator 目前支持 Helm Charts 快速部署，配置文件目录为：apache-shardingsphere-operator-charts。用户可以根据需要采用在线安装或源码安装。

### 在线安装

```shell
 kubectl create ns shardingsphere-operator
 helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
 helm repo update
 helm install shardingsphere-cluster shardingsphere/apache-shardingsphere-operator-charts -n shardingsphere-operator
```

### 源码安装

```shell
kubectl create ns shardingsphere-operator
cd charts/apache-shardingsphere-operator-charts/
helm dependency build
cd ../
helm install shardingsphere-cluster apache-shardingsphere-operator-charts -n shardingsphere-operator
```

### Charts 参数说明

#### 通用参数
| Name              | Description                                                                                               | Value                                 |
|-------------------|-----------------------------------------------------------------------------------------------------------|---------------------------------------|
| `nameOverride`    | nameOverride String to partially override common.names.fullname template (will maintain the release name) | `shardingsphere-proxy` |

#### ShardingSphere Operator 参数
| Name                              | Description                                                                                                | Value                                                                   |
|-----------------------------------|------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| `operator.replicaCount`           | operator replica count                                                                                     | `2`                                                                     |
| `operator.image.repository`       | operator image name                                                                                        | `apache/shardingsphere-operator` |
| `operator.image.pullPolicy`       | image pull policy                                                                                          | `IfNotPresent`                                                          |
| `operator.image.tag`              | image tag                                                                                                  | `0.3.0`                                                                 |
| `operator.imagePullSecrets`       | image pull secret of private repository                                                                    | `[]`                                                                    |
| `operator.resources`              | operator Resources required by the operator                                                                | `{}`                                                                    |
| `operator.health.healthProbePort` | operator health check port                                                                                 | `8080`                                                                  |

## CRD 介绍

### ComputeNode 

ComputeNode 用来描述 ShardingSphere 集群中的计算节点，通常指的是 Proxy。由于 ShardingSphere Proxy 是无状态应用，所以可以利用 Kubernetes 原生的工作负载 Deployment 进行管理，同时使用 ConfigMap 和 Service 实现对于启动配置和服务发现的配置。利用 ComputeNode 不仅可以将 Deployment、ConfigMap 和 Service 中的关键配置统一，还匹配了 ShardingSphere 的语义，帮助 Operator 快速锁定工作负载。如图：

![]()

#### Operator 配置

目前 Operator 想要使用 ComputeNode 需要打开相应的 FeatureGate：

```shell
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts --set operator.featureGates.computeNode=true --set proxyCluster.enabled=false
```

#### 字段说明

##### 必填配置 

配置项 |  描述 | 类型 | 示例 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`metadata.name` | 计划部署的名称 |  string | `foo` 
`metadata.namespace` | 计划部署的命名空间，默认为 default | string |                                      | `shardingsphere-system`
`spec.storageNodeConnector.type`     | 后端驱动类型 | string | `mysql`
`spec.storageNodeConnector.version`  | 后端驱动版本| string  | `5.1.47`
`spec.serverVersion`                 | ShardingSphere Proxy 版本 | string | `5.4.0`
`spec.replicas `     | 计划部署的实例数量 |  number | `3`
`spec.selectors`     | 实例选择器,同 Deployment.Spec.Selectors |  number | `3`
`spec.portBindings[0].name`          | 对外暴露的端口名称 | string |                                                                         | `3307`
`spec.portBindings[0].containerPort` | 对外暴露的容器端口号| number |`3307`
`spec.portBindings[0].servicePort`   | 对外暴露的服务端口号 | number                                                                 | `3307`
`spec.portBindings[0].procotol`      | 对外暴露的端口协议 | string|  `TCP`
`spec.serviceType`                   | 对外暴露的服务类型| string                                                                     | `ClusterIP`
`spec.bootstrap.serverConfig.authority.privilege.type`    | 计算节点权限设置，默认值为 ALL_PRIVILEGES_PERMITTED  | string                                                                        | `ALL_PRIVILEGES_PERMITTED` 
`spec.bootstrap.serverConfig.authority.users[0].user`     | 计算节点用户名，格式: <username>@<hostname> ，将 hostname 设置为 % 或为空表示不关心来源主机|string |`root@%`
`spec.bootstrap.serverConfig.authority.users[0].password` | 计算节点密码 |string                                                                                                                     | `root`
`spec.bootstrap.serverConfig.mode.type`                                          | 运行模式配置，支持 Standalone 和 Cluster           | string | `Cluster`
`spec.bootstrap.serverConfig.mode.repository.type`                               | 治理中心类型，支持 ZooKeeper 和 Etcd  |string              | `ZooKeeper`
`spec.bootstrap.serverConfig.mode.repository.props`            | 治理中心属性配置，可以参考[常用的 ServerConfig Repository Props](#常用的\ ServerConfig\ Repository\ Props\ 配置)  | map[string]string                                    | 

##### 常用的 ServerConfig Repository Props 配置
配置项 |  描述 | 示例 
------------------ | -------------------------------------------------------------------------------- | ----------------------------------------
`spec.bootstrap.serverConfig.mode.repository.props.timeToLiveSeconds`            | TTL                                        | `600`
`spec.bootstrap.serverConfig.mode.repository.props.serverlists`                 | 治理中心列表                                     | `zookeeper.default:2181` 
`spec.bootstrap.serverConfig.mode.repository.props.retryIntervalMilliseconds`    | 重试间隔                                      | `500`
`spec.bootstrap.serverConfig.mode.repository.props.operationTimeoutMilliseconds` | 超时时间                                   | `5000`
`spec.bootstrap.serverConfig.mode.repository.props.namespace`                    | 治理中心命名空间（非 K8s 命名空间）                                        | `governance_ds`
`spec.bootstrap.serverConfig.mode.repository.props.maxRetries`                   | 客户端最大重试次数                                   | `3`


##### 选填配置 

配置项 |  描述 | 类型 | 示例 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`spec.probes.livenessProbe` | 健康检查探针 |  corev1.Probe | 
`spec.probes.readinessProbe` | 就绪检查探针 |  corev1.Probe | 
`spec.probes.startupProbe` | 启动检查探针 |  corev1.Probe | 
`spec.imgaePullSecrets ` | 镜像密钥 | corev1.LocalObjectReference  | 
`spec.env` | 环境变量 | corev1.Env | 
`spec.resources` | 资源声明 | corev1.ResourceRequirements | 
`spec.bootstrap.agentConfig.plugins.logging.file.props` | Agent 日志插件配置属性| map[string]string |
`spec.bootstrap.agentConfig.plugins.metrics.prometheus.host` | Agent 指标插件配置主机| map[string]string |
`spec.bootstrap.agentConfig.plugins.metrics.prometheus.port` | Agent 指标插件配置端口| map[string]string |
`spec.bootstrap.agentConfig.plugins.metrics.prometheus.props` | Agent 指标插件配置属性| map[string]string |
`spec.bootstrap.agentConfig.plugins.tracing.openTracing.props` | Agent 追踪插件配置属性| map[string]string |
`spec.bootstrap.agentConfig.plugins.tracing.openTelemetry.props` | Agent 追踪插件配置属性| map[string]string |

#### 示例

以下是一个基础的 ComputeNode CRD 配置示例，可以拉起一个3节点的 ShardingSphere Proxy 集群。

```yaml
apiVersion: shardingsphere.apache.org/v1alpha1
kind: ComputeNode
metadata:
  labels:
    app: foo
  name: foo
spec:
  storageNodeConnector:
    type: mysql
    version: 5.1.47
  serverVersion: 5.3.1
  replicas: 3
  selector:
    matchLabels:
      app: foo
  portBindings:
  - name: server
    containerPort: 3307
    servicePort: 3307
    protocol: TCP
  serviceType: ClusterIP
  bootstrap:
    serverConfig:
      authority:
        privilege:
          type: ALL_PERMITTED
        users:
        - user: root@%
          password: root
      mode:
        type: Cluster
        repository:
          type: ZooKeeper
          props:
            timeToLiveSeconds: "600"
            server-lists: ${PLEASE_REPLACE_THIS_WITH_YOUR_ZOOKEEPER_SERVICE}
            retryIntervalMilliseconds: "500"
            operationTimeoutMilliseconds: "5000"
            namespace: governance_ds
            maxRetries: "3"
      props:
        proxy-frontend-database-protocol-type: MySQL
```
注意：请先准备一个可以正常运行的 ZooKeeper 集群

### StorageNode

StorageNode 是 Operator 对于数据源的描述，提供对数据源的生命周期管理。它的使用需要配合 StorageProvider，现在支持 AWS RDS 和 CloudNative PG 。如图：

![]()

注意：StorageNode 是可选 CRD，用户可根据实际场景决定是否需要通过 StorageNode 管理数据源。

#### Operator 配置

目前 Operator 想要使用 StorageNode 需要打开相应的 FeatureGate：

```shell
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts --set operator.featureGates.storageNode=true
```

#### 字段说明

##### 必填配置 

配置项 |  描述 | 类型 | 示例 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`spec.storageProviderName` | StorageProvider 名称 |  string  | `aws-rds-instance` 

##### 选填配置

配置项 |  描述 | 类型 | 示例 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`spec.storageProviderSchema` | 初始化 Schema  | string | `sharding_db`
`spec.replicas` | Aurora 集群规模  | number | 2

#### 示例

以下是一个 AWS RDS Aurora 对应的 StorageNode 配置说明，可以拉起相应的 Aurora 集群：

```yaml
apiVersion: shardingsphere.apache.org/v1alpha1
kind: StorageNode
metadata:
  name: storage-node-with-aurora-example
  annotations:
    "storageproviders.shardingsphere.apache.org/cluster-identifier": "storage-node-with-aurora-example"
    "storageproviders.shardingsphere.apache.org/instance-db-name": "test_db"
    # 以下是自动注册所需的 Annotations 
    "shardingsphere.apache.org/register-storage-unit-enabled": "false" # 如果需要自动注册，请设置为 `true` 
    "shardingsphere.apache.org/logic-database-name": "sharding_db"
    "shardingsphere.apache.org/compute-node-name": "shardingsphere-operator-shardingsphere-proxy"
spec:
  schema: "test_db"
  storageProviderName: aws-aurora-cluster-mysql-5.7
  replicas: 2 # 目前仅 Aurora 有效
```
### StorageProvider

StorageProvider 声明了不同的 StorageNode 提供方，比如 AWS RDS 和 CloudNative PG。

#### 字段说明

##### 必填配置 

配置项 |  描述 | 类型 | 示例 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`spec.storageProviderName` | StorageProvider 名称 |  string  | `aws-rds-instance` 

#### 示例

以下声明了一个 AWS Aurora 的 StorageProvider，其中包含了对相关属性的设置：

```yaml
apiVersion: shardingsphere.apache.org/v1alpha1
kind: StorageProvider
metadata:
  name: aws-aurora-cluster-mysql-5.7
spec:
  provisioner: storageproviders.shardingsphere.apache.org/aws-aurora
  reclaimPolicy: Delete
  parameters:
    masterUsername: "root"
    masterUserPassword: "root123456"
    instanceClass: "db.t3.small"
    engine: "aurora-mysql"
    engineVersion: "5.7"
```

## 清理

```shell
helm uninstall shardingsphere-cluster -n shardingsphere-operator
```

## 下一步
为了使用创建好的 ShardingSphere Proxy 集群，您需要使用 [DistSQL](https://shardingsphere.apache.org/document/current/cn/user-manual/shardingsphere-proxy/distsql/usage/) 去配置相应的资源及规则，如数据库资源，分片规则 等等。
