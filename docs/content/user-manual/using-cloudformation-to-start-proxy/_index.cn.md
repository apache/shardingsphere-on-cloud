+++
pre = "<b>3.3 </b>"
title = "利用 CloudFormation 启动 ShardingSphere Proxy"
weight = 3
chapter = true
+++

AWS CloudFormation 是一个以基础设施即代码的方式配置和启动任何环境和基础设施的简易工具。通过 AWS CloudFormation Stack 模板可以帮助在 AWS 上快速启动 Apache ShardingSphere。

## 前置条件

开始之前，需要确认以下的检查列表清单：

- [ ] 选择区域为 ap-north-1（北京），当前保护 Apache ShardingSphere Proxy 的 AMI 和相关组件仅在 ap-north-1 区域有效
- [ ] 一个已存在的 VPC 用于部署 Apache ShardingSphere Proxy
- [ ] 该 VPC 下一个已规划的 CIDR 和对应子网
- [ ] 允许应用访问数据库（比如 3307 端口）和控制流量（比如 22 端口）的安全组配置
- [ ] 可以用于访问该实例资源的密钥对 
- [ ] 对该 CloudFormation Stack 涉及资源设计的标签

## 启动 ShardingSphere Proxy 集群

### 1. 利用新资源创建 CloudFormation 堆栈

如下图所示：

![](../../../img/operation-guide/1.PNG)

![](../../../img/operation-guide/2.PNG)

### 2. 上传本仓库中的模板文件

上传本地文件 `cloudformation/apache-shardingsphere-5.2.0.json` 到 CloudFormation，然后点击 `Next`。

![](../../../img/operation-guide/3.PNG)

![](../../../img/operation-guide/4.PNG)

### 3. 指定 CloudFormation 堆栈细节

填写本页中的空白项，必填项已在前置条件中就绪。

![](../../../img/operation-guide/5.PNG)

### 4. 配置堆栈选项

为该堆栈添加标签，有助于后续成本分析。

![](../../../img/operation-guide/6.PNG)

### 5. 回顾和确认配置

回顾配置内容，在提交前确认所有内容符合期望。

![](../../../img/operation-guide/7.PNG)

### 6. 检查 EC2 实例

几分钟后，EC2 实例已经启动。

![](../../../img/operation-guide/8.PNG)

### 7. 检查 ShardingSphere Proxy 和 ZooKeeper 状态

使用 `systemctl status shardingsphere-proxy` 和 `./bin/zkServer.sh status` 来检查组件的运行状态。

![](../../../img/operation-guide/9.PNG)

![](../../../img/operation-guide/10.PNG)

### 8. 测试简单的分片示例

创建数据库 `sharding_db`，以及添加两个独立的数据库实例 `resources`。然后创建逻辑表 `t_order` 并插入两行数据。如下检查结果：

![](../../../img/operation-guide/11.PNG)

![](../../../img/operation-guide/12.PNG)

![](../../../img/operation-guide/13.PNG)
