

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


 <pre class="file">
 Pod status running
 Nodes with status ready
 </pre>


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
get_helm.sh
```{{execute}}
