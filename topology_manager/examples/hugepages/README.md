# Feature: Topology Manager - Huge Pages Demonstration for OpenShift on Power


# grep Hugepagesize /proc/meminfo
Hugepagesize:      16384 kB


To allocate our 2048 Huge Pages we can use:

# echo 2048 > /proc/sys/vm/nr_hugepages

    Please Note: Before allocating a big number of Hugepage on a system that is running Virtual Machines or other memory hungry applications, make sure to shutdown your Virtual Machines and any memory hungry application before executing the previous command otherwise the execution may take a long time to complete.

$ ls /sys/kernel/mm/hugepages
hugepages-16384kB  hugepages-16777216kB


$ grep HugePages_ /proc/meminfo


To quickly and temporarly allocate them, or we can use:

# sysctl -w vm.nr_hugepages=2048


# hugeadm --pool-list
      Size  Minimum  Current  Maximum  Default
  16777216      676      676      676        *
17179869184        0        0        0         

ipcs -m

$ grep -i huge /proc/mounts
cgroup /sys/fs/cgroup/hugetlb cgroup rw,seclabel,nosuid,nodev,noexec,relatime,hugetlb 0 0
hugetlbfs /dev/hugepages hugetlbfs rw,seclabel,relatime,pagesize=16M 0 0

<hr>

# References

- [Red Hat Customer Portal: [RHEL] How do I check for hugepages usage and what is using it?](https://access.redhat.com/solutions/320303)
- [Povilas: Go Memory Management](https://povilasv.me/go-memory-management/)
- [Red Hat Universal Base Image 8 Minimal](https://catalog.redhat.com/software/containers/ubi8/ubi-minimal/5c359a62bed8bd75a2c3fba8?architecture=ppc64le&container-tabs=gti)
- [Linux: How to force any application to use Hugepages without modifying the source code](https://paolozaino.wordpress.com/2016/10/02/how-to-force-any-linux-application-to-use-hugepages-without-modifying-the-source-code/)

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

# Is this a Red Hat or IBM supported solution?

No. This is only a proof of concept that serves as a good starting point to understand how the Descheduler Profiles works in OpenShift.