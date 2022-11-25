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

data "aws_availability_zones" "available" {
  state = "available"
}

data "aws_vpc" "vpc" {
  id = var.vpc_id
}

resource "aws_network_interface" "ss" {
  count           = var.cluster_size
  subnet_id       = element(var.subnet_ids, count.index)
  security_groups = var.security_groups
}

resource "aws_launch_template" "ss" {
  count                                = var.cluster_size
  name                                 = "ss-${element(data.aws_availability_zones.available.names, count.index)}"
  image_id                             = var.image_id
  instance_initiated_shutdown_behavior = "terminate"
  instance_type                        = var.instance_type
  key_name                             = var.key_name

  user_data = base64encode(templatefile("${path.module}/cloud-init.yml", {
    version    = var.shardingsphere_version
    zk_servers = join(",", var.zk_servers)
  }))

  metadata_options {
    http_endpoint               = "enabled"
    http_tokens                 = "required"
    http_put_response_hop_limit = 1
    instance_metadata_tags      = "enabled"
  }

  monitoring {
    enabled = true
  }

  network_interfaces {
    delete_on_termination = false
    device_index          = 0
    network_interface_id  = element(aws_network_interface.ss.*.id, count.index)
  }

  tag_specifications {
    resource_type = "instance"

    tags = {
      Name = "ss-${count.index + 1}"
    }
  }
}

resource "aws_autoscaling_group" "ss" {
  count                     = var.cluster_size
  name                      = "ss-${count.index + 1}"
  availability_zones        = [element(data.aws_availability_zones.available.names, count.index)]
  desired_capacity          = 1
  max_size                  = 1
  min_size                  = 1
  health_check_grace_period = 300
  health_check_type         = "EC2"

  launch_template {
    id = element(aws_launch_template.ss.*.id,
      index(
        aws_launch_template.ss.*.name,
        "ss-${element(data.aws_availability_zones.available.names, count.index)}"
      )
    )
    version = "$Latest"
  }

  lifecycle {
    ignore_changes = [load_balancers, target_group_arns]
  }
}

resource "aws_lb" "ss" {
  name               = "ss-internal-lb"
  internal           = true
  load_balancer_type = "network"

  enable_deletion_protection = false

  dynamic "subnet_mapping" {
    for_each = var.subnet_ids
    content {
      subnet_id = subnet_mapping.value
    }
  }

  tags = {
    Name = "shardingsphere"
  }
}

resource "aws_lb_target_group" "ss_tg" {
  name               = "shardingsphere-lb-tg"
  port               = var.lb_listener_port
  protocol           = "TCP"
  vpc_id             = var.vpc_id
  preserve_client_ip = false

  tags = {
    Name = "shardingsphere"
  }
}

resource "aws_autoscaling_attachment" "asg_attachment_lb" {
  count                  = var.cluster_size
  autoscaling_group_name = element(aws_autoscaling_group.ss.*.id, count.index)
  lb_target_group_arn    = aws_lb_target_group.ss_tg.arn
}


resource "aws_lb_listener" "ss" {
  load_balancer_arn = aws_lb.ss.arn
  port              = var.lb_listener_port
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ss_tg.arn
  }

  tags = {
    Name = "shardingsphere"
  }
}

data "aws_route53_zone" "zone" {
  name         = "${var.hosted_zone_name}."
  private_zone = true
}

resource "aws_route53_record" "ss" {
  zone_id = data.aws_route53_zone.zone.zone_id
  name    = "proxy"
  type    = "A"

  alias {
    name                   = aws_lb.ss.dns_name
    zone_id                = aws_lb.ss.zone_id
    evaluate_target_health = true
  }
}