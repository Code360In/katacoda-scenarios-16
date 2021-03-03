
# Deploy Loki



``` 
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
```{{execute}}


``` 
helm upgrade --install loki grafana/loki-stack -n mon
```{{execute}}
