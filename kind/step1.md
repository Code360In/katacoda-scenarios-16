
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)  is a tool for running local [Kubernetes](https://kubernetes.io/) 
 clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.


# Install kind:

```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.10.1/kind-linux-amd64
```{{execute}}


```
chmod +x ./kind
```{{execute}}


```
mv ./kind /sbin/kind
```{{execute}}


# Creating a Kubernetes Cluster:
```
kind create cluster --name tx-cluster-k8s
```{{execute}}

Verify:
```
kind get clusters
```{{execute}}

