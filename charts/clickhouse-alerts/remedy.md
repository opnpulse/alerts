## TODOs on ClickHouse Critical Alerts

### Database Alerts

- #### ClickhouseInstanceDown
  - Describe the `ClickHouse` CR, check the reason from conditions and try restarting the pods
  - Contact AppsCode team
- #### ClickhouseTooManyConnections
  - Increase ClickHouse `max_connections` setting in users.xml or via KubeDB configuration
  - Check for application connection leaks
- #### ClickhouseSlowQueries
  - Check for long-running queries using `system.processes` table
  - Optimize slow queries or adjust `max_execution_time` settings
- #### ClickhouseReplicationLag
  - Check network connectivity between replicas
  - Verify ZooKeeper/ClickHouse Keeper cluster health
  - Check `system.replication_queue` for stuck tasks
- #### ClickhouseMergeErrors
  - Check `system.merges` and `system.part_log` for errors
  - Ensure sufficient disk space for merge operations
  - Check `system.errors` for relevant error codes
- #### ClickhouseExporterError
  - Check ClickHouse exporter logs for configuration errors
  - Verify exporter can connect to ClickHouse server
- #### ClickhouseRestarted
  - Check if `ClickHouse` CR is in Ready status
  - Contact AppsCode team if status is not updated.

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `ClickHouseOpsRequest` is ongoing on same database, Wait until it completes.
  - Check pod logs for errors
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If any `ClickHouseOpsRequest` is ongoing on same database, Wait until it completes.
  - Contact AppsCode team
- #### OpsRequestFailed
  - Describe the OpsRequest and Check the conditions in it
  - Contact AppsCode team

### KubeStash Alerts
- #### ClickHouseKubeStashBackupSessionFailed
  - Describe the BackupSession
  - Check the conditions in the BackupSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the BackupSession
  - View the Backup Job log
  - Check if the `INTEGRITY` of Repository is `true`
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### ClickHouseKubeStashRestoreSessionFailed
  - Describe the RestoreSession
  - Check the conditions in the RestoreSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the RestoreSession
  - View the Restore Job log
  - Check if the `INTEGRITY` of Repository is `true`
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### ClickHouseKubeStashNoBackupSessionForTooLong
  - Check if the BackupConfiguration is not `Paused`
  - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
  - Describe the BackupConfiguration
  - Check the conditions of BackupConfiguration
  - Check the reasons of the `false` conditions (if any)
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### ClickHouseKubeStashRepositoryCorrupted
  - Check if the `INTEGRITY` of `repository` is `true`
  - Contact AppsCode team
- #### ClickHouseKubeStashRepositoryStorageRunningLow
  - Increase the volume size of `repository` backend
  - Update RetentionPolicy to free up storage
- #### ClickHouseKubeStashBackupSessionPeriodTooLong | ClickHouseKubeStashRestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of `repository` is `true`
  - Check the `ClickHouse` CRs status
  - Contact AppsCode team
