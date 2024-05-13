+++
pre = "<b>4.2 </b>"
title = "ShardingSphere Operator User Manual"
weight = 2
chapter = true
+++

## Overview 

ShardingSphere Operator is a practical implementation of the Kubernetes Operator model. It transforms the maintenance experience of ShardingSphere Proxy into an executable program and leverages Kubernetes's declarative and "reconcile" features for implementation.

ShardingSphere Operator abstracts computing nodes, storage nodes, and even chaos faults as Kubernetes Custom Resource Definitions (CRDs). Users are responsible for writing corresponding CRD configurations, while the Operator executives and ensures the desired state.

Please refer to the 'Operator Installation' section to get installed and try it out, and refer to the 'CRD Introduction' to get a better understanding of CRD's configuration.

## Operator Installation

Operator currently supports Helm Charts rapid deployment, configuration content and configuration file directory is: apache-shardingsphere-operator-charts. Users can adopt online installation or source code installation depending on their needs. 

### Online Installation

```shell
 kubectl create ns shardingsphere-operator
 helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
 helm repo update
 helm install shardingsphere-cluster shardingsphere/apache-shardingsphere-operator-charts -n shardingsphere-operator --set zookeeper.persistence.enabled=false
```

### Source Code Installation

```shell
kubectl create ns shardingsphere-operator
cd charts/apache-shardingsphere-operator-charts/
helm dependency build
cd ../
helm install shardingsphere-cluster apache-shardingsphere-operator-charts -n shardingsphere-operator --set zookeeper.persistence.enabled=false
```

### Charts Parameters Instruction

#### Common Parameters
| Name |  Description  | Value |
|-------------------|-----------------------------------------------------------------------------------------------------------|---------------------------------------|
| `nameOverride`    | nameOverride String to partially override common.names.fullname template will maintain the release name | `shardingsphere-proxy` |

#### ShardingSphere Operator Parameters
| Name | Description | Value|
|-----------------------------------|------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| `operator.replicaCount`           | Operstor replica count| `1`                                                                     |
| `operator.image.repository`       | Operator image name| `apache/shardingsphere-operator` |
| `operator.image.pullPolicy`       | Image pull policy                                                                                         | `IfNotPresent`                                                          |
| `operator.image.tag`              | Image tag| `0.3.0`                                                                 |
| `operator.imagePullSecrets`       | Image pull secret of private repository| `[]`                                                                    |
| `operator.resources`              | Operator resources required by the operator| `{}`                                                                    |
| `operator.health.healthProbePort` | Operator health check pork| `8080`                                                                  |

Users can choose whether to install the supporting management center depending on their needs when using Operator Charts for installation. The relevant parameters are as follows:

| Name | Description | Value|
| ------------------------------------ | ---------------------------------------------------- | ------------------- |
| `zookeeper.enabled`                  | Switch to eable or disable the ZooKeeper helm chart| `true`              |
| `zookeeper.replicaCount`             | Zookeeper replica count                            | `1`                 |
| `zookeeper.persistence.enabled`      | Enable persistence on Zookeeper using PVC(s)         | `false`             |
| `zookeeper.persistence.storageClass` | Persistent Volume Storage Class | `""`                |
| `zookeeper.persistence.accessModes`  | Persistent Volume access modes| `["ReadWriteOnce"]` |
| `zookeeper.persistence.size`         | Persistent Volume size| `8Gi`               |

Note: currently the persist repository installed by Charts, supports Bitnami Zookeeper Charts only.

## CRD Introduction

### ShardingSphereProxy 

ShardingSphereProxy and ShardingSphereProxyServerConfig provide a basic description of deployment and configuration of ShardingSphereProxy, Operator submits and adds configuration provided by CRD convert to Kubernetes workload. ShardingSphereProxy effect relevant configuration of basic resource, and ShardingSphereProxyServerConfig effect `server.yaml`.

Note: ShardingSphereProxy and ShardingSphereProxyServerConfig planning end of support since 0.4.0 version.

 
#### Column Comment


##### Programmatic Configuration 

ShardingSphereProxy

Configuration item |  Description | Type | Examples 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`.spec.version`  | ShardingSphere-Proxy version | string | `5.5.0`
`.spec.serviceType.type` | Service type |  string | `NodePort`
`.spec.serviceType.nodePort` | Node Port service | number | `33307`
`.spec.replicas` | Operstor replica count | number | `3` 
`.spec.proxyConfigName` | Mounted configuration  | string  |
`.spec.port` | Exposed port  | number |

##### Optional Configuration 

Configuration item |  Description | Type | Examples 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`.spec.automaticScaling.enable` | Automatic scaling enable  | bool | `false` 
`.spec.automaticScaling.scaleUpWindows` |  Maximum automatic scale limit  | number | 
`.spec.automaticScaling.scaleDownWindows` | Minimum automatic scale limit | number |
`.spec.automaticScaling.target` | Automatic scaling target | number |
`.spec.automaticScaling.maxInstance` | Automatic scaling maxmum instance  | number |
`.spec.automaticScaling.minInstance` | Automatic scaling minimum instance | number |
`.spec.customMetrics` | Custom metrics | []autoscalingv2beta2.MetricSpec | 
`.spec.imagePullSecrets` | Image pull secrets  | v1.Local,ObjectReference | 
`.spec.mySQLDriver.version` | MySQL driver version | string |  
`.spec.resources` | Resources configuration| v1.ResourceRequirements | 
`.spec.livenssProbe` | Liveness probe | v1.Probe |
`.spec.readinessProbe` | Readness probe | v1.Probe |
`.spec.startupProbe` |  Startup probe| v1.Probe |


ShardingSphereProxyServerConfig

Configuration item |  Description | Type |  Examples
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`.spec.mode.type` | string | Type of mode configuration, supports Standalone and Cluster | string | `Cluster`
`.spec.mode.repository.type` | string  | Type of persist repository, supports ZooKeeper and Etcd  |string              | `ZooKeeper`
`.spec.mode.repository.props.namespace` | string  |Namespace of registry center(Not namespace of K8s) | `governance_ds`
`.spec.mode.repository.props.server-lists` |  string  | Server lists of registry center                                    | `zookeeper.default:2181` 
`.spec.mode.repository.props.retryIntervalMilliseconds` | number  | Milliseconds of retry interval                                      | `500`
`.spec.mode.repository.props.maxRetries` | number  | Max retries of client connection                                  | `3`
`.spec.mode.repository.props.timeToLiveSeconds` | number  | TTL                                        | `600`
`.spec.mode.repository.props.operationTimetoutMilliseconds` | number |  Milliseconds of operation timeout                                 | `5000`
`.spec.mode.repository.props.digest` | Abstract | string |  
`.spec.authority.users[0].user` |  Username, authorized host for compute node. Format: <username>@<hostname>, hostname is % or empty string means do not care about authorized host|string |`root@%`
`.spec.authority.users[0].password` | Username, authorized host for compute node. Format: <username>@<hostname>, hostname is % or empty string means do not care about authorized host|string |`root@%`
`.spec.authority.priviliege.type`  | Authority priviliege for compute node, the default value is ALL_PRIVILEGES_PERMITTED  | string                                                                        | `ALL_PRIVILEGES_PERMITTED` 
`.spec.props.kernel-executor-size` | Kernel executor size | number | 
`.spec.props.check-table-metadata-enabled` | Check table metadata enabled | bool | 
`.spec.props.proxy-backend-query-fetch-size` | Back end query fetch size  | number | 
`.spec.props.check-duplicate-table-enabled`|Check duplicate table enable | bool| 
`.spec.props.proxy-frontend-executeor-size` | Front end executor size | number | 
`.spec.props.proxy-backend-executor-suitable` |Back end executor suitable | string | 
`.spec.props.proxy-backend-driver-type` |Back end driver type | string | 
`.spec.props.proxy-frontend-database-protocol-type` | Front end database protocol type | string | 

#### Examples

ShardingSphereProxy example:

```yaml
apiVersion: shardingsphere.apache.org/v1alpha1
kind: ShardingSphereProxy
metadata:
  name: shardingsphere-cluster-shardingsphere-proxy
  namespace: shardingsphere-operator
spec:
  version: 5.5.0
  serviceType:
    type: ClusterIP
  replicas: 3
  proxyConfigName: "shardingsphere-cluster-shardingsphere-proxy-configuration"
  port: 3307
  mySQLDriver:
    version: "5.1.47"
```

ShardingSphereProxyServerConfig example:

```yaml
apiVersion: shardingsphere.apache.org/v1alpha1
kind: ShardingSphereProxyServerConfig
metadata:
  name: shardingsphere-cluster-shardingsphere-proxy-configuration
  namespace: shardingsphere-operator
spec:
  authority:
    privilege:
      type: ALL_PERMITTED
    users:
    - password: root
      user: root@%
  mode:
    repository:
      props:
        maxRetries: 3
        namespace: governance_ds
        operationTimeoutMilliseconds: 5000
        retryIntervalMilliseconds: 500
        server-lists: 'shardingsphere-cluster-zookeeper.shardingsphere-operator:2181'
        timeToLiveSeconds: 600
      type: ZooKeeper
    type: Cluster
  props:
    proxy-frontend-database-protocol-type: MySQL
```


### ComputeNode 

ComputeNode describes the computing nodes in the ShardingSphere cluster, usually referred to as Proxy. Since ShardingSphere Proxy is a stateless application, it can be managed using Kubernetes' native workload Deployment, meanwhile using ConfigMap and Service to configure startup configuration and service discovery. Using ComputeNode can not only unify key configurations in Deployment, ConfigMap, and Service, but also match the semantics of ShardingSphere, helping Operators quickly lock workloads. As shown in the picture:

![](../../../img/user-manual/cn-concepts-1.png)

#### Operator Configuration

Currently, the Operator use ComputeNode needs to open featureGate that is relevant:

```shell
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts --set operator.featureGates.computeNode=true --set proxyCluster.enabled=false
```

#### Column Comment

##### Programmatic Configuration

Configuration item|  Description | Type | Examples 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`metadata.name` | Name of deployment plan name |  string | `foo` 
`metadata.namespace` | Default namespace of deployment plan | string |                                      | `shardingsphere-system`
`spec.storageNodeConnector.type`     | Back end driver type | string | `mysql`
`spec.storageNodeConnector.version`  | Back end driver version| string  | `5.1.47`
`spec.serverVersion`                 | ShardingSphere-Proxy version | string | `5.5.0`
`spec.replicas `     | Deployment plan instance |  number | `3`
`spec.selectors`     | Instance selector, same as Deployment.Spec.Selectors |  number | `3`
`spec.portBindings[0].name`          | Name of exposed port  | string |                                                                         | `3307`
`spec.portBindings[0].containerPort` | Exposed container port | number |`3307`
`spec.portBindings[0].servicePort`   | Exposed container port | number                                                                 | `3307`
`spec.portBindings[0].procotol`      | Exposed port procotol| string|  `TCP`
`spec.serviceType`                   | Exposed port type | string                                                                     | `ClusterIP`
`spec.bootstrap.serverConfig.authority.privilege.type`    | Authority priviliege for compute node, the default value is ALL_PRIVILEGES_PERMITTED | string                                                                        | `ALL_PRIVILEGES_PERMITTED` 
`spec.bootstrap.serverConfig.authority.users[0].user`     | Username, authorized host for compute node. Format: <username>@<hostname> hostname is % or empty string means do not care about authorized host|string |`root@%`
`spec.bootstrap.serverConfig.authority.users[0].password` | Password of compute node |string                                                                                                                     | `root`
`spec.bootstrap.serverConfig.mode.type`                                          | Type of mode configuration, supports Standalone and Cluster          | string | `Cluster`
`spec.bootstrap.serverConfig.mode.repository.type`                               | Type of persist repository, supports ZooKeeper and Etcd  |string              | `ZooKeeper`
`spec.bootstrap.serverConfig.mode.repository.props`            |Registry center properties configuration, refer to [Common ServerConfig Repository Props](#Common\ ServerConfig\ Repository\ Props\ Configuration)  | map[string]string                                    | 

##### Common ServerConfig Repository Props Configuration
Configuration item |  Description | Examples 
------------------ | -------------------------------------------------------------------------------- | ----------------------------------------
`spec.bootstrap.serverConfig.mode.repository.props.timeToLiveSeconds`            | TTL                                        | `600`
`spec.bootstrap.serverConfig.mode.repository.props.serverlists`                 | Server lists of registry center                                    | `zookeeper.default:2181` 
`spec.bootstrap.serverConfig.mode.repository.props.retryIntervalMilliseconds`    | Milliseconds of retry interval                                       | `500`
`spec.bootstrap.serverConfig.mode.repository.props.operationTimeoutMilliseconds` | Millisecond of operation timeout                                   | `5000`
`spec.bootstrap.serverConfig.mode.repository.props.namespace`                    | Namespace of registry center(Not namespace of K8s)                                       | `governance_ds`
`spec.bootstrap.serverConfig.mode.repository.props.maxRetries`                   | Max retries of client connection                                   | `3`


##### Optional Configuration  

Configuration item |  Description | Type | Examples 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`spec.probes.livenessProbe` | Liveness probe |  corev1.Probe | 
`spec.probes.readinessProbe` | Readiness probe |  corev1.Probe | 
`spec.probes.startupProbe` | Startup probe |  corev1.Probe | 
`spec.imgaePullSecrets ` | Image pull secrets | corev1.LocalObjectReference  | 
`spec.env` | Environment variable | corev1.Env | 
`spec.resources` | Resources| corev1.ResourceRequirements | 
`spec.bootstrap.agentConfig.plugins.logging.file.props` | Agent configuration plugins logging file properties| map[string]string |
`spec.bootstrap.agentConfig.plugins.metrics.prometheus.host` | Agent configuration plugins metrics prometheus host| map[string]string |
`spec.bootstrap.agentConfig.plugins.metrics.prometheus.port` | Agent configuration plugins metrics prometheus port| map[string]string |
`spec.bootstrap.agentConfig.plugins.metrics.prometheus.props` | Agent configuration plugins metrics prometheus properties| map[string]string |
`spec.bootstrap.agentConfig.plugins.tracing.openTracing.props` | Agent configuration plugins tracing opentracing properties| map[string]string |
`spec.bootstrap.agentConfig.plugins.tracing.openTelemetry.props` | Agent configuration plugins tracing opentelemetry properties| map[string]string |

#### Instance Configuration

The following is a basic instance configuration of ComputeNode CRD, which can pull up a three-node file server cluster in ShardingSphere Proxy.

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
  serverVersion: 5.5.0
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
Note:  A ZooKeeper cluster in normal operation is a prerequisite.

### StorageNode

StorageNode is the Operator's description of the data source and provides data source lifecycle management. Its use needs to cooperate with StorageProvider, and now supports AWS RDS and CloudNative PG. As shown in the picture:

![](../../../img/user-manual/sn-concepts-1.png)

Note: StorageNode is an optional CRD, and users can decide whether to manage data sources through StorageNode depending on real situation.

#### Operator Configuration

Currently, the Operator use StorageNode needs to open featureGate that is relevant:

```shell
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts --set operator.featureGates.storageNode=true --set operator.storageNodeProviders.aws.region='' --set operator.storageNodeProviders.aws.accessKeyId='' --set operator.storageNodeProviders.aws.secretAccessKey='' --set operator.storageNodeProviders.aws.enabled=true
```

#### Column Comment

##### Programmatic Configuration 

Configuration items |  Description | Type | Examples 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`metadata.name` | Name of deployment plan |  string | `foo` 
`metadata.namespace` | Default namespace of deployment plan | string |                                      | `shardingsphere-system`
`spec.storageProviderName` | Name of provisioner |  string  | `aws-rds-instance` 

##### Optional Configuration

Configuration item |  Description | Type | Examples 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`spec.storageProviderSchema` |  Schema initialize | string | `sharding_db`
`spec.replicas` | Aurora cluster size  | number | 2

#### Examples

The following is a StorageNode configuration introduction for AWS RDS Aurora, which can pull up Aurora cluster:

```yaml
apiVersion: shardingsphere.apache.org/v1alpha1
kind: StorageNode
metadata:
  name: storage-node-with-aurora-example
  annotations:
    "storageproviders.shardingsphere.apache.org/cluster-identifier": "storage-node-with-aurora-example"
    "storageproviders.shardingsphere.apache.org/instance-db-name": "test_db"
    # The following annotations are required for auto registration.
    "shardingsphere.apache.org/register-storage-unit-enabled": "false" # If it needs auto registration, please set up 'true'.
    "shardingsphere.apache.org/logic-database-name": "sharding_db"
    "shardingsphere.apache.org/compute-node-name": "shardingsphere-operator-shardingsphere-proxy"
spec:
  schema: "test_db"
  storageProviderName: aws-aurora-cluster-mysql-5.7
  replicas: 2 # Currently, only AWS Aurora is efficient.
```
### StorageProvider

StorageProvider declares some different suppliers of StorageNode, such as AWS RDS and CloudNative PG.  

#### Column Comment

##### Programmatic Configuration

Configuration item |  Descrption | Type | Examples 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`metadata.name` | Name of deployment plan |  string | `foo` 
`spec.storageProviderName` | Name of provisioner |  string  | `aws-rds-instance` 

#### Examples

The following declares a StorageProvider in AWS Aurora, including the setting of relevant features:

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

## Clean

```shell
helm uninstall shardingsphere-cluster -n shardingsphere-operator
```

## Next
To use the created ShardingSphere-Proxy cluster, you need to use [DistSQL](https://shardingsphere.apache.org/document/current/cn/user-manual/shardingsphere-proxy/distsql/usage/) to configure corresponding resources and rules, such as database resources, sharding rules, etc.
