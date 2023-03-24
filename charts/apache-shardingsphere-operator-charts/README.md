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

## Try ComputeNode
Use the following command to install:
```shell
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts --set operator.featureGates.computeNode=true --set proxyCluster.enabled=false
```

## Parameters
### Common parameters
| Name              | Description                                                                                               | Value                  |
|-------------------|-----------------------------------------------------------------------------------------------------------|------------------------|
| `nameOverride`    | nameOverride String to partially override common.names.fullname template (will maintain the release name) | `shardingsphere-proxy` |

### ShardingSphere Operator Parameters
| Name                              | Description                                 | Value                                                                   |
|-----------------------------------| ------------------------------------------- |-------------------------------------------------------------------------|
| `operator.replicaCount`           | operator replica count                      | `2`                                                                     |
| `operator.image.repository`       | operator image name                         | `apache/shardingsphere-operator` |
| `operator.image.pullPolicy`       | image pull policy                           | `IfNotPresent`                                                          |
| `operator.image.tag`              | image tag                                   | `0.2.0`                                                                 |
| `operator.imagePullSecrets`       | image pull secret of private repository     | `[]`                                                                    |
| `operator.resources`              | operator Resources required by the operator | `{}`                                                                    |
| `operator.health.healthProbePort` | operator health check port                  | `8081`                                                                  |

### ShardingSphere ProxyCluster Parameters

| Name                                             | Description                                                                                                                                                                                        | Value       |
|--------------------------------------------------| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |-------------|
| `proxyCluster.replicaCount`                      | ShardingSphere-Proxy cluster starts the number of replicas, Note: After you enable automaticScaling, this parameter will no longer take effect                                                     | `3`         |
| `proxyCluster.proxyVersion`                      | ShardingSphere-Proxy cluster version                                                                                                                                                               | `5.3.1`     |
| `proxyCluster.automaticScaling.enable`           | ShardingSphere-Proxy Whether the ShardingSphere-Proxy cluster has auto-scaling enabled                                                                                                             | `false`     |
| `proxyCluster.automaticScaling.scaleUpWindows`   | ShardingSphere-Proxy automatically scales the stable window                                                                                                                                        | `30`        |
| `proxyCluster.automaticScaling.scaleDownWindows` | ShardingSphere-Proxy automatically shrinks the stabilized window                                                                                                                                   | `30`        |
| `proxyCluster.automaticScaling.target`           | ShardingSphere-Proxy auto-scaling threshold, the value is a percentage, note: at this stage, only cpu is supported as a metric for scaling                                                         | `20`        |
| `proxyCluster.automaticScaling.maxInstance`      | ShardingSphere-Proxy maximum number of scaled-out replicas                                                                                                                                         | `4`         |
| `proxyCluster.automaticScaling.minInstance`      | ShardingSphere-Proxy has a minimum number of boot replicas, and the shrinkage will not be less than this number of replicas                                                                        | `1`         |
| `proxyCluster.resources`                         | ShardingSphere-Proxy starts the requirement resource, and after opening automaticScaling, the resource of the request multiplied by the percentage of target is used to trigger the scaling action | `{}`        |
| `proxyCluster.service.type`                      | ShardingSphere-Proxy external exposure mode                                                                                                                                                        | `ClusterIP` |
| `proxyCluster.service.port`                      | ShardingSphere-Proxy exposes  port                                                                                                                                                                 | `3307`      |
| `proxyCluster.startPort`                         | ShardingSphere-Proxy boot port                                                                                                                                                                     | `3307`      |
| `proxyCluster.mySQLDriver.version`               | ShardingSphere-Proxy The ShardingSphere-Proxy mysql driver version will not be downloaded if it is empty                                                                                           | `5.1.47`    |


### ShardingSphere ProxyCluster ServerConfiguration Authority Parameters

| Name                                                    | Description                                                                                                                                    | Value                      |
|---------------------------------------------------------| ---------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------- |
| `proxyCluster.serverConfig.authority.privilege.type`    | authority provider for storage node, the default value is ALL_PERMITTED                                                                        | `ALL_PRIVILEGES_PERMITTED` |
| `proxyCluster.serverConfig.authority.users[0].password` | Password for compute node.                                                                                                                     | `root`                     |
| `proxyCluster.serverConfig.authority.users[0].user`     | Username,authorized host for compute node. Format: <username>@<hostname> hostname is % or empty string means do not care about authorized host | `root@%`                   |


### ShardingSphere ProxyCluster ServerConfiguration Mode Parameters

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


### ZooKeeper Parameters

| Name                                 | Description                                          | Value               |
| ------------------------------------ | ---------------------------------------------------- | ------------------- |
| `zookeeper.enabled`                  | Switch to enable or disable the ZooKeeper helm chart | `true`              |
| `zookeeper.replicaCount`             | Number of ZooKeeper nodes                            | `1`                 |
| `zookeeper.persistence.enabled`      | Enable persistence on ZooKeeper using PVC(s)         | `false`             |
| `zookeeper.persistence.storageClass` | Persistent Volume storage class                      | `""`                |
| `zookeeper.persistence.accessModes`  | Persistent Volume access modes                       | `["ReadWriteOnce"]` |
| `zookeeper.persistence.size`         | Persistent Volume size                               | `8Gi`               |


### ShardingSphere ComputeNode Parameters

| Name                                        | Description                                                                                            | Value               |
| --------------------------------------------| ------------------------------------------------------------------------------------------------------ | ------------------- |
| `computeNode.storageNodeConnector.type`     | ShardingSphere-Proxy driver type                                                                       | `mysql`             |
| `computeNode.storageNodeConnector.version`  | ShardingSphere-Proxy driver version. The MySQL driver need to be downloaded according to this version  | `5.1.47`            |
| `computeNode.serverVersion`                 | ShardingSphere-Proxy cluster version                                                                   | `5.3.1`             |
| `computeNode.portBindings[0].name`          | ShardingSphere-Proxy port name                                                                         | `3307`              |
| `computeNode.portBindings[0].containerPort` | ShardingSphere-Proxy port for container                                                                | `3307`              |
| `computeNode.portBindings[0].servicePort`   | ShardingSphere-Proxy port for service                                                                  | `3307`              |
| `computeNode.portBindings[0].procotol`      | ShardingSphere-Proxy port protocol                                                                     | `TCP`               |
| `computeNode.serviceType`                   | ShardingSphere-Proxy service type                                                                      | `ClusterIP`         |


### ShardingSphere ComputeNode Bootstrap Parameters

| Name                                                                           | Description                                                         | Value                                                                  |
|--------------------------------------------------------------------------------| ------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `computeNode.bootstrap.serverConfig.authority.privilege.type`    | authority provider for storage node, the default value is ALL_PERMITTED                                                                        | `ALL_PRIVILEGES_PERMITTED` |
| `computeNode.bootstrap.serverConfig.authority.users[0].user`     | Username,authorized host for compute node. Format: <username>@<hostname> hostname is % or empty string means do not care about authorized host | `root@%`                   |
| `computeNode.bootstrap.serverConfig.authority.users[0].password` | Password for compute node.                                                                                                                     | `root`                     |
| `computeNode.bootstrap.serverConfig.mode.type`                                          | Type of mode configuration. Now only support Cluster mode           | `Cluster`                                                              |
| `computeNode.bootstrap.serverConfig.mode.repository.type`                               | Type of persist repository. Now only support ZooKeeper              | `ZooKeeper`                                                            |
| `computeNode.bootstrap.mode.repository.props.timeToLiveSeconds`            | Seconds of ephemeral data live                                      | `600`                                                                  |
| `computeNode.bootstrap.serverConfig.mode.repository.props.serverlists`                 | Server lists of registry center                                     | `{{ printf "%s-zookeeper.%s:2181" .Release.Name .Release.Namespace }}` |
| `computeNode.bootstrap.serverConfig.mode.repository.props.retryIntervalMilliseconds`    | Milliseconds of retry interval                                      | `500`                                                                  |
| `computeNode.bootstrap.serverConfig.mode.repository.props.operationTimeoutMilliseconds` | Milliseconds of operation timeout                                   | `5000`                                                                 |
| `computeNode.bootstrap.serverConfig.mode.repository.props.namespace`                    | Namespace of registry center                                        | `governance_ds`                                                        |
| `computeNode.bootstrap.serverConfig.mode.repository.props.maxRetries`                   | Max retries of client connection                                    | `3`                                                                    |
| `computeNode.bootstrap.serverConfig.mode.overwrite`                                     | Whether overwrite persistent configuration with local configuration | `true`                                                                 |
| `computeNode.bootstrap.serverConfig.props.proxy-frontend-database-protocol-type`        | Default startup protocol                                            | `MySQL`                                                                |
