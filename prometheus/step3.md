# Deploy Prometheus Federate

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus.yml`{{open}}


Start prometheus with command:



```
docker run --net=host -d -p 9091:9091 \
    -v $PWD/prometheus-federate.yml:/etc/prometheus/prometheus.yml \
    --name prometheus-federate \
    prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.enable-admin-api --web.listen-address=:9091
```{{execute}}


You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/targets


![](prom.png)


Prometheus Metrics:

https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/metrics


test vy curl:

federate?match[]={__name__=~".*"}
```
curl http://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D

```{{execute}}
