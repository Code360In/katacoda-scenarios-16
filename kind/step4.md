
Deploy [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)

 <pre class="file">
Installs the kube-prometheus stack, a collection of Kubernetes manifests, Grafana dashboards, and Prometheus rules combined with documentation and scripts to provide easy to operate end-to-end Kubernetes cluster monitoring with Prometheus using the Prometheus Operator.
 </pre>

``` 
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```{{execute}}

``` 
helm repo update
```{{execute}}


Create namespace mon:
``` 
kubectl create namespace mon
```{{execute}}

``` 
helm upgrade --install  mon prometheus-community/kube-prometheus-stack  --set grafana.image.tag="7.4.1-ubuntu" --set prometheus.image.tag="v2.25.0-rc.0" --set defaultRules.rules.time=false -n mon
```{{execute}}


Verify Pods runnimg:
``` 
kubectl get pods -n mon
kubectl get sv -n mon
kubectl get deployments -n mon
kubectl get daemonset -n mon
```{{execute}}