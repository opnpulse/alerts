# Elasticsearch alerts

[Elasticsearch alerts by AppsCode](https://github.com/appscode/alerts) - Elasticsearch alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/elasticsearch-alerts --version=v2026.7.14
$ helm upgrade -i elasticsearch appscode/elasticsearch-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Elasticsearch alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `elasticsearch`:

```bash
$ helm upgrade -i elasticsearch appscode/elasticsearch-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Elasticsearch alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `elasticsearch`:

```bash
$ helm uninstall elasticsearch -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `elasticsearch-alerts` chart and their default values.

|                                   Parameter                                   |                   Description                   |                     Default                      |
|-------------------------------------------------------------------------------|-------------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                                 | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                                 | <code>Elasticsearch</code>                       |
| metadata.resource.name                                                        |                                                 | <code>elasticsearches</code>                     |
| metadata.resource.scope                                                       |                                                 | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                                 | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                    | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                               | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                  | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                      | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules                 | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts   | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                                 | <code>critical</code>                            |
| form.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.enabled        |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.val            | The heap usage is over 90%                      | <code>90</code>                                  |
| form.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.duration       |                                                 | <code>"2m"</code>                                |
| form.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.severity       |                                                 | <code>critical</code>                            |
| form.alert.groups.database.rules.elasticsearchHeapUsageWarning.enabled        |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchHeapUsageWarning.val            | The heap usage is over 80%                      | <code>80</code>                                  |
| form.alert.groups.database.rules.elasticsearchHeapUsageWarning.duration       |                                                 | <code>"2m"</code>                                |
| form.alert.groups.database.rules.elasticsearchHeapUsageWarning.severity       |                                                 | <code>warning</code>                             |
| form.alert.groups.database.rules.elasticsearchDiskOutOfSpace.enabled          |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchDiskOutOfSpace.val              | The disk usage is over 90%. Value range: 0-100. | <code>90</code>                                  |
| form.alert.groups.database.rules.elasticsearchDiskOutOfSpace.duration         |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchDiskOutOfSpace.severity         |                                                 | <code>critical</code>                            |
| form.alert.groups.database.rules.elasticsearchDiskSpaceLow.enabled            |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchDiskSpaceLow.val                | The disk usage is over 80%. Value range: 0-100. | <code>80</code>                                  |
| form.alert.groups.database.rules.elasticsearchDiskSpaceLow.duration           |                                                 | <code>"2m"</code>                                |
| form.alert.groups.database.rules.elasticsearchDiskSpaceLow.severity           |                                                 | <code>warning</code>                             |
| form.alert.groups.database.rules.elasticsearchClusterRed.enabled              |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchClusterRed.duration             |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchClusterRed.severity             |                                                 | <code>critical</code>                            |
| form.alert.groups.database.rules.elasticsearchClusterYellow.enabled           |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchClusterYellow.duration          |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchClusterYellow.severity          |                                                 | <code>warning</code>                             |
| form.alert.groups.database.rules.elasticsearchHealthyNodes.enabled            |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchHealthyNodes.val                | should have at least 3 healthy nodes            | <code>3</code>                                   |
| form.alert.groups.database.rules.elasticsearchHealthyNodes.duration           |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchHealthyNodes.severity           |                                                 | <code>critical</code>                            |
| form.alert.groups.database.rules.elasticsearchHealthyDataNodes.enabled        |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchHealthyDataNodes.val            | should have at least 3 healthy data nodes       | <code>3</code>                                   |
| form.alert.groups.database.rules.elasticsearchHealthyDataNodes.duration       |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchHealthyDataNodes.severity       |                                                 | <code>critical</code>                            |
| form.alert.groups.database.rules.elasticsearchRelocatingShards.enabled        |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchRelocatingShards.duration       |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchRelocatingShards.severity       |                                                 | <code>info</code>                                |
| form.alert.groups.database.rules.elasticsearchInitializingShards.enabled      |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchInitializingShards.duration     |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchInitializingShards.severity     |                                                 | <code>info</code>                                |
| form.alert.groups.database.rules.elasticsearchUnassignedShards.enabled        |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchUnassignedShards.duration       |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchUnassignedShards.severity       |                                                 | <code>critical</code>                            |
| form.alert.groups.database.rules.elasticsearchPendingTasks.enabled            |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.elasticsearchPendingTasks.duration           |                                                 | <code>"15m"</code>                               |
| form.alert.groups.database.rules.elasticsearchPendingTasks.severity           |                                                 | <code>warning</code>                             |
| form.alert.groups.database.rules.elasticsearchNoNewDocuments10m.enabled       |                                                 | <code>false</code>                               |
| form.alert.groups.database.rules.elasticsearchNoNewDocuments10m.duration      |                                                 | <code>"0m"</code>                                |
| form.alert.groups.database.rules.elasticsearchNoNewDocuments10m.severity      |                                                 | <code>info</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.enabled                        |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.val                            |                                                 | <code>80</code>                                  |
| form.alert.groups.database.rules.diskUsageHigh.duration                       |                                                 | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.severity                       |                                                 | <code>warning</code>                             |
| form.alert.groups.database.rules.diskAlmostFull.enabled                       |                                                 | <code>true</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.val                           |                                                 | <code>95</code>                                  |
| form.alert.groups.database.rules.diskAlmostFull.duration                      |                                                 | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.severity                      |                                                 | <code>critical</code>                            |
| form.alert.groups.provisioner.enabled                                         |                                                 | <code>warning</code>                             |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                                 | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                                 | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                                 | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                                 | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                                 | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                                 | <code>warning</code>                             |
| form.alert.groups.opsManager.enabled                                          |                                                 | <code>warning</code>                             |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                                 | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                                 | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                                 | <code>info</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                                 | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                                 | <code>"30m"</code>                               |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                                 | <code>critical</code>                            |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                                 | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                                 | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                                 | <code>critical</code>                            |
| form.alert.groups.stash.enabled                                               |                                                 | <code>warning</code>                             |
| form.alert.groups.stash.rules.backupSessionFailed.enabled                     |                                                 | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.duration                    |                                                 | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.severity                    |                                                 | <code>critical</code>                            |
| form.alert.groups.stash.rules.restoreSessionFailed.enabled                    |                                                 | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.duration                   |                                                 | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.severity                   |                                                 | <code>critical</code>                            |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.enabled               |                                                 | <code>true</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.duration              |                                                 | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.val                   |                                                 | <code>18000</code>                               |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.severity              |                                                 | <code>warning</code>                             |
| form.alert.groups.stash.rules.repositoryCorrupted.enabled                     |                                                 | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.duration                    |                                                 | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.severity                    |                                                 | <code>critical</code>                            |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.enabled             |                                                 | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.duration            |                                                 | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.val                 |                                                 | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.severity            |                                                 | <code>warning</code>                             |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.enabled              |                                                 | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.duration             |                                                 | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.val                  |                                                 | <code>1800 # 30 minute</code>                    |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.severity             |                                                 | <code>warning</code>                             |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.enabled             |                                                 | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.duration            |                                                 | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.val                 |                                                 | <code>1800 # 30 minute</code>                    |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.severity            |                                                 | <code>warning</code>                             |
| form.alert.groups.kubeStash.enabled                                           |                                                 | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.backupSessionFailed.enabled                 |                                                 | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionFailed.duration                |                                                 | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionFailed.severity                |                                                 | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.enabled                |                                                 | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.duration               |                                                 | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.severity               |                                                 | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.enabled           |                                                 | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.duration          |                                                 | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.val               |                                                 | <code>18000</code>                               |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.severity          |                                                 | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.enabled                 |                                                 | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.duration                |                                                 | <code>"5m"</code>                                |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.severity                |                                                 | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.enabled         |                                                 | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.duration        |                                                 | <code>"5m"</code>                                |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.val             |                                                 | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.severity        |                                                 | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.enabled          |                                                 | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.duration         |                                                 | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.val              |                                                 | <code>1800 # 30 minute</code>                    |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.severity         |                                                 | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.enabled         |                                                 | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.duration        |                                                 | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.val             |                                                 | <code>1800 # 30 minute</code>                    |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.severity        |                                                 | <code>warning</code>                             |
| grafana.enabled                                                               |                                                 | <code>false</code>                               |
| grafana.version                                                               |                                                 | <code>7.5.5</code>                               |
| grafana.jobName                                                               |                                                 | <code>kubedb-databases</code>                    |
| grafana.url                                                                   |                                                 | <code>""</code>                                  |
| grafana.apikey                                                                |                                                 | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i elasticsearch appscode/elasticsearch-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i elasticsearch appscode/elasticsearch-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
