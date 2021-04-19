# GO HTTP2


# server-http2.go

```
cd /root
more server-http2.go
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
more  prom-exporter-http2.go

```{{execute T1}}

```
go build  prom-exporter-http2.go
./prom-exporter-http2.go
```{{execute T1}}

run:
```
./prom-exporter-http2.go
```{{execute T1}}

test by curl:

only http2 --http2-prior-knowledge :

```
curl -vso /dev/null --http2-prior-knowledge --cacert /root/certs/prometheus.crt  https://localhost:8443/metrics

```{{execute T2 }


http/2 or fallback to http/1.1

```
curl -kvso /dev/null --http2 --cacert /root/certs/prometheus.crt  https://localhost:8443/metrics

```{{execute T2}}
