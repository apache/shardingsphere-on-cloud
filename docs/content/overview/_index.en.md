+++
pre = "<b>1. </b>"
title = "Overview"
weight = 1
chapter = true
+++

## What is ShardingSphere-on-Cloud?

The ShardingSphere-on-Cloud project is a collection of tools and best practices to take Apache ShardingSphere into the cloud. It helps users to transform any database into a distributed database system, and enhance it with sharding, elastic scaling, encryption features & more. 

## Feature Highlights 

### Out-of-box deployment in cloud

* For cloud vendors like AWS, it provides CloudFormation stack templates and Terraform provider scripts.
* For Kubernetes, it provides Helm Charts and Operator.

### Database Reliability Engineering

Leveraing Operator to accomplish ShardingSphere automatically operation, such as automatically horizontal scaling, high-availability deployment and observabilites.

### Chaos Engineering

Injecting any kind of chaos into the ShardingSphere system with the help of Operator, which includes network chaos, Pod chaos, stresses like CPU and memory, to help improve the fault tolerances level in similar scenarios.

### Ecosystem Extensions

ShardingSphere supports Java SPI extensions, at the meanwhile, support non-Java compiled WebAssembly extensions.

## Application Scenarios

1. If you want to quickly understand, verify or use the features of ShardingSphere Proxy, you can use AWS CloudFormation and Terraform scripts, or use Helm Charts and Operators in Kubernetes environments.
2. If you want to implement data sharding and manage databases on Kubernetes, you can use ShardingSphere Operator.
3. If you want to solve DBRE related issues, you can use ShardingSphere Operator.
4. If you want to extend ShardingSphere with your own plugins, you can write WebAssembly extensions.
