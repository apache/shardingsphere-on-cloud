# ZooKeeper Cluster Terraform module
A terraform module to create an ZooKeeper Cluster on HuaweiCloud.

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_huaweicloud"></a> [huaweicloud](#requirement\_huaweicloud) | 1.43.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_huaweicloud"></a> [huaweicloud](#provider\_huaweicloud) | 1.43.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [huaweicloud_compute_eip_associate.associated](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/compute_eip_associate) | resource |
| [huaweicloud_compute_instance.zk](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/compute_instance) | resource |
| [huaweicloud_dns_recordset.zk](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/dns_recordset) | resource |
| [huaweicloud_vpc_eip.zkeip](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/resources/vpc_eip) | resource |
| [huaweicloud_images_image.myimage](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/images_image) | data source |
| [huaweicloud_vpc_subnet.zknet](https://registry.terraform.io/providers/huaweicloud/huaweicloud/1.43.0/docs/data-sources/vpc_subnet) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cluster_size"></a> [cluster\_size](#input\_cluster\_size) | The cluster size that same size as available\_zones | `number` | n/a | yes |
| <a name="input_flavor_id"></a> [flavor\_id](#input\_flavor\_id) | The ECS instance type | `string` | n/a | yes |
| <a name="input_image_id"></a> [image\_id](#input\_image\_id) | The image id for the ECS instance | `string` | `""` | no |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | The ssh keypair for remote connection | `string` | n/a | yes |
| <a name="input_security_groups"></a> [security\_groups](#input\_security\_groups) | List of the Security Group, it must be allow access 2181, 2888, 3888 port | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of zk instance resource, the default tag is Name=zk-${count.idx} | `map(any)` | `{}` | no |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | The id of VPC | `string` | n/a | yes |
| <a name="input_zk_config"></a> [zk\_config](#input\_zk\_config) | The default config of zookeeper server | `map` | <pre>{<br>  "client_port": 2181,<br>  "zk_heap": 1024<br>}</pre> | no |
| <a name="input_zk_version"></a> [zk\_version](#input\_zk\_version) | The zookeeper version | `string` | `"3.7.1"` | no |
| <a name="input_zone_id"></a> [zone\_id](#input\_zone\_id) | The id of the private zone | `string` | `""` | no |
| <a name="input_zone_name"></a> [zone\_name](#input\_zone\_name) | The name of the private zone | `string` | `"shardingsphere.org"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_zk_node_domain"></a> [zk\_node\_domain](#output\_zk\_node\_domain) | The private domain names of zookeeper instances for use by ShardingSphere Proxy |
| <a name="output_zk_node_private_ip"></a> [zk\_node\_private\_ip](#output\_zk\_node\_private\_ip) | The private ips of zookeeper instances |
