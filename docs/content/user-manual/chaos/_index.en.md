+++
pre = "<b>4.6 </b>"
title = "Chaos Engineering in ShardingSphere"
weight = 6
chapter = true
+++

## Abstract

For automatically chaos engineering, ShardingSphere Operator supports a CRD named Chaos and build corresponding chaos platform CRD manifest with the attributions contained by the Chaos CRD. It supports Chaos Mesh currently, and is going to support LitmusChaos in future releases.

![](../../../img/user-manual/chaos-concepts-1.png)

## Installation of Operator

Please refer to ShardingSphere Operator user manual's Operator Installation chapter.

## CRD Introduction

### Chaos

#### Operator Configuration

It need to be activated by enabling the responding FeatureGate:

```shell
helm install [RELEASE_NAME] shardingsphere/apache-shardingsphere-operator-charts --set operator.featureGates.chaos=true
```

#### Parameters 

##### Required Paramaters

Name |  Description| Type |Example 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`metadata.name` | Name |  string | `foo` 
`metadata.namespace` | Namespace，default 'default '| string |                                      | `shardingsphere-system`

##### Optional Parameters

Name |  Description | Type | Example 
------------------ | --------------------------|------------------------------------------------------ | ----------------------------------------
`spec.podChaos.selector.namespaces` | Pod Selector: Namespace|  []string | 
`spec.podChaos.selector.labelSelectors` | Pod Selector：Labels|  map[string]string | 
`spec.podChaos.selector.annotationSelectors` | Pod Selector：Annotations|  map[string]string | 
`spec.podChaos.selector.nodes` | Pod Selector：Node|  []string | 
`spec.podChaos.selector.pods` | Pod Selector：Pod | map[string][]string| 
`spec.podChaos.selector.nodeSelectors` | Pod Selector：NodeSelector| map[string]string | 
`spec.podChaos.selector.expressionSelectors` | Pod Selector：ExpressionSelector|  []metav1.LabelSelectorRequirement | 
`spec.podChaos.action` | PodChaos Type，including PodFailure、ContainerKill、PodKill、CPUStress、MemoryStress|  PodChaosAction | `PodFailure` 
`spec.podChaos.params.podFailure.duration` | PodFailure Duration| string  |  `1m`
`spec.podChaos.params.containerKill.containerNames` | ContainerKill target container names| []string   | `shardingsphere-proxy` 
`spec.podChaos.params.podKill.gracePeriod` | PodKill graceful period time |  number | `0` 
`spec.podChaos.params.cpuStress.duration` | CPU duration | string  |  `1m`
`spec.podChaos.params.cpuStress.cores` | CPU cores| number | `2` 
`spec.podChaos.params.cpuStress.load` | CPU load| number  | `50` 
`spec.podChaos.params.memoryStress.duration` | Memory duration|  string | `1m` 
`spec.podChaos.params.memoryStress.workers` | Memory workers |  numbers | `2`
`spec.podChaos.params.memoryStress.consumption` | Memory consumption|  string | `50`
`spec.networkChaos.source.namespaces` | Pod Selector：namespace|  []string | 
`spec.networkChaos.soruce.labelSelectors` | Pod Selector: labels|  map[string]string | 
`spec.networkChaos.source.annotationSelectors` | Pod Selector: annotations|  map[string]string | 
`spec.networkChaos.source.nodes` | Pod Selector: node|  []string | 
`spec.networkChaos.source.pods` | Pod Selector：Pod | map[string][]string| 
`spec.networkChaos.source.nodeSelectors` | Pod Selector：NodeSlector| map[string]string | 
`spec.networkChaos.source.expressionSelectors` | Pod Selector：ExpressionSelector|  []metav1.LabelSelectorRequirement | 
`spec.networkChaos.target.namespaces` | Pod Selector：Namespace|  []string | 
`spec.networkChaos.target.labelSelectors` | Pod Selector：Labels|  map[string]string | 
`spec.networkChaos.target.annotationSelectors` | Pod Selector：Annotations|  map[string]string | 
`spec.networkChaos.target.nodes` | Pod Selector：Node|  []string | 
`spec.networkChaos.target.pods` | Pod Selector：Pod | map[string][]string| 
`spec.networkChaos.target.nodeSelectors` | Pod Selector：NodeSelector| map[string]string | 
`spec.networkChaos.target.expressionSelectors` | Pod Selector：ExpressionSelector|  []metav1.LabelSelectorRequirement | 
`spec.networkChaos.action.` | NetworkChaos type，including Delay，Loss，Duplication，Corruption，Partition，Bandwidth  |  string | `50`
`spec.networkChaos.duration.` | Duration|  string | `1m`
`spec.networkChaos.direction.` | Traffic direction，including to、from 和 both |  string | `both`
`spec.networkChaos.params.deplay.latency` | Packet delay|  string | `100`
`spec.networkChaos.params.loss.loss` | Packet loss |  string | `80`
`spec.networkChaos.params.duplicate.duplicate` | Packet duplication |  string | `80`
`spec.networkChaos.params.corrupt.corrupt` | Packet Corrupt|  string | `80`

##### Annotations Introduction 

While using PodChaos and NetworkChaos, some parameters need to be setup with annotations according to the difference of chaos platform, such as:

* Select target ComputeNode: selector.chaos-mesh.org/mode: one
* Select target traffic: target-selector.chaos-mesh.org/mode: all

#### Example 

Here is a example of PodChaos which injects CPU Stress:

```yaml
apiVersion: shardingsphere.apache.org/v1alpha1
kind: Chaos
metadata:
  name: cpu-chaos
  annotations:
    selector.chaos-mesh.org/mode: one
spec:
  podChaos:
    selector:
      labelSelectors:
        app: foo
      namespaces: 
      - "default"
    params:
      cpuStress:
        duration: 1m
        cores: 2
        load: 50
    action: "CPUStress"
```
