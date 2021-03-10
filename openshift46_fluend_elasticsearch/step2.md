
# Deploy Grafana



```
oc new-app grafana/grafana:latest-ubuntu -n elastic-system 
```{{execute}}

Expose service:

```
oc expose svc/grafana -n elastic-system 
```{{execute}}

To access Grafana:
https://grafana-elastic-system.[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/login

admin/admin 
skip


add datasource elasticsearch

