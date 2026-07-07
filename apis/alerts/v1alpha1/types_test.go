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

package v1alpha1_test

import (
	"os"
	"testing"

	"go.appscode.dev/alerts/apis/alerts/v1alpha1"

	sc "kmodules.xyz/schema-checker"
)

func TestDefaultValues(t *testing.T) {
	checker := sc.New(os.DirFS("../../.."),
		sc.TestCase{Obj: v1alpha1.B3AlertsSpec{}, File: "charts/b3-alerts/values.yaml"},
		sc.TestCase{Obj: v1alpha1.CassandraAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.ClickhouseAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.ConnectClusterAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.DruidAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.ElasticsearchAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.HanadbAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.HazelcastAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.IgniteAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.KafkaAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.MariadbAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.MemcachedAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.MilvusAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.MongodbAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.MSSQLServerAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.MysqlAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.NATSAlertsSpec{}, File: "charts/nats-alerts/values.yaml"},
		sc.TestCase{Obj: v1alpha1.Neo4jAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.OpenFGAAlertsSpec{}, File: "charts/openfga-alerts/values.yaml"},
		sc.TestCase{Obj: v1alpha1.OracleAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.PerconaxtradbAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.PgbouncerAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.PgpoolAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.PostgresAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.ProxysqlAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.QdrantAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.RabbitmqAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.RedisAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.SinglestoreAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.SolrAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.WeaviateAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.VaultserverAlertsSpec{}},
		sc.TestCase{Obj: v1alpha1.ZookeeperAlertsSpec{}},
	)
	checker.TestAll(t)
}
