# Logs

### Deploy Elastic



Install custom resource definitions and the operator with its RBAC rules:

``` 
kubectl apply -f https://download.elastic.co/downloads/eck/1.4.0/all-in-one.yaml --namespace=elastic-system
```{{execute}}


Monitor the operator logs:
``` 
kubectl -n elastic-system logs -f statefulset.apps/elastic-operator
```{{execute}}


Apply a simple Elasticsearch cluster specification, with one Elasticsearch node:

``` 
cat <<EOF | kubectl apply  -n elastic-system -f -
apiVersion: elasticsearch.k8s.elastic.co/v1
kind: Elasticsearch
metadata:
  name: quickstart
spec:
  version: 7.11.1
  nodeSets:
  - name: default
    count: 1
    config:
      node.store.allow_mmap: false
  http:
    tls:
      selfSignedCertificate:
        disabled: true
EOF
```{{execute}}



Get an overview of the current Elasticsearch clusters in the Kubernetes cluster, including health, version and number of nodes:
``` 
kubectl get elasticsearch -n elastic-system
```{{execute}}


You can see that one Pod is in the process of being started:
``` 
kubectl get pods --selector='elasticsearch.k8s.elastic.co/cluster-name=quickstart'  -n elastic-system
```{{execute}}

Access the logs for that Pod:
``` 
kubectl logs -f quickstart-es-default-0 -n elastic-system 
```{{execute}}



Decode quickstart-es-elastic-user secret username/password:
``` 
kubectl -n elastic-system get secret
kubectl -n elastic-system get secret/quickstart-es-elastic-user -o go-template='{{range $k,$v := .data}}{{printf "%s: " $k}}{{if not $v}}{{$v}}{{else}}{{$v | base64decode}}{{end}}{{"\n"}}{{end}}'

```{{execute}}



results: user/password

elastic: 3AlO1z58gb1801Xl86hVPvCW


Verify Pods running:
``` 
kubectl get pods -n elastic-system
kubectl get svc -n elastic-system
kubectl get deployments -n elastic-system
kubectl get daemonset -n elastic-system
```{{execute}}




### Deploy Daemon Fluentd

https://docs.fluentd.org/container-deployment/kubernetes

``` 
git clone https://github.com/fluent/fluentd-kubernetes-daemonset
```{{execute}}

``` 
 cd fluentd-kubernetes-daemonset
```{{execute}}


``` 
kubectl get svc  -n elastic-system 
```{{execute}}

hostname/dns: "quickstart-es-default.elastic-system"
port: 9200
user:elastic 
password: 3AlO1z58gb1801Xl86hVPvCW

Modify fluentd-daemonset-elasticsearch-rbac.yaml

results:
`
- name: fluentd
        image: fluent/fluentd-kubernetes-daemonset:v1-debian-elasticsearch
        env:
          - name:  FLUENT_ELASTICSEARCH_HOST
            value: "quickstart-es-default.elastic-system"
          - name:  FLUENT_ELASTICSEARCH_PORT
            value: "9200"
          - name: FLUENT_ELASTICSEARCH_SCHEME
            value: "http"
          # Option to configure elasticsearch plugin with self signed certs
          # ================================================================
          - name: FLUENT_ELASTICSEARCH_SSL_VERIFY
            value: "true"
          # Option to configure elasticsearch plugin with tls
          # ================================================================
          - name: FLUENT_ELASTICSEARCH_SSL_VERSION
            value: "TLSv1_2"
          # X-Pack Authentication
          # =====================
          - name: FLUENT_ELASTICSEARCH_USER
            value: "elastic"
          - name: FLUENT_ELASTICSEARCH_PASSWORD
            value: "3AlO1z58gb1801Xl86hVPvCW"
`


``` 
kubectl apply -f fluentd-daemonset-elasticsearch-rbac.yaml
```{{execute}}

Verify Pods running:
``` 
kubectl get pods -n kube-system
```{{execute}}

results:
`
NAME            READY   STATUS    RESTARTS   AGE
fluentd-5vvdj   1/1     Running   0          5s
`


