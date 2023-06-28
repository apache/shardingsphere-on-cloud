+++
pre = "<b>3.1 </b>"
title = "Deployment on Cloud"
weight = 1
chapter = true
+++

## Overview

Cloud computing has evolved over the years from IaaS to PaaS, and then to SaaS. It not only changed infrastructure compositions but also upgraded software development concepts. 
With Kubernetes leading the cloud-native wave, an increasing number of applications, including ShardingSphere, are being deployed using cloud-native technology stacks. To deploy ShardingSphere in a cloud environment, we recommend adopting Infrastructure as Code (IaC).

### AWS One-Click Deployment

To deploy ShardingSphere on AWS, you should first familiarize yourself with various AWS resources and services such as VPC, subnet, security group, elastic load balancer, domain name, EC2, RDS, and CloudWatch. You can adopt IaC, like AWS's CloudFormation, to quickly describe and deploy a complete set of ShardingSphere structures. 

CloudFormation uses json or yaml templates to describe and combine various resources required for abstract deployment. It's interpreted and executed by related services. You only need to write relevant descriptions using version control tools, such as Git, to manage and maintain the deployed code. 

Currently, Apache ShardingSphere's CloudFormation is hosted in the ShardingSphere on Cloud repo. You can get the corresponding AMI information on the AWS Marketplace by clicking [HERE](https://us-east-1.console.aws.amazon.com/marketplace/home?region=ap-southeast-1#/subscriptions/ef146e06-20ca-4da4-8954-78a7c51b3c5a).

See [Quick Start](https://shardingsphere.apache.org/document/current/en/quick-start/) to learn how to start a ShardingSphere Proxy cluster on AWS with CloudFormation's minimal configuration. If you want to learn more about CloudFormation parameters or are familiar with Terraform, please refer to [User Manual](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-jdbc/).

### Kubernetes One-Click Deployment

Deploying ShardingSphere on Kubernetes has never been easier, thanks to our one-click deployment feature that utilizes the Helm package manager. This tool enables users to describe the deployment structure using a set of templates and Charts comprised of variable declarations. 

Resource objects involved in the deployment include Kubernetes workloads such as Deployment, Service, and ConfigMap. You can produce Charts packages for each version release and submit them to public product repos like ArtifactHub. 

Currently, we offer this feature with the relevant source code hosted on the ShardingSphere on Cloud repo.

See [Quick Start](https://shardingsphere.apache.org/document/current/en/quick-start/) to learn how to start a ShardingSphere Proxy cluster on Kubernetes with Helm Charts's minimal configuration. If you want to know more about Charts parameters or are familiar with Operator, please refer to [User Manual](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-jdbc/).

## Applicable Scenarios

You can use the one-click deployment mode for testing purposes. If you plan to use ShardingSphere Proxy in a production environment, please refer to [User Manual](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-jdbc/). It is crucial to learn relevant parameters before configuring and deploying. 



