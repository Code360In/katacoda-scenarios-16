
# Deploy Loki



``` 
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
```{{execute}}


``` 
helm upgrade --install loki grafana/loki-stack -n mon
```{{execute}}


Verify Pods running:
``` 
kubectl get pods -n mon
kubectl get svc -n mon
kubectl get deployments -n mon
kubectl get daemonset -n mon
kubectl get statefulset -n mon
```{{execute}}



Let's define a datasource with Loki configured before:

The URL for Loki should be:

`http://loki:3100`{{copy}}

For all the other fields you can use the default values

Test and save your datasource.



