# Deploy Prometheus

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus.yml`{{open}}


```
ls
cat /root/prometheus.yml
echo
```{{execute}}



Start prometheus with command:



```
docker run  -d -p 9090:9090 \
    --net=host  -u root  \
    -v $PWD/prometheus.yml:/etc/prometheus/prometheus.yml \
    --name prometheus \
    prom/prometheus
```{{execute}}




Logs:
```
docker logs prometheus
```{{execute}}



You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/targets




Prometheus Metrics:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/metrics
