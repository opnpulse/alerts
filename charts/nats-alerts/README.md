# NATS alerts

[NATS alerts by AppsCode](https://github.com/appscode/alerts) - NATS alerts for NATS clusters

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/nats-alerts --version=v2026.2.24
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2026.2.24
```

## Introduction

This chart deploys NATS alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+
- Prometheus operator

## Installing the Chart

To install/upgrade the chart with the release name `nats-alerts`:

```bash
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2026.2.24
```

The command deploys NATS alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `nats-alerts`:

```bash
$ helm uninstall nats-alerts -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `nats-alerts` chart and their default values.

|                                        Parameter                                        |                  Description                  |                     Default                      |
|-----------------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                                 |                                               | <code>apps</code>                                |
| metadata.resource.kind                                                                  |                                               | <code>StatefulSet</code>                         |
| metadata.resource.name                                                                  |                                               | <code>statefulsets</code>                        |
| metadata.resource.scope                                                                 |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                               |                                               | <code>v1</code>                                  |
| metadata.release.name                                                                   | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                              | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                                      | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                                       | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                                  | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                                         | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.core.enabled                                                          |                                               | <code>warning</code>                             |
| form.alert.groups.core.rules.natsDown.enabled                                           |                                               | <code>true</code>                                |
| form.alert.groups.core.rules.natsDown.duration                                          |                                               | <code>"15m"</code>                               |
| form.alert.groups.core.rules.natsDown.severity                                          |                                               | <code>critical</code>                            |
| form.alert.groups.core.rules.natsReplicasNotReady.enabled                               |                                               | <code>true</code>                                |
| form.alert.groups.core.rules.natsReplicasNotReady.duration                              |                                               | <code>"15m"</code>                               |
| form.alert.groups.core.rules.natsReplicasNotReady.severity                              |                                               | <code>critical</code>                            |
| form.alert.groups.resourceUtilization.enabled                                           |                                               | <code>warning</code>                             |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighMemoryUsage.enabled        |                                               | <code>true</code>                                |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighMemoryUsage.duration       |                                               | <code>"5m"</code>                                |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighMemoryUsage.severity       |                                               | <code>critical</code>                            |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighMemoryUsage.val            |                                               | <code>90</code>                                  |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighStorageUsage.enabled       |                                               | <code>true</code>                                |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighStorageUsage.duration      |                                               | <code>"5m"</code>                                |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighStorageUsage.severity      |                                               | <code>critical</code>                            |
| form.alert.groups.resourceUtilization.rules.natsJetStreamHighStorageUsage.val           |                                               | <code>90</code>                                  |
| form.alert.groups.consumerHealth.enabled                                                |                                               | <code>warning</code>                             |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessages.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessages.duration        |                                               | <code>"10m"</code>                               |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessages.severity        |                                               | <code>critical</code>                            |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessages.val             |                                               | <code>5000</code>                                |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessagesWarning.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessagesWarning.duration |                                               | <code>"10m"</code>                               |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessagesWarning.severity |                                               | <code>warning</code>                             |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighPendingMessagesWarning.val      |                                               | <code>2000</code>                                |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighAckPending.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighAckPending.duration             |                                               | <code>"10m"</code>                               |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighAckPending.severity             |                                               | <code>critical</code>                            |
| form.alert.groups.consumerHealth.rules.natsJetStreamHighAckPending.val                  |                                               | <code>2000</code>                                |
| form.alert.groups.consumerHealth.rules.natsJetStreamBacklogNoProgress.enabled           |                                               | <code>true</code>                                |
| form.alert.groups.consumerHealth.rules.natsJetStreamBacklogNoProgress.duration          |                                               | <code>"15m"</code>                               |
| form.alert.groups.consumerHealth.rules.natsJetStreamBacklogNoProgress.severity          |                                               | <code>warning</code>                             |
| form.alert.groups.consumerHealth.rules.natsJetStreamBacklogNoProgress.val               |                                               | <code>100</code>                                 |
| form.alert.groups.streamPerformance.enabled                                             |                                               | <code>warning</code>                             |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearMessageLimit.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearMessageLimit.duration        |                                               | <code>"15m"</code>                               |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearMessageLimit.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearMessageLimit.val             |                                               | <code>80</code>                                  |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearByteLimit.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearByteLimit.duration           |                                               | <code>"15m"</code>                               |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearByteLimit.severity           |                                               | <code>warning</code>                             |
| form.alert.groups.streamPerformance.rules.natsJetStreamNearByteLimit.val                |                                               | <code>80</code>                                  |
| form.alert.groups.connectivity.enabled                                                  |                                               | <code>warning</code>                             |
| form.alert.groups.connectivity.rules.natsJetStreamDisabled.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.connectivity.rules.natsJetStreamDisabled.duration                     |                                               | <code>"2m"</code>                                |
| form.alert.groups.connectivity.rules.natsJetStreamDisabled.severity                     |                                               | <code>critical</code>                            |
| form.alert.groups.connectivity.rules.natsSlowConsumers.enabled                          |                                               | <code>true</code>                                |
| form.alert.groups.connectivity.rules.natsSlowConsumers.duration                         |                                               | <code>"2m"</code>                                |
| form.alert.groups.connectivity.rules.natsSlowConsumers.severity                         |                                               | <code>warning</code>                             |
| form.alert.groups.connectivity.rules.natsSlowConsumers.val                              |                                               | <code>5</code>                                   |
| form.alert.groups.connectivity.rules.natsStalledClients.enabled                         |                                               | <code>true</code>                                |
| form.alert.groups.connectivity.rules.natsStalledClients.duration                        |                                               | <code>"2m"</code>                                |
| form.alert.groups.connectivity.rules.natsStalledClients.severity                        |                                               | <code>warning</code>                             |
| form.alert.groups.connectivity.rules.natsStalledClients.val                             |                                               | <code>5</code>                                   |
| form.alert.groups.connectivity.rules.natsStaleConnections.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.connectivity.rules.natsStaleConnections.duration                      |                                               | <code>"5m"</code>                                |
| form.alert.groups.connectivity.rules.natsStaleConnections.severity                      |                                               | <code>warning</code>                             |
| form.alert.groups.connectivity.rules.natsStaleConnections.val                           |                                               | <code>8</code>                                   |
| form.alert.groups.connectivity.rules.natsHighActiveConnections.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.connectivity.rules.natsHighActiveConnections.duration                 |                                               | <code>"5m"</code>                                |
| form.alert.groups.connectivity.rules.natsHighActiveConnections.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.connectivity.rules.natsHighActiveConnections.val                      |                                               | <code>5000</code>                                |
| form.alert.groups.consumerManagement.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.consumerManagement.rules.natsHighTotalConsumers.enabled               |                                               | <code>false</code>                               |
| form.alert.groups.consumerManagement.rules.natsHighTotalConsumers.duration              |                                               | <code>"15m"</code>                               |
| form.alert.groups.consumerManagement.rules.natsHighTotalConsumers.severity              |                                               | <code>warning</code>                             |
| form.alert.groups.consumerManagement.rules.natsHighTotalConsumers.val                   |                                               | <code>100</code>                                 |
| grafana.enabled                                                                         |                                               | <code>false</code>                               |
| grafana.version                                                                         |                                               | <code>8.2.3</code>                               |
| grafana.jobName                                                                         |                                               | <code>nats</code>                                |
| grafana.url                                                                             |                                               | <code>""</code>                                  |
| grafana.apikey                                                                          |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2026.2.24 --set metadata.resource.group=apps
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2026.2.24 --values values.yaml
```
