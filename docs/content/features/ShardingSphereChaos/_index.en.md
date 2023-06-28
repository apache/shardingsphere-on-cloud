+++
pre = "<b>3.3 </b>"
title = "ShardingSphere Chaos"
weight = 3
chapter = true
+++


## Overview

System availability is a critical metric for evaluating service reliability. There are numerous techniques to ensure availability, such as engineering resilience, anti-fragility, and others. 

However, disruptions in hardware and software can still occur, resulting in potential damage to the availability and robustness of the system. 

Chaos Engineering is a practice that aims to enhance system robustness by detecting the weaknesses in software systems, ultimately optimizing the ability to react to stresses and failures. According to the definition from [principleofchaos.org](https://principleofchaos.org/): 
> *Chaos Engineering is the discipline of experimenting on a system in order to build confidence in the system’s capability to withstand turbulent conditions in production.*

## General Principle

Chaos engineering generally involves five steps, which can be repeated if necessary: 
- defining a steady-state
- formulating hypotheses about the steady-state
- running chaos experiments
- verifying the results
- fixing the issue if necessary

To save time and increase teams' productivity, we suggest using Continuous Verification (CV) in chaos experiments, similar to Continuous Integration (CI). 

We also recommend introducing a diverse range of real-world events into the chaos experiments. While conducting experiments, minimize the blast radius to contain negative impact on a larger group of customers. 

## CustomResourceDefinitions (CRD) Chaos

ShardingSphere Operator supports `CustomResourceDefinitions` (CRD) chaos. The Operator supports multiple types of fault injection, for example, PodChaos including experiment actions like Pod Kill, Pod Failure, CPU Stress and Memory Stress, and NetworkChaos including network delay and loss. Once the basic parameters have been defined, Operator converts them into corresponding chaos experiments. For example:

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

If you are using Chaos Mesh as the Chaos Engineering platform, you will need to deploy it in Kubernetes as the test environment prior to creating and submitting ShardingSphere Chaos configuration files. For further information, please refer to the user manual.
