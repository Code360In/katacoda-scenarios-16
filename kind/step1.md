
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)  is a tool for running local Kubernetes clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.


Quick Start

Install on Linux:

```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.10.0/kind-linux-amd64
```{{execute}}


```
chmod +x ./kind
```{{execute}}


```
mv ./kind /sbin/kind
```{{execute}}


Creating a Cluster:


```
kind create cluster --name kind-2
```{{execute}}


```
kind get clusters
```{{execute}}


```       
kubectl cluster-info --context kind-2
```{{execute}}