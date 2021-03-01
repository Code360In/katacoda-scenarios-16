
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)  is a tool for running local [Kubernetes](https://kubernetes.io/) 
 clusters using Docker container “nodes”.
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




Install kubectl

```       
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```{{execute}}


```       
chmod +x ./kubectl
```{{execute}}


```       
mv ./kubectl /sbin/kubectl
```{{execute}}



Verify context:

```       
kubectl cluster-info --context kind-2
```{{execute}}

