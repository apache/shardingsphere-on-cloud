# Apache ShardingSphere Terraform Module
A terraform module to create an Apache ShardingSphere Proxy Cluster on HuaweiCloud.

Root module calls these modules which can also be used separately to create independent resources:
* [zk](./modules/zk) - Creates Zookeeper Cluster

## Usage
We assume that you have moved all the content in the current directory to the `shardingsphere` directory.
So `source = "./shardingsphere"` in your module.

```hcl
module "shardingsphere_proxy" {
  source                       = "./shardingsphere"
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
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_huaweicloud"></a> [huaweicloud](#requirement\_huaweicloud) | 1.43.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_huaweicloud"></a> [huaweicloud](#provider\_huaweicloud) | 1.43.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_zk"></a> [zk](#module\_zk) | ./modules/zk | n/a |

## Resources

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

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_flavor_id"></a> [flavor\_id](#input\_flavor\_id) | The flavor id of the ECS | `string` | n/a | yes |
| <a name="input_image_id"></a> [image\_id](#input\_image\_id) | The image id | `string` | `""` | no |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | the ssh keypair for remote connection | `string` | n/a | yes |
| <a name="input_lb_listener_port"></a> [lb\_listener\_port](#input\_lb\_listener\_port) | The lb listener port | `string` | n/a | yes |
| <a name="input_security_groups"></a> [security\_groups](#input\_security\_groups) | List of The Security group IDs | `list(string)` | `[]` | no |
| <a name="input_shardingsphere_proxy_as_desired_number"></a> [shardingsphere\_proxy\_as\_desired\_number](#input\_shardingsphere\_proxy\_as\_desired\_number) | The initial expected number of ShardSphere Proxy Auto Scaling. The default value is 3 | `number` | `3` | no |
| <a name="input_shardingsphere_proxy_as_healthcheck_grace_period"></a> [shardingsphere\_proxy\_as\_healthcheck\_grace\_period](#input\_shardingsphere\_proxy\_as\_healthcheck\_grace\_period) | The health check grace period for instances, in seconds | `number` | `120` | no |
| <a name="input_shardingsphere_proxy_as_max_number"></a> [shardingsphere\_proxy\_as\_max\_number](#input\_shardingsphere\_proxy\_as\_max\_number) | The maximum size of ShardingSphere Proxy Auto Scaling. The default value is 6 | `number` | `6` | no |
| <a name="input_shardingsphere_proxy_doamin_prefix_name"></a> [shardingsphere\_proxy\_doamin\_prefix\_name](#input\_shardingsphere\_proxy\_doamin\_prefix\_name) | The prefix name of the shardinsphere domain, the final generated name will be [prefix\_name].[zone\_name], the default value is proxy. | `string` | `"proxy"` | no |
| <a name="input_shardingsphere_proxy_version"></a> [shardingsphere\_proxy\_version](#input\_shardingsphere\_proxy\_version) | The shardingsphere proxy version | `string` | n/a | yes |
| <a name="input_subnet_ids"></a> [subnet\_ids](#input\_subnet\_ids) | List of subnets sorted by availability zone in your VPC | `list(string)` | n/a | yes |
| <a name="input_vip_subnet_id"></a> [vip\_subnet\_id](#input\_vip\_subnet\_id) | The IPv4 subnet ID of the subnet where the load balancer works | `string` | `""` | no |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | The id of your VPC | `string` | n/a | yes |
| <a name="input_zk_cluster_size"></a> [zk\_cluster\_size](#input\_zk\_cluster\_size) | The Zookeeper cluster size | `number` | `3` | no |
| <a name="input_zk_flavor_id"></a> [zk\_flavor\_id](#input\_zk\_flavor\_id) | The ECS instance type | `string` | n/a | yes |
| <a name="input_zk_servers"></a> [zk\_servers](#input\_zk\_servers) | The Zookeeper servers | `list(string)` | `[]` | no |
| <a name="input_zone_id"></a> [zone\_id](#input\_zone\_id) | The id of the private zone | `string` | `""` | no |
| <a name="input_zone_name"></a> [zone\_name](#input\_zone\_name) | The name of the private zone | `string` | `"shardingsphere.org"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_shardingsphere_domain"></a> [shardingsphere\_domain](#output\_shardingsphere\_domain) | The domain name of the ShardingSphere Proxy Cluster for use by other services |
| <a name="output_zk_node_domain"></a> [zk\_node\_domain](#output\_zk\_node\_domain) | The domain of zookeeper instances |

