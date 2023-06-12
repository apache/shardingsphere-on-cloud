+++
pre = "<b>3.1 </b>"
title = "云上一键部署"
weight = 1
chapter = true
+++

## 概述

云计算的概念从 IaaS 到 PaaS 再到 SaaS 的转换，不仅是基础设施构成的变化，也是软件开发理念的升级。随着 Kubernetes 掀起的云原生浪潮，越来越多的应用开始采用云原生技术栈进行部署。对于 ShardingSphere 来说亦如此。想要在云环境下使用和体验 ShardingSphere，首先需要解决的是部署的问题，而云环境下的部署推荐使用基础设施即代码的方式进行实现。

### AWS 一键部署

在 AWS 上进行部署，首先需要了解 AWS 的各类资源和服务，比如 VPC、子网、安全组、弹性负载均衡器、域名、EC2、RDS、CloudWatch 等服务。为了快速描述清楚一套完整的 ShardingSphere 部署结构，可以采用基础设施即代码的方式进行描述，比如 AWS 官方推出的 CloudFormation 服务。该服务通过 json 或 yaml
模板的形式，对抽象的部署所需要的各类资源进行定义和组合，并由相关服务进行解释和执行。用户只需要编写相关描述文件，并且可以利用 Git 等版本控制工具将部署的代码进行管理和维护。

目前 Apache ShardingSphere 的 CloudFormation 托管在 ShardingSphere on Cloud 的仓库中。在 AWS Marketplace 上，可以获取相应的 AMI 信息，详见：[链接](https://us-east-1.console.aws.amazon.com/marketplace/home?region=ap-southeast-1#/subscriptions/ef146e06-20ca-4da4-8954-78a7c51b3c5a)。

在快速使用的章节中提供了利用 CloudFormation 最简配置在 AWS 启动一套 ShardingSphere Proxy 集群的方式。如果想了解更多 CloudFormation 的参数或者对 Terraform 比较熟悉，可以阅读用户手册相关章节进行了解。

### Kubernetes 一键部署

在 Kubernetes 上进行部署可以利用包管理器 Helm 完成。Helm 通过一组模板和变量声明组成的 Charts 描述部署结构，里面涉及的资源对象都是 Kubernetes 领域的工作负载，如 Deployment，Service，ConfigMap 等。每次版本更新都可以制作相应的 Charts 包，并且这些包可以提交到公开的制品仓库里，如 ArtifactHub 等。

目前 Apache ShardingSphere 每次发布新版本都会构建相应的 Helm Charts 包，相关的源码托管在 ShardingSphere on Cloud 仓库中。

在快速使用的章节中提供了利用 Helm Charts 最简配置在 Kubernetes 上启动一套 ShardingSphere Proxy 集群的方式，如果想了解更多 Charts 参数或者对 Operator 比较熟悉，可以阅读用户手册相关章节进行了解。

## 适用场景

在快速使用的章节中提供了在 AWS 和 Kubernetes 上启动一套 ShardingSphere Proxy 集群的方法，如果希望在测试环境中进行试用和功能验证，可以选择一键部署模式。如果希望在生产环境中使用，请先阅读用户手册章节并了解相关参数后再进行配置部署。
