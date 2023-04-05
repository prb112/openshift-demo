To wrap ansible for a demo

1. Create your anisble-check container

```
❯ podman build -t quay.io/pbastide_rh/demo:ansible-check .
...
Successfully tagged quay.io/pbastide_rh/demo:ansible-check
f2a111264e30a0338fa76cd07aae86b4089c1974d676e8cfc467841926d5fc71
```

2. Add a layer 

```
COPY --chown=1001:0 /opt/ol/wlp/usr /opt/ol/wlp/usr
```

3. Run the container

```
❯ podman container run quay.io/pbastide_rh/demo:ansible-check 
ansible-playbook [core 2.14.4]
  config file = None
  configured module search path = ['/root/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
  ansible python module location = /root/.local/lib/python3.9/site-packages/ansible
  ansible collection location = /root/.ansible/collections:/usr/share/ansible/collections
  executable location = /root/.local/bin/ansible-playbook
  python version = 3.9.16 (main, Dec  8 2022, 00:00:00) [GCC 11.3.1 20221121 (Red Hat 11.3.1-4)] (/usr/bin/python3)
  jinja version = 3.1.2
  libyaml = True
```