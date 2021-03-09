### Logs

https://docs.fluentd.org/container-deployment/kubernetes

``` 
git clone https://github.com/fluent/fluentd-kubernetes-daemonset

```{{execute}}



kubectl create namespace elastic
helm repo add elastic https://helm.elastic.co
helm pull elastic/elasticsearch

helm install elasticsearch ./elasticsearch -n elastic

helm uninstall elasticsearch  -n elastic
kubectl delete namespace elastic