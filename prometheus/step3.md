# Deploy Prometheus Federate

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus.yml`{{open}}


Start prometheus with command:


Generate tls keys:

You can generate a self-signed certificate and private key using this command:

```
mkdir -p /root/certs/localhost && cd /root/certs/localhost
openssl req \
  -x509 \
  -newkey rsa:4096 \
  -nodes \
  -keyout tls.key \
  -out tls.crt
```{{execute}}


```
docker run --net=host -d -p 9091:9091 \
    -v $PWD/certs/localhost:/certs/localhost \
    -v $PWD/prometheus-federate.yml:/etc/prometheus/prometheus.yml \
    -v $PWD/prometheus-federate-tls.yml:/etc/prometheus/prometheus-tls.yml \
    --name prometheus-federate \
    prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.config.file=/etc/prometheus/prometheus-tls.yml --web.enable-admin-api --web.listen-address=:9091
```{{execute}}


You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/targets


![](prom.png)


Prometheus Metrics:

https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/metrics


test vy curl:

federate?match[]={__name__=~".*"}
```
curl http://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D

```{{execute}}
