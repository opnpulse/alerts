# ConnectCluster alerts

[ConnectCluster alerts by AppsCode](https://github.com/appscode/alerts) - ConnectCluster alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/connectcluster-alerts --version=v2026.7.14
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys ConnectCluster alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `connectcluster`:

```bash
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys ConnectCluster alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `connectcluster`:

```bash
$ helm uninstall connectcluster -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `connectcluster-alerts` chart and their default values.

|                                     Parameter                                     |                  Description                  |                     Default                      |
|-----------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                           |                                               | <code>kafka.kubedb.com</code>                    |
| metadata.resource.kind                                                            |                                               | <code>ConnectCluster</code>                      |
| metadata.resource.name                                                            |                                               | <code>connectclusters</code>                     |
| metadata.resource.scope                                                           |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                         |                                               | <code>v1alpha1</code>                            |
| metadata.release.name                                                             | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                        | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                                | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                                 | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                            | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                                   | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.connect.enabled                                                 |                                               | <code>warning</code>                             |
| form.alert.groups.connect.rules.connectClusterWorkerDown.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.connect.rules.connectClusterWorkerDown.duration                 |                                               | <code>"1m"</code>                                |
| form.alert.groups.connect.rules.connectClusterWorkerDown.severity                 |                                               | <code>critical</code>                            |
| form.alert.groups.connect.rules.connectClusterTooManyConnections.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.connect.rules.connectClusterTooManyConnections.duration         |                                               | <code>"2m"</code>                                |
| form.alert.groups.connect.rules.connectClusterTooManyConnections.val              |                                               | <code>80</code>                                  |
| form.alert.groups.connect.rules.connectClusterTooManyConnections.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.connect.rules.connectClusterConnectorCount.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.connect.rules.connectClusterConnectorCount.duration             |                                               | <code>"1m"</code>                                |
| form.alert.groups.connect.rules.connectClusterConnectorCount.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.connect.rules.connectClusterConnectorCount.val                  |                                               | <code>50</code>                                  |
| form.alert.groups.connect.rules.connectClusterCoordinatorRebalanceFailed.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.connect.rules.connectClusterCoordinatorRebalanceFailed.duration |                                               | <code>"1m"</code>                                |
| form.alert.groups.connect.rules.connectClusterCoordinatorRebalanceFailed.severity |                                               | <code>critical</code>                            |
| form.alert.groups.connect.rules.connectClusterCoordinatorRebalanceFailed.val      |                                               | <code>10</code>                                  |
| form.alert.groups.task.enabled                                                    |                                               | <code>warning</code>                             |
| form.alert.groups.task.rules.connectClusterTaskErrorTotalRetries.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.task.rules.connectClusterTaskErrorTotalRetries.duration         |                                               | <code>"1m"</code>                                |
| form.alert.groups.task.rules.connectClusterTaskErrorTotalRetries.severity         |                                               | <code>critical</code>                            |
| form.alert.groups.task.rules.connectClusterTaskErrorTotalRetries.val              |                                               | <code>5</code>                                   |
| form.alert.groups.task.rules.connectClusterTaskTotal.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.task.rules.connectClusterTaskTotal.duration                     |                                               | <code>"1m"</code>                                |
| form.alert.groups.task.rules.connectClusterTaskTotal.severity                     |                                               | <code>warning</code>                             |
| form.alert.groups.task.rules.connectClusterTaskTotal.val                          |                                               | <code>150</code>                                 |
| form.alert.groups.task.rules.connectClusterTaskTotalFailed.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.task.rules.connectClusterTaskTotalFailed.duration               |                                               | <code>"1m"</code>                                |
| form.alert.groups.task.rules.connectClusterTaskTotalFailed.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.task.rules.connectClusterTaskTotalDestroyed.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.task.rules.connectClusterTaskTotalDestroyed.duration            |                                               | <code>"1m"</code>                                |
| form.alert.groups.task.rules.connectClusterTaskTotalDestroyed.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.enabled                                             |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                     |                                               | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                     |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                     |                                               | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                     |                                               | <code>warning</code>                             |
| grafana.enabled                                                                   |                                               | <code>false</code>                               |
| grafana.version                                                                   |                                               | <code>7.5.5</code>                               |
| grafana.jobName                                                                   |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                       |                                               | <code>""</code>                                  |
| grafana.apikey                                                                    |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kafka.kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
