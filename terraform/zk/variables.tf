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

variable "key_name" {
  type        = string
  description = "The ssh keypair for remote connection"
}

variable "instance_type" {
  type        = string
  description = "The EC2 instance type"
}

variable "vpc_id" {
  type        = string
  description = "The id of VPC"
}

variable "subnet_ids" {
  type        = list(string)
  description = "List of subnets sorted by availability zone in your VPC"
}

variable "security_groups" {
  type        = list(string)
  default     = []
  description = "List of the Security Group, it must be allow access 2181, 2888, 3888 port"
}


variable "hosted_zone_name" {
  type        = string
  default     = "shardingsphere.org"
  description = "The name of the hosted private zone"
}

variable "tags" {
  type        = map(any)
  description = "A map of zk instance resource, the default tag is Name=zk-$${count.idx}"
  default     = {}
}

variable "zk_version" {
  type        = string
  description = "The zookeeper version"
  default     = "3.7.1"
}

variable "zk_config" {
  default = {
    client_port = 2181
    zk_heap     = 1024
  }

  description = "The default config of zookeeper server"
}