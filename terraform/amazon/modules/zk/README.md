# Terraform Module for ZooKeeper Cluster

## Requirements

## Requirements
| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | 4.37.0 |


## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 4.37.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_instance.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance) | resource |
| [aws_network_interface.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface) | resource |
| [aws_route53_record.zk](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record) | resource |
| [aws_ami.base](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ami) | data source |
| [aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones) | data source |
| [aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cluster_size"></a> [cluster\_size](#input\_cluster\_size) | The cluster size that same size as available\_zones | `number` | n/a | yes |
| <a name="input_hosted_zone_name"></a> [hosted\_zone\_name](#input\_hosted\_zone\_name) | The name of the hosted private zone | `string` | `"shardingsphere.org"` | no |
| <a name="input_instance_type"></a> [instance\_type](#input\_instance\_type) | The EC2 instance type | `string` | n/a | yes |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | The ssh keypair for remote connection | `string` | n/a | yes |
| <a name="input_security_groups"></a> [security\_groups](#input\_security\_groups) | List of the Security Group, it must be allow access 2181, 2888, 3888 port | `list(string)` | `[]` | no |
| <a name="input_subnet_ids"></a> [subnet\_ids](#input\_subnet\_ids) | List of subnets sorted by availability zone in your VPC | `list(string)` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of zk instance resource, the default tag is Name=zk-${count.idx} | `map(any)` | `{}` | no |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | The id of VPC | `string` | n/a | yes |
| <a name="input_zk_config"></a> [zk\_config](#input\_zk\_config) | The default config of zookeeper server | `map` | <pre>{<br>  "client_port": 2181,<br>  "zk_heap": 1024<br>}</pre> | no |
| <a name="input_zk_version"></a> [zk\_version](#input\_zk\_version) | The zookeeper version | `string` | `"3.7.1"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_zk_node_domain"></a> [zk\_node\_domain](#output\_zk\_node\_domain) | The private domain names of zookeeper instances for use by ShardingSphere Proxy |
| <a name="output_zk_node_private_ip"></a> [zk\_node\_private\_ip](#output\_zk\_node\_private\_ip) | The private ips of zookeeper instances |
