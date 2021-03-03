### OpenTelemetry/Jaeger

`https://opentelemetry.io/`

`https://github.com/open-telemetry/opentelemetry-helm-charts`

`https://github.com/open-telemetry/opentelemetry-helm-charts/tree/main/charts/opentelemetry-collector`


``` 
kubectl create namespace otel
```{{execute}}



Add OpenTelemetry Helm repository:
``` 
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
```{{execute}}


To install the chart with the release name my-opentelemetry-collector, run the following command:

``` 
helm upgrade --install  -n otel  opentelemetry-collector open-telemetry/opentelemetry-collector --set standaloneCollector.enabled=false
```{{execute}}



Patch Service from  LoadBalancer to ClusterIP:
``` 
kubectl patch svc jaeger-query -n otel --type='json' -p '[{"op":"replace","path":"/spec/type","value":"ClusterIP"}]'
```{{execute}}


Verify Pods running:
``` 
kubectl get pods -n otel
kubectl get svc -n otel
kubectl get deployments -n otel
kubectl get daemonset -n otel
```{{execute}}



Expose Jaeger:
``` 
kubectl -n otel port-forward service/jaeger-query 16686:16686  --address 0.0.0.0 &
kubectl get pods -A
```{{execute}}


Access Jaeger Gui:

https://[[HOST_SUBDOMAIN]]-16686-[[KATACODA_HOST]].environments.katacoda.com


