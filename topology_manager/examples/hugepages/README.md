# Feature: OpenShift on Power: Topology Manager - Hugepages Demonstration

This demonstration shows the Hugepages allocation in OpenShift and some of the finer debug points.

*Note*: The allocated Hugepages memory is not deallocated.

1. Login to OpenShift 

```
$ oc login
```

2. List the Nodes

```
$ oc get nodes -l node-role.kubernetes.io/worker
NAME                    STATUS   ROLES    AGE   VERSION
lon06-worker-0.xip.io   Ready    worker   9d    v1.23.5+3afdacb
lon06-worker-1.xip.io   Ready    worker   9d    v1.23.5+3afdacb
```

3. Check one of your worker nodes by starting a terminal session

```
$ ssh core@lon06-worker-0.xip.io
```

4. Check the Hugepagesize to verify it exists

```
$ grep Hugepagesize /proc/meminfo
Hugepagesize:      16384 kB
```

5. Check the `vm.nr_hugepages`, if it's zero we'll need to set it up.

```
$ sysctl vm.nr_hugepages
vm.nr_hugepages = 0
```

6. You can manually set until reboot (you'll need to do this one each worker) or use a MachineConfig

```
$ sudo sysctl -w vm.nr_hugepages=256
vm.nr_hugepages = 256
```

7. Use a MachineConfig to set `vm.nr_hugepages`

```
$ oc apply -f machineconfig.yaml
machineconfig.machineconfiguration.openshift.io/99-sysctl-nr-hugepages created
```

8. Wait for an update to the MachineConfig

```
$ oc wait mcp/worker --for condition=updated --timeout=25m
```

9. Create the Hugepages demonstration pod

```
$ oc apply -f pod.yaml 
pod/hugepages-demo created
```

10. Check the hugepages output is correct. You are looking for Page Size 16M and total size.

```
$ oc exec -it hugepages-demo -- grep hugepages/demo /proc/1/smaps -A3
7efff8000000-7f0000000000 rw-s 00000000 00:1de 197766                    /dev/hugepages/demo
Size:             131072 kB
KernelPageSize:    16384 kB
MMUPageSize:       16384 kB
```

11. Find the node the pod is running on: 

```
$ oc get pods -o wide
NAME             READY   STATUS    RESTARTS   AGE     IP       NODE     NOMINATED NODE   READINESS GATES
hugepages-demo   1/1     Running   0          5m30s   10.128.2.12   lon06-worker-1.xip.io   <none> <none>
```

12. Check the HugePages value. 

```
$ grep HugePages_ /proc/meminfo
HugePages_Total:      20
HugePages_Free:       17
HugePages_Rsvd:        5
HugePages_Surp:        0
```

13. Check the Allocatable HugePages allocation is non-zero

```
$ oc get node lon06-worker-1.xip.io   -o jsonpath="{.status.allocatable}" | jq -r .  
{
  "cpu": "7500m",
  "ephemeral-storage": "115586611009",
  "hugepages-16Gi": "0",
  "hugepages-16Mi": "4Gi",
  "memory": "28069952Ki",
  "pods": "250"
}
```

14. Check that there are allocated Hugepages (on one of the Worker nodes)

1. Switch to the root user

```
$ sudo -s 
```

2. Find the hugepaged process

```
ps -ef | grep hugepaged
```

3. Verify the allocated data 

```
$ PROC=27693
$ grep -A3 'demo' /proc/${PROC}/smaps
7efff8000000-7f0000000000 rw-s 00000000 00:1de 197766                    /dev/hugepages/demo
Size:             131072 kB
KernelPageSize:    16384 kB
MMUPageSize:       16384 kB
```

# Summary

You have seen how to configure a node to support hugepages and deploy a Pod with Hugepages support, and confirm Hugepages are used.

<hr>

# References

- [Red Hat OpenShift 4.10: How huge pages are consumed by apps](https://docs.openshift.com/container-platform/4.10/scalability_and_performance/what-huge-pages-do-and-how-they-are-consumed-by-apps.html)
- [Red Hat Customer Portal: [RHEL] How do I check for hugepages usage and what is using it?](https://access.redhat.com/solutions/320303)
- [Red Hat Customer Portal: How to use, monitor, and disable transparent hugepages in Red Hat Enterprise Linux 6 and 7?](https://access.redhat.com/solutions/46111)
- [Povilas: Go Memory Management](https://povilasv.me/go-memory-management/)
- [Red Hat Universal Base Image 8](https://catalog.redhat.com/software/containers/ubi8/ubi/5c359854d70cc534b3a3784e)
- [Linux: How to force any application to use Hugepages without modifying the source code](https://paolozaino.wordpress.com/2016/10/02/how-to-force-any-linux-application-to-use-hugepages-without-modifying-the-source-code/)
- [Kernel: HugeTLB Pages](https://www.kernel.org/doc/html/latest/admin-guide/mm/hugetlbpage.html)
- [Linux Kernel: SelfTests Hugepage example using mmap](https://github.com/torvalds/linux/blob/master/tools/testing/selftests/vm/hugepage-mmap.c)
- [StackOverflow: Why mmap cannot allocate memory?](https://stackoverflow.com/questions/27634109/why-mmap-cannot-allocate-memory)
- [Erik Rigtorp: Using huge pages on Linux](https://rigtorp.se/hugepages/)

# Appendix: Install Hugepages Tools

1. Check that hugectl is provided by your package manager

```
yum whatprovides hugectl
```

2. Install the Hugectl tools

```
yum -y install libhugetlbfs-utils libhugetlbfs
```

# Appendix: Install Build Tools

The minimum build tools required to build this sample project are `make` and `golang`.

```
 yum install -y golang make
```

# Appendix: [mmap allocation issues](https://stackoverflow.com/questions/27634109/why-mmap-cannot-allocate-memory)

If you see `mmap: Cannot allocate memory`, then you should run `echo 20 > /proc/sys/vm/nr_hugepages` which switches the hugepages to non-zero.

# Appendix: Using Tuned to Configure the Node with Alternative Page Sizes

Create the Tuned configuration and wait for a node restart.

1. Create a tuned.yaml file with the additional page size and the number of pages per the manifests/tuned.yaml

```
cmdline_openshift_node_hugepages=hugepagesz=2M hugepages=50 
```

2. Apply the configuration 

```
$ oc apply -f manifests/tuned.yaml 
tuned/hugepages created
```

3.  Wait while the node is restarted and have fun.

# Appendix: huge tools

1. Install `yum install libhugetlbfs-utils -y`

2. Check the mounts

```
$  hugeadm  --list-all-mounts 
Mount Point                      Options
/dev/hugepages                   rw,seclabel,relatime,pagesize=16M
/var/lib/hugetlbfs/pagesize-16MB rw,seclabel,relatime,pagesize=16M
/var/lib/hugetlbfs/pagesize-16GB rw,seclabel,relatime,pagesize=16384M
```

3. Check the Pools and what setting they have

```
$ hugeadm  --pool-list
      Size  Minimum  Current  Maximum  Default
  16777216       20       20       20        *
17179869184        0        0        0         
```

4. Start an application written with glibc to transparently use hugepages. 

```
$ hugectl myapp
```

Note, it does not work with non-glibc apps (e.g. golang)

6. Create the mounts automatically:

```
$ hugeadm --create-mounts
```

And then you can check the mounts 

```
$ mount | grep pagesize
none on /var/lib/hugetlbfs/pagesize-16MB type hugetlbfs (rw,relatime,seclabel,pagesize=16M)
none on /var/lib/hugetlbfs/pagesize-16GB type hugetlbfs (rw,relatime,seclabel,pagesize=16384M)
```

# Appendix: Manual Creation of the Hugepages FS

To mount 64KB pages (if the system hardware supports it):

```
mkdir -p /mnt/hugetlbfs-16K
mount -t hugetlbfs none -opagesize=16777216 /mnt/hugetlbfs-16K
```

```
mkdir /dev/hugepages16G
mount -t hugetlbfs -o pagesize=17179869184 none /dev/hugepages16G
```

# Appendix: Check Hugepage Sizes

Check Hugepage sizes

```
$ ls /sys/kernel/mm/hugepages
hugepages-16384kB  hugepages-16777216kB
```

# Appendix: List sysctl settings

```
$ sysctl -a | grep hugepages
vm.nr_hugepages = 642
vm.nr_hugepages_mempolicy = 642
vm.nr_overcommit_hugepages = 676
```

To allocate our 2048 Huge Pages we can use:

```
$ echo 2048 > /proc/sys/vm/nr_hugepages
```

To disable transparent huge pages

```
echo never > /sys/kernel/mm/transparent_hugepage/enabled
```

# Appendix: Check Hugepage Mount usage for a process

1. Check Hugepage Mount usage for a process (you need to know the process id and the filesystem name)

```
$ PROC=4991
$ grep -A3 '/var/lib/hugetlbfs/pagesize-16MB/demo' /proc/${PROC}/smaps
7efff8000000-7f0000000000 rw-s 00000000 00:2f 63535                      /var/lib/hugetlbfs/pagesize-16MB/demo
Size:             131072 kB
KernelPageSize:    16384 kB
MMUPageSize:       16384 kB
```

# Is this a Red Hat or IBM supported solution?

No. This is only a proof of concept that serves as a good starting point to understand how the Descheduler Profiles works in OpenShift.