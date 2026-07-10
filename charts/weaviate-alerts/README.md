# Weaviate alerts

[Weaviate alerts by AppsCode](https://github.com/appscode/alerts) - Weaviate alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/weaviate-alerts --version=v2026.7.14
$ helm upgrade -i weaviate appscode/weaviate-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Weaviate alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `weaviate`:

```bash
$ helm upgrade -i weaviate appscode/weaviate-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Weaviate alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `weaviate`:

```bash
$ helm uninstall weaviate -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `weaviate-alerts` chart and their default values.

|                                 Parameter                                 |                  Description                  |                Default                |
|---------------------------------------------------------------------------|-----------------------------------------------|---------------------------------------|
| metadata.resource.group                                                   |                                               | <code>kubedb.com</code>               |
| metadata.resource.kind                                                    |                                               | <code>Weaviate</code>                 |
| metadata.resource.name                                                    |                                               | <code>weaviates</code>                |
| metadata.resource.scope                                                   |                                               | <code>Namespaced</code>               |
| metadata.resource.version                                                 |                                               | <code>v1alpha2</code>                 |
| metadata.release.name                                                     | Release name                                  | <code>""</code>                       |
| metadata.release.namespace                                                | Release namespace                             | <code>""</code>                       |
| form.alert.enabled                                                        | # Enable PrometheusRule alerts                | <code>warning</code>                  |
| form.alert.labels                                                         | # Labels for default rules                    | <code>{"release":"prometheus"}</code> |
| form.alert.annotations                                                    | # Annotations for default rules               | <code>{}</code>                       |
| form.alert.additionalRuleLabels                                           | # Additional labels for PrometheusRule alerts | <code>{}</code>                       |
| form.alert.groups.database.enabled                                        |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateInstanceDown.enabled             |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateInstanceDown.duration            |                                               | <code>30s</code>                      |
| form.alert.groups.database.rules.weaviateInstanceDown.severity            |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.weaviateRestarted.enabled                |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateRestarted.duration               |                                               | <code>1m</code>                       |
| form.alert.groups.database.rules.weaviateRestarted.severity               |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateRestarted.val                    |                                               | <code>180</code>                      |
| form.alert.groups.database.rules.weaviateHighCPUUsage.enabled             |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateHighCPUUsage.duration            |                                               | <code>1m</code>                       |
| form.alert.groups.database.rules.weaviateHighCPUUsage.severity            |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateHighCPUUsage.val                 |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.weaviateHighMemoryUsage.enabled          |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateHighMemoryUsage.duration         |                                               | <code>1m</code>                       |
| form.alert.groups.database.rules.weaviateHighMemoryUsage.severity         |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateHighMemoryUsage.val              |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.weaviateHighProcessMemoryUsage.enabled   |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateHighProcessMemoryUsage.duration  |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateHighProcessMemoryUsage.severity  |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateHighProcessMemoryUsage.val       |                                               | <code>1073741824</code>               |
| form.alert.groups.database.rules.weaviateGoroutinesExplosion.enabled      |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateGoroutinesExplosion.duration     |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateGoroutinesExplosion.severity     |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateGoroutinesExplosion.val          |                                               | <code>1000</code>                     |
| form.alert.groups.database.rules.weaviateHighThreadPressure.enabled       |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateHighThreadPressure.duration      |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateHighThreadPressure.severity      |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateHighThreadPressure.val           |                                               | <code>100</code>                      |
| form.alert.groups.database.rules.weaviateHighFDsUsage.enabled             |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateHighFDsUsage.duration            |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateHighFDsUsage.severity            |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateHighFDsUsage.val                 |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.enabled                    |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.diskUsageHigh.val                        |                                               | <code>80</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.duration                   |                                               | <code>1m</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.severity                   |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.diskAlmostFull.enabled                   |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.diskAlmostFull.val                       |                                               | <code>95</code>                       |
| form.alert.groups.database.rules.diskAlmostFull.duration                  |                                               | <code>1m</code>                       |
| form.alert.groups.database.rules.diskAlmostFull.severity                  |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.weaviateHTTPErrorRateHigh.enabled        |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateHTTPErrorRateHigh.val            |                                               | <code>5</code>                        |
| form.alert.groups.database.rules.weaviateHTTPErrorRateHigh.duration       |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateHTTPErrorRateHigh.severity       |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateHTTPP95LatencyHigh.enabled       |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateHTTPP95LatencyHigh.val           |                                               | <code>1</code>                        |
| form.alert.groups.database.rules.weaviateHTTPP95LatencyHigh.duration      |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateHTTPP95LatencyHigh.severity      |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateGRPCErrorRateHigh.enabled        |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateGRPCErrorRateHigh.val            |                                               | <code>5</code>                        |
| form.alert.groups.database.rules.weaviateGRPCErrorRateHigh.duration       |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateGRPCErrorRateHigh.severity       |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateGRPCP95LatencyHigh.enabled       |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateGRPCP95LatencyHigh.val           |                                               | <code>1</code>                        |
| form.alert.groups.database.rules.weaviateGRPCP95LatencyHigh.duration      |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateGRPCP95LatencyHigh.severity      |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateReplicationEngineDown.enabled    |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateReplicationEngineDown.duration   |                                               | <code>1m</code>                       |
| form.alert.groups.database.rules.weaviateReplicationEngineDown.severity   |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.weaviateReplicationQueueHigh.enabled     |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateReplicationQueueHigh.val         |                                               | <code>10</code>                       |
| form.alert.groups.database.rules.weaviateReplicationQueueHigh.duration    |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateReplicationQueueHigh.severity    |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.weaviateReplicationFailuresHigh.enabled  |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.weaviateReplicationFailuresHigh.val      |                                               | <code>0</code>                        |
| form.alert.groups.database.rules.weaviateReplicationFailuresHigh.duration |                                               | <code>5m</code>                       |
| form.alert.groups.database.rules.weaviateReplicationFailuresHigh.severity |                                               | <code>warning</code>                  |
| form.alert.groups.provisioner.enabled                                     |                                               | <code>warning</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled              |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration             |                                               | <code>1m</code>                       |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity             |                                               | <code>critical</code>                 |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled              |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration             |                                               | <code>15m</code>                      |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity             |                                               | <code>warning</code>                  |
| grafana.enabled                                                           |                                               | <code>false</code>                    |
| grafana.version                                                           |                                               | <code>7.5.5</code>                    |
| grafana.jobName                                                           |                                               | <code>kubedb-databases</code>         |
| grafana.url                                                               |                                               | <code>""</code>                       |
| grafana.apikey                                                            |                                               | <code>""</code>                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i weaviate appscode/weaviate-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i weaviate appscode/weaviate-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
