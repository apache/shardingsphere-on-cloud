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
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = "1.43.0"
    }
  }
}

provider "huaweicloud" {
  region     = "cn-north-9"
}

locals {
  create_zk_servers   = length(var.zk_servers) == 0
  zk_servers          = local.create_zk_servers ? module.zk.zk_node_domain : var.zk_servers
  create_private_zone = length(var.zone_id) == 0
  zone_id             = local.create_private_zone ? resource.huaweicloud_dns_zone.private_zone[0].id : var.zone_id
}

resource "huaweicloud_dns_zone" "private_zone" {
  count       = local.create_private_zone ? 1 : 0
  name        = "${var.zone_name}."
  description = "The shardingsphere proxy zone"
  ttl         = 300
  zone_type   = "private"

  router {
    router_id = var.vpc_id
  }
}

data "huaweicloud_availability_zones" "zones" {}

module "zk" {
  source          = "./modules/zk"
  cluster_size    = var.zk_cluster_size
  key_name        = var.key_name
  flavor_id       = var.zk_flavor_id
  vpc_id          = var.vpc_id
  zone_id         = local.zone_id
  security_groups = var.security_groups
}

data "huaweicloud_images_image" "myimage" {
  name        = "Ubuntu 22.04 server 64bit"
  most_recent = true
}

data "huaweicloud_vpc_subnet" "vipnet" {
  #4b9db05b-4d57-464d-a9fe-83da3de0a74c
  vpc_id = var.vpc_id
}

resource "huaweicloud_elb_loadbalancer" "ss" {
  name              = "shardingsphere-proxy-internal-lb"
  vpc_id            = var.vpc_id
  ipv4_subnet_id    = length(var.vip_subnet_id) == 0 ? data.huaweicloud_vpc_subnet.vipnet.ipv4_subnet_id : var.vip_subnet_id
  l4_flavor_id      = "02308113-fee6-4f06-997d-1cbd0e0b3f04"
  cross_vpc_backend = true
  availability_zone = data.huaweicloud_availability_zones.zones.names

  lifecycle {
    ignore_changes = [
      ipv6_bandwidth_id, iptype, bandwidth_charge_mode, sharetype, bandwidth_size,
    ]
  }
}

resource "huaweicloud_elb_listener" "ss" {
  name            = "shardingsphere-proxy-internal-lb-listener"
  protocol        = "TCP"
  protocol_port   = var.lb_listener_port
  loadbalancer_id = huaweicloud_elb_loadbalancer.ss.id
}

resource "huaweicloud_elb_pool" "ss" {
  name        = "shardingsphere-proxy-internal-lb-pool"
  protocol    = "TCP"
  lb_method   = "ROUND_ROBIN"
  listener_id = huaweicloud_elb_listener.ss.id
}

resource "huaweicloud_elb_monitor" "ss" {
  protocol    = "TCP"
  interval    = 10
  timeout     = 5
  max_retries = 3
  port        = var.lb_listener_port
  pool_id     = huaweicloud_elb_pool.ss.id
}


resource "huaweicloud_as_configuration" "ss" {
  scaling_configuration_name = "shardingsphere-proxy-as-config"

  instance_config {
    flavor             = var.flavor_id
    image              = length(var.image_id) == 0 ? data.huaweicloud_images_image.myimage.id : var.image_id
    key_name           = var.key_name
    security_group_ids = var.security_groups

    disk {
      size        = 40
      volume_type = "SSD"
      disk_type   = "SYS"
    }

    public_ip {
      eip {
        ip_type = "5_bgp"
        bandwidth {
          share_type    = "PER"
          charging_mode = "traffic"
          size          = 300
        }
      }
    }

    user_data = templatefile("./cloud-init.yml", {
      version       = var.shardingsphere_proxy_version
      version_elems = split(".", var.shardingsphere_proxy_version)
      zk_servers    = join(",", local.zk_servers)
    })
  }

  lifecycle {
    ignore_changes = [
      instance_config.0.user_data,
    ]
  }
}


resource "huaweicloud_as_group" "ss" {
  scaling_group_name       = "shardingsphere-proxy-asg"
  scaling_configuration_id = resource.huaweicloud_as_configuration.ss.id
  desire_instance_number   = var.shardingsphere_proxy_as_desired_number
  min_instance_number      = 3
  max_instance_number      = var.shardingsphere_proxy_as_max_number
  vpc_id                   = var.vpc_id
  availability_zones       = data.huaweicloud_availability_zones.zones.names

  health_periodic_audit_method       = "ELB_AUDIT"
  health_periodic_audit_time         = 1
  health_periodic_audit_grace_period = var.shardingsphere_proxy_as_healthcheck_grace_period

  delete_publicip  = true
  delete_instances = "yes"
  force_delete     = true

  dynamic "networks" {
    for_each = var.subnet_ids
    content {
      id = networks.value
    }
  }

  lbaas_listeners {
    pool_id       = huaweicloud_elb_pool.ss.id
    protocol_port = huaweicloud_elb_listener.ss.protocol_port
  }
}

resource "huaweicloud_dns_recordset" "ss" {
  zone_id     = local.zone_id
  name        = "${var.shardingsphere_proxy_doamin_prefix_name}.${var.zone_name}."
  description = "The shardingsphere-proxy internal domain name"
  ttl         = 300
  type        = "A"
  records     = [resource.huaweicloud_elb_loadbalancer.ss.ipv4_address]
}

