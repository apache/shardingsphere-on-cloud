# Apache ShardingSphere-Proxy Charts

This Chart is used to quickly install ShardingSphere-Proxy Cluster without ShardingSphere-Operator.

## Install

Use the following command to install:

```
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-proxy-charts --version 0.3.0
```

## Uninstall

Use the following command to uninstall:

```
helm unstall [RELEASE_NAME]
```

## Parameters

### Governance-Node parameters


| Name                 | Description                                           | Value  |
| ---------------------- | ------------------------------------------------------- | -------- |
| `governance.enabled` | Switch to enable or disable the governance helm chart | `true` |

### Governance-Node ZooKeeper parameters


| Name                                             | Description                                          | Value               |
| -------------------------------------------------- | ------------------------------------------------------ | --------------------- |
| `governance.zookeeper.enabled`                   | Switch to enable or disable the ZooKeeper helm chart | `true`              |
| `governance.zookeeper.replicaCount`              | Number of ZooKeeper nodes                            | `1`                 |
| `governance.zookeeper.persistence.enabled`       | Enable persistence on ZooKeeper using PVC(s)         | `false`             |
| `governance.zookeeper.persistence.storageClass`  | Persistent Volume storage class                      | `""`                |
| `governance.zookeeper.persistence.accessModes`   | Persistent Volume access modes                       | `["ReadWriteOnce"]` |
| `governance.zookeeper.persistence.size`          | Persistent Volume size                               | `8Gi`               |
| `governance.zookeeper.resources.limits`          | The resources limits for the ZooKeeper containers    | `{}`                |
| `governance.zookeeper.resources.requests.memory` | The requested memory for the ZooKeeper containers    | `256Mi`             |
| `governance.zookeeper.resources.requests.cpu`    | The requested cpu for the ZooKeeper containers       | `250m`              |

### Compute-Node ShardingSphere-Proxy parameters


| Name                                | Description                                                  | Value                         |
| ------------------------------------- | -------------------------------------------------------------- | ------------------------------- |
| `compute.image.repository`          | Image name of ShardingSphere-Proxy.                          | `apache/shardingsphere-proxy` |
| `compute.image.pullPolicy`          | The policy for pulling ShardingSphere-Proxy image            | `IfNotPresent`                |
| `compute.image.tag`                 | ShardingSphere-Proxy image tag                               | `5.5.0`                       |
| `compute.imagePullSecrets`          | Specify docker-registry secret names as an array             | `[]`                          |
| `compute.resources.limits`          | The resources limits for the ShardingSphere-Proxy containers | `{}`                          |
| `compute.resources.requests.memory` | The requested memory for the ShardingSphere-Proxy containers | `2Gi`                         |
| `compute.resources.requests.cpu`    | The requested cpu for the ShardingSphere-Proxy containers    | `200m`                        |
| `compute.replicas`                  | Number of cluster replicas                                   | `3`                           |
| `compute.service.type`              | ShardingSphere-Proxy network mode                            | `ClusterIP`                   |
| `compute.service.port`              | ShardingSphere-Proxy expose port                             | `3307`                        |
| `compute.mysqlConnector.version`    | MySQL connector version                                      | `5.1.49`                      |
| `compute.startPort`                 | ShardingSphere-Proxy start port                              | `3307`                        |
| `compute.serverConfig`              | Server Configuration file for ShardingSphere-Proxy           | `""`                          |
| `compute.agent.enabled`             | switch to enable or disable the agent metrics                | `false`                       |
| `compute.agent.config`              | agent Configuration file for ShardingSphere-Proxy agent      | `""`                          |
