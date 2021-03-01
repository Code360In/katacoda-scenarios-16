
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)  is a tool for running local [Kubernetes](https://kubernetes.io/) 
 clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.


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
kind create cluster --name tx-cluster-k8s
```{{execute}}


```
kind get clusters
```{{execute}}




Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/):

```       
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```{{execute}}


```       
chmod +x ./kubectl
```{{execute}}


```       
mv ./kubectl /sbin/kubectl
```{{execute}}


Verify cluster:
```       
kubectl cluster-info
```{{execute}}


Display nodes:
```       
kubectl get nodes
```{{execute}}

Display Pods:
```       
kubectl get pods -A
```{{execute}}


Install [helm](https://helm.sh/docs/intro/install/):

```       
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
```{{execute}}

```       
chmod 700 get_helm.sh
```{{execute}}

```       
./get_helm.sh
```{{execute}}

Verify:

```       
helmn ls```{{execute}}

