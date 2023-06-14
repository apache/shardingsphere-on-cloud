+++
pre = "<b>2.1 </b>"
title = "ShardingSphere CloudFormation Quick Start"
weight = 1
chapter = true
+++

## Prodecure

### Requirements

To create a ShardingSphere Proxy highly available cluster, you need to prepare the following resources:
1. An ssh keypair used to remotely connect EC2 instances.
2. One VPC.
3. The subnet of each AZ.
4. A SecurityGroup can release the 2888, 3888, 2181 ports used by ZooKeeper Server.
5. An intranet HostedZone.
6. A common AMI image, Amazon linux2.
7. Prepare CloudFormation [configuration file](https://raw.githubusercontent.com/apache/shardingsphere-on-cloud/main/cloudformation/multi-az/cf.json).

### Steps 

1. Enter Amazon's CloudFormation service and create Stacks.

![](../../../img/operation-guide/4-1.PNG)

Click `Choose File` button to upload the prepared CloudFormation configuration.

![](../../../img/operation-guide/4-2.PNG)

Click Next after uploading.

2. Fill the resources you have prepared into the relevant locations below.

![](../../../img/operation-guide/4-3.PNG)

![](../../../img/operation-guide/4-4.PNG)

After filling in the corresponding parameters, click Next.

3. Configure 'stack' related parameters according to your context.

![](../../../img/operation-guide/4-5.PNG)

![](../../../img/operation-guide/4-6.PNG)

Click Next after configuration.

4. Configure 'Review'.

![](../../../img/operation-guide/4-7.PNG)

![](../../../img/operation-guide/4-8.PNG)

![](../../../img/operation-guide/4-9.PNG)

Confirm and click `Submit`.

5. After performing the above operations, you will enter the creation phase.

![](../../../img/operation-guide/4-10.PNG)

![](../../../img/operation-guide/4-11.PNG)

![](../../../img/operation-guide/4-12.PNG)

6. Wait for a bit, and after the creation, enter the 'Outputs' tab as shown in the following figure.

![](../../../img/operation-guide/4-13.PNG)

The value corresponding to 'ssinernaldomain' is the domain name we need.

The internal domain name created by default is [proxy.shardingsphere.org](proxy.shardingsphere.org), the port is 3307, and the username and password are root.

