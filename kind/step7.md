### OpenTelemetry

https://opentelemetry.io/

``` 
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```{{execute}}

``` 
kubectl create namespace otel
```{{execute}}

``` 
kubectl apply -n otel -f https://raw.githubusercontent.com/open-telemetry/opentelemetry-collector/main/examples/k8s/otel-config.yaml
``` 


Verify Pods running:
``` 
kubectl get pods -n otel
kubectl get svc -n otel
kubectl get deployments -n otel
kubectl get daemonset -n otel
kubectl get statefulset -n otel
```{{execute}}