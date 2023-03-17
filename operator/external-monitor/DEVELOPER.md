1. Scaffold

operator-sdk init --domain external-monitor.ocp-power.xyz --repo github.com/prb112/openshift-demo/operator/external-monitor --plugins=go/v4-alpha --license apache2

2. Create the API 

operator-sdk create api --group external-monitor.ocp-power.xyz --version v1alpha1 --kind Monitor --resource --controller

Thought about using the plugin... however, it's a bit more complicated, as we want to Run a Job. 

~operator-sdk create api --group external-monitor.ocp-power.xyz --version v1alpha1 --kind Monitor --plugins="deploy-image/v1-alpha" --image=memcached:1.4.36-alpine --image-container-command="memcached,-m=64,modern,-v" --run-as-user="1001"~

3. Create the manifests

❯ make manifests

