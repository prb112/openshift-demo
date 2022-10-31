# Descheduler Profile - SoftTopologyAndDuplicates - RemovePodsViolatingTopologySpreadConstraint Strategy

Before proceeding, the Descheduler Operator must be installed on a Cluster with 3 Worker Nodes.

**WARNING**
Do not run this in a LIVE cluster, this should be dedicated to the specific tests, as it will EVICT running pods every 3 minutes.
**WARNING**

> [RemovePodsViolatingTopologySpreadConstraint](https://github.com/kubernetes-sigs/descheduler#removepodsviolatingtopologyspreadconstraint): This strategy makes sure that pods violating topology spread constraints are evicted from nodes. Specifically, it tries to evict the minimum number of pods required to balance topology domains to within each constraint's maxSkew.

Since this is run as part of the SoftTopologyAndDuplicates profile there is an implication of `RemoveDuplicates` which makes it harder to exercise without three nodes and at most a 1 Pod in a Kubernetes (Deployment, ReplicaSet, DaemonSet et cetra) as when the labels change in the environment the `RemoveDuplicates` will be executed first.

Note, in Kubernetes 1.24, they introduced Built-in default constraints. [link](https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#internal-default-constraints) as such I have used `custom` instead of the node label `topology.kubernetes.io/zone`.

## Steps

0. Check you have Three Nodes

```
$ oc get nodes -lnode-role.kubernetes.io/worker
NAME                                                 STATUS   ROLES    AGE   VERSION
worker-0.xip.io   Ready    worker   17h   v1.23.5+70fb84c
worker-1.xip.io   Ready    worker   17h   v1.23.5+70fb84c
worker-2.xip.io   Ready    worker   17h   v1.23.5+70fb84c
```

1. Update the Topology Policy

```
$ oc apply -n openshift-kube-descheduler-operator -f files/4_SoftTopologyAndDuplicates_RemovePodsViolatingTopologySpreadConstraint.yml
kubedescheduler.operator.openshift.io/cluster created
```

2. Check the configmap to see the Descheduler Policy. 

```
$ oc -n openshift-kube-descheduler-operator get cm cluster -o=yaml
```

This ConfigMap should show the excluded namespaces and `strategies.RemovePodsViolatingTopologySpreadConstraint.includeSoftConstraints: true` is configured.

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

Note, you'll see `Error from server (NotFound): namespaces "test" not found` if its the first time the namespace is being created.


5. Label the zones so it's unbalanced (a,b)

a. `worker-0`

```
$ oc label node 'worker-0.xip.io' custom=a
node/worker-0.xip.io labeled
```

b. `worker-1`

```
$ oc label node 'worker-1.xip.io' custom=b
node/worker-1.xip.io labeled
```

c. `worker-2`

```
$ oc label node 'worker-2.xip.io' custom=b
node/worker-2.xip.io labeled
```

5. Cordon the worker-2 node.

```
$ oc adm cordon worker-2.xip.io
node/worker-2.xip.io cordoned
```

6. Create the ReplicaSet

a. first replicaset

```
$ oc -n test apply -f files/4_SoftTopologyAndDuplicates_RemovePodsViolatingTopologySpreadConstraint_rs_a.yml
replicaset.apps/ua created
```

b. second replicaset

```
$ oc -n test apply -f files/4_SoftTopologyAndDuplicates_RemovePodsViolatingTopologySpreadConstraint_rs_b.yml
replicaset.apps/ub created
```

c. third replicaset

```
$ oc -n test apply -f files/4_SoftTopologyAndDuplicates_RemovePodsViolatingTopologySpreadConstraint_rs_c.yml
replicaset.apps/uc created
```

7. Verify the pods distributed between the two nodes.

```
$ oc -n test get pods -o=custom-columns='Name:metadata.name,NodeName:spec.nodeName'
Name               NodeName
ua-dwwwh   worker-0.xip.io
ua-gql9r   worker-1.xip.io
ub-7j6bw   worker-0.xip.io
ub-rhhx8   worker-1.xip.io
uc-4f52z   worker-1.xip.io
uc-kkffv   worker-0.xip.io
```

8. Uncordon the worker-2 node.

```
$ oc adm uncordon worker-2.xip.io
node/worker-2.xip.io cordoned
```

9. Reassign node labels 

```
oc label node worker-0.xip.io custom=a --overwrite=true
oc label node worker-1.xip.io custom=b --overwrite=true
oc label node worker-2.xip.io custom=b --overwrite=true
```

Note, you can use `oc -n test get pods -o=custom-columns='Name:metadata.name,NodeName:spec.nodeName' | grep -v NodeName | awk '{print $NF}' | sort | uniq -c` to see the node / pod distribution.

10. Check the Pod distribution to see the rebalancing.

```
$ oc -n test get pods -o=custom-columns='Name:metadata.name,NodeName:spec.nodeName'
Name       NodeName
ua-7z859   worker-0.xip.io
ua-fs8xl   worker-2.xip.io
ub-42qlf   worker-0.xip.io
ub-d2bgz   worker-2.xip.io
uc-cxlcn   worker-1.xip.io
uc-g22db   worker-0.xip.io
```

Note, you can use `oc -n test get pods -o=custom-columns='Name:metadata.name,NodeName:spec.nodeName' | grep -v NodeName | awk '{print $NF}' | sort | uniq -c` to see the node / pod distribution.

11. Verify the logs show an eviction based on the deschedulerPodTopologySpread

```
$ oc -n openshift-kube-descheduler-operator logs -l app=descheduler  --tail=20000 | grep deschedulerPodTop
I0928 20:09:22.750546       1 event.go:294] "Event occurred" object="test/ua-kj2wx" fieldPath="" kind="Pod" apiVersion="v1" type="Normal" reason="Descheduled" message="pod evicted by sigs.k8s.io/deschedulerPodTopologySpread"
```

12. Clean up replicaset

```
$ oc -n test delete rs ua ub uc                                                   
replicaset.apps "ua" deleted
replicaset.apps "ub" deleted
replicaset.apps "uc" deleted
```

# Summary
This is a simple realignment of pods using a topology spread constraint and Soft Topology and Duplicates.


# Notes

1. Check the Node-Pod Distribution

```
$ oc -n test get pods -o=custom-columns='Name:metadata.name,NodeName:spec.nodeName' | grep -v NodeName | awk '{print $NF}'  | sort | uniq -c
   4 worker-1.xip.io
   2 worker-2.xip.io
   1 worker-0.xip.io
```

2. Check the Node Labels for a worker

```
$ oc get nodes --show-labels -lnode-role.kubernetes.io/worker
NAME                                                 STATUS   ROLES    AGE   VERSION           LABELS
worker-0.xip.io   Ready    worker   20h   v1.23.5+70fb84c   beta.kubernetes.io/arch=ppc64le,beta.kubernetes.io/os=linux,kubernetes.io/arch=ppc64le,kubernetes.io/hostname=worker-0.xip.io,kubernetes.io/os=linux,node-role.kubernetes.io/worker=,node.openshift.io/os_id=rhcos,topology.kubernetes.io/zone=b
worker-1.xip.io   Ready    worker   20h   v1.23.5+70fb84c   beta.kubernetes.io/arch=ppc64le,beta.kubernetes.io/os=linux,kubernetes.io/arch=ppc64le,kubernetes.io/hostname=worker-1.xip.io,kubernetes.io/os=linux,node-role.kubernetes.io/worker=,node.openshift.io/os_id=rhcos,node=a,topology.kubernetes.io/zone=a
worker-2.xip.io   Ready    worker   20h   v1.23.5+70fb84c   beta.kubernetes.io/arch=ppc64le,beta.kubernetes.io/os=linux,kubernetes.io/arch=ppc64le,kubernetes.io/hostname=worker-2.xip.io,kubernetes.io/os=linux,node-role.kubernetes.io/worker=,node.openshift.io/os_id=rhcos,node=a,topology.kubernetes.io/zone=b
```

# References

- https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
- https://kubernetes.io/blog/2020/05/introducing-podtopologyspread/
- https://kubernetes.io/docs/reference/labels-annotations-taints/

# Summary

This Profile shows the SoftTopologyAndDuplicates and how it causes descheduling and scheduling on the node with a Soft Topology Constraint.
