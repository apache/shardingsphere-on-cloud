+++
pre = "<b>3.5 </b>"
title = "Terraform Deploys ShardingSphere Cluster in Multiple AZs On AWS"
weight = 5
chapter = true
+++

## Background

Terraform is an open source automated infrastructure orchestration tool, which uses the concept of "infrastructure as code" to manage infrastructure changes. Public cloud providers such as AWS, Google Cloud Platform (GCP), Azure, Alibaba Cloud, and Ucloud support it, as well as various community-supported providers have become a standard in the field of "infrastructure is code".

Terraform has the following advantages:

- Support multi-cloud deployment

  Terraform is applicable to multi-cloud solutions. It deploys similar infrastructure to Alibaba Cloud, other cloud providers or local data centers. Developers can use the same tools and similar configuration files to manage the resources of different cloud providers at the same time.

- Automated manage infrastructure

  Terraform can create modules that can be reused to reduce deployment and management errors caused by human factors.

- Infrastructure as Code

  You can use code to manage maintenance resources. Allows you to save infrastructure state, enabling you to track changes to different components in your system (infrastructure as code) and share these configurations with others.

- Reduce development costs

  You reduce costs by creating development and deployment environments on demand. Also, you can evaluate before the system changes.

## Goal

You can use Terraform to create a ShardingSphere high availability cluster on Amazon AWS. The cluster architecture is shown below. More cloud providers will be supported in the near future.

![](../../../img/overview/terraform.png)

The Amazon resources created are the following:
1. One ZooKeeper instance per AZ.
2. One Auto Scaling Group.
3. A Launch Template, which is used by the Auto Scaling Group to start the SharedingSphere Proxy instance.
4. An intranet Network LoadBalancer for applications.

## Quick Start

### Requirements

To create a ShardingSphere Proxy highly available cluster, you need to prepare the following resources in advance:
1. An ssh keypair used to remotely connect EC2 instances.
2. One VPC.
3. The subnet of each AZ.
4. A SecurityGroup can release the 2888, 3888, and 2181 ports used by ZooKeeper Server.
5. An intranet HostedZone.
6. A common AMI image, such as Amazon linux2.

Modify the parameters in `main.tf` according to the above prepared resources.

### Procedure

1. Enter the terraform directory, modify the parameters in `main.tf`  according to the above prepared resources.

```shell
git clone --depth=1 https://github.com/apache/shardingsphere-on-cloud.git
cd shardingsphere-on-cloud/terraform/amazon
```

The commands mentioned below need to be executed in the 'amazon' directory.

2. Run the following command to initialize the terraform.

```shell
terraform init
```

![](../../../img/operation-guide/5-1.png)

3. Run the following command to check the plan that terraform will execute, and check whether it meets your expectations.

```shell
terraform plan
```

![](../../../img/operation-guide/5-2.png)

4. If the plan is as expected, you can execute the following command to create a cluster.

```shell
terraform apply
```

![](../../../img/operation-guide/5-3.png)

After creation, the following outputs will be available:

![](../../../img/operation-guide/5-4.png)

You need to record the value corresponding to `shardingsphere_domain`. Applications can access the proxy by connecting to the domain name.

5. If you want to delete the created cluster, you can run the following command:

```shell
terraform destroy
```

## User Manual

### Dependency

|Name|Version|
|---|----|
|aws|4.37.0|

### Module List

|Name|Source|
|----|--|
|[zk](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#ZK)|./zk|
|[shardingsphere](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#shardingsphere)|./shardingsphere|

### Outputs

| Name                                                                                                                               |Type|Description|
|------------------------------------------------------------------------------------------------------------------------------------|----|---|
| [shardingsphere_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#output_shardingsphere_domain) |string|The final SharidngSphere proxy domain name provided internally, through which other applications can connect to the proxy|
| [zk_node_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#output_zk_node_domain)               |list(string)|List of domain names corresponding to ZooKeeper service|

### Module Details

#### ZK

**Internal resource list**

|Name|Type|
|----|---|
|[aws_instance.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance)|resource|
|[aws_network_interface.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface)|resource|
|[aws_route53_record.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record)|resource|
|[aws_ami.base](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami)|data source|
|[aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones)|data source|
|[aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone)|data source|

**Inputs**

| Name                                                                                                                    |Description|Type|Default Value|Dependent on or not|
|-------------------------------------------------------------------------------------------------------------------------|---|----|------|-------|
| [cluster_size](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_cluster_size)         |Cluster size of the same number as the availability zone|number|n/a|yes|
| [hosted_zone_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_hosted_zone_name) |Private zone name|string|"shardingsphere.org"|no|
| [instance_type](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_instance_type)       |EC2 instance type|string|n/a|yes|
| [key_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_key_name)                 |SSH key pair|string|n/a|yes|
| [security_groups](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_security_groups)   |Security Group list. 2181, 2888, 3888 ports must be released|list(string)|[]|no|
| [subnet_ids](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_subnet_ids)             |Subnet list sorted by AZ in VPC|list(string)|n/a|yes|
| [tags](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_tags)                         |ZooKeeper Server instance tags The default is: Name=zk-${count.idx}"|map(any)|	{}|no|
| [vpc_id](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_vpc_id)                     |VPC id|string|n/a|yes|
| [zk_config](https://github.com/apache/shardingsphere-on-cloud/tree/main/terrafor/amazonm#input_zk_config)               |Default configuration of ZooKeeper Server|map|{ <br>"client_port": 2181, <br>"zk_heap": 1024 <br>}"|no|
| [zk_version](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_zk_version)             |ZooKeeper Server version|string	|"3.7.1"|no|

**Outputs**

| Name                                                                                                                         |Description|
|------------------------------------------------------------------------------------------------------------------------------|----|
| [zk_node_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#output_zk_node_domain)         |List of domain names corresponding to ZooKeeper Server|
| [zk_node_private_ip](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#output_zk_node_private_ip) |The intranet IP address of the ZooKeeper Server example|

#### ShardingSphere

**Internal resource list**

|Name|Type|
|----|----|
|[aws_autoscaling_attachment.asg_attachment_lb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_attachment)|resource|
|[aws_autoscaling_group.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_group)|resource|
|[aws_launch_template.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/launch_template)|resource|
|[aws_lb.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb)|resource|
|[aws_lb_listener.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_listener)|resource|
|[aws_lb_target_group.ss_tg](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_target_group)|resource|
|[aws_network_interface.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface)|resource|
|[aws_route53_record.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record)|resource|
|[aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones)|data source|
|[aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone)|data source|
|[aws_vpc.vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/vpc)|data source|

**Inputs**

| Name                                                                                                                                |Description|Type|Default Value|Dependent on or not|
|-------------------------------------------------------------------------------------------------------------------------------------|---|----|------|-------|
| [cluster_size](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_cluster_size)                     |Cluster size of the same number as the availability zone|number|n/a|yes|
| [hosted_zone_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_hosted_zone_name)             |Private zone name|string|"shardingsphere.org"|no|
| [image_id](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_image_id)                             |AMI iamge ID|string|n/a|yes|
| [instance_type](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_instance_type)                   |EC2 instance type|string|n/a|yes|
| [key_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_key_name)                             |SSH key pair|string|n/a|yes|
| [lb_listener_port](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_lb_listener_port)             |ShardingSphere Proxy startup port|string|n/a|yes|
| [security_groups](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_security_groups)               |Security Group list|list(string)|[]|no|
| [shardingsphere_version](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_shardingsphere_version) |ShardingSphere Proxy version|string|n/a|yes|
| [subnet_ids](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_subnet_ids)                         |Subnet list sorted by AZ in VPC|list(string)|n/a|yes|
| [vpc_id](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_vpc_id)                                 |VPC ID|string|n/a|yes|
| [zk_servers](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#input_zk_servers)                         |Zookeeper Servers|list(string)|n/a|yes|

**Outputs**

| Name                                                                                                                               |Description|
|------------------------------------------------------------------------------------------------------------------------------------|----|
| [shardingsphere_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/amazon#output_shardingsphere_domain) |The domain name provided by the shardingSphere Proxy cluster. Other applications can connect to the proxy through this domain name.|

## DevOps

By default, ZooKeeper and ShardingSphere Proxy services created using our Terraform configuration can be managed using systemd.

### ZooKeeper 

#### Start

```shell
systemctl start zookeeper
```

#### Stop

```shell
systemctl stop zookeeper
```

#### Restart

```shell
systemctl restart zookeeper
```

### ShardingSphere Proxy

#### Start

```shell
systemctl start shardingsphere-proxy
```

#### Stop

```shell
systemctl stop shardingsphere-proxy
```

#### Restart

```shell
systemctl restart shardingsphere-proxy
```
