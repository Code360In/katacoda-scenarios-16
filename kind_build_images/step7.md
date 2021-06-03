### OpenTelemetry/Jaeger

`https://opentelemetry.io/`



``` 
kubectl create namespace otel
```{{execute}}

Install otel collector
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


# Jaeger-all-in-one not ready.