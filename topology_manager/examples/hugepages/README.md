# Feature: Topology Manager - Huge Pages Demonstration for OpenShift on Power



$ ls /sys/kernel/mm/hugepages
hugepages-16384kB  hugepages-16777216kB


$ grep HugePages_ /proc/meminfo


ipcs -m

$ grep -i huge /proc/mounts
cgroup /sys/fs/cgroup/hugetlb cgroup rw,seclabel,nosuid,nodev,noexec,relatime,hugetlb 0 0
hugetlbfs /dev/hugepages hugetlbfs rw,seclabel,relatime,pagesize=16M 0 0

# References

- [Red Hat Customer Portal: [RHEL] How do I check for hugepages usage and what is using it?](https://access.redhat.com/solutions/320303)
- [Povilas: Go Memory Management](https://povilasv.me/go-memory-management/)
- [Red Hat Universal Base Image 8 Minimal](https://catalog.redhat.com/software/containers/ubi8/ubi-minimal/5c359a62bed8bd75a2c3fba8?architecture=ppc64le&container-tabs=gti)
- [Linux: How to force any application to use Hugepages without modifying the source code](https://paolozaino.wordpress.com/2016/10/02/how-to-force-any-linux-application-to-use-hugepages-without-modifying-the-source-code/)

# Appendix: Install Hugepages Tools

```
yum whatprovides hugectl
```

```
yum -y install libhugetlbfs-utils libhugetlbfs
```

# Is this a Red Hat or IBM supported solution?

No. This is only a proof of concept that serves as a good starting point to understand how the Descheduler Profiles works in OpenShift.