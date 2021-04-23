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
 --name etcd quay.io/coreos/etcd:v2.3.8 \
 -name etcd0 \
 -advertise-client-urls http://${HostIP}:2379,http://${HostIP}:4001 \
 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -initial-advertise-peer-urls http://${HostIP}:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd0=http://${HostIP}:2380 \
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
ETCDCTL_API=3 etcdctl endpoint status --write-out=table --endpoints=https://[[HOST_IP]]:2379
```{{execute}}
