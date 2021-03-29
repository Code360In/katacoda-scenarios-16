
# tobs - The Observability Stack for Kubernetes

https://github.com/timescale/tobs
https://blog.timescale.com/blog/introducing-tobs-deploy-a-full-observability-suite-for-kubernetes-in-two-minutes/

```
curl --proto '=https' --tlsv1.2 -sSLf  https://tsdb.co/install-tobs-sh |sh
```{{execute}}



```
 sudo cp /root/.tobs/bin/tobs /usr/local/bin
```{{execute}}



```
 tobs install
```{{execute}}



TimescaleDB stores and analyzes observability data over the long-term.

Promscale provides the power of PromQL for data stored in TimescaleDB.

Prometheus collects metric data across your cluster.

Grafana visualizes your data through customizable graphs and dashboards.

PromLens helps you build the PromQL queries you need to understand your system.

Kube-State-Metrics exposes metrics about the Kubernetes cluster itself.

Prometheus-Node-Exporter exposes metric data for nodes in your Kubernetes cluster.


![](tobs.png)