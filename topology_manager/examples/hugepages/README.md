# Feature: Topology Manager - Huge Pages Demonstration for OpenShift on Power


$ sysctl -w vm.nr_hugepages=256

$ grep Hugepagesize /proc/meminfo
Hugepagesize:      16384 kB


mmap: Cannot allocate memory
echo 20 > /proc/sys/vm/nr_hugepages
https://stackoverflow.com/questions/27634109/why-mmap-cannot-allocate-memory

Please Note: Before allocating a big number of Hugepage on a system that is running Virtual Machines or other memory hungry applications, make sure to shutdown your Virtual Machines and any memory hungry application before executing the previous command otherwise the execution may take a long time to complete.

$  hugeadm  --list-all-mounts 
Mount Point                      Options
/dev/hugepages                   rw,seclabel,relatime,pagesize=16M
/var/lib/hugetlbfs/pagesize-16MB rw,seclabel,relatime,pagesize=16M
/var/lib/hugetlbfs/pagesize-16GB rw,seclabel,relatime,pagesize=16384M

$ hugeadm  --pool-list
      Size  Minimum  Current  Maximum  Default
  16777216       20       20       20        *
17179869184        0        0        0         


$ hugectl bin/hugepagesd &


$ PROC=4991
$ grep -A3 '/var/lib/hugetlbfs/pagesize-16MB/demo' /proc/${PROC}/smaps
7efff8000000-7f0000000000 rw-s 00000000 00:2f 63535                      /var/lib/hugetlbfs/pagesize-16MB/demo
Size:             131072 kB
KernelPageSize:    16384 kB
MMUPageSize:       16384 kB

$ grep HugePages_ /proc/meminfo
HugePages_Total:      20
HugePages_Free:       17
HugePages_Rsvd:        5
HugePages_Surp:        0

<hr>







To allocate our 2048 Huge Pages we can use:

# echo 2048 > /proc/sys/vm/nr_hugepages

    

$ ls /sys/kernel/mm/hugepages
hugepages-16384kB  hugepages-16777216kB


$ grep HugePages_ /proc/meminfo


To quickly and temporarly allocate them, or we can use:



if test -f /sys/kernel/mm/transparent_hugepage/enabled; then
   echo never > /sys/kernel/mm/transparent_hugepage/enabled
fi

# hugeadm --pool-list
      Size  Minimum  Current  Maximum  Default
  16777216      676      676      676        *
17179869184        0        0        0         

ipcs -m

$ grep -i huge /proc/mounts
cgroup /sys/fs/cgroup/hugetlb cgroup rw,seclabel,nosuid,nodev,noexec,relatime,hugetlb 0 0
hugetlbfs /dev/hugepages hugetlbfs rw,seclabel,relatime,pagesize=16M 0 0

grep -B 11 'KernelPageSize: 2048 kB' /proc/3831/smaps | grep "^Size:" | awk 'BEGIN{sum=0}{sum+=$2}END{print sum/1024}'

# echo always >/sys/kernel/mm/transparent_hugepage/enabled

sysctl vm.nr_hugepages

sysctl -a | grep hugepages
vm.nr_hugepages = 642
vm.nr_hugepages_mempolicy = 642
vm.nr_overcommit_hugepages = 676


mount -t hugetlbfs \
      -o uid=<value>,gid=<value>,mode=<value>,pagesize=<value>,size=<value>,\
      min_size=<value>,nr_inodes=<value> none /mnt/huge

To mount the default huge page size:

# mkdir -p /mnt/hugetlbfs
# mount -t hugetlbfs none /mnt/hugetlbfs

hugeadm --create-mounts

mount | grep pagesize
none on /var/lib/hugetlbfs/pagesize-16MB type hugetlbfs (rw,relatime,seclabel,pagesize=16M)
none on /var/lib/hugetlbfs/pagesize-16GB type hugetlbfs (rw,relatime,seclabel,pagesize=16384M)

gcc-c++

hugeadm --list-all-mounts 
Mount Point          Options
/dev/hugepages       rw,seclabel,relatime,pagesize=16M

16K
# hugeadm --pool-list
      Size  Minimum  Current  Maximum  Default
  16777216      642      642     1318        *
17179869184        0        0        0        

To mount 64KB pages (if the system hardware supports it):

mkdir -p /mnt/hugetlbfs-16K
mount -t hugetlbfs none -opagesize=16777216 /mnt/hugetlbfs-16K

mkdir /dev/hugepages16G
mount -t hugetlbfs -o pagesize=17179869184 none /dev/hugepages16G

<hr>

# References

- [Red Hat OpenShift 4.10: How huge pages are consumed by apps](https://docs.openshift.com/container-platform/4.10/scalability_and_performance/what-huge-pages-do-and-how-they-are-consumed-by-apps.html)
- [Red Hat Customer Portal: [RHEL] How do I check for hugepages usage and what is using it?](https://access.redhat.com/solutions/320303)
- [Red Hat Customer Portal: How to use, monitor, and disable transparent hugepages in Red Hat Enterprise Linux 6 and 7?](https://access.redhat.com/solutions/46111)
- [Povilas: Go Memory Management](https://povilasv.me/go-memory-management/)
- [Red Hat Universal Base Image 8 Minimal](https://catalog.redhat.com/software/containers/ubi8/ubi-minimal/5c359a62bed8bd75a2c3fba8?architecture=ppc64le&container-tabs=gti)
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

# Is this a Red Hat or IBM supported solution?

No. This is only a proof of concept that serves as a good starting point to understand how the Descheduler Profiles works in OpenShift.