
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)  is a tool for running local [Kubernetes](https://kubernetes.io/) 
 clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.


# Install kind on Linux:

```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.10.0/kind-linux-amd64
```{{execute}}


```
chmod +x ./kind
```{{execute}}


```
mv ./kind /sbin/kind
```{{execute}}


# Creating a Cluster:
```
kind create cluster --name tx-cluster-k8s
```{{execute}}

Verify:
```
kind get clusters
```{{execute}}


# Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/):

```       
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```{{execute}}

Change to executable:
```       
chmod +x ./kubectl
```{{execute}}

Move to /sbin
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



# Install Kubernetes dashboard [dashboard](https://helm.sh/docs/intro/install/):

Add kubernetes-dashboard repository
``` 
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
```{{execute}}

Deploy a Helm Release named "dashboard" using the kubernetes-dashboard chart
``` 
kubectl create namespace dashboard
helm install dashboard kubernetes-dashboard/kubernetes-dashboard -n dashboard
```{{execute}}



Create a new account that can be used to authenticate to the dashboard:
``` 
$ kubectl create serviceaccount me -n kube-system
```{{execute}}

 
Make that new account have the right permissions:
``` 
$  kubectl create clusterrolebinding me -n kube-system --clusterrole=cluster-admin --serviceaccount=kube-system:me
```{{execute}}
 

Show the token needed to log in to the dashboard:
``` 
$ kubectl get secret $(kubectl get serviceaccount me -n kube-system -o jsonpath="{.secrets[0].name}") -o jsonpath="{.data.token}" -n kube-system | base64 --decode; echo
```{{execute}}