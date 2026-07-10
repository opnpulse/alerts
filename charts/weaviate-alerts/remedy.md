## TODOs on Weaviate Critical Alerts

### Database Alerts

- #### WeaviateInstanceDown
  - Check the ServiceMonitor and the `weaviate-stats` service endpoints.
  - Describe the `Weaviate` CR and inspect status conditions.
  - Check Weaviate pod logs.
- #### WeaviateRestarted
  - Check pod events and previous container logs.
  - Verify whether an OpsRequest, rollout, or resource pressure caused the restart.
- #### WeaviateHighCPUUsage
  - Check query and import workload.
  - Scale replicas or increase CPU requests/limits using a WeaviateOpsRequest.
- #### WeaviateHighMemoryUsage / WeaviateHighProcessMemoryUsage
  - Check import/query workload and memory-heavy modules.
  - Increase memory limits using a WeaviateOpsRequest if pressure persists.
- #### WeaviateGoroutinesExplosion / WeaviateHighThreadPressure / WeaviateHighFDsUsage
  - Check API traffic, client connection behavior, and pod logs.
  - Verify there is no repeated retry loop from clients or controllers.
- #### DiskUsageHigh / DiskAlmostFull
  - Expand storage with a Weaviate volume expansion OpsRequest.
  - Remove unnecessary data only after confirming retention and backup requirements.
- #### WeaviateHTTPErrorRateHigh / WeaviateHTTPP95LatencyHigh
  - Check HTTP request routes with high error or latency in the Weaviate dashboard.
  - Inspect application client errors and Weaviate pod logs.
- #### WeaviateGRPCErrorRateHigh / WeaviateGRPCP95LatencyHigh
  - Check gRPC service and method labels in the Weaviate dashboard.
  - Inspect client retries and Weaviate pod logs.
- #### WeaviateReplicationEngineDown
  - Check whether replication is enabled and expected for the cluster.
  - Inspect pod logs for replication engine startup or shutdown errors.
  - Verify peer connectivity between Weaviate pods.
- #### WeaviateReplicationQueueHigh / WeaviateReplicationFailuresHigh
  - Check cluster health, peer connectivity, and replication metrics by node.
  - Inspect pod logs for replication or network errors.

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Describe the `Weaviate` CR and inspect status conditions.
  - Check KubeDB provisioner logs and Weaviate pod events.
- #### AppPhaseCritical
  - If a WeaviateOpsRequest is progressing, wait for it to complete or inspect its conditions.
  - Check pods that are not Ready and review their events and logs.
