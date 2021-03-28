
# Alertmanager

https://github.com/prometheus/alertmanager
https://github.com/timescale/promscale/blob/master/docs/alerting-recording.md

```
docker run  --name alertmanager -d  -v /root/alertmanager.yml:/alertmanager.yml  --net=pnet -p 9093:9093 quay.io/prometheus/alertmanager --config.file=/alertmanager.yml
```{{execute}}



Verify:

```
docker ps -a
```{{execute}}

```
docker logs prometheus
```{{execute}}




AlertManager:

https://[[HOST_SUBDOMAIN]]-9093-[[KATACODA_HOST]].environments.katacoda.com


Rules:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/rules

![](prom-rules.png)

Alerts:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/alerts

![](prom-alerts.png)