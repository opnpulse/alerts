# Milvus alerts

[Milvus alerts by AppsCode](https://github.com/appscode/alerts) - Milvus alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/milvus-alerts --version=v2026.7.14
$ helm upgrade -i milvus appscode/milvus-alerts -n kubedb --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Milvus alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `milvus`:

```bash
$ helm upgrade -i milvus appscode/milvus-alerts -n kubedb --create-namespace --version=v2026.7.14
```

The command deploys Milvus alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `milvus`:

```bash
$ helm uninstall milvus -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `milvus-alerts` chart and their default values.

|                               Parameter                                |                  Description                  |                Default                |
|------------------------------------------------------------------------|-----------------------------------------------|---------------------------------------|
| metadata.resource.group                                                |                                               | <code>kubedb.com</code>               |
| metadata.resource.kind                                                 |                                               | <code>Milvus</code>                   |
| metadata.resource.name                                                 |                                               | <code>milvuss</code>                  |
| metadata.resource.scope                                                |                                               | <code>Namespaced</code>               |
| metadata.resource.version                                              |                                               | <code>v1alpha2</code>                 |
| metadata.release.name                                                  | Release name                                  | <code>""</code>                       |
| metadata.release.namespace                                             | Release namespace                             | <code>""</code>                       |
| form.alert.enabled                                                     | # Enable PrometheusRule alerts                | <code>warning</code>                  |
| form.alert.labels                                                      | # Labels for default rules                    | <code>{"release":"prometheus"}</code> |
| form.alert.annotations                                                 | # Annotations for default rules               | <code>{}</code>                       |
| form.alert.additionalRuleLabels                                        | # Additional labels for PrometheusRule alerts | <code>{}</code>                       |
| form.alert.groups.database.enabled                                     |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusHighCPUUsage.enabled            |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusHighCPUUsage.duration           |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.milvusHighCPUUsage.severity           |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusHighCPUUsage.val                |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.milvusHighMemoryUsage.enabled         |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusHighMemoryUsage.duration        |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.milvusHighMemoryUsage.severity        |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusHighMemoryUsage.val             |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.enabled                 |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.diskUsageHigh.val                     |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.duration                |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.diskUsageHigh.severity                |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.diskAlmostFull.enabled                |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.diskAlmostFull.val                    |                                               | <code>95</code>                       |
| form.alert.groups.database.rules.diskAlmostFull.duration               |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.diskAlmostFull.severity               |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.milvusInstanceDown.enabled            |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusInstanceDown.duration           |                                               | <code>"30s"</code>                    |
| form.alert.groups.database.rules.milvusInstanceDown.severity           |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.milvusRestarted.enabled               |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusRestarted.duration              |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.milvusRestarted.severity              |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusRestarted.val                   |                                               | <code>180</code>                      |
| form.alert.groups.database.rules.milvusHighProcessMemoryUsage.enabled  |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusHighProcessMemoryUsage.duration |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.milvusHighProcessMemoryUsage.severity |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusHighProcessMemoryUsage.val      |                                               | <code>800000000</code>                |
| form.alert.groups.database.rules.milvusGoroutinesExplosion.enabled     |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusGoroutinesExplosion.duration    |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.milvusGoroutinesExplosion.severity    |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusGoroutinesExplosion.val         |                                               | <code>1000</code>                     |
| form.alert.groups.database.rules.milvusHighThreadPressure.enabled      |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusHighThreadPressure.duration     |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.milvusHighThreadPressure.severity     |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusHighThreadPressure.val          |                                               | <code>60</code>                       |
| form.alert.groups.database.rules.milvusHighFDsUsage.enabled            |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.milvusHighFDsUsage.duration           |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.milvusHighFDsUsage.severity           |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.milvusHighFDsUsage.val                |                                               | <code>110</code>                      |
| form.alert.groups.provisioner.enabled                                  |                                               | <code>warning</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled           |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration          |                                               | <code>"1m"</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity          |                                               | <code>critical</code>                 |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled           |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration          |                                               | <code>"15m"</code>                    |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity          |                                               | <code>warning</code>                  |
| grafana.enabled                                                        |                                               | <code>false</code>                    |
| grafana.version                                                        |                                               | <code>7.5.5</code>                    |
| grafana.jobName                                                        |                                               | <code>kubedb-databases</code>         |
| grafana.url                                                            |                                               | <code>""</code>                       |
| grafana.apikey                                                         |                                               | <code>""</code>                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i milvus appscode/milvus-alerts -n kubedb --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i milvus appscode/milvus-alerts -n kubedb --create-namespace --version=v2026.7.14 --values values.yaml
```
