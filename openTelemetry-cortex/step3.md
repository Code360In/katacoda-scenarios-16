# Deploy Prometheus

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus.yml`{{open}}


```
ls
cat /root/prometheus.yml
cd /root/
```{{execute}}



Start prometheus with command:



```
docker run -u root -d -p 9090:9090 \
    -v /root/prometheus.yml:/etc/prometheus/prometheus.yml \
    --name prometheus \
    prom/prometheus --config.file=/etc/prometheus/prometheus.yml  --enable-feature=remote-write-receiver --enable-feature=exemplar-storage
```{{execute}}



``` 
docker network connect np prometheus
```{{execute}}




You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/targets



Prometheus Metrics:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/metrics


Add datasource to Grafana

url  ```http://prometheus:9090```{{copy}}


to use Prometheus as remote-write-receiver  on otel.

change config.yam url to url: ```http://localhost:9090/api/v1/write```{{copy}}