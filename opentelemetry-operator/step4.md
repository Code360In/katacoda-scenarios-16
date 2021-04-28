
# Deploy cert-manager

https://cert-manager.io/docs/installation/kubernetes/


``` 
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.3.1/cert-manager.yaml
```{{execute}}


Verify if Pod is running
``` 
kubectl get pods -A
```{{execute}}


# Deploy opentelemetry-operator

``` 
kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml
```{{execute}}


Verify if Pod is running
``` 
kubectl get pods -A
```{{execute}}


Once the opentelemetry-operator deployment is ready, create an OpenTelemetry Collector (otelcol) instance, like:

``` 
kubectl apply -f - <<EOF
apiVersion: opentelemetry.io/v1alpha1
kind: OpenTelemetryCollector
metadata:
  name: simplest
spec:
  config: |
    receivers:
      jaeger:
        protocols:
          grpc:
    processors:

    exporters:
      logging:

    service:
      pipelines:
        traces:
          receivers: [jaeger]
          processors: []
          exporters: [logging]
EOF
```{{execute}}

