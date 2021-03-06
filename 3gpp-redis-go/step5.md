# Deploy Prometheus

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus.yml`{{open}}


Start prometheus with command:

```
docker run --net=host -d -p 9090:9090 \
    -v /root/prometheus.yml:/etc/prometheus/prometheus.yml \
    --name prometheus-server \
    prom/prometheus
```{{execute T4}}

You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/targets


![](prom.png)
