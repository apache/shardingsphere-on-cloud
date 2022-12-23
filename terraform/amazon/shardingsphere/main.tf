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

resource "aws_iam_role" "sts" {
  name = "shardingsphere-proxy-sts-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "ss" {
  name = "sharidngsphere-proxy-policy"
  role = aws_iam_role.sts.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "cloudwatch:PutMetricData",
        "ec2:DescribeTags",
        "logs:PutLogEvents",
        "logs:DescribeLogStreams",
        "logs:DescribeLogGroups",
        "logs:CreateLogStream",
        "logs:CreateLogGroup"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_instance_profile" "ss" {
  name = "shardingsphere-proxy-instance-profile"
  role = aws_iam_role.sts.name
}

resource "aws_launch_template" "ss" {
  name                                 = "shardingsphere-proxy-launch-template"
  image_id                             = var.image_id
  instance_initiated_shutdown_behavior = "terminate"
  instance_type                        = var.instance_type
  key_name                             = var.key_name
  iam_instance_profile {
    name = aws_iam_instance_profile.ss.name
  }

  user_data = base64encode(templatefile("${path.module}/cloud-init.yml", {
    version    = var.shardingsphere_proxy_version
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

  vpc_security_group_ids = var.security_groups

  tag_specifications {
    resource_type = "instance"

    tags = {
      Name = "shardingsphere-proxy"
    }
  }
}

resource "aws_autoscaling_group" "ss" {
  name                      = "shardingsphere-proxy-asg"
  availability_zones        = data.aws_availability_zones.available.names
  desired_capacity          = var.shardingsphere_proxy_asg_desired_capacity
  min_size                  = 1
  max_size                  = var.shardingsphere_proxy_asg_max_size
  health_check_grace_period = var.shardingsphere_proxy_asg_healthcheck_grace_period
  health_check_type         = "ELB"

  launch_template {
    id      = aws_launch_template.ss.id
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
    Name = "shardingsphere-proxy"
  }
}

resource "aws_lb_target_group" "ss_tg" {
  name               = "shardingsphere-proxy-lb-tg"
  port               = var.lb_listener_port
  protocol           = "TCP"
  vpc_id             = var.vpc_id
  preserve_client_ip = false

  health_check {
  	protocol = "TCP"
    healthy_threshold = 2
    unhealthy_threshold = 2
  }

  tags = {
    Name = "shardingsphere-proxy"
  }
}

resource "aws_autoscaling_attachment" "asg_attachment_lb" {
  autoscaling_group_name = aws_autoscaling_group.ss.id
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
    Name = "shardingsphere-proxy"
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
