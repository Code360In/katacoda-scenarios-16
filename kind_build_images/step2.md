

# Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/):


```       
cp /root/go/src/k8s.io/kubernetes/_output/dockerized/bin/linux/amd64/kubectl  /sbin/kubectl
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
rm get_helm.sh
```{{execute}}
