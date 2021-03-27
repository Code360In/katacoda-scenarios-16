# [cAdvisor](https://github.com/google/cadvisor)


```
docker run -d \
  --net=pnet \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:ro \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --volume=/dev/disk/:/dev/disk:ro \
  --publish=8080:8080 \
  --detach=true \
  --name=cadvisor \
  --privileged \
  --device=/dev/kmsg \
  gcr.io/cadvisor/cadvisor:latest
```{{execute T2}}


If you're interested in seeing the raw metrics, they can be viewed with 

```
curl localhost:8080/metrics
```{{execute T2}}


## Build a dashboard

In order to build a dashboard you can build one from scratch or you can use an existing one, for example:
https://grafana.com/grafana/dashboards/893

Let's use this existing dashboard. Copy json , and use the option `Import`.


`/root/grafana/grafana-dashboard-cnt.json`{{open}}


Select **Promscale** as the data source and Import.


![](cnt.png)