+++
pre = "<b>2.2 </b>"
title = "ShardingSphere CloudFormation 一键安装"
weight = 1
chapter = true
+++

## 操作说明

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

![](../../../img/operation-guide/4-1.PNG)

点击 `Choose File` 按钮 上传准备好的 CloudFormation 配置。

![](../../../img/operation-guide/4-2.PNG)

上传好后点击 `Next` 按钮。

2. 将您准备好的资源填入以下对应的相关位置。

![](../../../img/operation-guide/4-3.PNG)

![](../../../img/operation-guide/4-4.PNG)

填入相应参数后，点击 `Next`  按钮。

3. 按您实际情况配置 `stack` 相关参数。

![](../../../img/operation-guide/4-5.PNG)

![](../../../img/operation-guide/4-6.PNG)

配置好后点击 `Next` 按钮。

4. 进行配置 `Review`。

![](../../../img/operation-guide/4-7.PNG)

![](../../../img/operation-guide/4-8.PNG)

![](../../../img/operation-guide/4-9.PNG)

确认好点击 `Submit` 按钮。

5. 在上述操作后，将进入创建阶段。

![](../../../img/operation-guide/4-10.PNG)

![](../../../img/operation-guide/4-11.PNG)

![](../../../img/operation-guide/4-12.PNG)

6. 等待一段时间，创建完成后，进入 `Outputs` 标签页，如下图。

![](../../../img/operation-guide/4-13.PNG)

其中 `ssinernaldomain` 对应的值就是我们需要的域名。

默认创建的内部域名为 [proxy.shardingsphere.org](proxy.shardingsphere.org)，端口为 3307，用户名和密码为 root。

7. 应用配置数据库连接信息为 ShardingSphere Proxy
