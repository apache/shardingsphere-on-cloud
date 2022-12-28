+++
pre = "<b>3.5 </b>"
title = "Terraform Deploys ShardingSphere Cluster in Multiple AZs On Huawei Cloud"
weight = 5
chapter = true
+++

## Background
For an introduction to Terraform, please refer to [link](../using-terraform-to-deploy-multi-az-on-aws/_index.en.md).

## Goal

You can use Terraform to create a ShardingSphere high availability cluster on Huawei Cloud. The cluster architecture is shown below. More cloud providers will be supported in the near future.

![](../../../img/overview/terraform.png)

The HuaweiCloud resources created are the following:
1. One ZooKeeper instance per AZ.
2. One Auto Scaling Group and One Auto Scaling Configuration.
3. An intranet Network LoadBalancer for ShardingSphere Proxy Cluster.
4. An intranet domain for applications.

## Quick Start

### Requirements

To create a ShardingSphere Proxy highly available cluster, you need to prepare the following resources in advance:
1. An ssh keypair used to remotely connect ECS instances.
2. One VPC.
3. One subnet.
4. A SecurityGroup can release the 2888, 3888, and 2181 ports used by ZooKeeper Server.
5. An intranet Zone.
6. AK/SK of the Huawei Cloud account.

### Procedure

1. Enter the terraform directory, create the `terraform.tfvars` file  according to the above prepared resources.

```shell
git clone --depth=1 https://github.com/apache/shardingsphere-on-cloud.git
cd shardingsphere-on-cloud/terraform/huawei
```
The `terraform.tfvars` sample content is as follows:
```hcl
shardingsphere_proxy_version = "5.3.0"
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

2. RUn the following command to set AK/SK and Region.
```shell
export HW_ACCESS_KEY="AK"
export HW_SECRET_KEY="SK"
export HW_REGION_NAME="REGION"
```

3. Under the `huawei` directory, run the following command to deploy ShardingSphere-Proxy Cluster.

```shell
terraform init
terraform plan  -var-file=terraform.tfvars
terraform apply  -var-file=terraform.tfvars
```

## User Manual
### Requirements

| Name | Version |
|------|---------|
|[huaweicloud](#requirement\_huaweicloud) | 1.43.0 |

### Modules

| Name | Source | Version |
|------|--------|---------|
|[zk](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform/huawei/modules/zk) | ./modules/zk | n/a |

### Resources

| Name | Type |
|------|------|
| [huaweicloud_as_configuration.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/as_configuration) | resource |
| [huaweicloud_as_group.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/as_group) | resource |
| [huaweicloud_dns_recordset.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/dns_recordset) | resource |
| [huaweicloud_dns_zone.private_zone](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/dns_zone) | resource |
| [huaweicloud_elb_listener.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_listener) | resource |
| [huaweicloud_elb_loadbalancer.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_loadbalancer) | resource |
| [huaweicloud_elb_monitor.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_monitor) | resource |
| [huaweicloud_elb_pool.ss](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/elb_pool) | resource |
| [huaweicloud_availability_zones.zones](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/availability_zones) | data source |
| [huaweicloud_images_image.myimage](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/images_image) | data source |
| [huaweicloud_vpc_subnet.vipnet](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/vpc_subnet) | data source |

### Inputs

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

### Outputs

| Name | Description |
|------|-------------|
|[shardingsphere\_domain](#output\_shardingsphere\_domain) | The domain name of the ShardingSphere Proxy Cluster for use by other services |
|[zk\_node\_domain](#output\_zk\_node\_domain) | The domain of zookeeper instances |


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
