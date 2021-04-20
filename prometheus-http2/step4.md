# GO HTTP2 exporter OpenTelemetry


# server-http2.go

```
cd /root
cat server-http2.go
```{{execute}}

```
go build server-http2.go
```{{execute}}

run:



```
echo t2
```{{execute T2}}


```
./server-http2
```{{execute T2}}


test:


only http2 --http2-prior-knowledge :

```
curl -vso /dev/null --http2-prior-knowledge --cacert /root/certs/prometheus.crt  https://localhost:8443

```{{execute T1}}


http/2 or fallback to http/1.1

```
curl -kvso /dev/null --http2 --cacert  /root/certs/prometheus.crt  https://localhost:8443

```{{execute T1}}


Results:

`
< HTTP/2 200 
`



close  ./server-http2



# prom-exporter-http2.go


```
cat  prom-exporter-http2.go
```{{execute T1}}


```
go mod init prom-exporter-http2
go mod verify
```{{execute T1}}

```
go build  prom-exporter-http2.go
```{{execute T1}}


run:
```
./prom-exporter-http2
```{{execute T1}}

test by curl:

only http2 --http2-prior-knowledge :

```
curl -vso /dev/null --http2-prior-knowledge --cacert /root/certs/prometheus.crt  https://localhost:8443/metrics
```{{execute T2 }}


http/2 or fallback to http/1.1

```
curl -kvso /dev/null --http2 --cacert /root/certs/prometheus.crt  https://localhost:8443/metrics
```{{execute T2}}


see metrics:
```
curl --http2-prior-knowledge --cacert /root/certs/prometheus.crt  https://localhost:8443/metrics
```{{execute T2}}



# Prometheus have debug enable
```
docker logs prometheus-federate
```{{execute T2}}


Verify if both target are up and running

https://[[HOST_SUBDOMAIN]]-9091-[[KATACODA_HOST]].environments.katacoda.com/targets


a) federate by http/1.1
b) opentelemetry by http/2


Add Grafana datasource to http://localhost:9091 , with "Skip TLS Verify"

Explore metrics on Grafana: ex_*

https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/explore?orgId=1&left=%5B%22now-1h%22,%22now%22,%22Prometheus%22,%7B%22exemplar%22:false,%22expr%22:%22ex_com_one%22,%22requestId%22:%22Q-66af286b-c9d6-4220-858f-5a59b4cf845f-0A%22%7D%5D

