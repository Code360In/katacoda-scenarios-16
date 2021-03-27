# Deploy Prometheus

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus.yml`{{open}}


Start prometheus with command:


```
docker run  -d -p 9090:9090 \
    -v /root/config:/etc/prometheus \
    --name prometheus \
    --network promscale-timescaledb \
    prom/prometheus
```{{execute}}


You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/targets




Prometheus Metrics:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/metrics



Verify:

```
docker ps -a
```{{execute}}

```
docker logs prometheus
```{{execute}}
