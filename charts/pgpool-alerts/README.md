# Pgpool alerts

[Pgpool alerts by AppsCode](https://github.com/appscode/alerts) - Pgpool alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/pgpool-alerts --version=v2026.7.14
$ helm upgrade -i pgpool appscode/pgpool-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Pgpool alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `pgpool`:

```bash
$ helm upgrade -i pgpool appscode/pgpool-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Pgpool alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `pgpool`:

```bash
$ helm uninstall pgpool -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `pgpool-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                Default                |
|-------------------------------------------------------------------------------|-----------------------------------------------|---------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>               |
| metadata.resource.kind                                                        |                                               | <code>Pgpool</code>                   |
| metadata.resource.name                                                        |                                               | <code>pgpools</code>                  |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>               |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                 |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                       |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                       |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                  |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"prometheus"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                       |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                       |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.pgpoolTooManyConnections.enabled             |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolTooManyConnections.val                 |                                               | <code>.1 # 10%</code>                 |
| form.alert.groups.database.rules.pgpoolTooManyConnections.duration            |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.pgpoolTooManyConnections.severity            |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.pgpoolExporterLastScrapeError.enabled        |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolExporterLastScrapeError.duration       |                                               | <code>"0m"</code>                     |
| form.alert.groups.database.rules.pgpoolExporterLastScrapeError.severity       |                                               | <code>warning</code>                  |
| form.alert.groups.database.rules.pgpoolDown.enabled                           |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolDown.duration                          |                                               | <code>"0m"</code>                     |
| form.alert.groups.database.rules.pgpoolDown.severity                          |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.pgpoolPostgresHealthCheckFailure.enabled     |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolPostgresHealthCheckFailure.val         |                                               | <code>10</code>                       |
| form.alert.groups.database.rules.pgpoolPostgresHealthCheckFailure.duration    |                                               | <code>"0m"</code>                     |
| form.alert.groups.database.rules.pgpoolPostgresHealthCheckFailure.severity    |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.pgpoolBackendPanicMessageCount.enabled       |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolBackendPanicMessageCount.val           |                                               | <code>10</code>                       |
| form.alert.groups.database.rules.pgpoolBackendPanicMessageCount.duration      |                                               | <code>"0m"</code>                     |
| form.alert.groups.database.rules.pgpoolBackendPanicMessageCount.severity      |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.pgpoolBackendFatalMessageCount.enabled       |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolBackendFatalMessageCount.val           |                                               | <code>10</code>                       |
| form.alert.groups.database.rules.pgpoolBackendFatalMessageCount.duration      |                                               | <code>"0m"</code>                     |
| form.alert.groups.database.rules.pgpoolBackendFatalMessageCount.severity      |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.pgpoolBackendErrorMessageCount.enabled       |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolBackendErrorMessageCount.val           |                                               | <code>10</code>                       |
| form.alert.groups.database.rules.pgpoolBackendErrorMessageCount.duration      |                                               | <code>"0m"</code>                     |
| form.alert.groups.database.rules.pgpoolBackendErrorMessageCount.severity      |                                               | <code>critical</code>                 |
| form.alert.groups.database.rules.pgpoolLowCacheMemory.enabled                 |                                               | <code>true</code>                     |
| form.alert.groups.database.rules.pgpoolLowCacheMemory.val                     |                                               | <code>100 # 10mb</code>               |
| form.alert.groups.database.rules.pgpoolLowCacheMemory.duration                |                                               | <code>"1m"</code>                     |
| form.alert.groups.database.rules.pgpoolLowCacheMemory.severity                |                                               | <code>warning</code>                  |
| form.alert.groups.provisioner.enabled                                         |                                               | <code>warning</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                               | <code>"1m"</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                               | <code>critical</code>                 |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                               | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"15m"</code>                    |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                               | <code>warning</code>                  |
| form.alert.groups.opsManager.enabled                                          |                                               | <code>warning</code>                  |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                               | <code>true</code>                     |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                               | <code>"0m"</code>                     |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                               | <code>info</code>                     |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                               | <code>true</code>                     |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                               | <code>"30m"</code>                    |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                               | <code>critical</code>                 |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                               | <code>true</code>                     |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                               | <code>"0m"</code>                     |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                               | <code>critical</code>                 |
| grafana.enabled                                                               |                                               | <code>false</code>                    |
| grafana.version                                                               |                                               | <code>7.5.5</code>                    |
| grafana.jobName                                                               |                                               | <code>kubedb-databases</code>         |
| grafana.url                                                                   |                                               | <code>""</code>                       |
| grafana.apikey                                                                |                                               | <code>""</code>                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i pgpool appscode/pgpool-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i pgpool appscode/pgpool-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
