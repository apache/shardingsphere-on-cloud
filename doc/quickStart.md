# ShardingSphere-Operator 简明使用手册

## 配置
**Proxy.shardingsphere.sphere-ex.com/v1alpha1**

```yaml
apiVersion: shardingsphere.sphere-ex.com/v1alpha1
kind: Proxy
metadata:
  name: proxy-sample
spec:
  version: "5.1.2"
  serviceType:
    type: ClusterIP
  replicas: 1
  proxyConfigName: "sharding-proxy"
  port: 3307
  mySQLDriver:
    version: "5.1.47"
  resources:
    limits:
      cpu: "2"
      memory: "2Gi"
    requests:
      cpu: "0.2"
      memory: "1.6Gi"
```

**ProxyConfig.shardingsphere.sphere-ex.com/v1alpha1**

```yaml
apiVersion: shardingsphere.sphere-ex.com/v1alpha1
kind: ProxyConfig
metadata:
  name: proxyconfig-sample
spec:
  authority:
    users:
      - user: root@%
        password: root
      - user: sphere
        password: sphere
    privilege:
      type: ALL_PRIVILEGES_PERMITTED
  mode:
    type: Cluster
    repository:
      type: ZooKeeper
      props:
        namespace: "governance_ds"
        server-lists: "<your zkAddr>"
        retryIntervalMilliseconds: 500
        timeToLiveSeconds: 600
        maxRetries: 3
        operationTimeoutMilliseconds: 5000
    overwrite: true
```


<span id="001">**values.yaml**</span>
```yaml
# Default values for proxy-operator.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 2

image:
  repository: "shardingsphere-operator"
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: "0.0.1"
imagePullSecrets: []
service:
  type: ClusterIP
  port: 80
resources:
   requests:
     cpu: 100m
     memory: 128Mi
webhook:
  serviceName: shardingsphere-operator-admission-webhook
  port: 9443
health:
  healthProbePort: 8081
serviceAccount:
  name: shardingsphere-operator
```


## 安装 ShardingSphere-Operator 
按照[values.yaml](#001)中的配置完成对 charts/shardingsphere-operator/values.yaml  

执行 
```shell
kubectl create ns  shardingsphere-operator
helm install  shardingsphere-operator shardingsphere-operator -n shardingsphere-operator
```
## 安装 ShardingSphere-Proxy

```shell
kubectl apply -f shardingsphere-operator/deploy/samples/shardingsphere_v1alpha1_proxy.yaml
kubectl apply -f shardingsphere-operator/deploy/samples/shardingsphere_v1alpha1_proxyconfig.yaml
```
