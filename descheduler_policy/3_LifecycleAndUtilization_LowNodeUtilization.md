# Descheduler Profile - LifecycleAndUtilization - LowNodeUtilization Strategy

Before proceeding, the Descheduler Operator must be installed.

**WARNING**
Do not run this in a LIVE cluster, this should be dedicated to the specific tests, as it will EVICT running pods every 1 minute when the Pods are older than `5m`.
**WARNING**

> [LowNodeUtilization](https://github.com/kubernetes-sigs/descheduler/tree/master#lownodeutilization): All types of pods with the annotation descheduler.alpha.kubernetes.io/evict are eligible for eviction. This annotation is used to override checks which prevent eviction and users can select which pod is evicted. Users should know how and if the pod will be recreated.

## Steps 

1. Update the LifecycleAndUtilization Policy

```
$ oc apply -n openshift-kube-descheduler-operator -f files/3_LifecycleAndUtilization_LowNodeUtilization.yml
kubedescheduler.operator.openshift.io/cluster created
```

2. Check the configmap to see the Descheduler Policy. 

```
$ oc -n openshift-kube-descheduler-operator get cm cluster -o=yaml
```

This ConfigMap should show the excluded namespaces and the nodeResourceUtilizationThresholds.

```
nodeResourceUtilizationThresholds:
    targetThresholds:
      cpu: 50
      memory: 50
      pods: 50
    thresholds:
      cpu: 20
      memory: 20
      pods: 20
```

3. Check the descheduler cluster 

```
$ oc -n openshift-kube-descheduler-operator logs -l app=descheduler 
```

This log should show a started Descheduler.

4. Create a test namespace

```
$ oc get namespace test || oc create namespace test
namespace/test created
```

5. [Cordon](https://docs.openshift.com/container-platform/4.11/nodes/nodes/nodes-nodes-working.html) one of the workers so we unbalance the number of assigned pods. 

a. List the nodes

```
$ oc get nodes
```

b. Select a worker node, such as `worker-1.rdr-rhop.sslip.io` 

c. Cordon the node

```
$ oc adm cordon worker-1.rdr-rhop.sslip.io
node/worker-1.rdr-rhop.sslip.io cordoned
```

Note, if you have more than 2 worker nodes, cordon N-1 nodes.
You may also have to figure out which node has the MOST pods and cordon all of the other nodes.

6. Create a deployment

```
$ oc -n test apply -f files/3_LifecycleAndUtilization_LowNodeUtilization_dp.yml
```

7. Check the pods are all on the other node 

```
$ oc -n test get pods -o=custom-columns='DATA:metadata.name,DATA:spec.nodeName'
DATA                          DATA
unbalanced-6d757874c4-5bsml worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-bjxfr worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-bsmqk worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-ccns4 worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-kvx7q worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-ldvfn worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-m5tgb worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-nlfkm worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-nrxwg worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-rdkmx worker-0.rdr-rhop.sslip.io
```

8. Uncordon the original worker.

```
$ oc adm uncordon worker-1.rdr-rhop.sslip.io              
node/worker-1.rdr-rhop.sslip.io uncordoned
```

Note, you might have to uncordon all nodes.

9. Check the Logs until we see the processing

```
$ oc -n openshift-kube-descheduler-operator logs -l app=descheduler  --tail=200
I0511 20:53:59.692336       1 descheduler.go:287] "Number of evicted pods" totalEvicted=5
```

You might see the following

```
I1004 17:12:38.927608       1 lownodeutilization.go:118] "Criteria for a node under utilization" CPU=20 Mem=20 Pods=20
I1004 17:12:38.927620       1 lownodeutilization.go:119] "Number of underutilized nodes" totalNumber=2
I1004 17:12:38.927632       1 lownodeutilization.go:132] "Criteria for a node above target utilization" CPU=50 Mem=50 Pods=50
I1004 17:12:38.927642       1 lownodeutilization.go:133] "Number of overutilized nodes" totalNumber=0
I1004 17:12:38.927652       1 lownodeutilization.go:151] "All nodes are under target utilization, nothing to do here"
I1004 17:12:38.927664       1 descheduler.go:304] "Number of evicted pods" totalEvicted=0
```

It indicates you need to raise the memory utilization.

10. Check the pods are now redistributed. 

```
$ oc -n test get pods -o=custom-columns='DATA:metadata.name,DATA:spec.nodeName'
unbalanced-6d757874c4-5bsml worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-bjxfr worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-bsmqk worker-1.rdr-rhop.sslip.io
unbalanced-6d757874c4-ccns4 worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-kvx7q worker-1.rdr-rhop.sslip.io
unbalanced-6d757874c4-ldvfn worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-m5tgb worker-0.rdr-rhop.sslip.io
unbalanced-6d757874c4-nlfkm worker-1.rdr-rhop.sslip.io
unbalanced-6d757874c4-nrxwg worker-1.rdr-rhop.sslip.io
unbalanced-6d757874c4-rdkmx worker-1.rdr-rhop.sslip.io
```

11. Delete the deployment

```
oc -n test delete deployment.apps/unbalanced
deployment.apps "unbalanced" deleted
```

## Summary

This Profile shows the LowNodeUtilization and how it causes descheduling and scheduling on the uncordoned *or* underutilized node.

# Notes

1. Check the Node-Pod Distribution

```
$ $ oc get pods -A -o=custom-columns='Name:metadata.name,NodeName:spec.nodeName' | grep -v NodeName | awk '{print $NF}'  | sort | uniq -c
   4 master-1.xip.io
   2 master-2.xip.io
   1 master-0.xip.io
   4 worker-1.xip.io
   2 worker-2.xip.io
   1 worker-0.xip.io
```