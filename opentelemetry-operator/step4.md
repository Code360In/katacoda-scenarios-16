
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


<pre class="file">

$ kubectl get pods -A
NAMESPACE                       NAME                                                         READY   STATUS    RESTARTS   AGE
cert-manager                    cert-manager-7dd5854bb4-ptbw5                                1/1     Running   0          12m
cert-manager                    cert-manager-cainjector-64c949654c-tmhpf                     1/1     Running   0          12m
cert-manager                    cert-manager-webhook-6bdffc7c9d-x6jp8                        1/1     Running   0          12m
dashboard                       dashboard-kubernetes-dashboard-67f7f4ff48-gtvrb              1/1     Running   0          13m
kube-system                     coredns-74ff55c5b-bjmkb                                      1/1     Running   0          13m
kube-system                     coredns-74ff55c5b-tzfmn                                      1/1     Running   0          13m
kube-system                     etcd-tx-cluster-k8s-control-plane                            1/1     Running   0          13m
kube-system                     kindnet-4x8z9                                                1/1     Running   0          13m
kube-system                     kube-apiserver-tx-cluster-k8s-control-plane                  1/1     Running   0          13m
kube-system                     kube-controller-manager-tx-cluster-k8s-control-plane         1/1     Running   0          13m
kube-system                     kube-proxy-fjhj5                                             1/1     Running   0          13m
kube-system                     kube-scheduler-tx-cluster-k8s-control-plane                  1/1     Running   0          13m
local-path-storage              local-path-provisioner-78776bfc44-9bjsz                      1/1     Running   0          13m
opentelemetry-operator-system   opentelemetry-operator-controller-manager-7df489cb85-nwjwl   2/2     Running   0          23s
$ 

</pre>



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



