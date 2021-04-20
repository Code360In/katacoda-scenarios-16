# Deploy Prometheus Read

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus-read.yml`{{open}}



```
ls
cat /root/prometheus-read.yml
```{{execute}}




Start prometheus with command:



Deploy Prometheus Read:

```
cd /root/
docker run -u root --net=host -d -p 9091:9091 \
    -v $PWD/certs:/certs \
    -v $PWD/prometheus-read.yml:/etc/prometheus/prometheus.yml \
    --name prometheus-read \
    prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.enable-admin-api --web.listen-address=:9092 --log.level=debug
```{{execute}}


You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9092-[[KATACODA_HOST]].environments.katacoda.com/targets


![](prom.png)


Prometheus Metrics:

https://[[HOST_SUBDOMAIN]]-9092-[[KATACODA_HOST]].environments.katacoda.com/metrics
