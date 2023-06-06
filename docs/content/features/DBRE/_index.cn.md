+++
pre = "<b>3.2 </b>"
title = "数据库可靠性工程"
weight = 2
chapter = true
+++

## 概述

数据库可靠性工程（Database Reliability Engineering，DBRE）类似站点可靠性工程（Site Reliability Engineering，SRE），关注通过各种技术手段提升数据库相关服务的稳定性。对于部署在 Kubernetes 之上的 ShardingSphere 来说，需要借助 Operator 进一步实现 DBRE。

## 高可用部署

ShardingSphere Proxy 本身是无状态的，作为计算节点用来处理客户端发来的 SQL 并完成相关数据计算。因此，Operator 通过 ComputeNode 对 ShardingSphere Proxy 进行抽象和描述。目前，ShardingSphere Proxy 因其无状态性，可以使用 Deployment 部署模式进行。Deployment 是 Kubernetes 提供的基础部署方式，其所管理的 Pod 之间并无差异。ShardingSphere 通过 Deployment 进行部署，可以获得健康检查、就绪检查等基础自愈能力，以及滚动升级和历史版本回滚等特性。

ComputeNode 的定义中包含了部署 ShardingSphere Proxy 所需要的各类属性，如副本数、镜像仓库信息和版本信息、数据库驱动信息、健康检查和就绪检查探针等，还包含了端口映射规则、服务启动需要的 Server.yaml、logback.xml ，Agent 相关的配置等信息。这些信息在 Operator 的调谐过程中会分别渲染为 Kubernetes Deployment、Service 和 ConfigMap，并自动完成相关绑定和挂载动作。

借助于 Deployment 的能力，除了可以轻松实现多副本部署外，还可以实现亲和性、污点容忍等高级调度特征，这些都为 ShardingSphere Proxy 提供了基础高可用能力。

## 自动弹性扩容

Kubernetes 社区提供了水平弹性扩容控制器（Horizontal Pod Autoscaler，HPA），可以实现基于 CPU 和内存的自动水平弹性扩容，也可以配合 Prometheus Adapter 实现基于自定义指标的扩容能力。

## 可观测性

ShardingSphere Proxy 可以通过 ShardingSphere Agent 采集并暴露相关运行指标，ShardingSphere on Cloud 提供了相应的 Grafana 模板。Agent 相关信息可以参考[链接](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/observability/)。
