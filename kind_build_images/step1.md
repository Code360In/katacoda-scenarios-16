
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)  is a tool for running local [Kubernetes](https://kubernetes.io/) 
 clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.


# Install kind:

```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.11.1/kind-linux-amd64
```{{execute}}


```
chmod +x ./kind
```{{execute}}


```
mv ./kind /sbin/kind
```{{execute}}


# Build Kubernetes images


```
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone https://github.com/kubernetes/kubernetes
```{{execute}}

```
kind build node-image
```{{execute}}

wait for: 
"Image build completed."

Verify

```
docker images
```{{execute}}

<pre class="file">

REPOSITORY                           TAG                            IMAGE ID            CREATED             SIZE
kindest/node                         latest                         cd68e297a35c        2 minutes ago       1.3GB

</pre>

# Creating a Kubernetes Cluster:
```
kind create cluster --image kindest/node:latest 
```{{execute}}


Verify:
```
kubectl cluster-info --context kind-kind
```{{execute}}

