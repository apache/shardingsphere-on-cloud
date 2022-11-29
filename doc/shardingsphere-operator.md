# ShardingSphere-Operator Concise User Manual

## Install ShardingSphere-Operator

Configuration [below](#shardingsphere-operator-parameters) configuration content, configuration file location apache-shardingsphere-operator-charts/values.yaml
run

### Online Install
```shell
kubectl create ns shardingsphere-operator
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
helm install shardingsphere-operator shardingsphere/apache-shardingsphere-operator-charts --version 0.1.0 -n shardingsphere-operator
```

### Source Code Install 
```shell
kubectl create ns shardingsphere-operator
cd charts/apache-shardingsphere-operator-charts/
helm dependency build
cd ../
helm install shardingsphere-operator apache-shardingsphere-operator-charts -n shardingsphere-operator
```

## Install ShardingSphere-Proxy cluster

Configuration [below](#shardingsphere-proxy-cluster-parameters) configuration content, configuration file location apache-shardingsphere-operator-cluster-charts/values.yaml
run 

## Online Install
```shell
kubectl create ns shardingsphere
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
helm install shardingsphere shardingsphere/apache-shardingsphere-operator-cluster-charts --version 0.1.0 -n shardingsphere
```

### Source Code Install
```shell
kubectl create ns shardingsphere
cd charts/apache-shardingsphere-operator-cluster-charts
helm dependency build
cd ../
helm install shardingsphere apache-shardingsphere-operator-cluster-charts -n shardingsphere
```

## Online Install ShardingSphere-Proxy cluster && ShardingSphere-Operator
```shell
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
kubectl create ns  shardingsphere-operator
helm install shardingsphere-operator shardingsphere/apache-shardingsphere-operator-charts --version 0.1.0 -n shardingsphere-operator
kubectl create ns  shardingsphere
helm install shardingsphere shardingsphere/apache-shardingsphere-operator-cluster-charts --version 0.1.0 -n shardingsphere
```

##  Parameters
### ShardingSphere Operator Parameters

| Name                     | Description                                 | Value                     |
| ------------------------ | ------------------------------------------- | ------------------------- |
| `replicaCount`           | operator replica count                      | `2`                       |
| `image.repository`       | operator image name                         | `sahrdingsphere-operator` |
| `image.pullPolicy`       | image pull policy                           | `IfNotPresent`            |
| `image.tag`              | image tag                                   | `0.0.1`                   |
| `imagePullSecrets`       | image pull secret of private repository     | `[]`                      |
| `resources`              | operator Resources required by the operator | `{}`                      |
| `webhook.port`           | operator webhook boot port                  | `9443`                    |
| `health.healthProbePort` | operator health check port                  | `8081`                    |


### ShardingSphere-Proxy Cluster Parameters

| Name                                | Description                                                                                                                                                                                        | Value       |
| ----------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| `replicaCount`                      | ShardingSphere-Proxy cluster starts the number of replicas, Note: After you enable automaticScaling, this parameter will no longer take effect                                                     | `3`         |
| `proxyVersion`                      | ShardingSphere-Proxy cluster version                                                                                                                                                               | `5.2.0`     |
| `automaticScaling.enable`           | ShardingSphere-Proxy Whether the ShardingSphere-Proxy cluster has auto-scaling enabled                                                                                                             | `false`     |
| `automaticScaling.scaleUpWindows`   | ShardingSphere-Proxy automatically scales the stable window                                                                                                                                        | `30`        |
| `automaticScaling.scaleDownWindows` | ShardingSphere-Proxy automatically shrinks the stabilized window                                                                                                                                   | `30`        |
| `automaticScaling.target`           | ShardingSphere-Proxy auto-scaling threshold, the value is a percentage, note: at this stage, only cpu is supported as a metric for scaling                                                         | `20`        |
| `automaticScaling.maxInstance`      | ShardingSphere-Proxy maximum number of scaled-out replicas                                                                                                                                         | `4`         |
| `automaticScaling.minInstance`      | ShardingSphere-Proxy has a minimum number of boot replicas, and the shrinkage will not be less than this number of replicas                                                                        | `1`         |
| `resources`                         | ShardingSphere-Proxy starts the requirement resource, and after opening automaticScaling, the resource of the request multiplied by the percentage of target is used to trigger the scaling action | `{}`        |
| `service.type`                      | ShardingSphere-Proxy external exposure mode                                                                                                                                                        | `ClusterIP` |
| `service.port`                      | ShardingSphere-Proxy exposes  port                                                                                                                                                                 | `3307`      |
| `startPort`                         | ShardingSphere-Proxy boot port                                                                                                                                                                     | `3307`      |
| `mySQLDriver.version`               | ShardingSphere-Proxy The ShardingSphere-Proxy mysql driver version will not be downloaded if it is empty                                                                                           | `5.1.47`    |


### Compute-Node ShardingSphere-Proxy ServerConfiguration Authority Parameters

| Name                                       | Description                                                                                                                                    | Value                      |
| ------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------- |
| `serverConfig.authority.privilege.type`    | authority provider for storage node, the default value is ALL_PERMITTED                                                                        | `ALL_PRIVILEGES_PERMITTED` |
| `serverConfig.authority.users[0].password` | Password for compute node.                                                                                                                     | `root`                     |
| `serverConfig.authority.users[0].user`     | Username,authorized host for compute node. Format: <username>@<hostname> hostname is % or empty string means do not care about authorized host | `root@%`                   |


### Compute-Node ShardingSphere-Proxy ServerConfiguration Mode Configuration Parameters

| Name                                                              | Description                                                         | Value                                                                  |
| ----------------------------------------------------------------- | ------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `serverConfig.mode.type`                                          | Type of mode configuration. Now only support Cluster mode           | `Cluster`                                                              |
| `serverConfig.mode.repository.props.namespace`                    | Namespace of registry center                                        | `governance_ds`                                                        |
| `serverConfig.mode.repository.props.server-lists`                 | Server lists of registry center                                     | `{{ printf "%s-zookeeper.%s:2181" .Release.Name .Release.Namespace }}` |
| `serverConfig.mode.repository.props.maxRetries`                   | Max retries of client connection                                    | `3`                                                                    |
| `serverConfig.mode.repository.props.operationTimeoutMilliseconds` | Milliseconds of operation timeout                                   | `5000`                                                                 |
| `serverConfig.mode.repository.props.retryIntervalMilliseconds`    | Milliseconds of retry interval                                      | `500`                                                                  |
| `serverConfig.mode.repository.props.timeToLiveSeconds`            | Seconds of ephemeral data live                                      | `600`                                                                  |
| `serverConfig.mode.repository.type`                               | Type of persist repository. Now only support ZooKeeper              | `ZooKeeper`                                                            |
| `serverConfig.mode.overwrite`                                     | Whether overwrite persistent configuration with local configuration | `true`                                                                 |
| `serverConfig.props.proxy-frontend-database-protocol-type`        | Default startup protocol                                            | `MySQL`                                                                |


### ZooKeeper Chart Parameters

| Name                                 | Description                                          | Value               |
| ------------------------------------ | ---------------------------------------------------- | ------------------- |
| `zookeeper.enabled`                  | Switch to enable or disable the ZooKeeper helm chart | `true`              |
| `zookeeper.replicaCount`             | Number of ZooKeeper nodes                            | `1`                 |
| `zookeeper.persistence.enabled`      | Enable persistence on ZooKeeper using PVC(s)         | `false`             |
| `zookeeper.persistence.storageClass` | Persistent Volume storage class                      | `""`                |
| `zookeeper.persistence.accessModes`  | Persistent Volume access modes                       | `["ReadWriteOnce"]` |
| `zookeeper.persistence.size`         | Persistent Volume size                               | `8Gi`               |


## Configuration Example

apache-shardingsphere-operator-charts/values.yaml

```yaml
## @section ShardingSphere-Proxy operator parameters
## @param replicaCount operator  replica count
##
replicaCount: 2
image:
  ## @param image.repository operator image name
  ##
  repository: "sahrdingsphere-operator"
  ## @param image.pullPolicy image pull policy
  ##
  pullPolicy: IfNotPresent
  ## @param image.tag image tag
  ##
  tag: "0.0.1"
## @param imagePullSecrets image pull secret of private repository
## e.g:
## imagePullSecrets:
##   - name: mysecret
##
imagePullSecrets: []
## @param resources operator Resources required by the operator
## e.g:
## resources:
##   limits:
##     cpu: 2
##   limits:
##     cpu: 2
##
resources: {}
## @param webhook.port operator webhook boot port
##
webhook:
  port: 9443
## @param health.healthProbePort operator health check port
##
health:
  healthProbePort: 8081
```

apache-shardingsphere-operator-cluster-charts/values.yaml

```yaml
# @section ShardingSphere-Proxy cluster parameters
## @param replicaCount ShardingSphere-Proxy cluster starts the number of replicas, Note: After you enable automaticScaling, this parameter will no longer take effect
## @param proxyVersion ShardingSphere-Proxy cluster version
##
replicaCount: "3"
proxyVersion: "5.2.0"
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
##   requests:
##     cpu: 2
##
resources:
  limits:
    cpu: '2'
  requests:
    cpu: '1'
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
      type: ALL_PRIVILEGES_PERMITTED
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
  ## @param serverConfig.mode.overwrite Whether overwrite persistent configuration with local configuration
  ##
  mode:
    overwrite: true
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
  replicaCount: 3
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
```



## Clean

```shell
helm uninstall shardingsphere -n shardingsphere
helm uninstall shardingsphere-operator -n shardingsphere-operator
kubectl delete crd shardingsphereproxies.shardingsphere.apache.org shardingsphereproxyserverconfigs.shardingsphere.apache.org
```
