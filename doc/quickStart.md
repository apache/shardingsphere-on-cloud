# ShardingSphere-Operator 简明使用手册

⚠️⚠️⚠️⚠️ 现阶段因为 shardingsphere-proxy 5.1.2 还未发布。proxyConfig 功能还不能使用，请使用前自行挂载现阶段版本支持的 server.yaml ⚠️⚠️⚠️⚠️
## 配置
**Proxy.shardingsphere.sphere-ex.com/v1alpha1**

```yaml
apiVersion: shardingsphere.sphere-ex.com/v1alpha1
kind: Proxy
metadata:
  name: proxy-sample
spec:
  version: "5.1.1"
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
<span id=002>**用于 v5.1.1 测试的配置文件**</span>
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: sharding-proxy
  namespace: default
data:
  server.yaml: |-
    mode:
        type: Cluster
        repository:
          type: ZooKeeper
          props:
            namespace: governance_ds
            server-lists: zookeeper.default:2181
            retryIntervalMilliseconds: 500
            timeToLiveSeconds: 600
            maxRetries: 3
            operationTimeoutMilliseconds: 5000
        overwrite: false
    rules:
      - !AUTHORITY
        users:
          - root@%:root
        provider:
          type: ALL_PRIVILEGES_PERMITTED

```

## 安装 ShardingSphere-Operator 
按照[values.yaml](#001)中的配置完成对 charts/shardingsphere-operator/values.yaml  

执行 
```shell
kubectl create ns  shardingsphere-operator
helm install  shardingsphere-operator shardingsphere-operator -n shardingsphere-operator
```
## 安装 ShardingSphere-Proxy


首先新建一个[configmap.yaml](#002)
```shell
kubectl apply -f deploy/samples/shardingsphere_v1alpha1_proxy.yaml
kubectl apply -f  configmap.yaml
```
