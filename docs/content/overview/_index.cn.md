+++
pre = "<b>1. </b>"
title = "概览"
weight = 1
chapter = true
+++

## 什么是 ShardingSphere-on-cloud ?

ShardingSphere-on-Cloud 项目是面向 Apache ShardingSphere 的云上解决方案集合，帮助用户在云环境下利用 ShardingSphere 将任意数据库转换为分布式数据库，并通过数据分片、弹性伸缩、加密等能力对原有数据库进行增强。

## 功能特性

### 云上一键部署

* 对于 AWS 等公有云环境，提供了如 CloudFormation Stack 模板、Terraform 等自动化部署脚本。
* 对于 Kubernetes 环境，提供了 Helm Charts、Operator 等   。

### 数据库可靠性工程 

利用 Operator 实现 ShardingSphere 自动化运维，提供自动弹性扩容、高可用部署、可观测性等能力。

### 混沌工程

通过 Operator 可以向 ShardingSphere 注入各类故障，如网络故障、Pod 故障、CPU 压力和内存压力等，从而验证在相关场景下的可用性。

### 生态扩展

ShardingSphere 同时支持 Java 语言实现 SPI 和非 Java 语言编写 WebAssembly 插件进行扩展。

## 应用场景

1. 希望快速了解、试用或验证 ShardingSphere Proxy 的功能特性，可以选择在 AWS 上使用 CloudFormation 和 Terraform 进行部署。或者选择在 Kubernetes 上使用 Helm Charts 和 Operator 进行部署。

2. 希望在 Kubernetes 环境能实现数据分片并管理数据库，可以选择 ShardingSphere Operator

3. 希望在 Kubernetes 解决数据库可靠性工程相关问题，可以选择 ShardingSphere Operator

4. 希望对 ShardingSphere 生态进行扩展，可以编写自定义的 WebAssembly 插件
