# Deploy Etcd


https://etcd.io/docs/v2.3/docker_guide/

## Running etcd in standalone mode

`
In order to expose the etcd API to clients outside of the Docker host you’ll need use the host IP address when configuring etc
`


```
export HostIP="[[HOST_IP]]"
```{{execute}}


`
This will run the latest release version of etcd. You can specify version if needed (e.g. quay.io/coreos/etcd:v2.2.0).
`


```
docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd quay.io/coreos/etcd:v3.2.32 \
 -name etcd \
 -advertise-client-urls http://${HostIP}:2379,http://${HostIP}:4001 \
 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -initial-advertise-peer-urls http://${HostIP}:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd=http://${HostIP}:2380 \
 -initial-cluster-state new
```{{execute}}



# Verify Metrics

```
curl http://localhost:2379/metrics
```{{execute}}


or

https://[[HOST_SUBDOMAIN]]-2379-[[KATACODA_HOST]].environments.katacoda.com/metrics


# Install etcdctl

```
apt install etcd-client
```{{execute}}


```
ETCDCTL_API=3 etcdctl endpoint status --write-out=table --endpoints=http://[[HOST_IP]]:2379  --insecure-skip-tls-verify
```{{execute}}


Verification:

```
export ETCDCTL_API=3
```{{execute}}


```
etcdctl version
```{{execute}}


httpie
https://httpie.io/docs

```
apt install httpie
```{{execute}}


```
 http POST http://localhost:2379/metrics
```{{execute}}


https://etcd.io/docs/v2.3/api/


```
http GET http://127.0.0.1:2379/version
```{{execute}}


<pre class="file">

HTTP/1.1 200 OK
Content-Length: 44
Content-Type: application/json
Date: Fri, 23 Apr 2021 16:03:10 GMT

{
    "etcdcluster": "2.3.0",
    "etcdserver": "2.3.8"
}

</pre>


Insert:
```
http -v PUT http://127.0.0.1:2379/v2/keys/message value="Hello world" 
```{{execute}}

<pre class="file">

PUT /v2/keys/message HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 24
Content-Type: application/json
Host: 127.0.0.1:2379
User-Agent: HTTPie/1.0.3

{
    "value": "Hello world"
}

HTTP/1.1 200 OK
Content-Length: 165
Content-Type: application/json
Date: Fri, 23 Apr 2021 16:05:11 GMT
X-Etcd-Cluster-Id: 77c65fb74b0d401d
X-Etcd-Index: 6
X-Raft-Index: 1192
X-Raft-Term: 2

{
    "action": "set",
    "node": {
        "createdIndex": 6,
        "key": "/message",
        "modifiedIndex": 6,
        "value": ""
    },
    "prevNode": {
        "createdIndex": 5,
        "key": "/message",
        "modifiedIndex": 5,
        "value": ""
    }
}

</pre>


GET
```
$ http -v  get http://127.0.0.1:2379/v2/keys/message
```{{execute}}

<pre class="file">

GET /v2/keys/message HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Host: 127.0.0.1:2379
User-Agent: HTTPie/1.0.3



HTTP/1.1 200 OK
Content-Length: 89
Content-Type: application/json
Date: Fri, 23 Apr 2021 16:06:59 GMT
X-Etcd-Cluster-Id: 77c65fb74b0d401d
X-Etcd-Index: 6
X-Raft-Index: 1406
X-Raft-Term: 2

{
    "action": "get",
    "node": {
        "createdIndex": 6,
        "key": "/message",
        "modifiedIndex": 6,
        "value": ""
    }
}
</pre>


Using key TTL
Keys in etcd can be set to expire after a specified number of seconds. 
You can do this by setting a TTL (time to live) on the key when sending a PUT request:

```
curl http://localhost:2379/v2/keys/foo -XPUT -d value=bar -d ttl=5
```{{execute}}



#  benchmark tool

https://etcd.io/docs/v3.2/op-guide/performance/#benchmarks

https://github.com/etcd-io/etcd/tree/master/tools/benchmark

```
go get go.etcd.io/etcd/tools/benchmark
```{{execute}}


# GOPATH should be set

```
ls $GOPATH/bin
benchmark
```{{execute}}




# write to leader

Number of connections 1:

```
benchmark --endpoints=[[HOST_IP]] --target-leader --conns=1 --clients=1 \
    put --key-size=8 --sequential-keys --total=10000 --val-size=256
```{{execute}}


Number of connections 100:

```
benchmark --endpoints=[[HOST_IP]] --target-leader  --conns=100 --clients=1000 \
    put --key-size=8 --sequential-keys --total=100000 --val-size=256
```{{execute}}

