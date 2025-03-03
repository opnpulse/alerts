## TODOs on MSSQLServer Critical Alerts

### Database Alerts

- #### MSSQLServerInstanceDown
  - Describe the `MSSQLServer` CR, check the reason from conditions and try restarting the pods
  - Check the KubeDB Operator logs for errors
  - Check pod status: `kubectl describe pod <mssql-pod> -n <namespace>`
  - Verify resource limits (CPU/Memory)
  - Command: `kubectl get mssqlserver <name> -n <namespace> -o yaml`
  - Contact AppsCode team
- #### MSSQLServerServiceDown
  - Describe the `MSSQLServer` CR and Try restarting all the pods
  - Check the KubeDB Operator logs for errors
  - Check pods status: `kubectl describe pod <mssql-pod> -n <namespace>`
  - Contact AppsCode team
- #### MSSQLServerTooManyConnections
  - Increase MSSQLServer max connections
- #### MSSQLServerRestarted
  - Check if `MSSQLServer` CR is in Ready status
  - Contact AppsCode team if status is not updated.

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `MSSQLServerOpsRequest` is ongoing on same database, Wait until it completes.
  - check for `mssql-coordinator` logs on the node that is not connected with `Primary` node.
  - If some nodes of the MSSQLServer Cluster are not `Up`, Try restarting those nodes one at a time.
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If any `MSSQLServerSQLOpsRequest` is ongoing on same database, Wait until it completes. For `UpdateVersion` or some other `OpsRequest` It may take longer then expectation.
  - Contact AppsCode team
- #### OpsRequestFailed
  - Describe the OpsRequest and Check the conditions in it
  - Contact AppsCode team

