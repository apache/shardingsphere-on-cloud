# Apache ShardingSphere-Operator Charts
The Apache ShardingSphere-Operator is used to quickly install an Apache ShardingSphere-Proxy Cluster. 

## Install
Use the following command to install:
```shell
helm repo add shardingsphere https://apache.github.io/shardingsphere-on-cloud
helm repo update
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts --version 0.1.0 
```

## Uninstall 
Use the following command to uninstall:
```shell
helm unstall [RELEASE_NAME]
```

## Parameters
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

## Next
For the next step, please search for `apache-shardingsphere-operator-cluster-charts` in artifacthub.io to complete the install of ShardingSphere Cluster, 
