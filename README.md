# Red Hat OpenShift - Demo

The purpose of this repository is to provide a set of simple proof-of-concept configuration scripts for OpenShift and various features on OpenShift on Power.

# Feature
1. [OpenShift `oc` plugin](oc_plugin) a sample plugin for `oc` running on mulitple architectures where the [`oc`](https://access.redhat.com/downloads/content/290/ver=4.10/rhel---8/4.10.6/x86_64/product-software) tool manages all the OpenShift resources with handy commands for OpenShift and Kubernetes. The [OpenShift Client CLI (oc)](https://github.com/openshift/oc) project is built on top of [`kubectl`](https://kubernetes.io/docs/reference/kubectl/) adding built-in features to simplify interactions with an OpenShift cluster.
1. [Topology Manager](topology_manager) aligns pod resources for optimal CPU, Memory, Device and other Topology placements in a cluster.


# Is this a Red Hat or IBM supported solution?

No. This is only a proof-of-concept that serves as a good starting point to understand how the various features works with OpenShift.