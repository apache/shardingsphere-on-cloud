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


output "zk_node_private_ip" {
  value       = aws_instance.zk.*.private_ip
  description = "The private ips of zookeeper instances"
}

output "zk_node_domain" {
  value       = [for v in aws_route53_record.zk.*.name : format("%s.%s", v, var.hosted_zone_name)]
  description = "The private domain names of zookeeper instances for use by ShardingSphere Proxy"
}