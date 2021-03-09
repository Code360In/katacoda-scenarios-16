
# Deploy Loki



``` 
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
```{{execute}}


``` 
helm upgrade --install loki grafana/loki-stack -n mon
```{{execute}}


Verify Pods running:
``` 
kubectl get pods -n mon
kubectl get svc -n mon
kubectl get deployments -n mon
kubectl get daemonset -n mon
kubectl get statefulset -n mon
```{{execute}}



Let's define a datasource for Loki on Grafana:

The URL for Loki should be:

`http://loki:3100`{{copy}}

For all the other fields you can use the default values

Test and save your datasource.


# Grafana Explore choose Loki.

Select pod/namespace/job/app/....

search for:

<pre class="file">
{pod=~".*"} |~  "lines"
{pod=~".*"} |~  "fail
{pod=~".*"} |~  "fault" 
{pod=~".*"} |~  "err"
{pod=~".*"} |~  "effects"
{pod=~".*"} |~  "panic"
{pod=~".*"} |~  "except"
{pod=~".*"} |~  "crash"
{pod=~".*"} |~  "kill"
{pod=~".*"} |~  "load"
{pod=~".*"} |~  "broken"
{pod=~".*"} |~  "down"
{pod=~".*"} |~  "mistakes"
{pod=~".*"} |~  "anomalies"
{pod=~".*"} |~  "bugs"
{pod=~".*"} |~  "glitches"
{pod=~".*"} |~  "critical"
{pod=~".*"} |~  "major"
{pod=~".*"} |~  "alert"
{pod=~".*"} |~  "warn"
{pod=~".*"} |~  "info"
</pre>
