# VaultServer alerts

[VaultServer alerts by AppsCode](https://github.com/appscode/alerts) - VaultServer alerts for KubeVault

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/vaultserver-alerts --version=v2026.7.14
$ helm upgrade -i vaultserver appscode/vaultserver-alerts -n demo --create-namespace --version=v2026.7.14
```

## Introduction

This chart deploys VaultServer alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `vaultserver`:

```bash
$ helm upgrade -i vaultserver appscode/vaultserver-alerts -n demo --create-namespace --version=v2026.7.14
```

The command deploys VaultServer alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `vaultserver`:

```bash
$ helm uninstall vaultserver -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `vaultserver-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                Default                |
|-------------------------------------------------------------------------------|-----------------------------------------------|---------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubevault.com</code>            |
| metadata.resource.kind                                                        |                                               | <code>VaultServer</code>              |
| metadata.resource.name                                                        |                                               | <code>vaultserver</code>              |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>               |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                 |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                       |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                       |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                  |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"prometheus"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                       |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                       |
| form.alert.groups.vault.enabled                                               |                                               | <code>warning</code>                  |
| form.alert.groups.vault.rules.vaultDown.enabled                               |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultDown.duration                              |                                               | <code>"0m"</code>                     |
| form.alert.groups.vault.rules.vaultDown.severity                              |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultSealed.enabled                             |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultSealed.duration                            |                                               | <code>"0m"</code>                     |
| form.alert.groups.vault.rules.vaultSealed.severity                            |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultAutoPilotNodeUnhealthy.enabled             |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultAutoPilotNodeUnhealthy.duration            |                                               | <code>"1m"</code>                     |
| form.alert.groups.vault.rules.vaultAutoPilotNodeUnhealthy.severity            |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultLeadershipLoss.enabled                     |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultLeadershipLoss.duration                    |                                               | <code>"1m"</code>                     |
| form.alert.groups.vault.rules.vaultLeadershipLoss.val                         |                                               | <code>5</code>                        |
| form.alert.groups.vault.rules.vaultLeadershipLoss.severity                    |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultLeadershipStepsDowns.enabled               |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultLeadershipStepsDowns.duration              |                                               | <code>"1m"</code>                     |
| form.alert.groups.vault.rules.vaultLeadershipStepsDowns.val                   |                                               | <code>5</code>                        |
| form.alert.groups.vault.rules.vaultLeadershipStepsDowns.severity              |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultLeadershipSetupFailures.enabled            |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultLeadershipSetupFailures.duration           |                                               | <code>"1m"</code>                     |
| form.alert.groups.vault.rules.vaultLeadershipSetupFailures.val                |                                               | <code>5</code>                        |
| form.alert.groups.vault.rules.vaultLeadershipSetupFailures.severity           |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultRequestFailures.enabled                    |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultRequestFailures.duration                   |                                               | <code>"15m"</code>                    |
| form.alert.groups.vault.rules.vaultRequestFailures.severity                   |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultResponseFailures.enabled                   |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultResponseFailures.duration                  |                                               | <code>"15m"</code>                    |
| form.alert.groups.vault.rules.vaultResponseFailures.severity                  |                                               | <code>critical</code>                 |
| form.alert.groups.vault.rules.vaultTooManyInfinityTokens.enabled              |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.vaultTooManyInfinityTokens.duration             |                                               | <code>"5m"</code>                     |
| form.alert.groups.vault.rules.vaultTooManyInfinityTokens.val                  |                                               | <code>3</code>                        |
| form.alert.groups.vault.rules.vaultTooManyInfinityTokens.severity             |                                               | <code>warning</code>                  |
| form.alert.groups.vault.rules.diskUsageHigh.enabled                           |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.diskUsageHigh.val                               |                                               | <code>80</code>                       |
| form.alert.groups.vault.rules.diskUsageHigh.duration                          |                                               | <code>"1m"</code>                     |
| form.alert.groups.vault.rules.diskUsageHigh.severity                          |                                               | <code>warning</code>                  |
| form.alert.groups.vault.rules.diskAlmostFull.enabled                          |                                               | <code>true</code>                     |
| form.alert.groups.vault.rules.diskAlmostFull.val                              |                                               | <code>95</code>                       |
| form.alert.groups.vault.rules.diskAlmostFull.duration                         |                                               | <code>"1m"</code>                     |
| form.alert.groups.vault.rules.diskAlmostFull.severity                         |                                               | <code>critical</code>                 |
| form.alert.groups.operator.enabled                                            |                                               | <code>warning</code>                  |
| form.alert.groups.operator.rules.appPhaseNotReady.enabled                     |                                               | <code>true</code>                     |
| form.alert.groups.operator.rules.appPhaseNotReady.duration                    |                                               | <code>"5m"</code>                     |
| form.alert.groups.operator.rules.appPhaseNotReady.severity                    |                                               | <code>critical</code>                 |
| form.alert.groups.operator.rules.appPhaseCritical.enabled                     |                                               | <code>true</code>                     |
| form.alert.groups.operator.rules.appPhaseCritical.duration                    |                                               | <code>"15m"</code>                    |
| form.alert.groups.operator.rules.appPhaseCritical.severity                    |                                               | <code>warning</code>                  |
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
| form.alert.groups.stash.enabled                                               |                                               | <code>warning</code>                  |
| form.alert.groups.stash.rules.backupSessionFailed.enabled                     |                                               | <code>true</code>                     |
| form.alert.groups.stash.rules.backupSessionFailed.duration                    |                                               | <code>"0m"</code>                     |
| form.alert.groups.stash.rules.backupSessionFailed.severity                    |                                               | <code>critical</code>                 |
| form.alert.groups.stash.rules.restoreSessionFailed.enabled                    |                                               | <code>true</code>                     |
| form.alert.groups.stash.rules.restoreSessionFailed.duration                   |                                               | <code>"0m"</code>                     |
| form.alert.groups.stash.rules.restoreSessionFailed.severity                   |                                               | <code>critical</code>                 |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.enabled               |                                               | <code>true</code>                     |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.duration              |                                               | <code>"0m"</code>                     |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.val                   |                                               | <code>18000</code>                    |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.severity              |                                               | <code>warning</code>                  |
| form.alert.groups.stash.rules.repositoryCorrupted.enabled                     |                                               | <code>true</code>                     |
| form.alert.groups.stash.rules.repositoryCorrupted.duration                    |                                               | <code>"5m"</code>                     |
| form.alert.groups.stash.rules.repositoryCorrupted.severity                    |                                               | <code>critical</code>                 |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.enabled             |                                               | <code>true</code>                     |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.duration            |                                               | <code>"5m"</code>                     |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.val                 |                                               | <code>10737418240 # 10GB</code>       |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.severity            |                                               | <code>warning</code>                  |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.enabled              |                                               | <code>true</code>                     |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.duration             |                                               | <code>"0m"</code>                     |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.val                  |                                               | <code>1800 # 30 minute</code>         |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.severity             |                                               | <code>warning</code>                  |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.enabled             |                                               | <code>true</code>                     |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.duration            |                                               | <code>"0m"</code>                     |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.val                 |                                               | <code>1800 # 30 minute</code>         |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.severity            |                                               | <code>warning</code>                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i vaultserver appscode/vaultserver-alerts -n demo --create-namespace --version=v2026.7.14 --set metadata.resource.group=kubevault.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i vaultserver appscode/vaultserver-alerts -n demo --create-namespace --version=v2026.7.14 --values values.yaml
```
