# Druid alerts

[Druid alerts by AppsCode](https://github.com/appscode/alerts) - Druid alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/druid-alerts --version=v2026.7.14
$ helm upgrade -i druid-sample appscode/druid-alerts -n druid --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Druid alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `druid-sample`:

```bash
$ helm upgrade -i druid-sample appscode/druid-alerts -n druid --create-namespace --version=v2026.7.14
```

The command deploys Druid alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `druid-sample`:

```bash
$ helm uninstall druid-sample -n druid
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `druid-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                Default                |
|-------------------------------------------------------------------------------|-----------------------------------------------|---------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>               |
| metadata.resource.kind                                                        |                                               | <code>Druid</code>                    |
| metadata.resource.name                                                        |                                               | <code>druids</code>                   |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>               |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                 |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                       |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                       |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                  |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"prometheus"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                       |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                       |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.druidDown.enabled                            |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.druidDown.duration                           |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.druidDown.severity                           |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.zkDisconnected.enabled                       |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.zkDisconnected.duration                      |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.zkDisconnected.severity                      |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.highQueryTime.enabled                        |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.highQueryTime.duration                       |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.highQueryTime.severity                       |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.highQueryWaitTime.enabled                    |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.highQueryWaitTime.duration                   |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.highQueryWaitTime.severity                   |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.highSegmentScanPending.enabled               |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.highSegmentScanPending.duration              |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.highSegmentScanPending.severity              |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.highSegmentScanPending.val                   |                                               | <code>2</code>                        |
| form.alert.groups.database.rules.highSegmentUsage.enabled                     |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.highSegmentUsage.duration                    |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.highSegmentUsage.severity                    |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.highSegmentUsage.val                         |                                               | <code>95</code>                       |
| form.alert.groups.database.rules.highJVMPoolUsage.enabled                     |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.highJVMPoolUsage.duration                    |                                               | <code>"30s"</code>                    |
| form.alert.groups.database.rules.highJVMPoolUsage.severity                    |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.highJVMPoolUsage.val                         |                                               | <code>95</code>                       |
| form.alert.groups.database.rules.highJVMMemoryUsage.enabled                   |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.highJVMMemoryUsage.duration                  |                                               | <code>"30s"</code>                    |
| form.alert.groups.database.rules.highJVMMemoryUsage.severity                  |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.highJVMMemoryUsage.val                       |                                               | <code>95</code>                       |
| form.alert.groups.provisioner.enabled                                         |                                               | <code>warning</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                               | <code>"1m"</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                               | <code>critical</code>                 |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"15m"</code>                    |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                               | <code>warning</code>                  |
| form.alert.groups.opsManager.enabled                                          |                                               | <code>warning</code>                  |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                               | <code>true</code>                     |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                               | <code>"0m"</code>                     |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                               | <code>info</code>                     |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                               | <code>true</code>                     |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                               | <code>"30m"</code>                    |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                               | <code>critical</code>                 |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                               | <code>true</code>                     |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                               | <code>"0m"</code>                     |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                               | <code>critical</code>                 |
| grafana.enabled                                                               |                                               | <code>false</code>                    |
| grafana.version                                                               |                                               | <code>7.5.5</code>                    |
| grafana.jobName                                                               |                                               | <code>kubedb-databases</code>         |
| grafana.url                                                                   |                                               | <code>""</code>                       |
| grafana.apikey                                                                |                                               | <code>""</code>                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i druid-sample appscode/druid-alerts -n druid --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i druid-sample appscode/druid-alerts -n druid --create-namespace --version=v2026.7.14 --values values.yaml
```
