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
curl -kvso /dev/null --http2 --cert    client.crt \ /root/certs/prometheus.crt  https://localhost:8443

```{{execute T1}}

