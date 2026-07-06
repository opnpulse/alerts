# Kafka alerts

[Kafka alerts by AppsCode](https://github.com/appscode/alerts) - Kafka alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/kafka-alerts --version=v2026.7.14
$ helm upgrade -i kafka appscode/kafka-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Kafka alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `kafka`:

```bash
$ helm upgrade -i kafka appscode/kafka-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Kafka alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kafka`:

```bash
$ helm uninstall kafka -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kafka-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Kafka</code>                               |
| metadata.resource.name                                                        |                                               | <code>kafkas</code>                              |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaUnderReplicatedPartitions.enabled       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaUnderReplicatedPartitions.duration      |                                               | <code>"10s"</code>                               |
| form.alert.groups.database.rules.kafkaUnderReplicatedPartitions.severity      |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaUnderReplicatedPartitions.val           |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.kafkaAbnormalControllerState.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaAbnormalControllerState.duration        |                                               | <code>"10s"</code>                               |
| form.alert.groups.database.rules.kafkaAbnormalControllerState.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaAbnormalControllerState.val             |                                               | <code>1</code>                                   |
| form.alert.groups.database.rules.kafkaOfflinePartitions.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaOfflinePartitions.duration              |                                               | <code>"10s"</code>                               |
| form.alert.groups.database.rules.kafkaOfflinePartitions.severity              |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaOfflinePartitions.val                   |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.kafkaUnderMinIsrPartitionCount.enabled       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaUnderMinIsrPartitionCount.duration      |                                               | <code>"10s"</code>                               |
| form.alert.groups.database.rules.kafkaUnderMinIsrPartitionCount.severity      |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaUnderMinIsrPartitionCount.val           |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.kafkaOfflineLogDirectoryCount.enabled        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaOfflineLogDirectoryCount.duration       |                                               | <code>"10s"</code>                               |
| form.alert.groups.database.rules.kafkaOfflineLogDirectoryCount.severity       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaOfflineLogDirectoryCount.val            |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.kafkaISRExpandRate.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaISRExpandRate.duration                  |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.kafkaISRExpandRate.severity                  |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaISRExpandRate.val                       |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.kafkaISRShrinkRate.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaISRShrinkRate.duration                  |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.kafkaISRShrinkRate.severity                  |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaISRShrinkRate.val                       |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.kafkaBrokerCount.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaBrokerCount.duration                    |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.kafkaBrokerCount.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.kafkaBrokerCount.val                         |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.kafkaNetworkProcessorIdlePercent.enabled     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaNetworkProcessorIdlePercent.duration    |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.kafkaNetworkProcessorIdlePercent.severity    |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.kafkaNetworkProcessorIdlePercent.val         |                                               | <code>30</code>                                  |
| form.alert.groups.database.rules.kafkaRequestHandlerIdlePercent.enabled       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaRequestHandlerIdlePercent.duration      |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.kafkaRequestHandlerIdlePercent.severity      |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.kafkaRequestHandlerIdlePercent.val           |                                               | <code>30</code>                                  |
| form.alert.groups.database.rules.kafkaReplicaFetcherManagerMaxLag.enabled     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaReplicaFetcherManagerMaxLag.duration    |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.kafkaReplicaFetcherManagerMaxLag.severity    |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.kafkaReplicaFetcherManagerMaxLag.val         |                                               | <code>50</code>                                  |
| form.alert.groups.database.rules.kafkaTopicCount.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaTopicCount.duration                     |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.kafkaTopicCount.severity                     |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaTopicCount.val                          |                                               | <code>1000</code>                                |
| form.alert.groups.database.rules.kafkaPhaseCritical.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaPhaseCritical.duration                  |                                               | <code>"3m"</code>                                |
| form.alert.groups.database.rules.kafkaPhaseCritical.severity                  |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.kafkaDown.enabled                            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.kafkaDown.duration                           |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.kafkaDown.severity                           |                                               | <code>critical</code>                            |
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
$ helm upgrade -i kafka appscode/kafka-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kafka appscode/kafka-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
