# Ghost

This demonstration uses [kustomize](https://kubectl.docs.kubernetes.io/references/kustomize/) to deploy [ghost](https://ghost.org/) to OpenShift. 

# Steps
1. Install [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/)

```
$ brew install kustomize
```

2. Login to your cluster using `oc`. 

3. Generate a randomized password

```
$ ENV_PASS=$(openssl rand -hex 10)
$ echo ${ENV_PASS}
```

Note, save the output... 

4. Generate the working url for the cluster/ghost app.

```
$ export WEB_DOMAIN=https://web-route-ghost.apps.$(oc get ingress.config.openshift.io cluster -o yaml | grep domain | awk '{print $NF}')
$ echo ${WEB_DOMAIN}
```

5. Create the secret for the database

```
$ cat secrets/01_db_secret.yml | sed "s|ENV_PASS|${ENV_PASS}|" | oc apply -f -
```

6. Create the configmap for the Ghost app URL. 

```
$ cat secrets/02_web_cm.yml | sed "s|WEB_DOMAIN|${WEB_DOMAIN}|" | oc apply -f -
```

7. Create the deployment for the website

```
$ oc apply -k overlays/dev
namespace/ghost configured
service/db-service unchanged
service/web unchanged
persistentvolumeclaim/db-pvc unchanged
persistentvolumeclaim/web-content unchanged
deployment.apps/ghost-db unchanged
deployment.apps/web unchanged
route.route.openshift.io/web-route unchanged
```

8. To clean it up you can run... 

```
$ oc delete -k overlays/dev
namespace "ghost" deleted
service "db-service" deleted
service "web" deleted
persistentvolumeclaim "db-pvc" deleted
persistentvolumeclaim "web-content" deleted
deployment.apps "ghost-db" deleted
deployment.apps "web" deleted
route.route.openshift.io "web-route" deleted
```

9. To see your website URL, you can grab the config map. 

```
$ oc get cm -o yaml
```

10. Navigate to the URL, such as https://web-route-ghost.apps.xyz.zzz.zyz.com/ghost/ to start setting up your site.

# References
1. https://elixm.com/how-to-deploy-ghost-blog-with-kubernetes/
1. https://hub.docker.com/_/ghost
1. https://hub.docker.com/_/mysql
1. https://github.com/openshift-cs/ghost-example/blob/master/ghost_template.yaml