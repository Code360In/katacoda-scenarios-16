#### Launch Cluster

`launch.sh`{{execute}}

This will create a two node Kubernetes cluster using WeaveNet for networking.

#### Health Check

`
kubectl cluster-info
`{{execute}}


# kubectl Cheat Sheet


https://kubernetes.io/docs/reference/kubectl/cheatsheet/


Kubectl autocomplete
BASH

```
source <(kubectl completion bash) # setup autocomplete in bash into the current shell, bash-completion package should be installed first.
echo "source <(kubectl completion bash)" >> ~/.bashrc # add autocomplete permanently to your bash shell.
```{{execute}}


Kubectl context and configuration
Set which Kubernetes cluster kubectl communicates with and modifies configuration information. See Authenticating Across Clusters with kubeconfig documentation for detailed config file information.

```
kubectl config view # Show Merged kubeconfig settings.
```{{execute}}



# get the password for the e2e user
```
kubectl config view -o jsonpath='{.users[?(@.name == "e2e")].user.password}'
```{{execute}}

```
kubectl config view -o jsonpath='{.users[].name}'    # display the first user
```{{execute}}

```
kubectl config view -o jsonpath='{.users[*].name}'   # get a list of users
```{{execute}}

```
kubectl config get-contexts                          # display list of contexts 
```{{execute}}


```
kubectl config current-context                       # display the current-context
```{{execute}}


```
kubectl config use-context my-cluster-name           # set the default context to my-cluster-name
```{{execute}}


# add a new user to your kubeconf that supports basic auth
```
kubectl config set-credentials kubeuser/foo.kubernetes.com --username=kubeuser --password=kubepassword
```{{execute}}

# permanently save the namespace for all subsequent kubectl commands in that context.
```
kubectl config set-context --current --namespace=ggckad-s2
```{{execute}}


# set a context utilizing a specific username and namespace.
```
kubectl config set-context gce --user=cluster-admin --namespace=foo \
  && kubectl config use-context gce
```{{execute}}


```
kubectl config unset users.foo                       # delete user foo
```{{execute}}

