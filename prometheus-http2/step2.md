# Deploy Prometheus Federate

Prometheus is an open-source monitoring solution.

Now, let's see the prometheus configuration:

`prometheus-federate.yml`{{open}}


Start prometheus with command:


Generate tls keys:

You can generate a self-signed certificate and private key using this command:


```
mkdir certs
cd certs
openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout prometheus.key -out prometheus.crt -subj "/C=BE/ST=Antwerp/L=Brasschaat/O=Inuits/CN=localhost" -addext "subjectAltName = DNS:localhost"  
```{{execute}}


Deploy Prometheus Federate:

```
cd /root/
docker run -u root --net=host -d -p 9091:9091 \
    -v $PWD/certs:/certs \
    -v $PWD/prometheus-federate.yml:/etc/prometheus/prometheus.yml \
    -v $PWD/prometheus-federate-tls.yml:/etc/prometheus/tls.yml \
    --name prometheus-federate \
    prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.config.file=/etc/prometheus/tls.yml --web.enable-admin-api --web.listen-address=:9091 --log.level=debug
```{{execute}}


You can access to the prometheus dashboard using this link:

https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/targets


![](prom.png)


Prometheus Metrics:

https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/metrics


test by curl:

federate?match[]={__name__=~".*"}

only http2 --http2-prior-knowledge :

```
curl -vso /dev/null --http2-prior-knowledge --cacert /root/certs/prometheus.crt  https://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D

```{{execute}}


http/2 or fallback to http/1.1

```
curl -kvso /dev/null --http2 --cacert /root/certs/prometheus.crt  https://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D

```{{execute}}
