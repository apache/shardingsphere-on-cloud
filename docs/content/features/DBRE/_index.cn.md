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

StorageNode 的定义中包含了部署公有云上的 RDS 数据库实例相关配置，通过 StorageProvider 来指定相应的公有云资源提供方，并完成对云上数据库实例的创建、自动注册、自动注销和删除的能力。现已支持

## 自动弹性扩容

Kubernetes 社区提供了水平弹性扩容控制器（Horizontal Pod Autoscaler，HPA），可以实现基于 CPU 和内存的自动水平弹性扩容，也可以配合 Prometheus Adapter 实现基于自定义指标的扩容能力。而对于 AWS EC2 虚机部署的场景，社区也提供了机遇 AutoScalingGroup 的自动扩容能力，并通过 TargetGroup 的探活机制来保证只有就绪（Ready）的实例可以接收到业务流量。
## 可观测性

ShardingSphere Proxy 可以通过 ShardingSphere Agent 采集并暴露相关运行指标，Agent 相关信息可以参考[链接](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/observability/)。ShardingSphere on Cloud 提供了相应的 Grafana 模板，其中包含了基础的资源监控、JVM 监控和 ShardingSphere 运行时指标。将不同层级的指标绘制在同一个 Dashboard 中有助于用户快速定位可能的性能问题。

## 混沌工程

混沌工程是验证系统健壮性的有效手段，也可以帮助开发人员发现未知的缺陷。ShardingSphere Operator 支持 CRD Chaos，并可以针对 ShardingSphere Proxy 注入 Pod 异常、CPU 压力、内存压力和网络异常等不同类型的故障。详细信息参见混沌工程章节。
