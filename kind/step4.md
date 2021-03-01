
Deploy Metrics




``` 
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```{{execute}}

``` 
helm repo update
```{{execute}}

``` 
helm upgrade --install  mon prometheus-community/kube-prometheus-stack  --set grafana.image.tag="7.4.1-ubuntu" --set prometheus.image.tag="v2.25.0-rc.0" --set defaultRules.rules.time=false
```{{execute}}