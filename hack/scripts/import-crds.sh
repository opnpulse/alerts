#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eou pipefail

PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR_TAG=${PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR_TAG:-v0.74.0}

for dir in charts/*/; do
    dir=${dir%*/}
    dir=${dir##*/}
    # echo $dir
    crd-importer \
        --input=https://github.com/prometheus-operator/prometheus-operator/raw/${PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR_TAG}/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml \
        --annotations 'config.kubernetes.io/local-config=true' \
        --out=./charts/$dir/crds
done
