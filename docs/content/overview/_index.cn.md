+++
pre = "<b>1. </b>"
title = "概览"
weight = 1
chapter = true
+++

## 什么是 ShardingSphere-on-cloud ?

ShardingSphere-on-Cloud 项目是面向 Apache ShardingSphere 的云上解决方案集合，帮助用户在云环境下利用 ShardingSphere 将任意数据库转换为分布式数据库，并通过数据分片、弹性伸缩、加密等能力对原有数据库进行增强。

### 功能特性

#### 云上一键部署

不管是公有云还是 Kubernetes，社区都提供了相应的一键部署方案：如果是 AWS 等公有云环境，提供了如 CloudFormation Stack 模板、Terraform 等自动化部署脚本。如果是 Kubernetes 环境，提供了 Helm Charts、Operator CRD 等声明式配置。

#### 数据库可靠性工程 

数据库可靠性工程是为了提高数据基础设施稳定性的一套工具和最佳实践的集合。用户可以利用 ShardingSphere 完成对数据库的增强，并利用 Operator 实现 ShardingSphere 自动化运维，实现自动弹性扩容、高可用部署、可观测性等能力。

#### 混沌工程

混沌工程通过破坏性测试发现系统脆弱环节。ShardingSphere Operator 可以向 ShardingSphere 注入各类故障，如网络故障、Pod 故障、CPU 压力和内存压力等，从而帮助用户验证在各类场景下的可用性。

#### 生态扩展

Java 语言实现的 ShardingSphere 有着良好的社区生态。ShardingSphere 同时支持 Java 语言实现 SPI 和非 Java 语言编写 WebAssembly 插件分片算法、加解密等进行扩展，满足各类自定义场景需求。

## 应用场景

如果以下场景刚好也是你所期待的，那么请赶快试用吧：

* 希望快速了解、试用或验证 ShardingSphere Proxy 的功能特性，可以选择在 AWS 上使用 CloudFormation 和 Terraform 进行部署。或者选择在 Kubernetes 上使用 Helm Charts 和 Operator 进行部署。

* 希望在 Kubernetes 环境能实现数据分片并管理数据库，可以选择 ShardingSphere Operator 为您同时管理计算节点和存储节点。

* 希望在 Kubernetes 解决数据库可靠性工程相关问题，可以选择 ShardingSphere Operator 提供的自动化运维能力。

* 希望对 ShardingSphere 能力进行自定义，编写自定义的 WebAssembly 插件实现扩展。
