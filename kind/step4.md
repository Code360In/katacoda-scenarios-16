
Deploy [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)

 <pre class="file">
Installs the kube-prometheus stack, a collection of Kubernetes manifests, Grafana dashboards, and Prometheus rules combined with documentation and scripts to provide easy to operate end-to-end Kubernetes cluster monitoring with Prometheus using the Prometheus Operator.
 </pre>


<pre class="file">
rules:
    alertmanager: true
    etcd: true
    general: true
    k8s: true
    kubeApiserver: true
    kubeApiserverAvailability: true
    kubeApiserverError: true
    kubeApiserverSlos: true
    kubelet: true
    kubePrometheusGeneral: true
    kubePrometheusNodeAlerting: true
    kubePrometheusNodeRecording: true
    kubernetesAbsent: true
    kubernetesApps: true
    kubernetesResources: true
    kubernetesStorage: true
    kubernetesSystem: true
    kubeScheduler: true
    kubeStateMetrics: true
    network: true
    node: true
    prometheus: true
    prometheusOperator: true
    time: true
 </pre>


Add helm package:
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

Run helm:
``` 
helm upgrade --install mon prometheus-community/kube-prometheus-stack -n mon \
--set grafana.service.type=NodePort --set grafana.service.nodePort=30080 \
--set prometheus.service.type=NodePort --set prometheus.service.nodePort=30090 \
--set defaultRules.rules.time=false
```{{execute}}


Verify Pods running:
``` 
kubectl get pods -n mon
kubectl get sv -n mon
kubectl get deployments -n mon
kubectl get daemonset -n mon
```{{execute}}

