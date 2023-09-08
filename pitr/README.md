# README

This is a cli tool for point-in-time recovery of Apache ShardingSphere and OpenGauss distributed database cluster.

## Prerequisition

### Servers

| | Role | Components |
|:-:|:-:|:-:|
|1| Pitr cli operation server | Pitr Cli + ShardingSphere Proxy + Zookeeper|
|2| OpenGauss Server 1 | OpenGauss Server + Pitr Agent |
|3| OpenGauss Server 2 | OpenGauss Server + Pitr Agent |

### Environment

- Apache ShardingSphere Proxy can access OpenGauss network
- External access to Apache ShardingSphere Proxy
- External access to OpenGauss Server via port 18080
- OpenGauss has user `omm` and database `omm` which can be accessed
- OpenGauss enables `cbm tracking`
- SSL key pairs for Pitr cli-agent secure communication

#### Compilation (optional)

If you want to compile Pitr tools yourself, you should using a recommanded Golang version 1.20 with Linux 3.10.0-957.el7.x86_64. Following the steps below to compile both Pitr agent and cli.

Step 1. Firstly clone the project

```shell
git clone git@github.com:apache/shardingsphere-on-cloud.git
```

Step 2. Compile Pitr agent

```shell
cd shardingsphere-on-cloud/pitr/agent
make build
```

Step 3. Compile Pitr cli

```shell
cd shardingsphere-on-cloud/pitr/cli
make build
```

### SSL Configurations

The communication of Pitr cli and Pitr agent is secured by TLS which needs a SSL key pair. You can either use any available keypairs or generate a new keypair, e.g.:

- tls.key
- tls.crt

#### Generate new TLS keypair (Optional)

Make sure you have a usable OpenSSL environment, check environment variable OPENSSL_CONF, generally it is set to `/etc/pki/tls`.

Then using the script under Pitr agent code directory, execute the commands below:

```shell
git clone git@github.com:apache/shardingsphere-on-cloud.git
cd shardingsphere-on-cloud/pitr/agent
make openssl-local
```

After that, the keypair files will be write to `./certs` in the current directory.


## Deployment

Pitr cli (aka `gs_pitr`) and Pitr agent (aka `pitr-agent`) binaries could be downloaded at Apache ShardingSphere on Cloud release page, or just compiled in your local development environment.

### Step 1: Get ShardingSphere Proxy Configurations

Substitute the OpenGauss server address below:

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

globalClock:
  enabled: true
  type: TSO
  provider: redis
  props:
    host: 127.0.0.1
    port: 6379

transaction:
  defaultType: XA
  providerType: Atomikos

props:
  proxy-frontend-database-protocol-type: openGauss

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

### Step 2: Get OpenGauss Configurations

Enable `cbm tracking` in postgres.conf

```shell
enable_cbm_tracking = on
```

### Step 3: Get Pitr tools

You can download pre-compiled Pitr tools binary release or compile them yourself from source code.

#### Get binary release

The binaries are packaged as .tar.gz file on [release page](https://github.com/apache/shardingsphere-on-cloud/releases). You can download expected version and uncompress the binary files `gs_pitr` and `pitr-agent`.

#### Compile it yourself

Please refer to the `Compilation` section in `Prerequsition` for detailed instructions.

### Step 4: Deploy Pitr Agent

1. Copy cert files

If the TLS keypair is compiled yourself, the cert files are located at `shardingsphere-on-cloud/pitr/agent/certs`. You should change directory to the cert directory before executing the command below:

```shell
scp tls.crt tls.key root@${OPENGAUSS_SERVER_1}:/home/omm/
scp tls.crt tls.key root@${OPENGAUSS_SERVER_2}:/home/omm/
```

2. Copy binary files

```shell
cd shardingsphere-on-cloud/pitr/agent

scp pitr-agent root@${OPENGAUSS_SERVER_1}:/home/omm/
scp pitr-agent root@${OPENGAUSS_SERVER_2}:/home/omm/
```

### Step 5: Start Pitr Agent

1. Login OpenGauss servers and change directory to `/home/omm`

Here are files under `/home/omm`:

```shell
$ ll
total 13M
drwx------  4 omm  omm    32 Mar  2 14:22 data
drwx------ 29 omm  omm  4.0K May 23 11:37 pgdata
-rwxr-xr-x  1 root root  13M May 16 18:25 pitr-agent
-rwxr-xr-x  1 root root 1.1K May 16 18:26 tls.crt
-rwxr-xr-x  1 root root 1.7K May 16 18:26 tls.key
```

2. Start Pitr agent

```shell
./pitr-agent -pgdata /data/data-glt/d1 -port 18080 -tls-crt tls.crt -tls-key tls.key -log-level debug
```

Parameters:
- pgdata: OpenGauss data storage path
- port: Pitr agent exposed port
- tls-crt: TLS crt file path
- tls-key: TLS key file path
- log-level: Pitr agent log level

## Test

### Prepare Test Data

You can connect to ShardingSphere Proxy with `gsql` and generate some data for testing.
```shell
gsql -h127.0.0.1 -p3307 -Usharding -Wsharding -d sharding_db
```
1. Check Storage Units

```SQL
SHOW STORAGE UNITS
```

2. Create Sharding Table Rule `t_user`

```SQL
CREATE SHARDING TABLE RULE t_user(
 STORAGE_UNITS(ds_0,ds_1),
 SHARDING_COLUMN=user_id,
 TYPE(NAME="hash_mod",PROPERTIES("sharding-count"="4"))
);
```

3. Check Sharding Table Rule

```SQL
SHOW SHARDING TABLE RULE t_user;
```

4. Create Table `t_user`
```SQL
CREATE TABLE t_user (
  user_id INT NOT NULL,
  order_id INT NOT NULL,
  status VARCHAR(45) NULL,
  PRIMARY KEY (user_id)
);
```

5. Check Sharding Table Nodes:
```SQL
SHOW SHARDING TABLE NODES;
```

6. Insert test data
```SQL
insert into t_user( user_id, order_id, status) values(1,1,1);
insert into t_user( user_id, order_id, status) values(2,2,2);
insert into t_user( user_id, order_id, status) values(3,3,3);
insert into t_user( user_id, order_id, status) values(4,4,4);

select * from t_user;
```

### Test Case

#### Backup

```Shell
./gs_pitr backup --host ${OPENGAUSS_SERVER_1} --password sharding --port 3307 --username sharding --agent-port 18080 --dn-threads-num 1 --dn-backup-path "/home/omm/data" -b FULL
```

Parameters:
- host: ShardingSphere Proxy server
- port: ShardingSphere Proxy port
- username: ShardingSphere Proxy user
- password: ShardingSphere Proxy password
- agent-port: Pitr agent port
- dn-threads-num: OpenGauss concurrent backup
- dn-threads-path: OpenGauss backup files path
- b: Backup mode

Check backups and get the backup id for recovery.

```Shell
./gs_pitr show
```

#### Recovery

You may need to delete some records of `t_user` first.
```SQL
delete from t_user where user_id=1;
delete from t_user where user_id=2;
```

Do recovery:
```Shell
./gs_pitr restore --host ${OPENGAUSS_SERVER_1} --password sharding --port 3307 --username sharding --agent-port 18080 --dn-backup-path "/home/omm/data" --id ${BACKUP_ID}
```

Parameters:
- host: ShardingSphere Proxy server
- port: ShardingSphere Proxy port
- username: ShardingSphere Proxy user
- password: ShardingSphere Proxy password
- agent-port: Pitr agent port
- dn-backup-path: OpenGauss backup files path
- id: Backup id

Verify data:
```SQL
select * from t_user;
```
