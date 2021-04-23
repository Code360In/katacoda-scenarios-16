# Deploy Etcd


https://etcd.io/docs/v3.2/op-guide/container/


## Running etcd in standalone mode

`
In order to expose the etcd API to clients outside of the Docker host you’ll need use the host IP address when configuring etc
`
```
host ip [[HOST_IP]]
```

```
export NODE1="127.0.0.1"
```{{execute}}




`
This will run the latest release version of etcd. You can specify version if needed (e.g. quay.io/coreos/etcd:v2.2.0).
`




```
docker run -d --net=host\
  -p 2379:2379 \
  -p 2380:2380 \
  --name etcd quay.io/coreos/etcd:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://${NODE1}:2380 \
  --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://${NODE1}:2379 \
  --initial-cluster node1=http://${NODE1}:2380

```{{execute}}

Logs:
```
docker logs etcd
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



List the cluster member:

```
etcdctl --endpoints=http://${NODE1}:2379 member list
```{{execute}}


status:

```
ETCDCTL_API=3 etcdctl endpoint status --write-out=table --endpoints=http://${NODE1}:2379  
```{{execute}}


<pre class="file">
+-----------------------+------------------+---------+---------+-----------+-----------+------------+
|       ENDPOINT        |        ID        | VERSION | DB SIZE | IS LEADER | RAFT TERM | RAFT INDEX |
+-----------------------+------------------+---------+---------+-----------+-----------+------------+
| http://127.0.0.1:2379 | b71f75320dc06a6c |   3.3.8 |   20 kB |      true |         2 |          4 |
+-----------------------+------------------+---------+---------+-----------+-----------+------------+
</pre>


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
Date: Fri, 23 Apr 2021 17:56:35 GMT

{
    "etcdcluster": "3.3.0",
    "etcdserver": "3.3.8"
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

HTTP/1.1 201 Created
Content-Length: 89
Content-Type: application/json
Date: Fri, 23 Apr 2021 17:57:17 GMT
X-Etcd-Cluster-Id: 1c45a069f3a1d796
X-Etcd-Index: 4
X-Raft-Index: 5
X-Raft-Term: 2

{
    "action": "set",
    "node": {
        "createdIndex": 4,
        "key": "/message",
        "modifiedIndex": 4,
        "value": ""
    }
}

</pre>


GET
```
http -v  get http://127.0.0.1:2379/v2/keys/message
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
Date: Fri, 23 Apr 2021 17:57:39 GMT
X-Etcd-Cluster-Id: 1c45a069f3a1d796
X-Etcd-Index: 4
X-Raft-Index: 5
X-Raft-Term: 2

{
    "action": "get",
    "node": {
        "createdIndex": 4,
        "key": "/message",
        "modifiedIndex": 4,
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
benchmark  --target-leader --conns=1 --clients=1 \
    put --key-size=8 --sequential-keys --total=10000 --val-size=256
```{{execute}}

results:
  <pre class="file">
 0 / 10000 B                                                                                           !   0.00%INFO: 2021/04/23 17:00:10 [core] Channel Connectivity change to CONNECTING
INFO: 2021/04/23 17:00:10 [core] Subchannel picks a new address "127.0.0.1:2379" to connect
INFO: 2021/04/23 17:00:10 [core] Subchannel Connectivity change to READY
INFO: 2021/04/23 17:00:10 [roundrobin] roundrobinPicker: newPicker called with info: {map[0xc000042a50:{{127.0.0.1:2379 127.0.0.1 <nil> 0 <nil>}}]}
INFO: 2021/04/23 17:00:10 [core] Channel Connectivity change to READY
 10000 / 10000 Booooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo! 100.00% 19s

Summary:
  Total:        19.0216 secs.
  Slowest:      0.4539 secs.
  Fastest:      0.0003 secs.
  Average:      0.0019 secs.
  Stddev:       0.0094 secs.
  Requests/sec: 525.7181

Response time histogram:
  0.0003 [1]    |
  0.0456 [9988] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.0910 [3]    |
  0.1364 [1]    |
  0.1817 [1]    |
  0.2271 [0]    |
  0.2724 [1]    |
  0.3178 [1]    |
  0.3632 [2]    |
  0.4085 [1]    |
  0.4539 [1]    |

Latency distribution:
  10% in 0.0004 secs.
  25% in 0.0005 secs.
  50% in 0.0006 secs.
  75% in 0.0008 secs.
  90% in 0.0033 secs.
  95% in 0.0090 secs.
  99% in 0.0181 secs.
  99.9% in 0.0559 secs.
  </pre>


Number of connections 100:

```
benchmark  --target-leader  --conns=100 --clients=1000 \
    put --key-size=8 --sequential-keys --total=100000 --val-size=256
```{{execute}}


results:
  <pre class="file">
INFO: 2021/04/23 17:02:27 [core] Channel Connectivity change to READY
 100000 / 100000 Booooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo! 100.00% 10s

Summary:
  Total:        10.6759 secs.
  Slowest:      0.4662 secs.
  Fastest:      0.0068 secs.
  Average:      0.1053 secs.
  Stddev:       0.0685 secs.
  Requests/sec: 9366.9291

Response time histogram:
  0.0068 [1]    |
  0.0527 [13254]        |∎∎∎∎∎∎∎∎∎∎
  0.0986 [48218]        |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.1446 [22766]        |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.1905 [7269] |∎∎∎∎∎∎
  0.2365 [2612] |∎∎
  0.2824 [1758] |∎
  0.3283 [1093] |
  0.3743 [1848] |∎
  0.4202 [1066] |
  0.4662 [115]  |

Latency distribution:
  10% in 0.0487 secs.
  25% in 0.0652 secs.
  50% in 0.0875 secs.
  75% in 0.1200 secs.
  90% in 0.1733 secs.
  95% in 0.2573 secs.
  99% in 0.3840 secs.
  99.9% in 0.4209 secs.

  </pre>
