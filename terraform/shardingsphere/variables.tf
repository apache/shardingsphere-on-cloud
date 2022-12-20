#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

variable "cluster_size" {
  type        = number
  description = "The cluster size that same size as available_zones"
}

variable "shardingsphere_proxy_version" {
  type        = string
  description = "The shardingsphere proxy version"
}

variable "shardingsphere_proxy_asg_desired_capacity" {
  type        = string
  default     = "3"
  description = "The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain. see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacitytype, The default value is 3"
}

variable "shardingsphere_proxy_asg_max_size" {
  type        = string
  default     = "6"
  description = "The maximum size of ShardingSphere Proxy Auto Scaling Group. The default values is 6"
}

variable "shardingsphere_proxy_asg_healthcheck_grace_period" {
  type        = number
  default     = 120
  description = "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check. see https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html"
}

variable "image_id" {
  type        = string
  description = "The AMI id"
}

variable "key_name" {
  type        = string
  description = "the ssh keypair for remote connection"
}

variable "instance_type" {
  type        = string
  description = "The EC2 instance type"
}

variable "vpc_id" {
  type        = string
  description = "The id of your VPC"
}

variable "subnet_ids" {
  type        = list(string)
  description = "List of subnets sorted by availability zone in your VPC"
}

variable "security_groups" {
  type        = list(string)
  default     = []
  description = "List of The Security group IDs"
}

variable "lb_listener_port" {
  type        = string
  description = "lb listener port"
}

variable "hosted_zone_name" {
  type        = string
  default     = "shardingsphere.org"
  description = "The name of the hosted private zone"
}

variable "zk_servers" {
  type        = list(string)
  description = "The Zookeeper servers"
}
