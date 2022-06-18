1. Confirm the labels for the MachineConfigPool for the `worker` nodes

```
$ oc get machineconfigpools worker -o jsonpath="{ .metadata.labels }" | jq
{
  "machineconfiguration.openshift.io/mco-built-in": "",
  "pools.operator.machineconfiguration.openshift.io/worker": ""
}
```

2. Label the nodes with worker-0

```
$ oc label mcp/worker cpumanager=enabled
machineconfigpool.machineconfiguration.openshift.io/worker labeled
```

3. Get the nodes for the workers

```
$ oc get nodes -l node-role.kubernetes.io/worker=
NAME                                                    STATUS   ROLES    AGE   VERSION
lon06-worker-0.ocp-topman-7537.158.176.147.235.xip.io   Ready    worker   11d   v1.23.5+3afdacb
lon06-worker-1.ocp-topman-7537.158.176.147.235.xip.io   Ready    worker   11d   v1.23.5+3afdacb
```

4. Label the node worker-0

```
$ oc label node/lon06-worker-0.ocp-topman-7537.158.176.147.235.xip.io cpumanager=enabled
node/lon06-worker-0.ocp-topman-7537.158.176.147.235.xip.io labeled
```

5. Label the node worker-1

```
$ oc label node/lon06-worker-1.ocp-topman-7537.158.176.147.235.xip.io cpumanager=enabled
node/lon06-worker-1.ocp-topman-7537.158.176.147.235.xip.io labeled
```

6. Create the KubeletConfig that enables `single-numa-node`

```
$ oc apply -f files/1_Topology_Manager_with_Hugepages_KubeletConfig.yml
kubeletconfig.machineconfiguration.openshift.io/cpumanager-enabled created
```

7. Wait for an update to the MachineConfig

```
$ oc wait mcp/worker --for condition=updated --timeout=25m
```

8. Get the nodes for the control plane

```
$ oc get nodes -l node-role.kubernetes.io/master=
NAME                                                    STATUS   ROLES    AGE   VERSION
lon06-master-0.ocp-topman-7537.158.176.147.235.xip.io   Ready    master   11d   v1.23.5+3afdacb
lon06-master-1.ocp-topman-7537.158.176.147.235.xip.io   Ready    master   11d   v1.23.5+3afdacb
lon06-master-2.ocp-topman-7537.158.176.147.235.xip.io   Ready    master   11d   v1.23.5+3afdacb
```

9. Select one of the Nodes to run the debug container on. (This can take a few moments to complete)

```
$ oc debug node/lon06-master-0.ocp-topman-7537.158.176.147.235.xip.io
sh-4.2# cat /host/etc/kubernetes/kubelet.conf | grep cpuManager
sh-4.2#
```

You shouldn't see any output. 

10. Get the nodes for the workers

```
$ oc get nodes -l node-role.kubernetes.io/worker=
NAME                                                    STATUS   ROLES    AGE   VERSION
lon06-worker-0.ocp-topman-7537.158.176.147.235.xip.io   Ready    worker   11d   v1.23.5+3afdacb
lon06-worker-1.ocp-topman-7537.158.176.147.235.xip.io   Ready    worker   11d   v1.23.5+3afdacb
```

11. Select one of the Nodes to run the debug container on. (This can take a few moments to complete)

```
$ oc debug node/lon06-worker-0.ocp-topman-7537.158.176.147.235.xip.io
sh-4.2# cat /host/etc/kubernetes/kubelet.conf | grep cpuManager -A1
  "cpuManagerPolicy": "static",
  "cpuManagerReconcilePeriod": "5s",
  "topologyManagerPolicy": "single-numa-node",
sh-4.2#
```

You should see the above three lines which confirms the Topology Manager is activated for the MachineConfigPool. 

12. Create the CPU 

```
$ oc apply -f files/1_Pod.yaml
pod/cpumanager created
```

13. Review the Pods for `QoS Class: Guaranteed` and `Node`

```
$ oc describe pod/cpumanager
Node:         lon06-worker-1.ocp-topman-7537.158.176.147.235.xip.io/192.168.100.151
QoS Class:                   Guaranteed
```

14. Using the Node from the previous step

```
$ oc debug node/lon06-worker-1.ocp-topman-7537.158.176.147.235.xip.io
# 
```

15. Change the root

```
$ chroot /host
```

16. Find the kubepods slice and crio id.

```
$ systemctl status kubepods.slice | grep ' /pause' -B2
           └─kubepods-podd22b422c_be99_491c_b43d_3fb4295ea426.slice
             ├─crio-83b6eddb89083766259b9e8f89ddf4d9f1e56360ad9a30af82d0ff7f57f75d19.scope
             │ └─22329 /pause
```

17. Switch the kubepods and crio folder cgroup. 

```
$ cd /sys/fs/cgroup/cpuset/kubepods.slice/kubepods-podd22b422c_be99_491c_b43d_3fb4295ea426.slice/crio-83b6eddb89083766259b9e8f89ddf4d9f1e56360ad9a30af82d0ff7f57f75d19.scope
```

18. Confirm there is one CPU assigned. 

```
$ for i in `ls cpuset.cpus tasks` ; do echo -n "$i "; cat $i ; done
cpuset.cpus 1
tasks 22329
```

19. Check the CPU Allowed List 

```
$ grep ^Cpus_allowed_list /proc/22329/status
Cpus_allowed_list:	1
```

20. Check the allowed cpus.

```
cat /sys/fs/cgroup/cpuset/kubepods.slice/kubepods-podd22b422c_be99_491c_b43d_3fb4295ea426.slice/crio-83b6eddb89083766259b9e8f89ddf4d9f1e56360ad9a30af82d0ff7f57f75d19.scope/cpuset.cpus
0
```

# References 

- [OpenShift 4.10: Controlling pod placement using node taints](https://docs.openshift.com/container-platform/4.10/nodes/scheduling/nodes-scheduler-taints-tolerations.html)