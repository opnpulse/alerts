# Qdrant alerts

[Qdrant alerts by AppsCode](https://github.com/appscode/alerts) - Qdrant alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/qdrant-alerts --version=v2026.7.14
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Qdrant alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `qdrant`:

```bash
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Qdrant alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `qdrant`:

```bash
$ helm uninstall qdrant -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `qdrant-alerts` chart and their default values.

|                               Parameter                               |                  Description                  |                Default                |
|-----------------------------------------------------------------------|-----------------------------------------------|---------------------------------------|
| metadata.resource.group                                               |                                               | <code>kubedb.com</code>               |
| metadata.resource.kind                                                |                                               | <code>Qdrant</code>                   |
| metadata.resource.name                                                |                                               | <code>qdrants</code>                  |
| metadata.resource.scope                                               |                                               | <code>Namespaced</code>               |
| metadata.resource.version                                             |                                               | <code>v1alpha2</code>                 |
| metadata.release.name                                                 | Release name                                  | <code>""</code>                       |
| metadata.release.namespace                                            | Release namespace                             | <code>""</code>                       |
| form.alert.enabled                                                    | # Enable PrometheusRule alerts                | <code>warning</code>                  |
| form.alert.labels                                                     | # Labels for default rules                    | <code>{"release":"prometheus"}</code> |
| form.alert.annotations                                                | # Annotations for default rules               | <code>{}</code>                       |
| form.alert.additionalRuleLabels                                       | # Additional labels for PrometheusRule alerts | <code>{}</code>                       |
| form.alert.groups.database.enabled                                    |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.qdrantInstanceDown.enabled           |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantInstanceDown.duration          |                                               | <code>"30s"</code>                    |
| form.alert.groups.database.rules.qdrantInstanceDown.severity          |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.qdrantPhaseCritical.enabled          |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantPhaseCritical.duration         |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.qdrantPhaseCritical.severity         |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.qdrantRestarted.enabled              |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantRestarted.duration             |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.qdrantRestarted.severity             |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.qdrantRestarted.val                  |                                               | <code>180</code>                      |
| form.alert.groups.database.rules.qdrantHighCPUUsage.enabled           |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantHighCPUUsage.duration          |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.qdrantHighCPUUsage.severity          |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.qdrantHighCPUUsage.val               |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.qdrantHighMemoryUsage.enabled        |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantHighMemoryUsage.duration       |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.qdrantHighMemoryUsage.severity       |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.qdrantHighMemoryUsage.val            |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.enabled                |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.diskUsageHigh.val                    |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.duration               |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.diskUsageHigh.severity               |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.diskAlmostFull.enabled               |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.diskAlmostFull.val                   |                                               | <code>95</code>                       |
| form.alert.groups.database.rules.diskAlmostFull.duration              |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.diskAlmostFull.severity              |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.qdrantHighPendingOperations.enabled  |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantHighPendingOperations.val      |                                               | <code>10</code>                       |
| form.alert.groups.database.rules.qdrantHighPendingOperations.duration |                                               | <code>"5m"</code>                     |
| form.alert.groups.database.rules.qdrantHighPendingOperations.severity |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.qdrantGrpcResponsesFailHigh.enabled  |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantGrpcResponsesFailHigh.val      |                                               | <code>5</code>                        |
| form.alert.groups.database.rules.qdrantGrpcResponsesFailHigh.duration |                                               | <code>"5m"</code>                     |
| form.alert.groups.database.rules.qdrantGrpcResponsesFailHigh.severity |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.qdrantRestResponsesFailHigh.enabled  |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.qdrantRestResponsesFailHigh.val      |                                               | <code>5</code>                        |
| form.alert.groups.database.rules.qdrantRestResponsesFailHigh.duration |                                               | <code>"5m"</code>                     |
| form.alert.groups.database.rules.qdrantRestResponsesFailHigh.severity |                                               | <code>critical</code>                 |
| form.alert.groups.provisioner.enabled                                 |                                               | <code>warning</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled          |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration         |                                               | <code>"1m"</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity         |                                               | <code>critical</code>                 |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled          |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration         |                                               | <code>"15m"</code>                    |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity         |                                               | <code>warning</code>                  |
| grafana.enabled                                                       |                                               | <code>false</code>                    |
| grafana.version                                                       |                                               | <code>7.5.5</code>                    |
| grafana.jobName                                                       |                                               | <code>kubedb-databases</code>         |
| grafana.url                                                           |                                               | <code>""</code>                       |
| grafana.apikey                                                        |                                               | <code>""</code>                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
