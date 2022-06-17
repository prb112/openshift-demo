This is a multi-architecture sample project. 

0. `make build` to create the multi-architecture build

1. Copy the output file to the test machine *postfixed with ppc64le*

2. Copy the File to a directory in the PATH

3. Rename the file to oc-multiarch

4. Change to executable 

chmod +x oc-multiarch

5. Run the test

```
[root@test-ocp-62a4-lon06-bastion-0 ~]# oc multiarch a b c d e f
OS Argument: /root/bin/oc-multiarch
OS Argument: a
OS Argument: b
OS Argument: c
OS Argument: d
OS Argument: e
OS Argument: f

The Home Directory for Kube /root
The list of the nodes:
lon06-master-0.test-ocp-62a4.158.176.141.246.xip.io
lon06-master-1.test-ocp-62a4.158.176.141.246.xip.io
lon06-master-2.test-ocp-62a4.158.176.141.246.xip.io
lon06-worker-0.test-ocp-62a4.158.176.141.246.xip.io
lon06-worker-1.test-ocp-62a4.158.176.141.246.xip.io
```
