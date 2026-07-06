# HanaDB alerts

[HanaDB alerts by AppsCode](https://github.com/appscode/alerts) - HanaDB alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/hanadb-alerts --version=v2026.7.14
$ helm upgrade -i hana-cluster appscode/hanadb-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys HanaDB alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `hana-cluster`:

```bash
$ helm upgrade -i hana-cluster appscode/hanadb-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys HanaDB alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `hana-cluster`:

```bash
$ helm uninstall hana-cluster -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `hanadb-alerts` chart and their default values.

|                              Parameter                               | Description |              Default               |
|----------------------------------------------------------------------|-------------|------------------------------------|
| metadata.resource.group                                              |             | <code>kubedb.com</code>            |
| metadata.resource.kind                                               |             | <code>HanaDB</code>                |
| metadata.resource.name                                               |             | <code>hanadbs</code>               |
| metadata.resource.scope                                              |             | <code>Namespaced</code>            |
| metadata.resource.version                                            |             | <code>v1alpha2</code>              |
| metadata.release.name                                                |             | <code>"hana-cluster"</code>        |
| metadata.release.namespace                                           |             | <code>"demo"</code>                |
| form.alert.enabled                                                   |             | <code>warning</code>               |
| form.alert.labels.release                                            |             | <code>kube-prometheus-stack</code> |
| form.alert.annotations                                               |             | <code>{}</code>                    |
| form.alert.additionalRuleLabels                                      |             | <code>{}</code>                    |
| form.alert.groups.database.enabled                                   |             | <code>warning</code>               |
| form.alert.groups.database.rules.hanadbInstanceDown.enabled          |             | <code>true</code>                  |
| form.alert.groups.database.rules.hanadbInstanceDown.duration         |             | <code>"0m"</code>                  |
| form.alert.groups.database.rules.hanadbInstanceDown.severity         |             | <code>critical</code>              |
| form.alert.groups.database.rules.hanadbServiceDown.enabled           |             | <code>true</code>                  |
| form.alert.groups.database.rules.hanadbServiceDown.duration          |             | <code>"0m"</code>                  |
| form.alert.groups.database.rules.hanadbServiceDown.severity          |             | <code>critical</code>              |
| form.alert.groups.database.rules.hanadbRestarted.enabled             |             | <code>true</code>                  |
| form.alert.groups.database.rules.hanadbRestarted.duration            |             | <code>"1m"</code>                  |
| form.alert.groups.database.rules.hanadbRestarted.severity            |             | <code>warning</code>               |
| form.alert.groups.database.rules.hanadbRestarted.val                 |             | <code>180</code>                   |
| form.alert.groups.database.rules.hanadbHighCPUUsage.enabled          |             | <code>true</code>                  |
| form.alert.groups.database.rules.hanadbHighCPUUsage.duration         |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.hanadbHighCPUUsage.severity         |             | <code>warning</code>               |
| form.alert.groups.database.rules.hanadbHighCPUUsage.val              |             | <code>80</code>                    |
| form.alert.groups.database.rules.hanadbHighMemoryUsage.enabled       |             | <code>true</code>                  |
| form.alert.groups.database.rules.hanadbHighMemoryUsage.duration      |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.hanadbHighMemoryUsage.severity      |             | <code>warning</code>               |
| form.alert.groups.database.rules.hanadbHighMemoryUsage.val           |             | <code>80</code>                    |
| form.alert.groups.database.rules.hanadbReplicationNotActive.enabled  |             | <code>true</code>                  |
| form.alert.groups.database.rules.hanadbReplicationNotActive.duration |             | <code>"2m"</code>                  |
| form.alert.groups.database.rules.hanadbReplicationNotActive.severity |             | <code>critical</code>              |
| form.alert.groups.database.rules.diskUsageHigh.enabled               |             | <code>true</code>                  |
| form.alert.groups.database.rules.diskUsageHigh.val                   |             | <code>80</code>                    |
| form.alert.groups.database.rules.diskUsageHigh.duration              |             | <code>"1m"</code>                  |
| form.alert.groups.database.rules.diskUsageHigh.severity              |             | <code>warning</code>               |
| form.alert.groups.database.rules.diskAlmostFull.enabled              |             | <code>true</code>                  |
| form.alert.groups.database.rules.diskAlmostFull.val                  |             | <code>95</code>                    |
| form.alert.groups.database.rules.diskAlmostFull.duration             |             | <code>"1m"</code>                  |
| form.alert.groups.database.rules.diskAlmostFull.severity             |             | <code>critical</code>              |
| form.alert.groups.provisioner.enabled                                |             | <code>warning</code>               |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled         |             | <code>true</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration        |             | <code>"1m"</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity        |             | <code>critical</code>              |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled         |             | <code>true</code>                  |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration        |             | <code>"15m"</code>                 |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity        |             | <code>warning</code>               |
| grafana.enabled                                                      |             | <code>false</code>                 |
| grafana.version                                                      |             | <code>7.5.5</code>                 |
| grafana.jobName                                                      |             | <code>kubedb-databases</code>      |
| grafana.url                                                          |             | <code>""</code>                    |
| grafana.apikey                                                       |             | <code>""</code>                    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i hana-cluster appscode/hanadb-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i hana-cluster appscode/hanadb-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
