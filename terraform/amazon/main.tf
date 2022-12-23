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



terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.37.0"
    }
  }
}

provider "aws" {
  region     = "ap-southeast-1"
}

module "zk" {
  source              = "./zk"
  cluster_size        = 3
  key_name            = "test-tf"
  instance_type       = "t2.nano"
  vpc_id              = "vpc-0ef2b7440d3ade8d5"
  subnet_ids          = ["subnet-0f388a6f23063b8c9", "subnet-0bc2cd85facb5ca06", "subnet-009077567350ef1b7"]
  security_groups     = ["sg-008e74936b3f9de19"]
}

module "shardingsphere" {
  depends_on                    = [module.zk]
  source                        = "./shardingsphere"
  cluster_size                  = 3
  shardingsphere_proxy_version  = "5.2.1"
  key_name                      = "test-tf"
  image_id                      = "ami-094bbd9e922dc515d"
  instance_type                 = "t3.medium"
  lb_listener_port              = 3307
  vpc_id                        = "vpc-0ef2b7440d3ade8d5"
  subnet_ids                    = ["subnet-0f388a6f23063b8c9", "subnet-0bc2cd85facb5ca06", "subnet-009077567350ef1b7"]
  security_groups               = ["sg-008e74936b3f9de19"]
  zk_servers                    = module.zk.zk_node_domain
}
