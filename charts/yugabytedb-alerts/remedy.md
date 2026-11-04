## TODOs on YugabyteDB Critical Alerts

### Database Alerts

- #### YugabyteDBInstanceDown
  - Check the ServiceMonitor and the `<name>-stats` service endpoints.
  - Describe the `YugabyteDB` CR and inspect status conditions.
  - Check the master/tserver pod logs.
- #### YugabyteDBRestarted
  - Check pod events and previous container logs.
  - Verify whether an OpsRequest, a lost volume, or resource pressure caused the restart.
- #### YugabyteDBTServerDead
  - Check whether the pod is Running; a dead tserver missed its heartbeat window
    (FLAGS_tserver_unresponsive_timeout_ms) on the masters.
  - If the pod lost its volume, the operator replaces a master automatically (see
    `pkg/controller` master-replacement logic); a dead tserver pod is recreated by its PetSet.
  - Check network connectivity between the masters and the tserver.
- #### YugabyteDBClockSkewHigh
  - Check NTP/chrony on the node the pod runs on.
  - If skew keeps climbing towards max_clock_skew_usec (500ms default), the server will start
    rejecting operations; fix the node's clock source before that happens.
- #### YugabyteDBFollowerLagHigh
  - Check whether the leader for the lagging tablet is under load or network-partitioned.
  - A consistently lagging follower may need its pod restarted, or its node investigated.
- #### YugabyteDBMemoryPressureRejectionsHigh
  - Check the tserver's memory usage against its limit.
  - Increase memory with a VerticalScaling YugabyteDBOpsRequest, or reduce load.
- #### YugabyteDBYSQLConnMgrQueueHigh
  - Check `ysql_conn_mgr_max_client_connections` and the application's connection pool size.
  - Increase `spec.ysql.connectionManager.maxClientConnections`, or raise the physical server
    count via HorizontalScaling.
- #### YugabyteDBHighCPUUsage / YugabyteDBHighMemoryUsage
  - Check query and tablet-compaction workload.
  - Scale resources with a VerticalScaling YugabyteDBOpsRequest.
- #### DiskUsageHigh / DiskAlmostFull
  - Expand storage with a VolumeExpansion YugabyteDBOpsRequest.
  - YugabyteDB itself starts rejecting writes below 5 percent free; treat DiskAlmostFull as urgent.

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If a `YugabyteDBOpsRequest` is ongoing on the same database, wait until it completes.
  - If some masters or tservers of the universe are not `Up`, check their pod logs and events.
  - A failed OpsRequest is not resumed automatically by KubeDB; it needs a manual look
    (see `ticket-template.md`'s [get pod logs](#get-pod-logs) section).
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If another `YugabyteDBOpsRequest` is ongoing on the same database, wait until it completes.
  - Add/remove TLS requests have a known, non-zero-downtime window while the universe mixes TLS
    and plaintext servers; that alone is not a bug.
  - Contact AppsCode team
- #### OpsRequestFailed
  - KubeDB OpsRequests do not self-heal from a failure; the on-call recovers by hand.
  - Check the OpsRequest's status conditions and the ops-manager log for the step that failed.
  - Contact AppsCode team

### KubeStash Alerts
- #### YugabyteDBKubeStashBackupSessionFailed
  - Describe the BackupSession
  - Check the conditions in the BackupSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the BackupSession
  - View the backup Job log
  - Check if the `INTEGRITY` of the Repository is `true`
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### YugabyteDBKubeStashRestoreSessionFailed
  - Describe the RestoreSession
  - Check the conditions in the RestoreSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the RestoreSession
  - View the restore Job log
  - Check if the `INTEGRITY` of the Repository is `true`
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### YugabyteDBKubeStashNoBackupSessionForTooLong
  - Check if the BackupConfiguration is not `Paused`
  - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
  - Describe the BackupConfiguration
  - Check the conditions of the BackupConfiguration
  - Check the reasons of the `false` conditions (if any)
  - Check the KubeStash operator log
  - Contact AppsCode team
- #### YugabyteDBKubeStashRepositoryCorrupted
  - Check if the `INTEGRITY` of the `repository` is `true`
  - Contact AppsCode team
- #### YugabyteDBKubeStashRepositoryStorageRunningLow
  - Increase the volume size of the `repository` backend
  - Update the RetentionPolicy to free up storage
- #### YugabyteDBKubeStashBackupSessionPeriodTooLong | YugabyteDBKubeStashRestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of the `repository` is `true`
  - Check the `YugabyteDB` CR's status
  - Contact AppsCode team
