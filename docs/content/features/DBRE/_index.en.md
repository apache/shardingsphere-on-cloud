+++
pre = "<b>3.2 </b>"
title = "Database Reliability Engineering"
weight = 2
chapter = true
+++

## Overview

Database Reliability Engineering (DBRE) aims to improve the stability of database-related services using different technical methods, similar to Site Reliability Engineering (SRE). In cases where ShardingSphere is deployed on Kubernetes, DBRE can be further implemented with the help of Operator.

## High Availability Deployment

Since ShardingSphere Proxy is stateless, it serves as a computing node that processes SQL sent by the client to complete relevant data calculations. Operator abstracts and describes ShardingSphere Proxy through ComputeNode. 

Currently, the Deployment mode is suitable for ShardingSphere Proxy, due to its statelessness. Deployment is a basic method provided by Kubernetes with no difference between the Pods it manages. ShardingSphere Proxy can be deployed through Deployment, which offers essential capabilities such as health checks, readiness checks, rolling upgrades, and version rollbacks.

ComputeNode encompasses various attributes that are essential for deploying ShardingSphere Proxy including the number of copies, mirror repo information, version information, database driver information, health check, and readiness check probes. It also includes port mapping rules, service startup requirements server.yaml, logback.xml, and information like Agent-related configuration. During the operator tuning process, these pieces of information will be rendered through Kubernetes Deployment, Service, and ConfigMap respectively and binding and mounting actions will automatically occur. 

Deployment’s capabilities result in easy multi-replicas deployment and advanced scheduling features such as affinity and taint tolerance, which in turn provide basic high availability capabilities for ShardingSphere Proxy.

StorageNode includes configurations related to deploying RDS database instances on the public cloud. It specifies the corresponding public cloud resource provider through StorageProvider and enables the ability to create, automatically register and unregister, and delete database instances on the cloud, with automatic elastic expansion now supported. 

## Automatic Elastic Expansion

The Kubernetes community provides a Horizontal Pod Autoscaler (HPA) that automatically expands based on CPU and memory, and can also be paired with Prometheus Adapter for expansion based on custom indicators. For AWS EC2 virtual machine deployment scenarios, the community also offers the option to expand the capacity of AutoScalingGroup and through the detection mechanism of TargetGroup, only Ready instances can receive business traffic.

## Observability

ShardingSphere Proxy, with the assistance of ShardingSphere Agent, can effectively compile and display necessary operational information. You can get more details on ShardingSphere Agent by clicking [here](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/observability/). Additionally, ShardingSphere on Cloud features Grafana templates that provide valuable insights into basic resource monitoring, JVM monitoring, and ShardingSphere runtime indicators. By plotting these indicators at different levels on a single Dashboard, users can easily pinpoint potential problems.

## Chaos Engineering

Chaos engineering allows us to verify the robustness of our system and uncover previously unknown issues. ShardingSphere Operator supports CRD Chaos that injects different types of faults such as Pod exceptions, CPU pressure, memory pressure, and network exceptions, directly into ShardingSphere Proxy. See Chaos Engineering for more details.

