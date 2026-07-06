# Ignite alerts

[Ignite alerts by AppsCode](https://github.com/appscode/alerts) - Ignite alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/ignite-alerts --version=v2026.7.14
$ helm upgrade -i ignite appscode/ignite-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Ignite alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `ignite`:

```bash
$ helm upgrade -i ignite appscode/ignite-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Ignite alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `ignite`:

```bash
$ helm uninstall ignite -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `ignite-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Ignite</code>                              |
| metadata.resource.name                                                        |                                               | <code>ignites</code>                             |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>"ignite-quickstart"</code>                 |
| metadata.release.namespace                                                    | Release namespace                             | <code>"demo"</code>                              |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteDown.enabled                           |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.igniteDown.duration                          |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.igniteDown.severity                          |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.ignitePhaseCritical.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.ignitePhaseCritical.duration                 |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.ignitePhaseCritical.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteClusterNoBaselineNode.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.igniteClusterNoBaselineNode.duration         |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.igniteClusterNoBaselineNode.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteClusterNoBaselineNode.val              |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.igniteRestarted.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.igniteRestarted.duration                     |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.igniteRestarted.severity                     |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteRestarted.val                          |                                               | <code>180</code>                                 |
| form.alert.groups.database.rules.igniteHighCPULoad.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.igniteHighCPULoad.duration                   |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.igniteHighCPULoad.severity                   |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteHighCPULoad.val                        |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.igniteHighHeapMemoryUsed.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.igniteHighHeapMemoryUsed.duration            |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.igniteHighHeapMemoryUsed.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteHighHeapMemoryUsed.val                 |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.igniteHighDataregionOffHeapUsed.enabled      |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.igniteHighDataregionOffHeapUsed.duration     |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.igniteHighDataregionOffHeapUsed.severity     |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteHighDataregionOffHeapUsed.val          |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.igniteJVMPausesTotalDuration.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.igniteJVMPausesTotalDuration.duration        |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.igniteJVMPausesTotalDuration.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.igniteJVMPausesTotalDuration.val             |                                               | <code>0</code>                                   |
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
$ helm upgrade -i ignite appscode/ignite-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i ignite appscode/ignite-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
