+++
pre = "<b>3.5 </b>"
title = "Terraform 在华为云上部署多可用区 ShardingSphere 集群"
weight = 5
chapter = true
+++

## 背景
有关 Terraform 的介绍可以参考 [链接](../using-terraform-to-deploy-multi-az-on-aws/_index.cn.md)

## 目标

能够使用 Terraform 在华为云上创建 ShardingSphere 高可用集群，创建的集群架构图如下，后续会支持更多的云厂商。

![](../../../img/overview/terraform.png)

创建的华为云资源如下：
1. 每个可用区一个 ZooKeeper 实例。
2. 一个弹性伸缩组和弹性伸缩配置。
3. 用于 ShardingSphere Proxy 集群的内网负载均衡。
4. 一个内网域名，给应用使用。

## 快速开始

### 前提条件

为创建 ShardingSphere Proxy 高可用集群，您需要事先准备如下资源：
1. 一个 ssh 密钥对，用于远程连接 ECS 实例。
2. 一个 VPC。
3. 一个 Subnet。
4. 一个 SecurityGroup，能够放行 ZooKeeper Server 使用的 2888，3888，2181 端口。
5. 一个内网 Zone。
6. 用于访问华为云 API 的 AK/SK。

### 步骤

1. 进入到 terraform 目录, 根据上述准备好的资源，创建 `terraform.tfvars` 文件。

```shell
git clone --depth=1 https://github.com/apache/shardingsphere-on-cloud.git
cd shardingsphere-on-cloud/terraform/huawei
```

`terraform.tfvars` 示例内容如下：

```hcl
shardingsphere_proxy_version = "5.3.1"
image_id                     = ""
key_name                     = "test-tf"
flavor_id                    = "c7.large.2"
vpc_id                       = "4b9db05b-4d57-464d-a9fe-83da3de0a74c"
vip_subnet_id                = ""
subnet_ids                   = ["6d6c57ed-5284-4a7b-b0e3-0b24aa6c9552"]
security_groups              = ["f5ad3525-dc9e-482e-afde-868ee330e7a5"]
lb_listener_port             = 3307
zk_flavor_id                 = "s6.medium.2"
```
2. 运行以下命令设置 AK/SK 及 Region
```shell
export HW_ACCESS_KEY="AK"
export HW_SECRET_KEY="SK"
export HW_REGION_NAME="REGION"
```

3. 在 `huawei` 目录下， 运行以下命令创建 ShardingSphere Proxy 集群。
```shell
terraform init
terraform plan  -var-file=terraform.tfvars
terraform apply  -var-file=terraform.tfvars
```

## 使用手册
### 依赖

| Name | Version |
|------|---------|
|[huaweicloud](#provider\_huaweicloud) | 1.43.0 |

### 模块列表

| 名称                                                                                             | 源 |
|------------------------------------------------------------------------------------------------|--------|
| [zk](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/huawei/modules/zk)  | ./modules/zk |

### 内部资源列表

| 名称                                                                                                                                                  | 类型 |
|-----------------------------------------------------------------------------------------------------------------------------------------------------|------|
| [huaweicloud_as_configuration.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/as_configuration)           | resource |
| [huaweicloud_as_group.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/as_group)                           | resource |
| [huaweicloud_dns_recordset.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/dns_recordset)                 | resource |
| [huaweicloud_dns_zone.private_zone](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/dns_zone)                 | resource |
| [huaweicloud_elb_listener.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_listener)                   | resource |
| [huaweicloud_elb_loadbalancer.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_loadbalancer)           | resource |
| [huaweicloud_elb_monitor.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_monitor)                     | resource |
| [huaweicloud_elb_pool.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_pool)                           | resource |
| [huaweicloud_availability_zones.zones](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/availability_zones) | data source |
| [huaweicloud_images_image.myimage](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/images_image)           | data source |
| [huaweicloud_vpc_subnet.vipnet](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/vpc_subnet)                | data source |

### 输入

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| [flavor\_id](#input\_flavor\_id) | The flavor id of the ECS | `string` | n/a | yes |
| [image\_id](#input\_image\_id) | The image id | `string` | `""` | no |
| [key\_name](#input\_key\_name) | the ssh keypair for remote connection | `string` | n/a | yes |
| [lb\_listener\_port](#input\_lb\_listener\_port) | The lb listener port | `string` | n/a | yes |
| [security\_groups](#input\_security\_groups) | List of The Security group IDs | `list(string)` | `[]` | no |
| [shardingsphere\_proxy\_as\_desired\_number](#input\_shardingsphere\_proxy\_as\_desired\_number) | The initial expected number of ShardSphere Proxy Auto Scaling. The default value is 3 | `number` | `3` | no |
| [shardingsphere\_proxy\_as\_healthcheck\_grace\_period](#input\_shardingsphere\_proxy\_as\_healthcheck\_grace\_period) | The health check grace period for instances, in seconds | `number` | `120` | no |
| [shardingsphere\_proxy\_as\_max\_number](#input\_shardingsphere\_proxy\_as\_max\_number) | The maximum size of ShardingSphere Proxy Auto Scaling. The default value is 6 | `number` | `6` | no |
| [shardingsphere\_proxy\_doamin\_prefix\_name](#input\_shardingsphere\_proxy\_doamin\_prefix\_name) | The prefix name of the shardinsphere domain, the final generated name will be [prefix\_name].[zone\_name], the default value is proxy. | `string` | `"proxy"` | no |
| [shardingsphere\_proxy\_version](#input\_shardingsphere\_proxy\_version) | The shardingsphere proxy version | `string` | n/a | yes |
| [subnet\_ids](#input\_subnet\_ids) | List of subnets sorted by availability zone in your VPC | `list(string)` | n/a | yes |
| [vip\_subnet\_id](#input\_vip\_subnet\_id) | The IPv4 subnet ID of the subnet where the load balancer works | `string` | `""` | no |
| [vpc\_id](#input\_vpc\_id) | The id of your VPC | `string` | n/a | yes |
| [zk\_cluster\_size](#input\_zk\_cluster\_size) | The Zookeeper cluster size | `number` | `3` | no |
| [zk\_flavor\_id](#input\_zk\_flavor\_id) | The ECS instance type | `string` | n/a | yes |
| [zk\_servers](#input\_zk\_servers) | The Zookeeper servers | `list(string)` | `[]` | no |
| [zone\_id](#input\_zone\_id) | The id of the private zone | `string` | `""` | no |
| [zone\_name](#input\_zone\_name) | The name of the private zone | `string` | `"shardingsphere.org"` | no |

### 输出

| 名称 | 描述 |
|--|-------------|
|[shardingsphere\_domain](#output\_shardingsphere\_domain) | The domain of the ShardingSphere Proxy Cluster for use by other services |
|[zk\_node\_domain](#output\_zk\_node\_domain) | The domain of zookeeper instances |

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
systemctl start shardingsphere-proxy
```

#### 停止

```shell
systemctl stop shardingsphere-proxy
```

#### 重启

```shell
systemctl restart shardingsphere-proxy
```
