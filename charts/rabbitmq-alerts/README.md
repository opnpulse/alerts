# RabbitMQ alerts

[RabbitMQ alerts by AppsCode](https://github.com/appscode/alerts) - RabbitMQ alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/rabbitmq-alerts --version=v2026.7.14
$ helm upgrade -i rabbitmq appscode/rabbitmq-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys RabbitMQ alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `rabbitmq`:

```bash
$ helm upgrade -i rabbitmq appscode/rabbitmq-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys RabbitMQ alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `rabbitmq`:

```bash
$ helm uninstall rabbitmq -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `rabbitmq-alerts` chart and their default values.

|                                            Parameter                                             |                  Description                  |                     Default                      |
|--------------------------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                                          |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                                           |                                               | <code>RabbitMQ</code>                            |
| metadata.resource.name                                                                           |                                               | <code>rabbitmqs</code>                           |
| metadata.resource.scope                                                                          |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                                        |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                                            | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                                       | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                                               | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                                                | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                                           | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                                                  | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                                               |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqPhaseCritical.enabled                                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqPhaseCritical.duration                                  |                                               | <code>"3m"</code>                                |
| form.alert.groups.database.rules.rabbitmqPhaseCritical.severity                                  |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqDown.enabled                                            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqDown.duration                                           |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqDown.severity                                           |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.rabbitmqFileDescriptorsNearLimit.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqFileDescriptorsNearLimit.duration                       |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqFileDescriptorsNearLimit.severity                       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqHighConnectionChurn.enabled                             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqHighConnectionChurn.duration                            |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqHighConnectionChurn.severity                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqInsufficientEstablishedErlangDistributionLinks.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqInsufficientEstablishedErlangDistributionLinks.duration |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqInsufficientEstablishedErlangDistributionLinks.severity |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqLowDiskWatermarkPredicted.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqLowDiskWatermarkPredicted.duration                      |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqLowDiskWatermarkPredicted.severity                      |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqUnroutableMessages.enabled                              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqUnroutableMessages.duration                             |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqUnroutableMessages.severity                             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqTCPSocketsNearLimit.enabled                             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqTCPSocketsNearLimit.duration                            |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqTCPSocketsNearLimit.severity                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.rabbitmqQueueIsGrowing.enabled                                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.rabbitmqQueueIsGrowing.duration                                 |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.rabbitmqQueueIsGrowing.severity                                 |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.diskUsageHigh.enabled                                           |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.val                                               |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.diskUsageHigh.duration                                          |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.severity                                          |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.diskAlmostFull.enabled                                          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.val                                              |                                               | <code>95</code>                                  |
| form.alert.groups.database.rules.diskAlmostFull.duration                                         |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.severity                                         |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.enabled                                                            |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                                     |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                                    |                                               | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                                    |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                                     |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                                    |                                               | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                                    |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.enabled                                                             |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled                                  |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration                                 |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity                                 |                                               | <code>info</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration                    |                                               | <code>"30m"</code>                               |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                                      |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                                     |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                                     |                                               | <code>critical</code>                            |
| grafana.enabled                                                                                  |                                               | <code>false</code>                               |
| grafana.version                                                                                  |                                               | <code>7.5.5</code>                               |
| grafana.jobName                                                                                  |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                                      |                                               | <code>""</code>                                  |
| grafana.apikey                                                                                   |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i rabbitmq appscode/rabbitmq-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i rabbitmq appscode/rabbitmq-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
