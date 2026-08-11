# Solr alerts

[Solr alerts by AppsCode](https://github.com/appscode/alerts) - Solr alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/solr-alerts --version=v2026.7.14
$ helm upgrade -i solr appscode/solr-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Solr alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `solr`:

```bash
$ helm upgrade -i solr appscode/solr-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Solr alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `solr`:

```bash
$ helm uninstall solr -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `solr-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Solr</code>                                |
| metadata.resource.name                                                        |                                               | <code>solrs</code>                               |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.solrDownShards.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.solrDownShards.duration                      |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.solrDownShards.severity                      |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.solrDownShards.val                           |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.solrRecoveryFailedShards.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.solrRecoveryFailedShards.duration            |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.solrRecoveryFailedShards.severity            |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.solrHighThreadRunning.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.solrHighThreadRunning.duration               |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.solrHighThreadRunning.val                    |                                               | <code>300</code>                                 |
| form.alert.groups.database.rules.solrHighThreadRunning.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.solrHighPoolSize.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.solrHighPoolSize.duration                    |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.solrHighPoolSize.val                         |                                               | <code>85</code>                                  |
| form.alert.groups.database.rules.solrHighPoolSize.severity                    |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.solrHighQPS.enabled                          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.solrHighQPS.duration                         |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.solrHighQPS.val                              |                                               | <code>1000</code>                                |
| form.alert.groups.database.rules.solrHighQPS.severity                         |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.solrHighHeapSize.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.solrHighHeapSize.duration                    |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.solrHighHeapSize.val                         |                                               | <code>85</code>                                  |
| form.alert.groups.database.rules.solrHighHeapSize.severity                    |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.solrHighBufferSize.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.solrHighBufferSize.duration                  |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.solrHighBufferSize.val                       |                                               | <code>25</code>                                  |
| form.alert.groups.database.rules.solrHighBufferSize.severity                  |                                               | <code>warning</code>                             |
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
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"1m"</code>                                |
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
$ helm upgrade -i solr appscode/solr-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i solr appscode/solr-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
