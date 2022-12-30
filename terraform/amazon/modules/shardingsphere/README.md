# Terraform Module for Apache ShardingSphere-Proxy Cluster

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
| [aws_autoscaling_attachment.asg_attachment_lb](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_attachment) | resource |
| [aws_autoscaling_group.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_group) | resource |
| [aws_iam_instance_profile.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_instance_profile) | resource |
| [aws_iam_role.sts](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_launch_template.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/launch_template) | resource |
| [aws_lb.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb) | resource |
| [aws_lb_listener.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_listener) | resource |
| [aws_lb_target_group.ss_tg](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_target_group) | resource |
| [aws_network_interface.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface) | resource |
| [aws_route53_record.ss](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record) | resource |
| [aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones) | data source |
| [aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone) | data source |
| [aws_vpc.vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/vpc) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cluster_size"></a> [cluster\_size](#input\_cluster\_size) | The cluster size that same size as available\_zones | `number` | n/a | yes |
| <a name="input_hosted_zone_name"></a> [hosted\_zone\_name](#input\_hosted\_zone\_name) | The name of the hosted private zone | `string` | `"shardingsphere.org"` | no |
| <a name="input_image_id"></a> [image\_id](#input\_image\_id) | The AMI id | `string` | n/a | yes |
| <a name="input_instance_type"></a> [instance\_type](#input\_instance\_type) | The EC2 instance type | `string` | n/a | yes |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | the ssh keypair for remote connection | `string` | n/a | yes |
| <a name="input_lb_listener_port"></a> [lb\_listener\_port](#input\_lb\_listener\_port) | lb listener port | `string` | n/a | yes |
| <a name="input_security_groups"></a> [security\_groups](#input\_security\_groups) | List of The Security group IDs | `list(string)` | `[]` | no |
| <a name="input_shardingsphere_proxy_asg_desired_capacity"></a> [shardingsphere\_proxy\_asg\_desired\_capacity](#input\_shardingsphere\_proxy\_asg\_desired\_capacity) | The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain. see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacitytype, The default value is 3 | `string` | `"3"` | no |
| <a name="input_shardingsphere_proxy_asg_healthcheck_grace_period"></a> [shardingsphere\_proxy\_asg\_healthcheck\_grace\_period](#input\_shardingsphere\_proxy\_asg\_healthcheck\_grace\_period) | The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check. see https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html | `number` | `120` | no |
| <a name="input_shardingsphere_proxy_asg_max_size"></a> [shardingsphere\_proxy\_asg\_max\_size](#input\_shardingsphere\_proxy\_asg\_max\_size) | The maximum size of ShardingSphere Proxy Auto Scaling Group. The default values is 6 | `string` | `"6"` | no |
| <a name="input_shardingsphere_proxy_version"></a> [shardingsphere\_proxy\_version](#input\_shardingsphere\_proxy\_version) | The shardingsphere proxy version | `string` | n/a | yes |
| <a name="input_subnet_ids"></a> [subnet\_ids](#input\_subnet\_ids) | List of subnets sorted by availability zone in your VPC | `list(string)` | n/a | yes |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | The id of your VPC | `string` | n/a | yes |
| <a name="input_zk_servers"></a> [zk\_servers](#input\_zk\_servers) | The Zookeeper servers | `list(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_shardingsphere_domain"></a> [shardingsphere\_domain](#output\_shardingsphere\_domain) | The domain name of the ShardingSphere Proxy Cluster for use by other services |


