
# Alertmanager

https://github.com/prometheus/alertmanager
https://github.com/timescale/promscale/blob/master/docs/alerting-recording.md

```
docker run  --name alertmanager -d  -v /root/alert/:/alert  --net=pnet -p 9093:9093 quay.io/prometheus/alertmanager -config.file=/alert/alertmanager.conf
```{{execute}}


AlertManager:

https://[[HOST_SUBDOMAIN]]-9093-[[KATACODA_HOST]].environments.katacoda.com


Rules:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/rules


Alerts:

https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com/alerts
