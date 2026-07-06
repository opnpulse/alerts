# Oracle alerts

[Oracle alerts by AppsCode](https://github.com/appscode/alerts) - Oracle alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/oracle-alerts --version=v2026.7.14
$ helm upgrade -i oracle appscode/oracle-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Oracle alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `oracle`:

```bash
$ helm upgrade -i oracle appscode/oracle-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Oracle alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `oracle`:

```bash
$ helm uninstall oracle -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `oracle-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                          Default                           |
|-------------------------------------------------------------------------------|-----------------------------------------------|------------------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                                    |
| metadata.resource.kind                                                        |                                               | <code>Oracle</code>                                        |
| metadata.resource.name                                                        |                                               | <code>oracles</code>                                       |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                                    |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                                      |
| metadata.release.name                                                         | Release name                                  | <code>"oracle"</code>                                      |
| metadata.release.namespace                                                    | Release namespace                             | <code>"exporter"</code>                                    |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                                       |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"prometheus"}</code>                      |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                            |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                            |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                                       |
| form.alert.groups.database.rules.oracleInstanceDown.enabled                   |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.oracleInstanceDown.duration                  |                                               | <code>"0m"</code>                                          |
| form.alert.groups.database.rules.oracleInstanceDown.severity                  |                                               | <code>critical</code>                                      |
| form.alert.groups.database.rules.oracleServiceDown.enabled                    |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.oracleServiceDown.duration                   |                                               | <code>"0m"</code>                                          |
| form.alert.groups.database.rules.oracleServiceDown.severity                   |                                               | <code>critical</code>                                      |
| form.alert.groups.database.rules.oracleRestarted.enabled                      |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.oracleRestarted.duration                     |                                               | <code>"0m"</code>                                          |
| form.alert.groups.database.rules.oracleRestarted.val                          |                                               | <code>60</code>                                            |
| form.alert.groups.database.rules.oracleRestarted.severity                     |                                               | <code>critical</code>                                      |
| form.alert.groups.database.rules.diskUsageHigh.enabled                        |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.diskUsageHigh.val                            |                                               | <code>80</code>                                            |
| form.alert.groups.database.rules.diskUsageHigh.duration                       |                                               | <code>"1m"</code>                                          |
| form.alert.groups.database.rules.diskUsageHigh.severity                       |                                               | <code>warning</code>                                       |
| form.alert.groups.database.rules.diskAlmostFull.enabled                       |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.diskAlmostFull.val                           |                                               | <code>95</code>                                            |
| form.alert.groups.database.rules.diskAlmostFull.duration                      |                                               | <code>"1m"</code>                                          |
| form.alert.groups.database.rules.diskAlmostFull.severity                      |                                               | <code>critical</code>                                      |
| form.alert.groups.provisioner.enabled                                         |                                               | <code>warning</code>                                       |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                               | <code>true</code>                                          |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                               | <code>"1m"</code>                                          |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                               | <code>critical</code>                                      |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                               | <code>true</code>                                          |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"15m"</code>                                         |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                               | <code>warning</code>                                       |
| form.alert.groups.opsManager.enabled                                          |                                               | <code>warning</code>                                       |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                               | <code>true</code>                                          |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                               | <code>"0m"</code>                                          |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                               | <code>info</code>                                          |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                               | <code>true</code>                                          |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                               | <code>"30m"</code>                                         |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                               | <code>critical</code>                                      |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                               | <code>true</code>                                          |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                               | <code>"0m"</code>                                          |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                               | <code>critical</code>                                      |
| grafana.enabled                                                               |                                               | <code>true</code>                                          |
| grafana.version                                                               |                                               | <code>7.5.5</code>                                         |
| grafana.jobName                                                               |                                               | <code>oracle-stats</code>                                  |
| grafana.url                                                                   |                                               | <code>"http://prometheus-grafana.monitoring.svc:80"</code> |
| grafana.apikey                                                                |                                               | <code>""</code>                                            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i oracle appscode/oracle-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i oracle appscode/oracle-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
