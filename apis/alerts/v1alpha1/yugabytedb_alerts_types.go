/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	api "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

const (
	ResourceKindYugabytedbAlerts = "YugabytedbAlerts"
	ResourceYugabytedbAlerts     = "yugabytedbalerts"
	ResourceYugabytedbAlertss    = "yugabytedbalertss"
)

// YugabytedbAlerts defines the schema for YugabyteDB Alerting Rules

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=yugabytedbalertss,singular=yugabytedbalerts,categories={kubedb,appscode}
type YugabytedbAlerts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              YugabytedbAlertsSpec `json:"spec,omitempty"`
}

// YugabytedbAlertsSpec is the schema for YugabyteDB alerts chart values file
type YugabytedbAlertsSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         YugabytedbAlertsSpecForm `json:"form"`
	Grafana      Grafana                  `json:"grafana"`
}

type YugabytedbAlertsSpecForm struct {
	Alert YugabytedbAlert `json:"alert"`
}

type YugabytedbAlert struct {
	Enabled mona.SeverityFlag `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string     `json:"additionalRuleLabels"`
	Groups               YugabytedbAlertGroups `json:"groups"`
}

type YugabytedbAlertGroups struct {
	Database    YugabytedbDatabaseAlert `json:"database"`
	Provisioner ProvisionerAlert        `json:"provisioner"`
	OpsManager  OpsManagerAlert         `json:"opsManager"`
	KubeStash   KubeStashAlert          `json:"kubeStash"`
}

type YugabytedbDatabaseAlert struct {
	Enabled mona.SeverityFlag            `json:"enabled"`
	Rules   YugabytedbDatabaseAlertRules `json:"rules"`
}

// YugabytedbDatabaseAlertRules. Every rule is scoped to the stats Service of one universe
// (master, tserver and, when configured, read replica pods).
type YugabytedbDatabaseAlertRules struct {
	// YugabyteDBInstanceDown fires when a master, tserver or read replica pod's metrics endpoint
	// stops responding.
	YugabyteDBInstanceDown FixedAlert `json:"yugabyteDBInstanceDown"`
	// YugabyteDBRestarted fires shortly after any pod restarts.
	YugabyteDBRestarted IntValAlert `json:"yugabyteDBRestarted"`
	// YugabyteDBTServerDead fires when a master reports a tserver as dead (missed its heartbeat
	// window, FLAGS_tserver_unresponsive_timeout_ms).
	YugabyteDBTServerDead FixedAlert `json:"yugabyteDBTServerDead"`
	// YugabyteDBClockSkewHigh fires when a server's clock skew approaches the universe's own
	// max_clock_skew_usec (500000us / 500ms by default; above it the server rejects operations).
	YugabyteDBClockSkewHigh IntValAlert `json:"yugabyteDBClockSkewHigh"`
	// YugabyteDBFollowerLagHigh fires when a tablet peer has gone too long without an
	// UpdateConsensus request from its leader.
	YugabyteDBFollowerLagHigh IntValAlert `json:"yugabyteDBFollowerLagHigh"`
	// YugabyteDBMemoryPressureRejectionsHigh fires when a tserver rejects operations for memory
	// pressure (operation, follower or leader rejections).
	YugabyteDBMemoryPressureRejectionsHigh IntValAlert `json:"yugabyteDBMemoryPressureRejectionsHigh"`
	// YugabyteDBYSQLConnMgrQueueHigh fires when YSQL clients queue in the Connection Manager
	// waiting for a physical connection.
	YugabyteDBYSQLConnMgrQueueHigh IntValAlert `json:"yugabyteDBYSQLConnMgrQueueHigh"`
	// YugabyteDBHighCPUUsage fires on high pod CPU usage against its requests.
	YugabyteDBHighCPUUsage IntValAlert `json:"yugabyteDBHighCPUUsage"`
	// YugabyteDBHighMemoryUsage fires on high pod memory usage against its limits.
	YugabyteDBHighMemoryUsage IntValAlert `json:"yugabyteDBHighMemoryUsage"`
	// DiskUsageHigh fires on high persistent volume usage.
	DiskUsageHigh IntValAlert `json:"diskUsageHigh"`
	// DiskAlmostFull fires near YugabyteDB's own 5 percent free space floor, below which the
	// universe starts rejecting writes.
	DiskAlmostFull IntValAlert `json:"diskAlmostFull"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// YugabytedbAlertsList is a list of YugabytedbAlertss
type YugabytedbAlertsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of YugabytedbAlerts CRD objects
	Items []YugabytedbAlerts `json:"items,omitempty"`
}
