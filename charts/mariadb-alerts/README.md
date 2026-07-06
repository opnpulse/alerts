# MariaDB alerts

[MariaDB alerts by AppsCode](https://github.com/appscode/alerts) - MariaDB alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/mariadb-alerts --version=v2026.7.14
$ helm upgrade -i mariadb appscode/mariadb-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys MariaDB alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `mariadb`:

```bash
$ helm upgrade -i mariadb appscode/mariadb-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys MariaDB alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `mariadb`:

```bash
$ helm uninstall mariadb -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `mariadb-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>MariaDB</code>                             |
| metadata.resource.name                                                        |                                               | <code>mariadbs</code>                            |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mysqlInstanceDown.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlInstanceDown.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mysqlInstanceDown.severity                   |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mysqlServiceDown.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlServiceDown.duration                    |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mysqlServiceDown.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mysqlTooManyConnections.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlTooManyConnections.duration             |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mysqlTooManyConnections.val                  |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.mysqlTooManyConnections.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mysqlHighThreadsRunning.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlHighThreadsRunning.duration             |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mysqlHighThreadsRunning.val                  |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.mysqlHighThreadsRunning.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mysqlSlowQueries.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlSlowQueries.duration                    |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mysqlSlowQueries.severity                    |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mysqlInnoDBLogWaits.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlInnoDBLogWaits.duration                 |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mysqlInnoDBLogWaits.val                      |                                               | <code>10</code>                                  |
| form.alert.groups.database.rules.mysqlInnoDBLogWaits.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mysqlRestarted.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlRestarted.duration                      |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mysqlRestarted.val                           |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.mysqlRestarted.severity                      |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mysqlHighQPS.enabled                         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlHighQPS.duration                        |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mysqlHighQPS.val                             |                                               | <code>1000</code>                                |
| form.alert.groups.database.rules.mysqlHighQPS.severity                        |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mysqlHighIncomingBytes.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlHighIncomingBytes.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mysqlHighIncomingBytes.val                   |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.mysqlHighIncomingBytes.severity              |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mysqlHighOutgoingBytes.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlHighOutgoingBytes.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mysqlHighOutgoingBytes.val                   |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.mysqlHighOutgoingBytes.severity              |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mysqlTooManyOpenFiles.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mysqlTooManyOpenFiles.duration               |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mysqlTooManyOpenFiles.val                    |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.mysqlTooManyOpenFiles.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.diskUsageHigh.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.val                            |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.diskUsageHigh.duration                       |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.severity                       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.diskAlmostFull.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.val                           |                                               | <code>95</code>                                  |
| form.alert.groups.database.rules.diskAlmostFull.duration                      |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.severity                      |                                               | <code>critical</code>                            |
| form.alert.groups.cluster.enabled                                             |                                               | <code>warning</code>                             |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.enabled       |                                               | <code>true</code>                                |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.val           |                                               | <code>0.1</code>                                 |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.duration      |                                               | <code>"5m"</code>                                |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.severity      |                                               | <code>warning</code>                             |
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
| form.alert.groups.stash.enabled                                               |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.backupSessionFailed.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.duration                    |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.restoreSessionFailed.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.severity                   |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.val                   |                                               | <code>18000</code>                               |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.severity              |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.repositoryCorrupted.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.duration                    |                                               | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.duration            |                                               | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.val                 |                                               | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.duration             |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.val                  |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.duration            |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.val                 |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.severity            |                                               | <code>warning</code>                             |
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
| form.alert.groups.schemaManager.enabled                                       |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.duration        |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.enabled      |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.duration     |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.severity     |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.enabled     |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.duration    |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.severity    |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaFailed.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaFailed.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.schemaManager.rules.schemaFailed.severity                   |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaExpired.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaExpired.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.schemaManager.rules.schemaExpired.severity                  |                                               | <code>warning</code>                             |
| grafana.enabled                                                               |                                               | <code>false</code>                               |
| grafana.version                                                               |                                               | <code>7.5.5</code>                               |
| grafana.jobName                                                               |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                   |                                               | <code>""</code>                                  |
| grafana.apikey                                                                |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i mariadb appscode/mariadb-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i mariadb appscode/mariadb-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
