+++
pre = "<b>3.5 </b>"
title = "Terraform 部署多可用区 ShardingSphere 集群"
weight = 5
chapter = true
+++

## 背景

Terraform 是一个开源的基础设施自动化编排工具，使用 “基础设施即代码” 的理念来管理基础设施的变更，AWS， gcp， Azure， aliyun， ucloud 等公有云厂商都支持，还有社区提供的各种各样的 provider，已成为 “基础设施即代码” 领域的一个事实上的标准。

Terraform 有以下优点：

- 支持多云部署

  Terraform 适用于多云方案，将类似的基础结构部署到阿里云、其他云提供商或者本地数据中心。开发人员能够使用相同的工具和相似的配置文件同时管理不同云提供商的资源。

- 自动化管理基础架构

  Terraform 能够创建模块，可重复使用，从而减少因人为因素导致的部署和管理错误。

- 基础架构即代码（Infrastructure as Code）

  可以用代码来管理维护资源。允许保存基础设施状态，从而使您能够跟踪对系统（基础设施即代码）中不同组件所做的更改，并与其他人共享这些配置。

- 降低开发成本

  您通过按需创建开发和部署环境来降低成本。并且，您可以在系统更改之前进行评估。

## 目标

能够使用 Terraform 在 Amazon 上创建 ShardingSphere 高可用集群，创建的集群架构图如下，后续会支持更多的云厂商。

![](../../../img/overview/terraform.png)

创建的 Amazon 资源如下：
1. 每个可用区一个 ZooKeeper 实例。
2. 每个可用区一个 Auto Scaling Group。
3. 每个可用区一个 Launch Template，用于 Auto Scaling Group 启动 ShardingSphere Proxy 实例。
4. 一个内网 Network LoadBalancer，给应用使用。

## 快速开始

### 前提条件

为创建 ShardingSphere Proxy 高可用集群，您需要事先准备如下资源：
1. 一个 ssh keypair，用于远程连接 EC2 实例。
2. 一个 VPC。
3. 每个可用区的 subnet。
4. 一个 SecurityGroup，能够放行 ZooKeeper Server 使用的 2888，3888，2181 端口。
5. 一个内网 HostedZone。
6. 一个通用的 AMI 镜像，Amazon linux2 即可。

根据上述准备好的资源，修改相应的 `main.tf` 中的参数。

### 步骤

1. 进入到 terraform 目录, 根据上述准备好的资源，修改相应的 `main.tf`  中的参数。

```shell
git clone --depth=1 https://github.com/apache/shardingsphere-on-cloud.git
cd shardingsphere-on-cloud/terraform
```

以下提到的命令都需要在 `terraform` 目录中执行。

2. 运行以下命令初始化 terraform。

```shell
terraform init
```

![](../../../img/operation-guide/5-1.png)

3. 运行以下命令查看 terraform 将要执行的计划，您可以观察是否符合预期。

```shell
terraform plan
```

![](../../../img/operation-guide/5-2.png)

4. 如果符合预期，可以执行以下命令创建集群。

```shell
terraform apply
```

![](../../../img/operation-guide/5-3.png)

创建完成后，会有以下输出：

![](../../../img/operation-guide/5-4.png)

您需要记录下 `shardingsphere_domain` 对应的值，应用可以通过连接该域名访问到 Proxy。

5. 如果您想删除创建的集群，可以运行以下命令：

```shell
terraform destroy
```

## 使用手册

### 依赖

|名称|版本|
|---|----|
|aws|4.37.0|

### 模块列表

|名称|源|
|----|--|
|[zk](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#ZK)|./zk|
|[shardingsphere](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#shardingsphere)|./shardingsphere|

### 输出

|名称 |类型|描述|
|---------------------|----|---|
|[shardingsphere_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#output_shardingsphere_domain)|string|最终的对内提供的 SharidngSphere Proxy 域名，其他应用可以通过此域名连接到 Proxy|
|[zk_node_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#output_zk_node_domain)       |list(string)|ZooKeeper 服务对应的域名列表|

### 模块详细信息

#### ZK

**内部资源列表**

|名称|类型|
|----|---|
|[aws_instance.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance)|resource|
|[aws_network_interface.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface)|resource|
|[aws_route53_record.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record)|resource|
|[aws_ami.base](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami)|data source|
|[aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones)|data source|
|[aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone)|data source|

**输入**

|名称|描述|类型|默认值|是否依赖|
|----|---|----|------|-------|
|[cluster_size](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_cluster_size)|与可用区相同数量的集群大小|number|n/a|yes|
|[hosted_zone_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_hosted_zone_name)|私有 zone 名称|string|"shardingsphere.org"|no|
|[instance_type](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_instance_type)|EC2 实例类型|string|n/a|yes|
|[key_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_key_name)|SSH 密钥对|string|n/a|yes|
|[security_groups](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_security_groups)|Security Group 列表, 必须放行 2181，2888，3888 端口|list(string)|[]|no|
|[subnet_ids](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_subnet_ids)|VPC 中按可用区排序的子网列表|list(string)|n/a|yes|
|[tags](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_tags)|ZooKeeper Server 实例 tags 默认是： Name=zk-${count.idx}"|map(any)|	{}|no|
|[vpc_id](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_vpc_id)|VPC id|string|n/a|yes|
|[zk_config](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_zk_config)|ZooKeeper Server 的默认配置|map|{ <br>"client_port": 2181, <br>"zk_heap": 1024 <br>}"|no|
|[zk_version](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_zk_version)|ZooKeeper Server 版本|string	|"3.7.1"|no|

**输出**

|名称|描述|
|----|----|
|[zk_node_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#output_zk_node_domain)|ZooKeeper Server 对应的域名列表|
|[zk_node_private_ip](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#output_zk_node_private_ip)|ZooKeeper Server 示例的内网 IP|

#### ShardingSphere

**内部资源列表**

|名称|类型|
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

**输入**

|名称|描述|类型|默认值|是否依赖|
|----|---|----|------|-------|
|[cluster_size](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_cluster_size)|与可用区相同数量的集群大小|number|n/a|yes|
|[hosted_zone_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_hosted_zone_name)|私有 zone 名称|string|"shardingsphere.org"|no|
|[image_id](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_image_id)|AMI 镜像 ID|string|n/a|yes|
|[instance_type](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_instance_type)|EC2 实例类型|string|n/a|yes|
|[key_name](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_key_name)|SSH 密钥对|string|n/a|yes|
|[lb_listener_port](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_lb_listener_port)|ShardingSphere Proxy 启动端口|string|n/a|yes|
|[security_groups](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_security_groups)|Security Group 列表|list(string)|[]|no|
|[shardingsphere_version](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_shardingsphere_version)|ShardingSphere Proxy 版本|string|n/a|yes|
|[subnet_ids](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_subnet_ids)|VPC 中按可用区排序的子网列表|list(string)|n/a|yes|
|[vpc_id](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_vpc_id)|VPC ID|string|n/a|yes|
|[zk_servers](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#input_zk_servers)|Zookeeper Servers|list(string)|n/a|yes|

**输出**

|名称|描述|
|----|----|
|[shardingsphere_domain](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform#output_shardingsphere_domain)|ShardingSphere Proxy 集群对内提供的域名，其他应用可以通过此域名连接到 Proxy|

## 运维

默认使用我们提供的 Terraform 配置创建的 ZooKeeper 和 ShardingSphere Proxy 服务可以使用 systemd 管理。

### ZooKeeper 

#### 启动

```shell
systemctl start zookeeper
```

#### 停止

```shell
systemctl stop zookeeper
```

#### 重启

```shell
systemctl restart zookeeper
```

### ShardingSphere Proxy

#### 启动

```shell
systemctl start shardingsphere
```

#### 停止

```shell
systemctl stop shardingsphere
```

#### 重启

```shell
systemctl restart shardingsphere
```
