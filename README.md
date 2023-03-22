# ShardingSphere on Cloud

[![GitHub release](https://img.shields.io/github/release/apache/shardingsphere-on-cloud.svg)](https://github.com/apache/shardingsphere-on-cloud/releases)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/ShardingSphere.svg?style=social&label=Follow%20%40ShardingSphere)](https://twitter.com/ShardingSphere)
[![Slack](https://img.shields.io/badge/%20Slack-ShardingSphere%20Channel-blueviolet)](https://join.slack.com/t/apacheshardingsphere/shared_invite/zt-sbdde7ie-SjDqo9~I4rYcR18bq0SYTg)
[![Gitter](https://badges.gitter.im/shardingsphere/shardingsphere.svg)](https://gitter.im/shardingsphere/Lobby)

This repository collects scripts, tools, manifests and documentations, and provides a home for [Apache ShardingSphere](https://shardingsphere.apache.org/) on cloud solutions.


## 🚀🚀  Spinning up a ShardingSphere cluster on Kubernetes in one minutes !

### A Demo of Starting `ShardingSphere Proxy` Using Helm Charts

 <p align="center">
    <a href="https://asciinema.org/a/569048" target="_blank"><img src="https://asciinema.org/a/569048.svg" /></a>
 <p>

### A Demo of Starting `ShardingSphere Operator` Using Helm Charts

 <p align="center">
    <a href="https://asciinema.org/a/569049" target="_blank"><img src="https://asciinema.org/a/569049.svg" /></a>
 <p>

## Current Status

The solutions currently included in this project are:

* Out-of-Box Deployment: 
    * [The ShardingSphere Helm Charts](https://github.com/apache/shardingsphere-on-cloud/tree/main/charts/apache-shardingsphere-proxy-charts): Deploy a ShardingSphere cluster on any Kubernetes distro.
    * [AWS CloudFormation Stack Template for ShardingSphere](https://github.com/apache/shardingsphere-on-cloud/tree/main/cloudformation): Deploy a ShardingSphere cluster on AWS resources with predefined CloudFormation.
    * [Terraform Configuration For ShardingSphere](https://github.com/apache/shardingsphere-on-cloud/tree/main/terraform): Using Terraform to implement IaC of ShardingSphere on any clouds.

* Database Reliability Engineering inspired Operator:
    * [The ShardingSphere Operator](https://github.com/apache/shardingsphere-on-cloud/tree/main/shardingsphere-operator): Deploy a ShardingSphere Operator on any Kubernetes distro, which provides advanced operation features.
    * [Grafana template For ShardingSphere](https://github.com/apache/shardingsphere-on-cloud/tree/main/grafana)

* Other Ecosystem Experiments:
    * [WebAssembly extension demo For ShardingSphere](https://github.com/apache/shardingsphere-on-cloud/tree/main/wasm)
    * [Point-in-Time-Recovery demo For ShardingSphere](https://github.com/apache/shardingsphere-on-cloud/tree/main/pitr)

## Get Started

**Apache ShardingSphere Official Website:** [https://shardingsphere.apache.org/](https://shardingsphere.apache.org/)

**Apache ShardingSphere-on-Cloud Official Website:** [https://shardingsphere.apache.org/oncloud/](https://shardingsphere.apache.org/oncloud/)

## Contributing

To contribute to this project, refer to [Contributing](CONTRIBUTING.md).

## Community & Support

Thank you for contributing to the ShardingSphere on Cloud project!

![Contributors](https://contrib.rocks/image?repo=apache/shardingsphere-on-cloud)

<hr>

:link: [ShardingSphere on Cloud Docs](https://shardingsphere.apache.org/oncloud/current/en/overview/). Best for: Manuals and best practices.

:link: [Mailing List](https://shardingsphere.apache.org/community/en/contribute/subscribe/). Best for: Apache community updates, releases, changes.

:link: [GitHub Issues](https://github.com/apache/shardingsphere-on-cloud/issues). Best for: larger systemic questions/bug reports or anything development related.

:link: [GitHub Discussions](https://github.com/apache/shardingsphere-on-cloud/discussions). Best for: technical questions & support, requesting new features, proposing new features.

:link: [Slack channel](https://join.slack.com/t/apacheshardingsphere/shared_invite/zt-sbdde7ie-SjDqo9~I4rYcR18bq0SYTg). Best for: instant communications and online meetings, sharing your applications.

:link: [Twitter](https://twitter.com/ShardingSphere). Best for: keeping up to date on everything ShardingSphere.


## License

Apache License 2.0, see [LICENSE](https://github.com/apache/shardingsphere-on-cloud/blob/main/LICENSE).
