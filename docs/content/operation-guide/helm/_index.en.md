+++
pre = "<b>3.1 </b>"
title = "ShardingSphere Helm Charts User Manual"
weight = 1
chapter = true
+++

## Procedure

### Online Installation

1. Add ShardingSphere-Proxy to the local Helm warehouse:

```shell
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
```

2. Install ShardingSphere-Proxy Charts:

```shell
helm install shardingsphere-proxy shardingsphere/apache-shardingsphere-proxy-charts 
```

### Source Code Installation

1. Charts can be configured and installed by default using the following command:

```shell
cd charts/apache-shardingsphere-proxy-charts/charts/governance
helm dependency build 
cd ../..
helm dependency build 
cd ..
helm install shardingsphere-proxy apache-shardingsphere-proxy-charts
```

Note: Please refer to the configuration description below for details.

2. Execute `helm list` to get the list of all installed releases.

### Uninstall

1. By default, all publishing records are deleted and can be retained by adding '-- keep history'.

```shell
helm uninstall shardingsphere-proxy
```

## Parameter Description

### Name parameters

| Name              | Description                                                                                                | Value                         |
|-------------------|------------------------------------------------------------------------------------------------------------|-------------------------------|
| `nameOverride   ` | nameOverride String to partially override common.names.fullname template (will maintain the release name)  | `apache-shardingsphere-proxy` |

### Governance Node Parameters

| Name                 | Description                                           | Value  |
| -------------------- | ----------------------------------------------------- | ------ |
| `governance.enabled` | Switch to enable or disable the governance helm chart | `true` |

### Governance Node ZooKeeper Parameters

| Name                                             | Description                                          | Value               |
| ------------------------------------------------ | ---------------------------------------------------- | ------------------- |
| `governance.zookeeper.enabled`                   | Switch to enable or disable the ZooKeeper helm chart | `true`              |
| `governance.zookeeper.replicaCount`              | Number of ZooKeeper nodes                            | `1`                 |
| `governance.zookeeper.persistence.enabled`       | Enable persistence on ZooKeeper using PVC(s)         | `false`             |
| `governance.zookeeper.persistence.storageClass`  | Persistent Volume storage class                      | `""`                |
| `governance.zookeeper.persistence.accessModes`   | Persistent Volume access modes                       | `["ReadWriteOnce"]` |
| `governance.zookeeper.persistence.size`          | Persistent Volume size                               | `8Gi`               |
| `governance.zookeeper.resources.limits`          | The resources limits for the ZooKeeper containers    | `{}`                |
| `governance.zookeeper.resources.requests.memory` | The requested memory for the ZooKeeper containers    | `256Mi`             |
| `governance.zookeeper.resources.requests.cpu`    | The requested cpu for the ZooKeeper containers       | `250m`              |

### Compute Node ShardingSphere-Proxy Parameters 

| Name                                | Description                                                  | Value                         |
| ----------------------------------- | ------------------------------------------------------------ |-------------------------------|
| `compute.image.repository`          | Image name of ShardingSphere-Proxy.                          | `apache/shardingsphere-proxy` |
| `compute.image.pullPolicy`          | The policy for pulling ShardingSphere-Proxy image            | `IfNotPresent`                |
| `compute.image.tag`                 | ShardingSphere-Proxy image tag                               | `5.2.0`                       |
| `compute.imagePullSecrets`          | Specify docker-registry secret names as an array             | `[]`                          |
| `compute.resources.limits`          | The resources limits for the ShardingSphere-Proxy containers | `{}`                          |
| `compute.resources.requests.memory` | The requested memory for the ShardingSphere-Proxy containers | `2Gi`                         |
| `compute.resources.requests.cpu`    | The requested cpu for the ShardingSphere-Proxy containers    | `200m`                        |
| `compute.replicas`                  | Number of cluster replicas                                   | `3`                           |
| `compute.service.type`              | ShardingSphere-Proxy network mode                            | `ClusterIP`                   |
| `compute.service.port`              | ShardingSphere-Proxy expose port                             | `3307`                        |
| `compute.mysqlConnector.version`    | MySQL connector version                                      | `5.1.49`                      |
| `compute.startPort`                 | ShardingSphere-Proxy start port                              | `3307`                        |
| `compute.serverConfig`              | Server Configuration file for ShardingSphere-Proxy            | `""`                          |

## Example

```yaml
## @section Name parameters
## @param nameOverride String to partially override common.names.fullname template (will maintain the release name)
nameOverride: apache-shardingsphere-proxy
## @section Governance-Node parameters
## @param governance.enabled Switch to enable or disable the governance helm chart
##
governance:
  enabled: true
  ## @section Governance-Node ZooKeeper parameters
  zookeeper:
    ## @param governance.zookeeper.enabled Switch to enable or disable the ZooKeeper helm chart
    ##
    enabled: true
    ## @param governance.zookeeper.replicaCount Number of ZooKeeper nodes
    ##
    replicaCount: 1
    ## ZooKeeper Persistence parameters
    ## ref: https://kubernetes.io/docs/user-guide/persistent-volumes/
    ## @param governance.zookeeper.persistence.enabled Enable persistence on ZooKeeper using PVC(s)
    ## @param governance.zookeeper.persistence.storageClass Persistent Volume storage class
    ## @param governance.zookeeper.persistence.accessModes Persistent Volume access modes
    ## @param governance.zookeeper.persistence.size Persistent Volume size
    ##
    persistence:
      enabled: false
      storageClass: ""
      accessModes:
        - ReadWriteOnce
      size: 8Gi
    ## ZooKeeper's resource requests and limits
    ## ref: https://kubernetes.io/docs/user-guide/compute-resources/
    ## @param governance.zookeeper.resources.limits The resources limits for the ZooKeeper containers
    ## @param governance.zookeeper.resources.requests.memory The requested memory for the ZooKeeper containers
    ## @param governance.zookeeper.resources.requests.cpu The requested cpu for the ZooKeeper containers
    ##
    resources:
      limits: {}
      requests:
        memory: 256Mi
        cpu: 250m

## @section Compute-Node parameters
## 
compute:
  ## @section Compute-Node ShardingSphere-Proxy parameters
  ## ref: https://kubernetes.io/docs/concepts/containers/images/
  ## @param compute.image.repository Image name of ShardingSphere-Proxy.
  ## @param compute.image.pullPolicy The policy for pulling ShardingSphere-Proxy image
  ## @param compute.image.tag ShardingSphere-Proxy image tag
  ##
  image:
    repository: "apache/shardingsphere-proxy"
    pullPolicy: IfNotPresent
    ## Overrides the image tag whose default is the chart appVersion.
    ##
    tag: "5.2.1"
  ## @param compute.imagePullSecrets Specify docker-registry secret names as an array
  ## e.g：
  ## imagePullSecrets:
  ##   - name: myRegistryKeySecretName
  ##
  imagePullSecrets: []
  ## ShardingSphere-Proxy resource requests and limits
  ## ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
  ## @param compute.resources.limits The resources limits for the ShardingSphere-Proxy containers
  ## @param compute.resources.requests.memory The requested memory for the ShardingSphere-Proxy containers
  ## @param compute.resources.requests.cpu The requested cpu for the ShardingSphere-Proxy containers
  ##
  resources:
    limits: {}
    requests:
      memory: 2Gi
      cpu: 200m
  ## ShardingSphere-Proxy Deployment Configuration
  ## ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/
  ## @param compute.replicas Number of cluster replicas
  ##
  replicas: 3
  ## @param compute.service.type ShardingSphere-Proxy network mode
  ## @param compute.service.port ShardingSphere-Proxy expose port
  ##
  service:
    type: ClusterIP
    port: 3307
  ## MySQL connector Configuration
  ## ref: https://shardingsphere.apache.org/document/current/en/quick-start/shardingsphere-proxy-quick-start/
  ## @param compute.mysqlConnector.version MySQL connector version
  ##
  mysqlConnector:
    version: "5.1.43"
  ## @param compute.startPort ShardingSphere-Proxy start port
  ## ShardingSphere-Proxy start port
  ## ref: https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/startup/docker/
  ##
  startPort: 3307
  ## @section Compute-Node ShardingSphere-Proxy ServerConfiguration parameters
  ## NOTE: If you use the sub-charts to deploy Zookeeper, the server-lists field must be "{{ printf \"%s-zookeeper.%s:2181\" .Release.Name .Release.Namespace }}",
  ## otherwise please fill in the correct zookeeper address
  ## The server.yaml is auto-generated based on this parameter.
  ## If it is empty, the server.yaml is also empty.
  ## ref: https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-jdbc/yaml-config/mode/
  ## ref: https://shardingsphere.apache.org/document/current/en/user-manual/common-config/builtin-algorithm/metadata-repository/
  ##
  serverConfig:
    ## @section Compute-Node ShardingSphere-Proxy ServerConfiguration authority parameters
    ## NOTE: It‘s used to set up initial user to login to the compute node, and  storage node authority data.
    ## ref: https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/yaml-config/authentication/
    ## @param compute.serverConfig.authority.privilege.type authority provider for storage node, the default value is ALL_PERMITTED
    ## @param compute.serverConfig.authority.users[0].password Password for compute node.
    ## @param compute.serverConfig.authority.users[0].user Username, authorized host for compute node. Format: <username>@<hostname> hostname is % or empty string means do not care about authorized host
    ##
    authority:
      privilege:
        type: ALL_PERMITTED
      users:
        - password: root
          user: root@%
    ## @section Compute-Node ShardingSphere-Proxy ServerConfiguration mode Configuration parameters
    ## @param compute.serverConfig.mode.type Type of mode configuration. Now only support Cluster mode
    ## @param compute.serverConfig.mode.repository.props.namespace Namespace of registry center
    ## @param compute.serverConfig.mode.repository.props.server-lists Server lists of registry center
    ## @param compute.serverConfig.mode.repository.props.maxRetries Max retries of client connection
    ## @param compute.serverConfig.mode.repository.props.operationTimeoutMilliseconds Milliseconds of operation timeout
    ## @param compute.serverConfig.mode.repository.props.retryIntervalMilliseconds Milliseconds of retry interval
    ## @param compute.serverConfig.mode.repository.props.timeToLiveSeconds Seconds of ephemeral data live
    ## @param compute.serverConfig.mode.repository.type Type of persist repository. Now only support ZooKeeper
    ##
    mode:
      type: Cluster
      repository:
        type: ZooKeeper
        props:
          maxRetries: 3
          namespace: governance_ds
          operationTimeoutMilliseconds: 5000
          retryIntervalMilliseconds: 500
          server-lists: "{{ printf \"%s-zookeeper.%s:2181\" .Release.Name .Release.Namespace }}"
          timeToLiveSeconds: 60
    ## @param compute.serverConfig.props.proxy-frontend-database-protocol-type proxy frontend database protocol type. Only support: PostgreSQL,openGauss,MariaDB,MySQL
    ##
    props:
      proxy-frontend-database-protocol-type: MySQL
```

