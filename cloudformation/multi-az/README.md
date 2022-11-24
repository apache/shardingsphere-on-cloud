The `cf.json` is used to quickly deploy a Apache ShardSphere Proxy Cluster.


## Develop

We use the [cfndsl]('https://github.com/cfndsl/cfndsl') to generate `CloudFormation` configuration. 

### Prerequisites
* `cfndsl`

***Please follow the steps provided by [cfndsl]('https://github.com/cfndsl/cfndsl') to install.***

### Step 1

Use the following command to init `cfndsl`. (only run once):
```shell
cfndsl -u 94.0.0
```

### Step 2

Use the following command to generate:
```shell
 cfndsl cf.rb -o cf.json --pretty
```