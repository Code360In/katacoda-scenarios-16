#### Launch Cluster

`launch.sh`{{execute}}

This will create a two node Kubernetes cluster using WeaveNet for networking.

#### Health Check

`
kubectl cluster-info
`{{execute}}

`
kubectl get nodes
`{{execute}}



# kubectl Cheat Sheet


https://kubernetes.io/docs/reference/kubectl/cheatsheet/


##  Kubectl autocomplete
BASH

setup autocomplete in bash into the current shell, bash-completion package should be installed first.
```
source <(kubectl completion bash) 
```{{execute}}


add autocomplete permanently to your bash shell.
```
echo "source <(kubectl completion bash)" >> ~/.bashrc 
```{{execute}}




## Kubectl context and configuration

Set which Kubernetes cluster kubectl communicates with and modifies configuration information. See Authenticating Across Clusters with kubeconfig documentation for detailed config file information.


Show Merged kubeconfig settings.
```
kubectl config view
```{{execute}}



get the password for the e2e user
```
kubectl config view -o jsonpath='{.users[?(@.name == "e2e")].user.password}'
```{{execute}}


display the first user
```
kubectl config view -o jsonpath='{.users[].name}'
```{{execute}}


get a list of users
```
kubectl config view -o jsonpath='{.users[*].name}'
```{{execute}}


display list of contexts 
```
kubectl config get-contexts                        
```{{execute}}



display the current-context
```
kubectl config current-context                      
```{{execute}}


set the default context to my-cluster-name
```
kubectl config use-context kubernetes-admin@kubernetes     
```{{execute}}


add a new user to your kubeconf that supports basic auth
```
kubectl config set-credentials kubeuser/foo.kubernetes.com --username=kubeuser --password=kubepassword
```{{execute}}

permanently save the namespace for all subsequent kubectl commands in that context.
```
kubectl config set-context --current --namespace=ggckad-s2
```{{execute}}


set a context utilizing a specific username and namespace.
```
kubectl config set-context gce --user=cluster-admin --namespace=foo \
  && kubectl config use-context gce
```{{execute}}


delete user foo
```
kubectl config unset users.foo                      
```{{execute}}


## Kubectl apply
apply manages applications through files defining Kubernetes resources. It creates and updates resources in a cluster through running kubectl apply. This is the recommended way of managing Kubernetes applications on production. See Kubectl Book.

Creating objects
Kubernetes manifests can be defined in YAML or JSON. The file extension .yaml, .yml, and .json can be used.


## Creating objects 

Kubernetes manifests can be defined in YAML or JSON. The file extension .yaml, .yml, and .json can be used.

create resource(s)

```
kubectl apply -f ./my-manifest.yaml
```{{execute}}

create from multiple files
```
kubectl apply -f ./my1.yaml -f ./my2.yaml
```{{execute}}


create resource(s) in all manifest files in dir
```
kubectl apply -f ./dir
```{{execute}}


create resource(s) from url
```
kubectl apply -f https://git.io/vPieo
```{{execute}}


start a single instance of nginx
```
kubectl create deployment nginx --image=nginx
```{{execute}}


create a Job which prints "Hello World"
```
kubectl create job hello --image=busybox -- echo "Hello World" 
```{{execute}}

create a CronJob that prints "Hello World" every minute
```
kubectl create cronjob hello --image=busybox   --schedule="*/1 * * * *" -- echo "Hello World"    
```{{execute}}


get the documentation for pod manifests
```
kubectl explain pods
```{{execute}}

Create multiple YAML objects from stdin
```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: busybox-sleep
spec:
  containers:
  - name: busybox
    image: busybox
    args:
    - sleep
    - "1000000"
---
apiVersion: v1
kind: Pod
metadata:
  name: busybox-sleep-less
spec:
  containers:
  - name: busybox
    image: busybox
    args:
    - sleep
    - "1000"
EOF
```{{execute}}


Create a secret with several keys
```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  password: $(echo -n "s33msi4" | base64 -w0)
  username: $(echo -n "jane" | base64 -w0)
EOF
```{{execute}}
