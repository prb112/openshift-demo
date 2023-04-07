# external-monitor
The External Monitor Operator

## Description


## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

❯ kubectl apply -f config/crd/bases/external-monitor.ocp-power.xyz_monitors.yaml
\customresourcedefinition.apiextensions.k8s.io/monitors.external-monitor.ocp-power.xyz created

openshift-demo/operator/external-monitor on  main [✘!?] via 🐹 v1.20.2
❯ kubectl apply -f config/samples/external-monitor.ocp-power.xyz_v1alpha1_monitor.yaml
monitor.external-monitor.ocp-power.xyz/monitor-sample created

❯ oc get monitor.external-monitor.ocp-power.xyz/monitor-sample
NAME             AGE
monitor-sample   26s

❯ oc get monitor.external-monitor.ocp-power.xyz/monitor-sample -o yaml
apiVersion: external-monitor.ocp-power.xyz/v1alpha1
kind: Monitor
metadata:
annotations:
kubectl.kubernetes.io/last-applied-configuration: |
{"apiVersion":"external-monitor.ocp-power.xyz/v1alpha1","kind":"Monitor","metadata":{"annotations":{},"labels":{"app.kubernetes.io/created-by":"external-monitor","app.kubernetes.io/instance":"monitor-sample","app.kubernetes.io/managed-by":"kustomize","app.kubernetes.io/name":"monitor","app.kubernetes.io/part-of":"external-monitor"},"name":"monitor-sample","namespace":"acmeair"},"spec":{"image":"quay.io/pbastide_rh/openshift-demo","name_value":["a=b","c=d"],"tag":"co-content-latest"}}
creationTimestamp: "2023-04-07T00:34:29Z"
generation: 1
labels:
app.kubernetes.io/created-by: external-monitor
app.kubernetes.io/instance: monitor-sample
app.kubernetes.io/managed-by: kustomize
app.kubernetes.io/name: monitor
app.kubernetes.io/part-of: external-monitor
name: monitor-sample
namespace: acmeair
resourceVersion: "1126786"
uid: 8c24e682-a382-403f-9a8d-a6eef93f0678
spec:
image: quay.io/pbastide_rh/openshift-demo
name_value:
- a=b
- c=d
  tag: co-content-latest

❯ make generate
test -s /Users/paulbastide/Desktop/work/multiarch_base_ocp/openshift-demo/operator/external-monitor/bin/controller-gen && /Users/paulbastide/Desktop/work/multiarch_base_ocp/openshift-demo/operator/external-monitor/bin/controller-gen --version | grep -q v0.11.1 || \
GOBIN=/Users/paulbastide/Desktop/work/multiarch_base_ocp/openshift-demo/operator/external-monitor/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.1
/Users/paulbastide/Desktop/work/multiarch_base_ocp/openshift-demo/operator/external-monitor/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."



2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/external-monitor:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/external-monitor:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

