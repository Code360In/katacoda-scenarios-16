# Cortex & dependencies with Prometheus in Kubernetes


``` 
kubectl apply -f k8s/
```{{execute}}


# Prometheus

``` 
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```{{execute}}

``` 
helm search repo prometheus-community
```{{execute}}

``` 
helm install prom-one prometheus-community/prometheus \
 --set server.global.external_labels.cluster=one \
 --set serverFiles."prometheus\.yml".remote_write[0].url=http://nginx.default.svc.cluster.local:80/api/prom/push
```{{execute}}



expose prometheus
``` 
export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace default port-forward $POD_NAME 9090:9090  --address 0.0.0.0 &
```{{execute}}


# Grafana

``` 
helm repo add grafana https://grafana.github.io/helm-charts
helm search repo grafana
```{{execute}}


``` 
helm install grafana grafana/grafana  \
 --set datasources."datasources\.yaml".apiVersion=1 \
 --set datasources."datasources\.yaml".datasources[0].name=cortex \
 --set datasources."datasources\.yaml".datasources[0].type=prometheus \
 --set datasources."datasources\.yaml".datasources[0].url=http://nginx.default.svc.cluster.local/api/prom \
 --set datasources."datasources\.yaml".datasources[0].access=proxy \
 --set datasources."datasources\.yaml".datasources[0].isDefault=true
```{{execute}}




expose grafana
``` 
export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}")

kubectl --namespace default port-forward $POD_NAME 3000:3000  --address 0.0.0.0 &

```{{execute}}



get secret
``` 
kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```{{execute}}


test
```
up{cluster="one"}
prometheus_tsdb_head_samples_appended_total{cluster="one"}
```



# HA Prometheus setup with deduplication


``` 

kubectl apply -f k8s-ha/

helm install prom-one stable/prometheus \
--set server.global.external_labels.cluster=one \
--set server.global.external_labels.replica=one \
--set serverFiles."prometheus\.yml".remote_write[0].url=http://nginx.default.svc.cluster.local:80/api/prom/push

helm install prom-two stable/prometheus \
--set server.global.external_labels.cluster=one \
--set server.global.external_labels.replica=two \
--set serverFiles."prometheus\.yml".remote_write[0].url=http://nginx.default.svc.cluster.local:80/api/prom/push

helm install grafana stable/grafana \
--set datasources."datasources\.yaml".apiVersion=1 \
--set datasources."datasources\.yaml".datasources[0].name=cortex \
--set datasources."datasources\.yaml".datasources[0].type=prometheus \
--set datasources."datasources\.yaml".datasources[0].url=http://nginx.default.svc.cluster.local/api/prom \
--set datasources."datasources\.yaml".datasources[0].access=proxy \
--set datasources."datasources\.yaml".datasources[0].isDefault=true

kubectl port-forward svc/grafana 3000:80
```{{execute}}


get secret
```
kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```{{execute}}


To test HA, we can try deleting one of the Prometheus pods.
```
kubectl delete pod -l app=prometheus,component=server,release=prom-one
```{{execute}}


