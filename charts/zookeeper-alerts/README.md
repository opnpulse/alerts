# ZooKeeper alerts

[ZooKeeper alerts by AppsCode](https://github.com/appscode/alerts) - ZooKeeper alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/zookeeper-alerts --version=v2026.7.14
$ helm upgrade -i zookeeper appscode/zookeeper-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys ZooKeeper alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `zookeeper`:

```bash
$ helm upgrade -i zookeeper appscode/zookeeper-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys ZooKeeper alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `zookeeper`:

```bash
$ helm uninstall zookeeper -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `zookeeper-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>ZooKeeper</code>                           |
| metadata.resource.name                                                        |                                               | <code>zookeepers</code>                          |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperDown.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperDown.duration                       |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperDown.severity                       |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.zookeeperTooManyNodes.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyNodes.duration               |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyNodes.val                    |                                               | <code>1000000</code>                             |
| form.alert.groups.database.rules.zookeeperTooManyNodes.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperTooBigMemory.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooBigMemory.duration               |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooBigMemory.val                    |                                               | <code>1</code>                                   |
| form.alert.groups.database.rules.zookeeperTooBigMemory.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperTooManyWatch.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyWatch.duration               |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyWatch.val                    |                                               | <code>10000</code>                               |
| form.alert.groups.database.rules.zookeeperTooManyWatch.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperTooManyConnections.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyConnections.duration         |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyConnections.val              |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.zookeeperTooManyConnections.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperLeaderElection.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperLeaderElection.duration             |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperLeaderElection.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperTooManyOpenFiles.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyOpenFiles.duration           |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooManyOpenFiles.val                |                                               | <code>300</code>                                 |
| form.alert.groups.database.rules.zookeeperTooManyOpenFiles.severity           |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperTooLongFsyncTime.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooLongFsyncTime.duration           |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooLongFsyncTime.val                |                                               | <code>100</code>                                 |
| form.alert.groups.database.rules.zookeeperTooLongFsyncTime.severity           |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperTooLongSnapshotTime.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooLongSnapshotTime.duration        |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooLongSnapshotTime.val             |                                               | <code>100</code>                                 |
| form.alert.groups.database.rules.zookeeperTooLongSnapshotTime.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperTooHighAvgLatency.enabled           |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperTooHighAvgLatency.duration          |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperTooHighAvgLatency.val               |                                               | <code>100</code>                                 |
| form.alert.groups.database.rules.zookeeperTooHighAvgLatency.severity          |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.zookeeperJvmMemoryFilingUp.enabled           |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.zookeeperJvmMemoryFilingUp.duration          |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.zookeeperJvmMemoryFilingUp.val               |                                               | <code>0.8</code>                                 |
| form.alert.groups.database.rules.zookeeperJvmMemoryFilingUp.severity          |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.diskUsageHigh.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.val                            |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.diskUsageHigh.duration                       |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.severity                       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.diskAlmostFull.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.val                           |                                               | <code>95</code>                                  |
| form.alert.groups.database.rules.diskAlmostFull.duration                      |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.severity                      |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.enabled                                         |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                               | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.enabled                                          |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                               | <code>info</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                               | <code>"30m"</code>                               |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                               | <code>critical</code>                            |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                               | <code>critical</code>                            |
| grafana.enabled                                                               |                                               | <code>false</code>                               |
| grafana.version                                                               |                                               | <code>7.5.5</code>                               |
| grafana.jobName                                                               |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                   |                                               | <code>""</code>                                  |
| grafana.apikey                                                                |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i zookeeper appscode/zookeeper-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i zookeeper appscode/zookeeper-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
