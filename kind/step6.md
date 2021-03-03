### Prometheus Adapter for Kubernetes Metrics APIs


# Deploy prometheus-community/prometheus-adapter


<pre class="file">
This [adapter](https://github.com/kubernetes-sigs/prometheus-adapter) is therefore suitable for use with the autoscaling/v2 Horizontal Pod Autoscaler in Kubernetes 1.6+.
It can also replace the metrics server on clusters that already run Prometheus and collect the appropriate metrics.
</pre>


``` 
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```{{execute}}

``` 
helm repo update
```{{execute}}

``` 
helm upgrade --install prometheus-adapter prometheus-community/prometheus-adapter -n mon \
--set prometheus-url="http://mon-kube-prometheus-stack-prometheus:9090"
``` 


Verify Pods running:
``` 
kubectl get pods -n mon
kubectl get svc -n mon
kubectl get deployments -n mon
kubectl get daemonset -n mon
kubectl get statefulset -n mon
```{{execute}}