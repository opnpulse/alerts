# Neo4j alerts

[Neo4j alerts by AppsCode](https://github.com/appscode/alerts) - Neo4j alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/neo4j-alerts --version=v2026.7.14
$ helm upgrade -i neo4j-alerts appscode/neo4j-alerts -n kubedb --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys Neo4j alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+
- Prometheus Operator 0.40+
- Grafana 7.0+ (optional, for dashboard visualization)
- KubeDB operator installed

## Installing the Chart

To install/upgrade the chart with the release name `neo4j-alerts`:

```bash
$ helm upgrade -i neo4j-alerts appscode/neo4j-alerts -n kubedb --create-namespace --version=v2026.7.14
```

The command deploys Neo4j alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `neo4j-alerts`:

```bash
$ helm uninstall neo4j-alerts -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `neo4j-alerts` chart and their default values.

|                               Parameter                                |                                             Description                                              |                Default                |
|------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------|---------------------------------------|
| metadata.resource.group                                                |                                                                                                      | <code>kubedb.com</code>               |
| metadata.resource.kind                                                 |                                                                                                      | <code>Neo4j</code>                    |
| metadata.resource.name                                                 |                                                                                                      | <code>neo4js</code>                   |
| metadata.resource.scope                                                |                                                                                                      | <code>Namespaced</code>               |
| metadata.resource.version                                              |                                                                                                      | <code>v1alpha2</code>                 |
| metadata.release.name                                                  | Release name                                                                                         | <code>""</code>                       |
| metadata.release.namespace                                             | Release namespace                                                                                    | <code>""</code>                       |
| form.alert.enabled                                                     | # Enable PrometheusRule alerts                                                                       | <code>warning</code>                  |
| form.alert.labels                                                      | # Labels for default rules                                                                           | <code>{"release":"prometheus"}</code> |
| form.alert.annotations                                                 | # Annotations for default rules                                                                      | <code>{}</code>                       |
| form.alert.additionalRuleLabels                                        | # Additional labels for PrometheusRule alerts                                                        | <code>{}</code>                       |
| form.alert.groups.database.enabled                                     |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.database.rules.neo4jHighCPUUsage.enabled             |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.neo4jHighCPUUsage.duration            |                                                                                                      | <code>"1m"</code>                     |
| form.alert.groups.database.rules.neo4jHighCPUUsage.severity            |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.database.rules.neo4jHighCPUUsage.val                 |                                                                                                      | <code>80</code>                       |
| form.alert.groups.database.rules.neo4jHighMemoryUsage.enabled          |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.neo4jHighMemoryUsage.duration         |                                                                                                      | <code>"1m"</code>                     |
| form.alert.groups.database.rules.neo4jHighMemoryUsage.severity         |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.database.rules.neo4jHighMemoryUsage.val              |                                                                                                      | <code>80</code>                       |
| form.alert.groups.database.rules.neo4jPageCacheUsageRatioHigh.enabled  |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.neo4jPageCacheUsageRatioHigh.duration |                                                                                                      | <code>"5m"</code>                     |
| form.alert.groups.database.rules.neo4jPageCacheUsageRatioHigh.severity |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.database.rules.neo4jPageCacheUsageRatioHigh.val      |                                                                                                      | <code>0.85</code>                     |
| form.alert.groups.database.rules.neo4jPageCacheHitRatioLow.enabled     |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.neo4jPageCacheHitRatioLow.duration    |                                                                                                      | <code>"5m"</code>                     |
| form.alert.groups.database.rules.neo4jPageCacheHitRatioLow.severity    |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.database.rules.neo4jPageCacheHitRatioLow.val         |                                                                                                      | <code>0.98</code>                     |
| form.alert.groups.database.rules.neo4jPageFaultsHigh.enabled           |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.neo4jPageFaultsHigh.duration          |                                                                                                      | <code>"5m"</code>                     |
| form.alert.groups.database.rules.neo4jPageFaultsHigh.severity          |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.database.rules.neo4jPageFaultsHigh.val               |                                                                                                      | <code>5000</code>                     |
| form.alert.groups.database.rules.neo4jPageFaultFailuresHigh.enabled    |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.neo4jPageFaultFailuresHigh.duration   |                                                                                                      | <code>"5m"</code>                     |
| form.alert.groups.database.rules.neo4jPageFaultFailuresHigh.severity   |                                                                                                      | <code>critical</code>                 |
| form.alert.groups.database.rules.neo4jPageFaultFailuresHigh.val        |                                                                                                      | <code>0</code>                        |
| form.alert.groups.database.rules.diskUsageHigh.enabled                 |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.diskUsageHigh.val                     |                                                                                                      | <code>80</code>                       |
| form.alert.groups.database.rules.diskUsageHigh.duration                |                                                                                                      | <code>"1m"</code>                     |
| form.alert.groups.database.rules.diskUsageHigh.severity                |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.database.rules.diskAlmostFull.enabled                |                                                                                                      | <code>true</code>                     |
| form.alert.groups.database.rules.diskAlmostFull.val                    |                                                                                                      | <code>95</code>                       |
| form.alert.groups.database.rules.diskAlmostFull.duration               |                                                                                                      | <code>"1m"</code>                     |
| form.alert.groups.database.rules.diskAlmostFull.severity               |                                                                                                      | <code>critical</code>                 |
| form.alert.groups.provisioner.enabled                                  |                                                                                                      | <code>warning</code>                  |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled           |                                                                                                      | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration          |                                                                                                      | <code>"1m"</code>                     |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity          |                                                                                                      | <code>critical</code>                 |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled           |                                                                                                      | <code>true</code>                     |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration          |                                                                                                      | <code>"15m"</code>                    |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity          |                                                                                                      | <code>warning</code>                  |
| grafana.enabled                                                        |                                                                                                      | <code>true</code>                     |
| grafana.version                                                        |                                                                                                      | <code>7.5.5</code>                    |
| grafana.jobName                                                        |                                                                                                      | <code>kubedb-databases</code>         |
| grafana.url                                                            | URL of Grafana instance where the dashboard will be imported Example: http://grafana.monitoring:3000 | <code>""</code>                       |
| grafana.apikey                                                         |                                                                                                      | <code>""</code>                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i neo4j-alerts appscode/neo4j-alerts -n kubedb --create-namespace --version=v2026.7.14 --set metadata:
  resource:
    group: kubedb.com
    kind: Neo4j
    name: neo4js
    scope: Namespaced
    version: v1alpha2
  release:
    name: neo4j
    namespace: demo

form:
  alert:
    enabled: warning
    labels:
      release: prometheus
    annotations: {}

    groups:
      database:
        enabled: true
        rules:
          neo4jDown:
            enabled: true
            duration: "30s"
            severity: critical

          neo4jPhaseCritical:
            enabled: true
            duration: "1m"
            severity: warning

          neo4jRestarted:
            enabled: true
            duration: "1m"
            severity: warning
            val: 0

          neo4jHighCPUUsage:
            enabled: true
            duration: "1m"
            severity: warning
            val: 80

          neo4jHighMemoryUsage:
            enabled: true
            duration: "1m"
            severity: warning
            val: 80

          neo4jHighPageCacheUsage:
            enabled: true
            duration: "1m"
            severity: warning
            val: 80

          neo4jTransactionFailures:
            enabled: true
            duration: "5m"
            severity: critical
            val: 5

          diskUsageHigh:
            enabled: true
            val: 80
            duration: "1m"
            severity: warning

          diskAlmostFull:
            enabled: true
            val: 95
            duration: "1m"
            severity: critical

      provisioner:
        enabled: true
        rules:
          appPhaseNotReady:
            enabled: true
            duration: "1m"
            severity: critical

          appPhaseCritical:
            enabled: true
            duration: "15m"
            severity: warning

grafana:
  enabled: true
  version: 7.5.5
  jobName: kubedb-databases
  url: "http://grafana.monitoring:3000"
  user: "admin"
  password: "your-password"

```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i neo4j-alerts appscode/neo4j-alerts -n kubedb --create-namespace --version=v2026.7.14 --values values.yaml
```
