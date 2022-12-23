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

variable "shardingsphere_proxy_version" {
  type        = string
  description = "The shardingsphere proxy version"
}

variable "shardingsphere_proxy_as_desired_number" {
  type        = number
  default     = 3
  description = "The initial expected number of ShardSphere Proxy Auto Scaling. The default value is 3"
}

variable "shardingsphere_proxy_as_max_number" {
  type        = number
  default     = 6
  description = "The maximum size of ShardingSphere Proxy Auto Scaling. The default value is 6"
}

variable "shardingsphere_proxy_as_healthcheck_grace_period" {
  type        = number
  default     = 120
  description = "The health check grace period for instances, in seconds"
}

variable "image_id" {
  type        = string
  default     = ""
  description = "The image id"
}

variable "key_name" {
  type        = string
  description = "the ssh keypair for remote connection"
}

variable "flavor_id" {
  type        = string
  description = "The flavor id of the ECS"
}

variable "vpc_id" {
  type        = string
  description = "The id of your VPC"
}

variable "vip_subnet_id" {
  type        = string
  default     = ""
  description = "The IPv4 subnet ID of the subnet where the load balancer works"
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
  description = "The lb listener port"
}

variable "zone_id" {
  type        = string
  default     = ""
  description = "The id of the private zone"
}

variable "zone_name" {
  type        = string
  default     = "shardingsphere.org"
  description = "The name of the private zone"
}

variable "shardingsphere_proxy_doamin_prefix_name" {
  type        = string
  default     = "proxy"
  description = "The prefix name of the shardinsphere domain, the final generated name will be [prefix_name].[zone_name], the default value is proxy."
}

variable "zk_servers" {
  type        = list(string)
  default     = []
  description = "The Zookeeper servers"
}

variable "zk_cluster_size" {
  type        = number
  default     = 3
  description = "The Zookeeper cluster size"
}

variable "zk_flavor_id" {
  type        = string
  description = "The ECS instance type"
}


