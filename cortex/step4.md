
# Deploy Cortex

following:
https://www.infracloud.io/blogs/cortex-prometheus-ha/

Install multiple Prometheus and Grafana to visualize the data.


``` 
git clone https://github.com/kanuahs/cortex-demo.git
cd cortex-demo
```{{execute}}


``` 
docker-compose -f docker-demo/docker-compose.yaml up
```{{execute}}


clean-up
```
docker-compose -f docker-demo/docker-compose.yaml down
```{{execute}}


# Cortex & dependencies with Prometheus in Kubernetes


``` 
kubectl apply -f k8s/
```{{execute}}


# Prometheus
``` 
helm install stable/prometheus \
 --name prom-one \
 --set server.global.external_labels.cluster=one \
 --set serverFiles."prometheus\.yml".remote_write[0].url=http://nginx.default.svc.cluster.local:80/api/prom/push

```{{execute}}


# Grafana

``` 
helm install stable/grafana --name=grafana \
 --set datasources."datasources\.yaml".apiVersion=1 \
 --set datasources."datasources\.yaml".datasources[0].name=cortex \
 --set datasources."datasources\.yaml".datasources[0].type=prometheus \
 --set datasources."datasources\.yaml".datasources[0].url=http://nginx.default.svc.cluster.local/api/prom \
 --set datasources."datasources\.yaml".datasources[0].access=proxy \
 --set datasources."datasources\.yaml".datasources[0].isDefault=true
```{{execute}}


expose
``` 
kubectl port-forward svc/grafana 3000:80
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

``` 
```{{execute}}



``` 
```{{execute}}


``` 
```{{execute}}


# HA Prometheus setup with deduplication


``` 

kubectl apply -f k8s-ha/

helm install stable/prometheus \
--name prom-one \
--set server.global.external_labels.cluster=one \
--set server.global.external_labels.replica=one \
--set serverFiles."prometheus\.yml".remote_write[0].url=http://nginx.default.svc.cluster.local:80/api/prom/push

helm install stable/prometheus \
--name prom-two \
--set server.global.external_labels.cluster=one \
--set server.global.external_labels.replica=two \
--set serverFiles."prometheus\.yml".remote_write[0].url=http://nginx.default.svc.cluster.local:80/api/prom/push

helm install stable/grafana --name=grafana \
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



# Use Cassandra as index and chunk storage

```
helm repo add incubator https://kubernetes-charts-incubator.storage.googleapis.com/
helm install --wait --name=cassie incubator/cassandra
```{{execute}}


Once Cassandra is ready, proceed with installing all the other services.

```
kubectl apply -f k8s-cassandra/

helm install stable/prometheus \
--name prom-one \
--set server.global.external_labels.cluster=one \
--set serverFiles."prometheus\.yml".remote_write[0].url=http://nginx.default.svc.cluster.local:80/api/prom/push

helm install stable/grafana --name=grafana \
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

test

```
up{cluster="one"}
prometheus_tsdb_head_samples_appended_total{cluster="one"}
```