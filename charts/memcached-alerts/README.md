# Memcached alerts

[Memcached alerts by AppsCode](https://github.com/appscode/alerts) - Memcached alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/memcached-alerts --version=v2026.7.14
$ helm upgrade -i memcached appscode/memcached-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Memcached alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `memcached`:

```bash
$ helm upgrade -i memcached appscode/memcached-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys Memcached alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `memcached`:

```bash
$ helm uninstall memcached -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `memcached-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Memcached</code>                           |
| metadata.resource.name                                                        |                                               | <code>memcacheds</code>                          |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1</code>                                  |
| metadata.release.name                                                         | Release name                                  | <code>"coreos-prom-memcd"</code>                 |
| metadata.release.namespace                                                    | Release namespace                             | <code>"demo"</code>                              |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.memcachedDown.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.memcachedDown.duration                       |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.memcachedDown.severity                       |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.memcachedServiceRespawn.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.memcachedServiceRespawn.duration             |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.memcachedServiceRespawn.val                  |                                               | <code>180</code>                                 |
| form.alert.groups.database.rules.memcachedServiceRespawn.severity             |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.memcachedConnectionThrottled.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.memcachedConnectionThrottled.duration        |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.memcachedConnectionThrottled.val             |                                               | <code>10</code>                                  |
| form.alert.groups.database.rules.memcachedConnectionThrottled.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.memcachedConnectionsNoneMinor.enabled        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.memcachedConnectionsNoneMinor.duration       |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.memcachedConnectionsNoneMinor.severity       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.memcachedItemsNoneMinor.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.memcachedItemsNoneMinor.duration             |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.memcachedItemsNoneMinor.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.memcachedEvictionsLimit.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.memcachedEvictionsLimit.duration             |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.memcachedEvictionsLimit.val                  |                                               | <code>10</code>                                  |
| form.alert.groups.database.rules.memcachedEvictionsLimit.severity             |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.memcachedMemoryLimit.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.memcachedMemoryLimit.duration                |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.memcachedMemoryLimit.val                     |                                               | <code>33554432 # 32MB</code>                     |
| form.alert.groups.database.rules.memcachedMemoryLimit.severity                |                                               | <code>critical</code>                            |
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
$ helm upgrade -i memcached appscode/memcached-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i memcached appscode/memcached-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
