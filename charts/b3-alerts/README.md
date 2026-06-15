# B3 alerts

[B3 alerts by AppsCode](https://github.com/appscode/alerts) - B3 platform alerts for monitoring CPU, memory, HTTP, database, messaging, and more

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/b3-alerts --version=v2026.2.24
$ helm upgrade -i b3-alerts appscode/b3-alerts -n monitoring --create-namespace --version=v2026.2.24
```

## Introduction

This chart deploys B3 alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+
- Prometheus operator
- kube-state-metrics (for CPU/memory limit-based alerts)

## Installing the Chart

To install/upgrade the chart with the release name `b3-alerts`:

```bash
$ helm upgrade -i b3-alerts appscode/b3-alerts -n monitoring --create-namespace --version=v2026.2.24
```

The command deploys B3 alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `b3-alerts`:

```bash
$ helm uninstall b3-alerts -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `b3-alerts` chart and their default values.

|                                  Parameter                                  | Description |                     Default                      |
|-----------------------------------------------------------------------------|-------------|--------------------------------------------------|
| metadata.resource.group                                                     |             | <code>apps</code>                                |
| metadata.resource.kind                                                      |             | <code>StatefulSet</code>                         |
| metadata.resource.name                                                      |             | <code>statefulsets</code>                        |
| metadata.resource.scope                                                     |             | <code>Namespaced</code>                          |
| metadata.resource.version                                                   |             | <code>v1</code>                                  |
| metadata.release.name                                                       |             | <code>""</code>                                  |
| metadata.release.namespace                                                  |             | <code>""</code>                                  |
| form.alert.enabled                                                          |             | <code>warning</code>                             |
| form.alert.labels                                                           |             | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                      |             | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                             |             | <code>{}</code>                                  |
| form.alert.groups.process.enabled                                           |             | <code>warning</code>                             |
| form.alert.groups.process.rules.b3Down.enabled                              |             | <code>true</code>                                |
| form.alert.groups.process.rules.b3Down.duration                             |             | <code>"15m"</code>                               |
| form.alert.groups.process.rules.b3Down.severity                             |             | <code>critical</code>                            |
| form.alert.groups.process.rules.b3ReplicasNotReady.enabled                  |             | <code>true</code>                                |
| form.alert.groups.process.rules.b3ReplicasNotReady.duration                 |             | <code>"15m"</code>                               |
| form.alert.groups.process.rules.b3ReplicasNotReady.severity                 |             | <code>critical</code>                            |
| form.alert.groups.process.rules.b3HighGcPauseDuration.enabled               |             | <code>true</code>                                |
| form.alert.groups.process.rules.b3HighGcPauseDuration.duration              |             | <code>"5m"</code>                                |
| form.alert.groups.process.rules.b3HighGcPauseDuration.severity              |             | <code>warning</code>                             |
| form.alert.groups.process.rules.b3HighGcPauseDuration.val                   |             | <code>0.2</code>                                 |
| form.alert.groups.process.rules.b3HighFileDescriptorUsage.enabled           |             | <code>true</code>                                |
| form.alert.groups.process.rules.b3HighFileDescriptorUsage.duration          |             | <code>"5m"</code>                                |
| form.alert.groups.process.rules.b3HighFileDescriptorUsage.severity          |             | <code>warning</code>                             |
| form.alert.groups.process.rules.b3HighFileDescriptorUsage.val               |             | <code>80</code>                                  |
| form.alert.groups.http.enabled                                              |             | <code>warning</code>                             |
| form.alert.groups.http.rules.b3HighHttpErrorRate.enabled                    |             | <code>true</code>                                |
| form.alert.groups.http.rules.b3HighHttpErrorRate.duration                   |             | <code>"5m"</code>                                |
| form.alert.groups.http.rules.b3HighHttpErrorRate.severity                   |             | <code>critical</code>                            |
| form.alert.groups.http.rules.b3HighHttpErrorRate.val                        |             | <code>5</code>                                   |
| form.alert.groups.http.rules.b3HighHttpLatency.enabled                      |             | <code>true</code>                                |
| form.alert.groups.http.rules.b3HighHttpLatency.duration                     |             | <code>"5m"</code>                                |
| form.alert.groups.http.rules.b3HighHttpLatency.severity                     |             | <code>warning</code>                             |
| form.alert.groups.http.rules.b3HighHttpLatency.val                          |             | <code>1</code>                                   |
| form.alert.groups.database.enabled                                          |             | <code>warning</code>                             |
| form.alert.groups.database.rules.b3HighDatabaseErrorRate.enabled            |             | <code>true</code>                                |
| form.alert.groups.database.rules.b3HighDatabaseErrorRate.duration           |             | <code>"5m"</code>                                |
| form.alert.groups.database.rules.b3HighDatabaseErrorRate.severity           |             | <code>critical</code>                            |
| form.alert.groups.database.rules.b3HighDatabaseErrorRate.val                |             | <code>5</code>                                   |
| form.alert.groups.database.rules.b3HighDatabaseLatency.enabled              |             | <code>true</code>                                |
| form.alert.groups.database.rules.b3HighDatabaseLatency.duration             |             | <code>"5m"</code>                                |
| form.alert.groups.database.rules.b3HighDatabaseLatency.severity             |             | <code>warning</code>                             |
| form.alert.groups.database.rules.b3HighDatabaseLatency.val                  |             | <code>0.5</code>                                 |
| form.alert.groups.database.rules.b3DatabaseConnectionErrors.enabled         |             | <code>true</code>                                |
| form.alert.groups.database.rules.b3DatabaseConnectionErrors.duration        |             | <code>"5m"</code>                                |
| form.alert.groups.database.rules.b3DatabaseConnectionErrors.severity        |             | <code>critical</code>                            |
| form.alert.groups.database.rules.b3DatabaseConnectionErrors.val             |             | <code>5</code>                                   |
| form.alert.groups.database.rules.b3HighDatabaseConnectionPoolUsage.enabled  |             | <code>true</code>                                |
| form.alert.groups.database.rules.b3HighDatabaseConnectionPoolUsage.duration |             | <code>"5m"</code>                                |
| form.alert.groups.database.rules.b3HighDatabaseConnectionPoolUsage.severity |             | <code>warning</code>                             |
| form.alert.groups.database.rules.b3HighDatabaseConnectionPoolUsage.val      |             | <code>80</code>                                  |
| form.alert.groups.auth.enabled                                              |             | <code>warning</code>                             |
<<<<<<< HEAD
| form.alert.groups.auth.rules.b3HighAuthzDenialRate.enabled                  |             | <code>true</code>                                |
| form.alert.groups.auth.rules.b3HighAuthzDenialRate.duration                 |             | <code>"5m"</code>                                |
| form.alert.groups.auth.rules.b3HighAuthzDenialRate.severity                 |             | <code>warning</code>                             |
| form.alert.groups.auth.rules.b3HighAuthzDenialRate.val                      |             | <code>10</code>                                  |
=======
| form.alert.groups.auth.rules                                                |             | <code></code>                                    |
>>>>>>> 983eddd8 (Remove unused types/values and run `make gen manifests fmt`)
| form.alert.groups.storage.enabled                                           |             | <code>warning</code>                             |
| form.alert.groups.storage.rules.b3HighBlobStorageErrorRate.enabled          |             | <code>true</code>                                |
| form.alert.groups.storage.rules.b3HighBlobStorageErrorRate.duration         |             | <code>"5m"</code>                                |
| form.alert.groups.storage.rules.b3HighBlobStorageErrorRate.severity         |             | <code>warning</code>                             |
| form.alert.groups.storage.rules.b3HighBlobStorageErrorRate.val              |             | <code>5</code>                                   |
| form.alert.groups.proxy.enabled                                             |             | <code>warning</code>                             |
| form.alert.groups.proxy.rules.b3HighProxyErrorRate.enabled                  |             | <code>true</code>                                |
| form.alert.groups.proxy.rules.b3HighProxyErrorRate.duration                 |             | <code>"5m"</code>                                |
| form.alert.groups.proxy.rules.b3HighProxyErrorRate.severity                 |             | <code>warning</code>                             |
| form.alert.groups.proxy.rules.b3HighProxyErrorRate.val                      |             | <code>5</code>                                   |
| form.alert.groups.proxy.rules.b3HighProxyLatency.enabled                    |             | <code>true</code>                                |
| form.alert.groups.proxy.rules.b3HighProxyLatency.duration                   |             | <code>"5m"</code>                                |
| form.alert.groups.proxy.rules.b3HighProxyLatency.severity                   |             | <code>warning</code>                             |
| form.alert.groups.proxy.rules.b3HighProxyLatency.val                        |             | <code>5</code>                                   |
| grafana.enabled                                                             |             | <code>false</code>                               |
| grafana.version                                                             |             | <code>8.2.3</code>                               |
| grafana.jobName                                                             |             | <code>ace-platform-api</code>                    |
| grafana.url                                                                 |             | <code>""</code>                                  |
| grafana.apikey                                                              |             | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i b3-alerts appscode/b3-alerts -n monitoring --create-namespace --version=v2026.2.24 --set metadata.resource.group=apps
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i b3-alerts appscode/b3-alerts -n monitoring --create-namespace --version=v2026.2.24 --values values.yaml
```
