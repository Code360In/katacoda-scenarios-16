### OpenTelemetry/Jaeger

`https://opentelemetry.io/`



``` 
kubectl create namespace otel
```{{execute}}



``` 
kubectl -n otel apply -f https://raw.githubusercontent.com/open-telemetry/opentelemetry-collector/main/examples/k8s/otel-config.yaml
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

