# Deploy Prometheus

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`/root/config/prometheus.yml`{{open}}


Start prometheus with command:


```
docker run  -d -p 9090:9090 \
    -v /root/config:/etc/prometheus \
    --name prometheus \
    --network pnet \
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




Remote Endpoints and Storage:
The remote write and remote read features of Prometheus allow transparently sending and receiving samples. This is primarily intended for long term storage. It is recommended that you perform careful evaluation of any solution in this space to confirm it can handle your data volumes.

https://prometheus.io/docs/operating/integrations/

<pre class="file">
remote_write:
  - url: "http://promscale:9201/write"
remote_read:
  - url: "http://promscale:9201/read"
    read_recent: true
</pre>
