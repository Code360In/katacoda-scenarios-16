# Deploy Node Exporter

To collect metrics related to a node it's required to run a Prometheus Node Exporter. Prometheus has many exporters that are designed to output metrics for a particular system, such as Postgres or MySQL.

Task
Launch the Node Exporter container. By mounting the host /proc and /sys directory, the container has accessed to the necessary information to report on.

https://github.com/prometheus/node_exporter 

```
docker run -d \
  --net=pnet -p 9100:9100 \
  --name=nodeexporter \
  prom/node-exporter:latest
```{{execute}}


If you're interested in seeing the raw metrics, they can be viewed with 

```
curl localhost:9100/metrics
```{{execute}}

or

https://[[HOST_SUBDOMAIN]]-9100-[[KATACODA_HOST]].environments.katacoda.com/metrics


## Build a dashboard

In order to build a dashboard you can build one from scratch or you can use an existing one, for example:
https://grafana.com/grafana/dashboards/1860

Let's use this existing dashboard. Copy the ID, and use the option `Import`.

`1860`{{copy}}

Select **Prometheus** as the data source and Import.


[View Dashboard for the targetCluster](https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]])


ref:
https://hub.docker.com/r/prom/node-exporter
