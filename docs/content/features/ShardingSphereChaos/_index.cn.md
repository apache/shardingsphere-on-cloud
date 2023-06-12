+++
pre = "<b>3.3 </b>"
title = "ShardingSphere 混沌工程"
weight = 3
chapter = true
+++

## 概述

系统可用性是衡量一个系统正确地对外提供服务（可工作）的能力，在实际生产中有很多方法来保证系统可用性： 韧性工程、反脆弱等。而故障是不可避免的，且可能在任意时间发生。不管是物理的、硬件的故障，还是软件的 Bug，最终都可能会损坏系统的可用性和健壮性。混沌工程是一种实践，它的目的是发现生产环境的弱点，并提升应对故障的能力，从而提高系统的健壮性。

混沌工程在 https://principleofchaos.org 的定义如下:
> 混沌工程是在系统上进行实验的学科， 目的是提升系统抵御生产环境中失控条件的能力以及增加信心

## 一般原则

混沌工程的执行一般需要按照稳态定义、提出假设、混沌实验、结果验证、修正提升五个步骤进行。整个实验的过程可以被重复进行，为了及时修正健壮性问题以及节约人力，推荐将混沌实验的过程设置为持续混沌验证，过程类似持续集成。

除此之外，混沌实验的执行还希望尽可能在生产环境进行，尽可能引入现实世界的事件，需要注意的是，在执行过程中要控制好爆炸半径，以免影响大量的生产用户。


## 自定义资源对象 Chaos

ShardingSphere Operator 支持自定义资源对象 Chaos，现支持 Pod Kill、Pod Failure 、CPU 压力、 Memory 压力等 Pod 混沌，同时支持延迟、丢包等网络混沌注入。Chaos 定义基本的混沌参数，并由 Operator 负责转换为对应的 Chaos 实现。例如：

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
      - foo-chaos
    params:
      cpuStress:
        duration: 1m
        cores: 2
        load: 50
    action: "CPUStress"
```

如果采用 Chaos Mesh 作为混沌平台，那么用户需要在用于测试的 Kubernetes 环境中预先部署 Chaos Mesh 组件，然后编写并提交 ShardingSphere Chaos 配置文件并执行实验。详细说明参见用户手册。
