+++
pre = "<b>1. </b>"
title = "Overview"
weight = 1
chapter = true
+++

## What is ShardingSphere-on-Cloud?

The ShardingSphere-on-Cloud project is a collection of tools and best practices to take Apache ShardingSphere into the cloud. It includes automated deployment scripts to virtual machines in AWS, Google Cloud Platform (GCP), Alibaba Cloud and other cloud environments such as CloudFormation Stack templates, and Terraform one click deployment scripts. 

It also includes Helm Charts, Operators, automatic horizontal scaling and other tools in the Kubernetes cloud native environment, as well as high availability, observability, security compliance and other aspects.

## Core Concept

Currently, the terms involved in this project are from common cloud service providers and other open-source projects, while maintaining concepts and definitions consistency.

- CloudFormation: a tool provided by AWS to facilitate the quick creation of cloud resources.
- CloudFormation Stack: a collection of AWS resources that can be managed as a unit.
- Terraform: an open source infrastructure management tool. Using the concept of "infrastructure as code", you can efficiently build, change and version infrastructure.
- Kubernetes: an open source container orchestration management platform, which can automatically deploy, manage and extend container applications.
- Operator: a Kubernetes extension software, which uses [custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) to manage applications and their components. Operator follows the concept of Kubernetes, especially in [controller](https://kubernetes.io/docs/concepts/architecture/controller/).
- Helm Charts: Helm is a package management tool for the Kubernetes application, while Charts is a collection of files that describe a group of related Kubernetes resources.

## Infrastructure

- ShardingSphere-Operator Diagram

![Operator Diagram](../../img/overview/operator.png)

- ShardingSphere-Terraform Diagram

![Terraform Diagram](../../img/overview/terraform.png)

## Features List

- ShardingSphere Proxy based on Helm Charts can be deployed in a Kubernetes environment with one click.
- One click deployment and automatic DevOps of ShardingSphere Proxy based on Operator in a Kubernetes environment.
- Rapid deployment of ShardingSphere Proxy based on AWS CloudFormation.
- Rapid deployment of ShardingSphere Proxy based on Terraform in an AWS environment.
- Point-in-Time-Recovery demo For ShardingSphere.
- Grafana template For ShardingSphere.
- WebAssembly extension demo For ShardingSphere.

## Application Scenarios

The following application scenarios are available for different deployment methods provided by SharidngSphere-On-Cloud:

1. If you want to quickly understand, verify or use the features of SharedingSphere Proxy, without a Kubernetes environment, you can use AWS CloudFormation or Terraform to deploy on demand.
2. If you want to deploy in a Kubernetes environment, you can leverage our Operator function, or you can directly deploy the native SharedingSphere Proxy through Helm Charts without using Operator.
