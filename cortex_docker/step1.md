
#### Launch  Kubernetes Cluster

`launch.sh`{{execute}}

This will create a two node Kubernetes cluster using WeaveNet for networking.

#### Health Check

`
kubectl cluster-info
`{{execute}}


# Install [helm](https://helm.sh/docs/intro/install/):


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
helm ls```{{execute}}


Remove sh file:
```       
rm get_helm.sh
```{{execute}}
