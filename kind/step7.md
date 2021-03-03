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

`
# Data sources: traces, metrics, logs
  memory_limiter:
    ballast_size_mib: 2000
    check_interval: 5s
    limit_mib: 4000
    spike_limit_mib: 500
`

``` 
helm upgrade --install  -n otel  opentelemetry-collector open-telemetry/opentelemetry-collector --set config.processors.memory_limiter.ballast_size_mib=2000 \
--set config.processors.memory_limiter.limit_mib=4000

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


`
processors:
  # Data sources: traces
  attributes:
    actions:
      - key: environment
        value: production
        action: insert
      - key: db.statement
        action: delete
      - key: email
        action: hash

  # Data sources: traces, metrics, logs
  batch:

  # Data sources: metrics
  filter:
    metrics:
      include:
        match_type: regexp
        metric_names:
        - prefix/.*
        - prefix_.*

  # Data sources: traces, metrics, logs
  memory_limiter:
    ballast_size_mib: 2000
    check_interval: 5s
    limit_mib: 4000
    spike_limit_mib: 500

  # Data sources: traces
  resource:
    attributes:
    - key: cloud.zone
      value: "zone-1"
      action: upsert
    - key: k8s.cluster.name
      from_attribute: k8s-cluster
      action: insert
    - key: redundant-attribute
      action: delete

  # Data sources: traces
  probabilistic_sampler:
    hash_seed: 22
    sampling_percentage: 15

  # Data sources: traces
  span:
    name:
      to_attributes:
        rules:
          - ^\/api\/v1\/document\/(?P<documentId>.*)\/update$
      from_attributes: ["db.svc", "operation"]
      separator: "::"
  `