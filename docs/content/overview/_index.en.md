+++
pre = "<b>1. </b>"
title = "Overview"
weight = 1
chapter = true
+++

## What is ShardingSphere-on-Cloud?

The ShardingSphere-on-Cloud project is a collection of cloud solutions for Apache ShardingSphere, including automated deployment scripts to virtual machines in AWS, GCP, Alibaba Cloud and other cloud environments. Such as CloudFormation Stack templates, Terraform one click deployment scripts. Helm Charts, Operators, automatic horizontal scaling and other tools in the Kubernetes cloud native environment, as well as high availability, observability, security compliance and other aspects.

## Core Concept

At present, the terms involved in this warehouse are from common cloud service providers and open source projects, and the relevant concepts and definitions are consistent.

- CloudFormation: It is a tool provided by AWS to help us quickly create cloud resources.
- CloudFormation Stack: It is a collection of AWS resources that can be managed as a unit.
- Terraform: It is an open source infrastructure management tool. Using the concept of "infrastructure as code", you can efficiently build, change and version infrastructure.
- Kubernetes: It is an open source container orchestration management platform, which can automatically deploy, manage and extend container applications.
- Operator: It is an extension software of Kubernetes, which uses [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) to manage applications and their components. The Operator follows the concept of Kubernetes, especially in [controller](https://kubernetes.io/docs/concepts/architecture/controller/).
- Helm Charts: Helm is a package management tool for the Kubernetes application, and Charts is a collection of files that describe a group of related Kubernetes resources.

## Infrastructure

- ShardingSphere-Operator Diagram

![Operator Diagram](../../../../img/overview/operator.png)

- ShardingSphere-Terraform Diagram

![Terraform Diagram](../../../../img/overview/terraform.png)

## Feature List

- ShardingSphere Proxy based on Helm Charts is deployed in the Kubernetes environment with one click
- One click deployment and automatic operation and maintenance of ShardingSphere Proxy based on Operator in the Kubernetes environment
- Rapid deployment of ShardingSphere Proxy based on AWS CloudFormation
- Rapid deployment of ShardingSphere Proxy based on Terraform in the AWS environment

## Application Scenario

The following application scenarios are available for different deployment schemes provided by SharidngSphere-On-Cloud:

1. If you want to quickly understand, verify or use the features of SharedingSphere Proxy, and there is no Kubernetes environment, you can use AWS CloudFormat or Terraform to deploy on demand.
2. If you want to deploy in the Kubernetes environment, you can experience the Operator function we provide, or you can directly deploy the native SharedingSphere Proxy through helm charts without using the Operator.
