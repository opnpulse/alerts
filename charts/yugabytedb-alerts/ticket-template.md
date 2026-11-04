This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern. To troubleshoot and find what to do let's get started with the database phase.

Let's say you have yugabytedb database in namespace demo.
```bash
Kubectl get yugabytedb -n <ns>     // will list all the database in a namesapce  
$ kubectl get yugabytedb -n demo
  NAME     VERSION    STATUS   AGE
  yb-tls   2025.2.8   Ready    6h51m
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

# Phases
## Ready
KubeDB Managed Database phase becomes Ready when every master and tserver (and, when configured, every read replica tserver) is in the universe and working properly. There's nothing to worry about.

## Provisioning
Usually, the Database Phase is set to `Provisioning` while bootstrapping for the first time. If you find the database phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the yugabytedb object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- [Get the yugabytedb object](#get-yugabytedb)
- [Describe yugabytedb object](#describe-yugabytedb)
- Describe the PetSet objects: there is one for the masters, one for the tservers, and (when
  configured) one for the read replica tservers. [List petsets under a yugabytedb](#list-petsets-under-a-yugabytedb),
  then [Describe petsets](#describe-petsets) for each of them.
- Describe the pods: if there are multiple pods, describe all of them. To do that, first
  [List pods under a yugabytedb](#list-pods-under-a-yugabytedb). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the auth secret is created](#check-secret)
- [Check pod logs](#get-pod-logs), for both the master and the tserver role, and the
  [init container's logs](#yugabytedb-init-logs) if spec.init.script is set.
- [Operator logs](#operator-logs)

## Critical
Database Phase Critical means some of the masters or tservers are not in the universe or are failing to synchronize with it.
The reasons could be a master or tserver leaving the universe, maybe for a restart, a lost volume, or an unexpected kill.
To resolve this, we need to find out which master or tserver pod is not in the universe by checking the logs, describing the database object, or querying a working server.

You can contact to AppsCode with the following things attached,
- [Get the yugabytedb object](#get-yugabytedb)
- [Describe yugabytedb object](#describe-yugabytedb)
- Describe the PetSet objects: there is one for the masters, one for the tservers, and (when
  configured) one for the read replica tservers. [List petsets under a yugabytedb](#list-petsets-under-a-yugabytedb),
  then [Describe petsets](#describe-petsets) for each of them.
- Describe the pods: if there are multiple pods, describe all of them. To do that, first
  [List pods under a yugabytedb](#list-pods-under-a-yugabytedb). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the auth secret is created](#check-secret)
- [Check pod logs](#get-pod-logs), for both the master and the tserver role.
- [Operator logs](#operator-logs)

## NotReady
Database Phase NotReady means none of the masters or tservers are working properly. There could be several possible reasons for that: misconfiguration,
a server being killed, consensus/replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from the operator and the pod containers, and describing the yugabytedb object and pods, is a recommended way to start debugging.
By KubeDB convention, a failed operation (OpsRequest) is not resumed automatically; recovery is done by hand, so human intervention may be needed to know the best way to resolve it.

You can contact to AppsCode with the following things attached,
- [Get the yugabytedb object](#get-yugabytedb)
- [Describe yugabytedb object](#describe-yugabytedb)
- Describe the PetSet objects: there is one for the masters, one for the tservers, and (when
  configured) one for the read replica tservers. [List petsets under a yugabytedb](#list-petsets-under-a-yugabytedb),
  then [Describe petsets](#describe-petsets) for each of them.
- Describe the pods: if there are multiple pods, describe all of them. To do that, first
  [List pods under a yugabytedb](#list-pods-under-a-yugabytedb). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the auth secret is created](#check-secret)
- [Check pod logs](#get-pod-logs), for both the master and the tserver role.
- [Operator logs](#operator-logs)

# How to ?
## Get yugabytedb
```bash
kubectl get yugabytedb -n <namespace> -oyaml <yugabytedb-object-Name>
```

## Describe yugabytedb
```bash
kubectl describe yugabytedb -n <namespace> <yugabytedb-object-Name>
```

## List petsets under a yugabytedb
```bash
kubectl get petsets -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=yugabytedbs.kubedb.com,app.kubernetes.io/instance=<yugabytedb-object-Name>
```

## Describe petsets
```bash
kubectl describe petsets -n <namespace> <petsets-Name>
```

## Check secret
```bash
# Get the secret name
kubectl get yugabytedb -n <namespace> <yugabytedb-object-Name> -o jsonpath="{.spec.authSecret.name}"
# Check secret's existence
kubectl get secrets -n <namespace> <yugabytedb-secret-name>
```

## List pods under a yugabytedb
```bash
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=yugabytedbs.kubedb.com,app.kubernetes.io/instance=<yugabytedb-object-Name>
```

## Describe pods
```bash
kubectl describe pods -n <namespace> <pod-Name>
```

## Get pod logs
```bash
# A master pod
kubectl logs -n <namespace> <yugabytedb-object-Name>-master-<ordinal> -c yb-master
# A tserver pod (or, for the read replica tservers, <yugabytedb-object-Name>-readreplica-<ordinal>)
kubectl logs -n <namespace> <yugabytedb-object-Name>-tserver-<ordinal> -c yb-tserver
```

## YugabyteDB init logs
Only present when spec.init.script is set: the Job that runs the initialization script once.
```bash
kubectl get jobs -n <namespace> -l=app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/instance=<yugabytedb-object-Name>
kubectl logs -n <namespace> <job-pod-name> -c yb-init
```

## Operator logs
```bash
kubectl logs -n <kubedb-ns> <provisioner-pod-name>
kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
```
