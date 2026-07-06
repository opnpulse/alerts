# Hazelcast alerts

[Hazelcast alerts by AppsCode](https://github.com/appscode/alerts) - Hazelcast alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/hazelcast-alerts --version=v2026.7.14
$ helm upgrade -i hazelcast appscode/hazelcast-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Hazelcast alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `hazelcast`:

```bash
$ helm upgrade -i hazelcast appscode/hazelcast-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Hazelcast alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `hazelcast`:

```bash
$ helm uninstall hazelcast -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `hazelcast-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Hazelcast</code>                           |
| metadata.resource.name                                                        |                                               | <code>hazelcasts</code>                          |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>"hazelcast-sample"</code>                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>"default"</code>                           |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastPhaseCritical.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastPhaseCritical.duration              |                                               | <code>"3m"</code>                                |
| form.alert.groups.database.rules.hazelcastPhaseCritical.severity              |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastDown.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastDown.duration                       |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.hazelcastDown.severity                       |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.hazelcastPartitionCountExceed.enabled        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastPartitionCountExceed.duration       |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.hazelcastPartitionCountExceed.severity       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastPartitionCountExceed.val            |                                               | <code>92</code>                                  |
| form.alert.groups.database.rules.hazelcastHighHeapPercentage.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastHighHeapPercentage.duration         |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.hazelcastHighHeapPercentage.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastHighHeapPercentage.val              |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.hazelcastHighMemoryUsage.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastHighMemoryUsage.duration            |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.hazelcastHighMemoryUsage.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastHighMemoryUsage.val                 |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.hazelcastHighPhysicalMemoryUsage.enabled     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastHighPhysicalMemoryUsage.duration    |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.hazelcastHighPhysicalMemoryUsage.severity    |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastHighPhysicalMemoryUsage.val         |                                               | <code>50</code>                                  |
| form.alert.groups.database.rules.hazelcastHighLatency.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastHighLatency.duration                |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.hazelcastHighLatency.severity                |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastHighLatency.val                     |                                               | <code>2.5</code>                                 |
| form.alert.groups.database.rules.hazelcastSystemCPULoadExceed.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.hazelcastSystemCPULoadExceed.duration        |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.hazelcastSystemCPULoadExceed.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.hazelcastSystemCPULoadExceed.val             |                                               | <code>5</code>                                   |
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
| grafana.enabled                                                               |                                               | <code>true</code>                                |
| grafana.version                                                               |                                               | <code>7.5.5</code>                               |
| grafana.jobName                                                               |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                   |                                               | <code>""</code>                                  |
| grafana.apikey                                                                |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i hazelcast appscode/hazelcast-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i hazelcast appscode/hazelcast-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
