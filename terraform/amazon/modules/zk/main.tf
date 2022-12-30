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

data "aws_ami" "base" {
  owners = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-ebs"]
  }

  most_recent = true
}

data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_network_interface" "zk" {
  count           = var.cluster_size
  subnet_id       = element(var.subnet_ids, count.index)
  security_groups = var.security_groups
}

resource "aws_instance" "zk" {
  count         = var.cluster_size
  ami           = data.aws_ami.base.id
  instance_type = var.instance_type
  key_name      = var.key_name

  network_interface {
    delete_on_termination = false
    device_index          = 0
    network_interface_id  = element(aws_network_interface.zk.*.id, count.index)
  }

  tags = merge(
    var.tags,
    {
      Name = "zk-${count.index}"
    }
  )

  user_data = base64encode(templatefile("${path.module}/cloud-init.yml", {
    version     = var.zk_version
    nodes       = range(1, var.cluster_size + 1)
    domain      = var.hosted_zone_name
    index       = count.index + 1
    client_port = var.zk_config["client_port"]
    zk_heap     = var.zk_config["zk_heap"]
  }))

  lifecycle {
    ignore_changes = [
      # Ignore changes to tags.
      tags
    ]
  }
}

data "aws_route53_zone" "zone" {
  name         = "${var.hosted_zone_name}."
  private_zone = true
}

resource "aws_route53_record" "zk" {
  count   = var.cluster_size
  zone_id = data.aws_route53_zone.zone.zone_id
  name    = "zk-${count.index + 1}"
  type    = "A"
  ttl     = 60
  records = element(aws_network_interface.zk.*.private_ips, count.index)
}
