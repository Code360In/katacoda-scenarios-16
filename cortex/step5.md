
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