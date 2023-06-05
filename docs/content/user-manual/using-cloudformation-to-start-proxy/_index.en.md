+++
pre = "<b>4.3 </b>"
title = "Start ShardingSphere Proxy with CloudFormation "
weight = 3
chapter = true
+++

AWS CloudFormation is a simple tool leveraging infrastructure as code to configure and start any environment and infrastructure. The AWS CloudFormation Stack template can help you to quickly start Apache ShardingSphere on AWS.

## Requirements

Before getting started, you need to confirm the following checklist:

- [ ] The selected region is ap-north-1 (Beijing). Currently, the AMI and related components that protect Apache ShardingSphere Proxy are only valid in the ap-north-1 region
- [ ] An existing VPC is used to deploy Apache ShardingSphere Proxy
- [ ] The next planned CIDR and corresponding subnet of the VPC
- [ ] Security group configuration that allows applications to access databases (such as port 3307) and control traffic (such as port 22)
- [ ] The key pair that can be used to access the instance resources 
- [ ] Label of resource design related to CloudFormation Stack

## Start ShardingSphere Proxy Cluster

### 1. Create CloudFormation stack with new resources

As shown in the figure below:

![](../../../img/operation-guide/1.PNG)

![](../../../img/operation-guide/2.PNG)

### 2. Upload the template file in this warehouse

Upload local file `cloudformation/apache-shardingsphere-5.2.0.json` to CloudFormation, and then click `Next`.

![](../../../img/operation-guide/3.PNG)

![](../../../img/operation-guide/4.PNG)

### 3. Specify CloudFormation stack details.

Fill in the blank items on this page. The required items are ready in the preconditions.

![](../../../img/operation-guide/5.PNG)

### 4. Configure stack options

Adding labels to the stack is helpful for subsequent cost analysis.

![](../../../img/operation-guide/6.PNG)

### 5. Review and confirm configuration

Review the configuration contents and confirm that all contents meet the expectations before submission.

![](../../../img/operation-guide/7.PNG)

### 6. Check EC2 instances

After a few minutes, the EC2 instance will have started.

![](../../../img/operation-guide/8.PNG)

### 7. Check the status of ShardingSphere Proxy and ZooKeeper

Use `systemctl status shardingsphere` and `./bin/zkServer.sh status` to check the running status of components.

![](../../../img/operation-guide/9.PNG)

![](../../../img/operation-guide/10.PNG)

### 8. Test simple sharding example

Create database `sharding_db` and add two independent database instances `resources`. Then create logical table `t_order` and insert two rows of data. The following inspection results:

![](../../../img/operation-guide/11.PNG)

![](../../../img/operation-guide/12.PNG)

![](../../../img/operation-guide/13.PNG)
