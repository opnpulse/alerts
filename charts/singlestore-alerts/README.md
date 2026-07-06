# Singlestore alerts

[Singlestore alerts by AppsCode](https://github.com/appscode/alerts) - Singlestore alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/singlestore-alerts --version=v2026.7.14
$ helm upgrade -i singlestore appscode/singlestore-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Singlestore alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `singlestore`:

```bash
$ helm upgrade -i singlestore appscode/singlestore-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Singlestore alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `singlestore`:

```bash
$ helm uninstall singlestore -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `singlestore-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Singlestore</code>                         |
| metadata.resource.name                                                        |                                               | <code>singlestores</code>                        |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.singlestoreInstanceDown.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreInstanceDown.duration             |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.singlestoreInstanceDown.severity             |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.singlestoreServiceDown.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreServiceDown.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.singlestoreServiceDown.severity              |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.singlestoreTooManyConnections.enabled        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreTooManyConnections.duration       |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.singlestoreTooManyConnections.val            |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.singlestoreTooManyConnections.severity       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.singlestoreHighThreadsRunning.enabled        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreHighThreadsRunning.duration       |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.singlestoreHighThreadsRunning.val            |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.singlestoreHighThreadsRunning.severity       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.singlestoreRestarted.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreRestarted.duration                |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.singlestoreRestarted.val                     |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.singlestoreRestarted.severity                |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.singlestoreHighQPS.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreHighQPS.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.singlestoreHighQPS.val                       |                                               | <code>1000</code>                                |
| form.alert.groups.database.rules.singlestoreHighQPS.severity                  |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.singlestoreHighIncomingBytes.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreHighIncomingBytes.duration        |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.singlestoreHighIncomingBytes.val             |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.singlestoreHighIncomingBytes.severity        |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.singlestoreHighOutgoingBytes.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.singlestoreHighOutgoingBytes.duration        |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.singlestoreHighOutgoingBytes.val             |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.singlestoreHighOutgoingBytes.severity        |                                               | <code>critical</code>                            |
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
| form.alert.groups.kubeStash.enabled                                           |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.backupSessionFailed.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionFailed.duration                |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionFailed.severity                |                                               | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.duration               |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.severity               |                                               | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.enabled           |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.duration          |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.val               |                                               | <code>18000</code>                               |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.severity          |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.duration                |                                               | <code>"5m"</code>                                |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.severity                |                                               | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.duration        |                                               | <code>"5m"</code>                                |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.val             |                                               | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.duration         |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.val              |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.duration        |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.val             |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.severity        |                                               | <code>warning</code>                             |
| grafana.enabled                                                               |                                               | <code>false</code>                               |
| grafana.version                                                               |                                               | <code>7.5.5</code>                               |
| grafana.jobName                                                               |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                   |                                               | <code>""</code>                                  |
| grafana.apikey                                                                |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i singlestore appscode/singlestore-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i singlestore appscode/singlestore-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
