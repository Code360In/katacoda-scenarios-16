# GO HTTP2



```
cd /root
more server-http2.go
```{{execute}}

```
go build server-http2.go
```{{execute}}

run:

```
./server-http2
```{{execute}}


test:


only http2 --http2-prior-knowledge :

```
curl -vso /dev/null --http2-prior-knowledge --cacert /root/certs/prometheus.crt  https://localhost:8443

```{{execute}}


http/2 or fallback to http/1.1

```
curl -kvso /dev/null --http2 --cert    client.crt \ /root/certs/prometheus.crt  https://localhost:8443

```{{execute}}

