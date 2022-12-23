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

# See https://discuss.hashicorp.com/t/using-a-non-hashicorp-provider-in-a-module/21841/2
terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = "1.43.0"
    }
  }
}

data "huaweicloud_vpc_subnet" "zknet" {
  name = "subnet-default"
}

data "huaweicloud_images_image" "myimage" {
  name        = "Ubuntu 22.04 server 64bit"
  most_recent = true
}


resource "huaweicloud_vpc_eip" "zkeip" {
  count = var.cluster_size

  publicip {
    type = "5_bgp"
  }

  bandwidth {
    name        = "zk-${count.index + 1}"
    size        = 300
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_compute_instance" "zk" {
  count              = var.cluster_size
  name               = "zk-${count.index + 1}"
  image_id           = length(var.image_id) == 0 ? data.huaweicloud_images_image.myimage.id : var.image_id
  flavor_id          = var.flavor_id
  key_pair           = var.key_name
  security_group_ids = var.security_groups

  network {
    uuid = data.huaweicloud_vpc_subnet.zknet.id
  }

  tags = merge(
    var.tags,
    {
      Name = "zk-${count.index + 1}"
    }
  )

  user_data = templatefile("${path.module}/cloud-init.yml", {
    version     = var.zk_version
    nodes       = range(1, var.cluster_size + 1)
    domain      = var.zone_name
    index       = count.index + 1
    client_port = var.zk_config["client_port"]
    zk_heap     = var.zk_config["zk_heap"]
  })

  lifecycle {
    ignore_changes = [
      user_data, data_disks,
    ]
  }
}

resource "huaweicloud_compute_eip_associate" "associated" {
  count       = var.cluster_size
  public_ip   = element(huaweicloud_vpc_eip.zkeip.*.address, count.index)
  instance_id = element(huaweicloud_compute_instance.zk.*.id, count.index)
}

resource "huaweicloud_dns_recordset" "zk" {
  count       = var.cluster_size
  zone_id     = var.zone_id
  name        = "zk-${count.index + 1}.${var.zone_name}."
  description = "The zk-${count.index + 1} record set"
  ttl         = 300
  type        = "A"
  records     = [element(huaweicloud_compute_instance.zk.*.access_ip_v4, count.index)]
}
