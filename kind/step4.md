
# Deploy [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)

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
--set defaultRules.rules.time=false \
--set defaultRules.rules.etcd=false
```{{execute}}


Verify Pods running:
``` 
kubectl get pods -n mon
kubectl get svc -n mon
kubectl get deployments -n mon
kubectl get daemonset -n mon
kubectl get statefulset -n mon
```{{execute}}


port-forward prometheus/grafana:
``` 
kubectl -n mon port-forward service/mon-kube-prometheus-stack-prometheus  9090:9090  --address 0.0.0.0 &
kubectl -n mon port-forward service/mon-grafana 3000:80  --address 0.0.0.0 &
kubectl get pods -A
```{{execute}}


Decode secret username/password:
``` 
kubectl -n mon get secret/mon-grafana -o go-template='{{range $k,$v := .data}}{{printf "%s: " $k}}{{if not $v}}{{$v}}{{else}}{{$v | base64decode}}{{end}}{{"\n"}}{{end}}'
```{{execute}}