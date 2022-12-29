+++
pre = "<b>3.2 </b>"
title = "ShardingSphere-Cluster Operator 简明用户手册"
weight = 2
chapter = true
+++

## 安装 ShardingSphere-Cluster Operator

如下配置内容和配置文件目录为：apache-shardingsphere-cluster-operator-charts/values.yaml。

### 在线安装

```shell
 kubectl create ns shardingsphere-operator
 helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
 helm repo update
 helm install shardingsphere-cluster shardingsphere/apache-shardingsphere-cluster-operator-charts -n shardingsphere-operator
```

### 源码安装

```shell
kubectl create ns shardingsphere-operator
cd charts/apache-shardingsphere-cluster-operator-charts/
helm dependency build
cd ../
helm install shardingsphere-cluster apache-shardingsphere-cluster-operator-charts -n shardingsphere-operator
```

## 参数

### 通用参数
| Name              | Description                                                                                               | Value                                 |
|-------------------|-----------------------------------------------------------------------------------------------------------|---------------------------------------|
| `nameOverride`    | nameOverride String to partially override common.names.fullname template (will maintain the release name) | `apache-shardingsphere-proxy-cluster` |

### ShardingSphere-Cluster Operator 参数
| Name                              | Description                                                                                                | Value                                                                   |
|-----------------------------------|------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| `operator.replicaCount`           | operator replica count                                                                                     | `2`                                                                     |
| `operator.image.repository`       | operator image name                                                                                        | `ghcr.io/apache/shardingsphere-on-cloud/apache-shardingsphere-operator` |
| `operator.image.pullPolicy`       | image pull policy                                                                                          | `IfNotPresent`                                                          |
| `operator.image.tag`              | image tag                                                                                                  | `0.1.1`                                                                 |
| `operator.imagePullSecrets`       | image pull secret of private repository                                                                    | `[]`                                                                    |
| `operator.resources`              | operator Resources required by the operator                                                                | `{}`                                                                    |
| `operator.health.healthProbePort` | operator health check port                                                                                 | `8081`                                                                  |

### ShardingSphere-Proxy Cluster 参数

| Name                                             | Description                                                                                                                                                                                        | Value                                    |
|--------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------|
| `proxyCluster.replicaCount`                      | ShardingSphere-Proxy cluster starts the number of replicas, Note: After you enable automaticScaling, this parameter will no longer take effect                                                     | `3`                                      |
| `proxyCluster.proxyVersion`                      | ShardingSphere-Proxy cluster version                                                                                                                                                               | `5.2.0`                                  |
| `proxyCluster.automaticScaling.enable`           | ShardingSphere-Proxy Whether the ShardingSphere-Proxy cluster has auto-scaling enabled                                                                                                             | `false`                                  |
| `proxyCluster.automaticScaling.scaleUpWindows`   | ShardingSphere-Proxy automatically scales the stable window                                                                                                                                        | `30`                                     |
| `proxyCluster.automaticScaling.scaleDownWindows` | ShardingSphere-Proxy automatically shrinks the stabilized window                                                                                                                                   | `30`                                     |
| `proxyCluster.automaticScaling.target`           | ShardingSphere-Proxy auto-scaling threshold, the value is a percentage, note: at this stage, only cpu is supported as a metric for scaling                                                         | `20`                                     |
| `proxyCluster.automaticScaling.maxInstance`      | ShardingSphere-Proxy maximum number of scaled-out replicas                                                                                                                                         | `4`                                      |
| `proxyCluster.automaticScaling.minInstance`      | ShardingSphere-Proxy has a minimum number of boot replicas, and the shrinkage will not be less than this number of replicas                                                                        | `1`                                      |
| `proxyCluster.resources`                         | ShardingSphere-Proxy starts the requirement resource, and after opening automaticScaling, the resource of the request multiplied by the percentage of target is used to trigger the scaling action | `{}`                                     |
| `proxyCluster.service.type`                      | ShardingSphere-Proxy external exposure mode                                                                                                                                                        | `ClusterIP`                              |
| `proxyCluster.service.port`                      | ShardingSphere-Proxy exposes  port                                                                                                                                                                 | `3307`                                   |
| `proxyCluster.startPort`                         | ShardingSphere-Proxy boot port                                                                                                                                                                     | `3307`                                   |
| `proxyCluster.mySQLDriver.version`               | ShardingSphere-Proxy The ShardingSphere-Proxy mysql driver version will not be downloaded if it is empty                                                                                           | `5.1.47`                                 |

### 计算节点 ShardingSphere-Proxy ServerConfig 权限相关参数

| Name                                                    | Description                                                                                                                                    | Value                      |
|---------------------------------------------------------| ---------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------- |
| `proxyCluster.serverConfig.authority.privilege.type`    | authority provider for storage node, the default value is ALL_PERMITTED                                                                        | `ALL_PRIVILEGES_PERMITTED` |
| `proxyCluster.serverConfig.authority.users[0].password` | Password for compute node.                                                                                                                     | `root`                     |
| `proxyCluster.serverConfig.authority.users[0].user`     | Username,authorized host for compute node. Format: <username>@<hostname> hostname is % or empty string means do not care about authorized host | `root@%`                   |

### 计算节点 ShardingSphere-Proxy ServerConfig Mode 相关参数

| Name                                                                           | Description                                                         | Value                                                                  |
|--------------------------------------------------------------------------------| ------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `proxyCluster.serverConfig.mode.type`                                          | Type of mode configuration. Now only support Cluster mode           | `Cluster`                                                              |
| `proxyCluster.serverConfig.mode.repository.props.namespace`                    | Namespace of registry center                                        | `governance_ds`                                                        |
| `proxyCluster.serverConfig.mode.repository.props.server-lists`                 | Server lists of registry center                                     | `{{ printf "%s-zookeeper.%s:2181" .Release.Name .Release.Namespace }}` |
| `proxyCluster.serverConfig.mode.repository.props.maxRetries`                   | Max retries of client connection                                    | `3`                                                                    |
| `proxyCluster.serverConfig.mode.repository.props.operationTimeoutMilliseconds` | Milliseconds of operation timeout                                   | `5000`                                                                 |
| `proxyCluster.serverConfig.mode.repository.props.retryIntervalMilliseconds`    | Milliseconds of retry interval                                      | `500`                                                                  |
| `proxyCluster.serverConfig.mode.repository.props.timeToLiveSeconds`            | Seconds of ephemeral data live                                      | `600`                                                                  |
| `proxyCluster.serverConfig.mode.repository.type`                               | Type of persist repository. Now only support ZooKeeper              | `ZooKeeper`                                                            |
| `proxyCluster.serverConfig.mode.overwrite`                                     | Whether overwrite persistent configuration with local configuration | `true`                                                                 |
| `proxyCluster.serverConfig.props.proxy-frontend-database-protocol-type`        | Default startup protocol                                            | `MySQL`                                                                |

### ZooKeeper Chart 参数

| Name                                 | Description                                          | Value               |
| ------------------------------------ | ---------------------------------------------------- | ------------------- |
| `zookeeper.enabled`                  | Switch to enable or disable the ZooKeeper helm chart | `true`              |
| `zookeeper.replicaCount`             | Number of ZooKeeper nodes                            | `1`                 |
| `zookeeper.persistence.enabled`      | Enable persistence on ZooKeeper using PVC(s)         | `false`             |
| `zookeeper.persistence.storageClass` | Persistent Volume storage class                      | `""`                |
| `zookeeper.persistence.accessModes`  | Persistent Volume access modes                       | `["ReadWriteOnce"]` |
| `zookeeper.persistence.size`         | Persistent Volume size                               | `8Gi`               |

## 配置示例

apache-shardingsphere-cluster-operator-charts/values.yaml

```yaml
## @section Name parameters
## @param nameOverride String to partially override common.names.fullname template (will maintain the release name)
##
nameOverride: apache-shardingsphere-proxy-cluster

## @section ShardingSphere operator parameters
operator:
  ## @param replicaCount operator replica count
  ##
  replicaCount: 2
  image:
    ## @param image.repository operator image name
    ##
    repository: "ghcr.io/apache/shardingsphere-on-cloud/apache-shardingsphere-operator"
    ## @param image.pullPolicy image pull policy
    ##
    pullPolicy: IfNotPresent
    ## @param image.tag image tag
    ##
    tag: "0.1.1"
  ## @param imagePullSecrets image pull secret of private repository
  ## e.g:
  ## imagePullSecrets:
  ##   - name: mysecret
  ##
  imagePullSecrets: {}
  ## @param resources operator Resources required by the operator
  ## e.g:
  ## resources:
  ##   limits:
  ##     cpu: 2
  ##   limits:
  ##     cpu: 2
  ##
  resources: {}
  ## @param health.healthProbePort operator health check port
  ##
  health:
    healthProbePort: 8081


## @section ShardingSphere-Proxy cluster parameters
proxyCluster:
  enabled: true
  ## @param replicaCount ShardingSphere-Proxy cluster starts the number of replicas, Note: After you enable automaticScaling, this parameter will no longer take effect
  ## @param proxyVersion ShardingSphere-Proxy cluster version
  ##
  replicaCount: "3"
  proxyVersion: "5.3.0"
  ## @param automaticScaling.enable ShardingSphere-Proxy Whether the ShardingSphere-Proxy cluster has auto-scaling enabled
  ## @param automaticScaling.scaleUpWindows ShardingSphere-Proxy automatically scales the stable window
  ## @param automaticScaling.scaleDownWindows ShardingSphere-Proxy automatically shrinks the stabilized window
  ## @param automaticScaling.target ShardingSphere-Proxy auto-scaling threshold, the value is a percentage, note: at this stage, only cpu is supported as a metric for scaling
  ## @param automaticScaling.maxInstance ShardingSphere-Proxy maximum number of scaled-out replicas
  ## @param automaticScaling.minInstance ShardingSphere-Proxy has a minimum number of boot replicas, and the shrinkage will not be less than this number of replicas
  ##
  automaticScaling:
    enable: false
    scaleUpWindows: 30
    scaleDownWindows: 30
    target: 20
    maxInstance: 4
    minInstance: 1
  ## @param resources ShardingSphere-Proxy starts the requirement resource, and after opening automaticScaling, the resource of the request multiplied by the percentage of target is used to trigger the scaling action
  ## e.g:
  ## resources:
  ##   limits:
  ##     cpu: 2
  ##     memory: 2Gi
  ##   requests:
  ##     cpu: 2
  ##     memory: 2Gi
  ##
  resources: {}
  ## @param service.type ShardingSphere-Proxy external exposure mode
  ## @param service.port ShardingSphere-Proxy exposes  port
  ##
  service:
    type: ClusterIP
    port: 3307
  ## @param startPort ShardingSphere-Proxy boot port
  ##
  startPort: 3307
  ## @param mySQLDriver.version ShardingSphere-Proxy The ShardingSphere-Proxy mysql driver version will not be downloaded if it is empty
  ##
  mySQLDriver:
    version: "5.1.47"
  ## @param imagePullSecrets ShardingSphere-Proxy pull private image repository key
  ## e.g:
  ## imagePullSecrets:
  ##   - name: mysecret
  ##
  imagePullSecrets: []
  ## @section  ShardingSphere-Proxy ServerConfiguration parameters
  ## NOTE: If you use the sub-charts to deploy Zookeeper, the server-lists field must be "{{ printf \"%s-zookeeper.%s:2181\" .Release.Name .Release.Namespace }}",
  ## otherwise please fill in the correct zookeeper address
  ## The server.yaml is auto-generated based on this parameter.
  ## If it is empty, the server.yaml is also empty.
  ## ref: https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-jdbc/yaml-config/mode/
  ## ref: https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-jdbc/builtin-algorithm/metadata-repository/
  ##
  serverConfig:
    ## @section Compute-Node ShardingSphere-Proxy ServerConfiguration authority parameters
    ## NOTE: It is used to set up initial user to login compute node, and authority data of storage node.
    ## @param serverConfig.authority.privilege.type authority provider for storage node, the default value is ALL_PERMITTED
    ## @param serverConfig.authority.users[0].password Password for compute node.
    ## @param serverConfig.authority.users[0].user Username,authorized host for compute node. Format: <username>@<hostname> hostname is % or empty string means do not care about authorized host
    ##
    authority:
      privilege:
        type: ALL_PERMITTED
      users:
        - password: root
          user: root@%
    ## @section Compute-Node ShardingSphere-Proxy ServerConfiguration mode Configuration parameters
    ## @param serverConfig.mode.type Type of mode configuration. Now only support Cluster mode
    ## @param serverConfig.mode.repository.props.namespace Namespace of registry center
    ## @param serverConfig.mode.repository.props.server-lists Server lists of registry center
    ## @param serverConfig.mode.repository.props.maxRetries Max retries of client connection
    ## @param serverConfig.mode.repository.props.operationTimeoutMilliseconds Milliseconds of operation timeout
    ## @param serverConfig.mode.repository.props.retryIntervalMilliseconds Milliseconds of retry interval
    ## @param serverConfig.mode.repository.props.timeToLiveSeconds Seconds of ephemeral data live
    ## @param serverConfig.mode.repository.type Type of persist repository. Now only support ZooKeeper
    ## @param serverConfig.props.proxy-frontend-database-protocol-type Default startup protocol
    mode:
      repository:
        props:
          maxRetries: 3
          namespace: governance_ds
          operationTimeoutMilliseconds: 5000
          retryIntervalMilliseconds: 500
          server-lists: "{{ printf \"%s-zookeeper.%s:2181\" .Release.Name .Release.Namespace }}"
          timeToLiveSeconds: 600
        type: ZooKeeper
      type: Cluster
    props:
      proxy-frontend-database-protocol-type: MySQL
  ## @section ZooKeeper chart parameters

## ZooKeeper chart configuration
## https://github.com/bitnami/charts/blob/master/bitnami/zookeeper/values.yaml
##
zookeeper:
  ## @param zookeeper.enabled Switch to enable or disable the ZooKeeper helm chart
  ##
  enabled: true
  ## @param zookeeper.replicaCount Number of ZooKeeper nodes
  ##
  replicaCount: 2
  ## ZooKeeper Persistence parameters
  ## ref: https://kubernetes.io/docs/user-guide/persistent-volumes/
  ## @param zookeeper.persistence.enabled Enable persistence on ZooKeeper using PVC(s)
  ## @param zookeeper.persistence.storageClass Persistent Volume storage class
  ## @param zookeeper.persistence.accessModes Persistent Volume access modes
  ## @param zookeeper.persistence.size Persistent Volume size
  ##
  persistence:
    enabled: false
    storageClass: ""
    accessModes:
      - ReadWriteOnce
    size: 8Gi

## 清理

```shell
helm uninstall shardingsphere-cluster -n shardingsphere-operator
kubectl delete crd shardingsphereproxies.shardingsphere.apache.org shardingsphereproxyserverconfigs.shardingsphere.apache.org
```

## 下一步
为了使用创建好的 shardingsphere-proxy cluster，您需要使用 [DistSQL](https://shardingsphere.apache.org/document/current/cn/user-manual/shardingsphere-proxy/distsql/usage/) 去配置相应的资源及规则，如数据库资源，分片规则 等等。
