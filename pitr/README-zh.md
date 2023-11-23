# 目录

* [使用说明](#使用说明)
  * [前置准备](#前置说明)
    * [服务器说明](#服务器说明)
    * [环境说明](#环境说明)
      * [编译说明（可选）](#编译说明可选)
    * [SSL 配置](#ssl-配置)
      * [生成新的 SSL 密钥对（可选）](#生成新的-ssl-密钥对可选)
  * [部署说明](#部署说明)
    * [步骤1：获取 Pitr 二进制](#步骤1-获取-pitr-二进制)
      * [下载二进制包](#下载二进制包)
      * [自行编译](#自行编译)
    * [步骤2：准备 ShardingSphere Proxy 配置](#步骤-2-准备-shardingsphere-proxy-配置)
    * [步骤3：配置 OpenGauss](#步骤-3-配置-opengauss)
    * [步骤4：为 Pitr Agent 部署 SSL 证书](#步骤-4-为-pitr-agent-部署-ssl-证书)
    * [步骤5：启动 Pitr Agent](#步骤-5-启动-pitr-agent)
  * [测试说明](#测试说明)
    * [准备测试数据](#准备测试数据)
    * [测试用例](#测试用例)
      * [备份](#备份)
      * [查看备份](#查看备份)
      * [恢复](#恢复)
      * [删除备份](#删除备份)
* [使用限制](#使用限制)

# 使用说明 

本工具是面向 Apache ShardingSphere 和 OpenGauss 构建的分布式数据库集群提供的按时间点恢复（PITR，Point-in-time Recovery）功能的命令行工具。

## 前置准备

在开始之前，你需要准备如下三台服务器，并设置相关运行环境和安装所需依赖软件。这三台服务器的拓扑结构示意如下：

```shell
+------------------------------+             +------------------+
|                              |             | OpenGauss Server |
| Apache ShardingSphere Proxy  |             | Pitr Agent       |
| Apache Zookeeper             | ----------> +------------------+
| Pitr Cli (aka `gs_pitr`)     |             | OpenGauss Server |
|                              |             | Pitr Agent       |
+------------------------------+             +------------------+
```

### 服务器说明

你需要至少准备三台服务器，其中一台作为 Pitr 命令行工具的执行环境，它上面还会部署 Apache ShardingSphere，另外两台则用来部署 OpenGauss 和 Pitr agent。

| | Role | Components |
|:-:|:-:|:-:|
|1| Pitr cli operation server | Pitr Cli + ShardingSphere Proxy + Zookeeper + GLT |
|2| OpenGauss Server 1 | OpenGauss Server + Pitr Agent |
|3| OpenGauss Server 2 | OpenGauss Server + Pitr Agent |

### 环境说明 

在服务器都准备就绪后，你需要检查和确认如下内容：

- Apache ShardingSphere 所在的服务器允许访问呢 OpenGauss 所在的服务器
- 允许从外部通过 3307 端口访问 Apache ShardingSphere
- 允许从外部通过 18080 端口访问 OpenGauss 服务器上的 Pitr Agent 
- 在 OpenGauss 服务器上设置如下环境变量：
  - export PGDATABASE=tpccdb
  - export PGPORT=13100
- OpenGauss 使用用户 `omm` 并且可以访问数据库 `omm`
- OpenGauss 开启了 `cbm tracking`
- SSL 密钥对。用来提供 Pitr 命令行工具和 Pitr Agent 之间的安全通信，可以使用任何有效的密钥对
- 需要手动在每个节点创建期望的备份数据路径，并保证多个节点路径是一致的
- 需要部署 GLT 服务，比如 Redis，用来向 ShardingSphere 和 OpenGauss 构成的分布式数据库提供全局 CSN

#### 编译说明（可选）

一般来说，Pitr 命令行工具，包括 cli 二进制和 agent 二进制都可以在 [Apache ShardingSphere-on-Cloud 的发布页](https://github.com/apache/shardingsphere-on-cloud/releases) 进行下载。

如果希望自行编译二进制程序，你可以使用 Golang 1.20（推荐）以及  Linux 3.10.0-957.el7.x86_64（至少）作为编译环境，并按照如下步骤编译得到 Pitr 命令行工具 cli 和 agent。

第一步：克隆项目

```shell
git clone git@github.com:apache/shardingsphere-on-cloud.git
```

第二步：编译 Pitr Agent

```shell
cd shardingsphere-on-cloud/pitr/agent
make build
```

第三步：编译 Pitr Cli

```shell
cd shardingsphere-on-cloud/pitr/cli
make build
```

### SSL 配置

Pitr 命令行工具 cli 和 agent 的通信由一对 SSL 密钥对进行加密，你可以选择使用任何可用的密钥对或者生成一对新的密钥对，例如：

- tls.key
- tls.crt

密钥对需要部署在 Pitr agent 所在的服务器上。

#### 生成新的 SSL 密钥对（可选）

如果你想要生成一对密钥对，请确保你拥有一个可用的 OpenSSL 环境。可以通过检查环境变量 OPENSSL_CONF 来查找 OpenSSL 使用的配置文件，默认的地址是 `/etc/pki/tls/openssl.cnf`。

然后使用如下 `pitr/agent` 目录下的脚本并执行如下命令进行生成：

```shell
cd shardingsphere-on-cloud/pitr/agent
make openssl-local
```

现在，正常情况下可以在 `./certs` 目录下得到一对新的密钥对。

## 部署说明

Pitr cli （即 `gs_pitr`）和 Pitr agent （即 `pitr-agent`）二进制都可以在 [Apache ShardingSphere on Cloud 的发布页](https://github.com/apache/shardingsphere-on-cloud/releases)下载，或者在你的环境中按前述步骤手动编译得到。

整个部署过程由如下两个步骤构成：

1. 根据步骤 1-2 部署 Apache ShardingSphere Proxy，Zookeeper 和 Pitr cli 
2. 根据步骤 3-5 部署 OpenGauss 和 Pitr agent

### 步骤1: 获取 Pitr 二进制 

You can download pre-compiled Pitr tools binary release or compile them yourself from source code.

#### 下载二进制包

二进制包以 .tar.gz 的文件存放在[发布页](https://github.com/apache/shardingsphere-on-cloud/releases)，你可以下载期望版本并解压得到二进制文件 `gs_pitr` 和 `pitr-agent`。

#### 自行编译 

根据`前置条件`章节的`编译说明`小节获取具体的操作说明。

在成功获取二进制文件后，你需要将 `gs_pitr` 放在和 Apache ShardingSphere Proxy 相同的服务器上，并将 `pitr-agent` 放在 OpenGauss 所在的所有服务器上。

### 步骤 2: 准备 ShardingSphere Proxy 配置

使用 OpenGauss 主机并替换如下配置中的 ${OPENGAUSS_SERVER_1} 和 ${OPENGAUSS_SERVER_2}：

server.yaml

```yaml
mode:
  type: Cluster
  repository:
    type: ZooKeeper
    props:
      namespace: governance
      server-lists: localhost:2181
      retryIntervalMilliseconds: 500
      timeToLiveSeconds: 60
      maxRetries: 3
      operationTimeoutMilliseconds: 500

authority:
  users:
    - user: root@%
      password: root
    - user: sharding
      password: sharding
  privilege:
    type: ALL_PERMITTED

transaction:
  defaultType: XA
  providerType: Atomikos

props:
  proxy-frontend-database-protocol-type: openGauss

# 以下配置为 GLT 相关配置
globalClock:
  enabled: true
  type: TSO
  provider: redis
  props:
    host: 127.0.0.1
    port: 6379

```

config-sharding.yaml
```yaml
databaseName: sharding_db
dataSources:
  ds_0:
    url: jdbc:opengauss://${OPENGAUSS_SERVER_1}:13100/tpccdb?useSSL=false
    username: root
    password: root
    connectionTimeoutMilliseconds: 30000
    idleTimeoutMilliseconds: 60000
    maxLifetimeMilliseconds: 1800000
    maxPoolSize: 50
    minPoolSize: 1

  ds_1:
    url: jdbc:opengauss://${OPENGAUSS_SERVER_2}:13100/tpccdb?useSSL=false
    username: root
    password: root
    connectionTimeoutMilliseconds: 30000
    idleTimeoutMilliseconds: 60000
    maxLifetimeMilliseconds: 1800000
    maxPoolSize: 50
    minPoolSize: 1
```

然后使用脚本 `bin/start.sh` 来启动 ShardingSphere Proxy。脚本可以在 apache-shardingsphere-{version}-shardingsphere-proxy-bin.tar.gz 中找到。

### 步骤 3: 配置 OpenGauss

a. 在 postgres.conf 中开启 `cbm tracking`

```shell
enable_cbm_tracking = on
```

b. 执行 `gs_probackup init -B ${backup-path}` 来设置期望的备份路径

然后可以启动所有的 OpenGauss 服务。


### 步骤 4: 为 Pitr Agent 部署 SSL 证书

在你启动 Pitr agent 之前，需要先为 Pitr agent 部署证书：

如果 SSL 密钥对是按前述步骤自行编译的，证书文件所在的目录为 `shardingsphere-on-cloud/pitr/agent/certs`。你可以切换目录到证书目录，并执行如下命令：

```shell
scp tls.crt tls.key root@${OPENGAUSS_SERVER_1}:/home/omm/
scp tls.crt tls.key root@${OPENGAUSS_SERVER_2}:/home/omm/
```

或者使用现有可用的密钥对同样需要将其部署到 OpenGauss 服务器上相同的路径下。

### 步骤 5: 启动 Pitr Agent 

1. 拷贝二进制文件 

```shell
cd shardingsphere-on-cloud/pitr/agent

scp pitr-agent root@${OPENGAUSS_SERVER_1}:/home/omm/
scp pitr-agent root@${OPENGAUSS_SERVER_2}:/home/omm/
```

2. 登录 OpenGauss 服务器并切换目录至 `/home/omm`

这里是 `/home/omm` 目录下面的文件：

```shell
$ ll
total 13M
drwx------  4 omm  omm    32 Mar  2 14:22 data
drwx------ 29 omm  omm  4.0K May 23 11:37 pgdata
-rwxr-xr-x  1 root root  13M May 16 18:25 pitr-agent
-rwxr-xr-x  1 root root 1.1K May 16 18:26 tls.crt
-rwxr-xr-x  1 root root 1.7K May 16 18:26 tls.key
```

3. 启动 Pitr Agent 

```shell
./pitr-agent -pgdata /data/data-glt/d1 -port 18080 -tls-crt tls.crt -tls-key tls.key -log-level debug
```

参数说明:
- pgdata: OpenGauss 数据存储路径
- port: Pitr agent 暴露端口
- tls-crt: TLS 证书文件路径
- tls-key: TLS 私钥文件路径
- log-level: Pitr agent 日志级别 

## 测试说明

### 准备测试数据 

你可以使用 `gspl` 连接到 ShardingSphere Proxy 并为测试生成数据：
```shell
gsql -h127.0.0.1 -p3307 -Usharding -Wsharding -d sharding_db
```
1. 检查存储节点

```SQL
SHOW STORAGE UNITS
```

2. 创建分片规则 `t_user`

```SQL
CREATE SHARDING TABLE RULE t_user(
 STORAGE_UNITS(ds_0,ds_1),
 SHARDING_COLUMN=user_id,
 TYPE(NAME="hash_mod",PROPERTIES("sharding-count"="4"))
);
```
3. 检查分片规则

```SQL
SHOW SHARDING TABLE RULE t_user;
```

4. 创建表 `t_user`
```SQL
CREATE TABLE t_user (
  user_id INT NOT NULL,
  order_id INT NOT NULL,
  status VARCHAR(45) NULL,
  PRIMARY KEY (user_id)
);
```

5. 检查分表节点
```SQL
SHOW SHARDING TABLE NODES;
```

6. 插入测试数据
```SQL
insert into t_user( user_id, order_id, status) values(1,1,1);
insert into t_user( user_id, order_id, status) values(2,2,2);
insert into t_user( user_id, order_id, status) values(3,3,3);
insert into t_user( user_id, order_id, status) values(4,4,4);

select * from t_user;
```

### 测试用例 

#### 备份 

执行备份：
```Shell
./gs_pitr backup --host ${OPENGAUSS_SERVER_1} --password sharding --port 3307 --username sharding --agent-port 18080 --dn-threads-num 10 --dn-backup-path "/home/omm/data" -b FULL
```

参数说明:
- host: SharidngSphere Proxy 服务器
- port: ShardingSphere Proxy 监听端口 
- username: ShardingSphere Proxy 连接用户名 
- password: ShardingSphere Proxy 连接密码
- agent-port: Pitr Agent 监听端口 
- dn-threads-num: OpenGauss 并发备份数量 
- dn-threads-path: OpenGauss 备份文件路径 
- b: 备份模式 

#### 查看备份 

查看备份：
```Shell
./gs_pitr show 
```

#### 恢复 

你需要先删除部分 `t_user` 表中的记录：
```SQL
delete from t_user where user_id=1;
delete from t_user where user_id=2;
```

执行恢复：
```Shell
./gs_pitr restore --host ${OPENGAUSS_SERVER_1} --password sharding --port 3307 --username sharding --agent-port 18080 --dn-threads-num 10 --dn-backup-path "/home/omm/data" --id ${BACKUP_ID}
```

参数说明:
- host: ShardingSphere Proxy 服务器 
- port: ShardingSphere Proxy 监听端口 
- username: ShardingSphere Proxy 连接用户名 
- password: ShardingSphere Proxy 连接密码
- agent-port: Pitr Agent 监听端口 
- dn-backup-path: OpenGauss 备份文件路径 
- dn-threads-num: OpenGauss 并发恢复数量 
- id: 备份 id 

验证数据:
```SQL
select * from t_user;
```

#### 删除备份

删除备份：
```shell
./gs_pitr delete --host ${OPENGAUSS_SERVER_1} --password sharding --port 3307 --username sharding --agent-port 18080 --dn-backup-path "/home/omm/data" --id ${BACKUP_ID}
```

参数说明：
- host: ShardingSphere Proxy 服务器 
- port: ShardingSphere Proxy 监听端口 
- username: ShardingSphere Proxy 连接用户名 
- password: ShardingSphere Proxy 连接密码
- agent-port: Pitr Agent 监听端口 
- dn-backup-path: OpenGauss 备份文件路径 
- id: 备份 id 

# 使用限制

- Pitr 备份恢复功能的使用需要开启 GLT，并在 ShardingSphere 中进行配置。如果没有 GLT，那么 Pitr 无法依据 CSN 保证一致性
- GLT 部署可以使用 Redis，无需对 Redis 进行额外配置
- 全局备份任务需要在没有进行中的事务的时间点进行开启，由 ShardingSphere 来加锁保证
- 备份开始后 ShardingSphere 会一直持有锁，当备份结束后才会释放锁
- 多个 Pitr cli 客户端同时操作，只有一个 Pitr cli 客户端可执行成功
- 恢复前后 OpenGauss 数据节点的 IP 地址和端口需保持不变，即和 ShardingSphere 中逻辑库注册的数据源保持一致
- 恢复时，保证 ShardingSphere 在备份时和恢复时使用的版本一致，确保元数据兼容
- 恢复操作需要停机，并且为同步操作，用户需保证完全恢复成功
- 当恢复失败时，OpenGauss 数据节点存在状态不一致，需用户重新发起恢复操作，保证最终恢复成功
- 当执行备份后，会在当前用户的 `$HOME` 下创建 `.gs_pitr/backup` 目录，并在该目录下存放备份元数据文件
- 如果需要另一台设备上需要恢复，需要复制该路径下的备份数据到对应设备的相同路径
- 当执行删除备份后，当前用户的 `$HOME/.gs_pitr/backup` 下的备份文件将被删除