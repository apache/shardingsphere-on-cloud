+++
pre = "<b>2.4 </b>"
title = "CloudFormation 部署多可用区 ShardingSphere Proxy 集群"
weight = 4
chapter = true
+++

## 背景

ShardingSphere Proxy 集群作为数据基础设施重要的一部分，集群自身的高可用性尤为重要，本部分内容将介绍使用 CloudFormation 在 Amazon 上从零搭建一套满足高可用的 ShardingSphere Proxy 集群。

## 目标

我们将创建如下架构图的 ShardingSphere Proxy 高可用集群：

![](../../../../img/overview/terraform.png)

创建的 Amazon 资源如下：
1. 每个可用区一个 ZooKeeper 实例。
2. 每个可用区一个 Auto Scaling Group。
3. 每个可用区一个 Launch Template, 用于给 Auto Scaling Group 启动 ShardingSphere Proxy 实例。
4. 一个内网 Network LoadBalancer, 给应用使用。

## 快速开始

### 前提条件

为创建 ShardingSphere Proxy 高可用集群，您需要事先准备如下资源：
1. 一个 ssh keypair，用于远程连接 EC2 实例。
2. 一个 VPC。
3. 每个可用区的 subnet。
4. 一个 SecurityGroup, 能够放行 ZooKeeper Server 使用的 2888，3888，2181 端口。
5. 一个内网 HostedZone。
6. 一个通用的 AMI 镜像， Amazon linux2 即可。
7. 最好准备好 CloudFormation [配置文件](https://raw.githubusercontent.com/apache/shardingsphere-on-cloud/main/cloudformation/multi-az/cf.json)。

### 步骤

1. 进入 Amazon CloudFormation 服务，创建 Stacks。

![](../../../../img/operation-guide/4-1.PNG)

点击 `Choose File` 按钮 上传准备好的 CloudFormation 配置。

![](../../../../img/operation-guide/4-2.PNG)

上传好后点击 `Next` 按钮。

2. 将您准备好的资源填入以下对应的相关位置。

![](../../../../img/operation-guide/4-3.PNG)

![](../../../../img/operation-guide/4-4.PNG)

填入相应参数后，点击 `Next`  按钮。

3. 按您实际情况配置 `stack` 相关参数。

![](../../../../img/operation-guide/4-5.PNG)

![](../../../../img/operation-guide/4-6.PNG)

配置好后点击 `Next` 按钮。

4. 进行配置 `Review`。

![](../../../../img/operation-guide/4-7.PNG)

![](../../../../img/operation-guide/4-8.PNG)

![](../../../../img/operation-guide/4-9.PNG)

确认好点击 `Submit` 按钮。

5. 在上述操作后，将进入创建阶段。

![](../../../../img/operation-guide/4-10.PNG)

![](../../../../img/operation-guide/4-11.PNG)

![](../../../../img/operation-guide/4-12.PNG)

6. 等待一段时间，创建完成后，进入 `Outputs` 标签页，如下图。

![](../../../../img/operation-guide/4-13.PNG)

其中 `ssinernaldomain` 对应的值就是我们需要的域名。

默认创建的内部域名为 [proxy.shardingsphere.org](proxy.shardingsphere.org)，端口为 3307，用户名和密码为 root。

## 使用手册

### CloudFormation 配置

#### 参数列表

|名称                      |描述                                                       |类型              |默认值|
|--------------------------|-----------------------------------------------------------|-----------------|------|
|HostedZoneId              |内网  HostedZone Id                                        |String               | |
|HostedZoneName            |内网 HostedZone 名称                                        |String          |[shardingsphere.org](shardingsphere.org)|
|ImageId                   |AMI Id， 需是Amazon Linux 2 类型或者包管理是 yum 的 Linux 系列|String          | |
|KeyName                   |SSH 密钥对                                                  |String          | |
|VpcId                     |VPC Id                                                     |String            | |
|Subnets                   |VPC 中的子网列表，顺序需要和按可用区字母排序的顺序一致          |CommaDelimitedList| |
|SecurityGroupIds          |安全组列表，需要放行 ZooKeeper Server 的 2181，2888，3888 端口|CommaDelimitedList| |
|ShardingSphereInstanceType|ShardingSphere Proxy Server 的 EC2 实例类型                 |String            | |
|ShardingSphereJavaMemOpts |ShardingSphere Proxy Server 的 jvm 内存参数                 |String            |-Xmx512m -Xms512m -Xmn128m|
|ShardingSpherePort        |ShardingSphere Proxy 的端口                                 |String            |3307|
|ShardingSphereVersion     |ShardingSphere Proxy 的版本                                 |String            |5.2.1|
|ZookeeperHeap             |Zookeeper Server 的 jvm Heap 大小，单位为 m                  |String            |512|
|ZookeeperInstanceType     |Zookeeper Server 的 EC2 实例类型                             |String            |t2.nano|
|ZookeeperVersion          |Zookeeper Server 版本号                                      |String            |3.7.1|

#### 输出列表

|名称|描述|导出名称|值|
|----|---|--------|--|
|ZK1|Zookeeper Server1 信息|{'Fn::Sub': '${AWS::StackName}-Zookeeper-Server-1'}|{'Fn::Join': [':', [{'Ref': 'ZK1'}, {'Fn::GetAtt': ['ZK1', 'PrivateIp']}, {'Fn::GetAtt': ['ZK1', 'AvailabilityZone']}]]}|
|ZK2|Zookeeper Server2 信息| {'Fn::Sub': '${AWS::StackName}-Zookeeper-Server-2'} |{'Fn::Join': [':', [{'Ref': 'ZK2'}, {'Fn::GetAtt': ['ZK2‘, 'PrivateIp']}, {'Fn::GetAtt': ['ZK2', 'AvailabilityZone']}]]}|
|ZK3|Zookeeper Server3 信息|{'Fn::Sub': '${AWS::StackName}-Zookeeper-Server-3'}|{'Fn::Join': [':', [{'Ref': 'ZK2'}, {'Fn::GetAtt': ['ZK2', 'PrivateIp']}, {'Fn::GetAtt': ['ZK2', 'AvailabilityZone']}]]}|
|zoneZK1|Zookeeper Server1 内部域名|{'Fn::Sub': '${AWS::StackName}-Zookeeper-Domain-1'}| {'Ref': 'zoneZK1'}|
|zoneZK2|Zookeeper Server2 内部域名| {'Fn::Sub': '${AWS::StackName}-Zookeeper-Domain-2'}|{'Ref': 'zoneZK2'}|
|zoneZK3|Zookeeper Server3 内部域名|{'Fn::Sub': '${AWS::StackName}-Zookeeper-Domain-3'}| {'Ref': 'zoneZK3'}|
|ssinternaldomain|ShardingSphere Proxy 对外使用的内部域名|{'Fn::Sub': '${AWS::StackName}-ShardingSphere-Internal-Domain'}|{'Ref': 'ssinternaldomain'}|

## 运维

默认使用我们提供的 CloudFormation 创建的 ZooKeeper 和 ShardingSphere Proxy 服务可以使用 Systemd 管理。

### ZooKeeper

#### 启动

```shell
systemctl start zookeeper
```

#### 停止

```shell
systemctl stop zookeeper
```

#### 重启

```shell
systemctl restart zookeeper
```

### ShardingSphere Proxy

#### 启动

```shell
systemctl start shardingsphere
```

#### 停止

```shell
systemctl stop shardingsphere
```

#### 重启

```shell
systemctl restart shardingsphere
```

## 开发手册

此 CloudFormation 涉及以下资源列表。

|资源名称        |类型|
|----------------|----|
|ZK1             |AWS::EC2::Instance|
|ZK2             |AWS::EC2::Instance|
|ZK3             |AWS::EC2::Instance|
|zoneZK1         |AWS::Route53::RecordSet|
|zoneZK2         |AWS::Route53::RecordSet|
|zoneZK3         |AWS::Route53::RecordSet|
|networkiface0   |AWS::EC2::NetworkInterface|
|networkiface1   |AWS::EC2::NetworkInterface|
|networkiface2   |AWS::EC2::NetworkInterface|
|launchtemplate0 |AWS::EC2::LaunchTemplate|
|launchtemplate1 |AWS::EC2::LaunchTemplate|
|launchtemplate2 |AWS::EC2::LaunchTemplate|
|ssinternallb    |AWS::ElasticLoadBalancingV2::LoadBalancer|
|sslbtg          |AWS::ElasticLoadBalancingV2::TargetGroup|
|autoscaling0    |AWS::AutoScaling::AutoScalingGroup |
|autoscaling1    |AWS::AutoScaling::AutoScalingGroup |
|autoscaling2    |AWS::AutoScaling::AutoScalingGroup |
|sslblistener    |AWS::ElasticLoadBalancingV2::Listener|
|ssinternaldomain|AWS::Route53::RecordSet|

### 依赖

我们使用 [cfndsl](https://github.com/cfndsl/cfndsl) 生成 CloudFormation 配置。

您需要按照 [cfndsl](https://github.com/cfndsl/cfndsl)  提供的步骤去安装。

### 步骤

1. 初始化 `cfndsl`，只需运行一次。

```shell
cfndsl -u 94.0.0
```

2. 修改 `cf.rb` 配置后，运行下面命令生成 CloudFormation 配置。

```shell
 cfndsl cf.rb -o cf.json --pretty
```
