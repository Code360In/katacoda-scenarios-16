
# Deploy Grafana



```
oc new-app grafana/grafana:latest-ubuntu -n elastic-system 
```{{execute}}

Expose service:

```
oc expose svc/grafana -n elastic-system 
```{{execute}}