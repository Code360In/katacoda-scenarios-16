# Interacting with running Pods

dump pod logs (stdout)
```
kubectl logs my-pod
```{{execute}}

dump pod logs, with label name=myLabel (stdout)
```
kubectl logs -l name=myLabel
```{{execute}}

dump pod logs (stdout) for a previous instantiation of a container
```
kubectl logs my-pod --previous
```{{execute}}

dump pod container logs (stdout, multi-container case)
```
kubectl logs my-pod -c my-container
```{{execute}}

dump pod logs, with label name=myLabel (stdout)
```
kubectl logs -l name=myLabel -c my-container
```{{execute}}


dump pod container logs (stdout, multi-container case) for a previous instantiation of a container
```
kubectl logs my-pod -c my-container --previous      
```{{execute}}


stream pod logs (stdout)
```
kubectl logs -f my-pod
```{{execute}}

stream pod container logs (stdout, multi-container case)
```
kubectl logs -f my-pod -c my-container
```{{execute}}

stream all pods logs with label name=myLabel (stdout)
```
kubectl logs -f -l name=myLabel --all-containers
```{{execute}}


Run pod as interactive shell
```
kubectl run -i --tty busybox --image=busybox -- sh  
```{{execute}}

Run pod nginx in a specific namespace
```
kubectl run nginx --image=nginx -n mynamespace
```{{execute}}


Run pod nginx and write its spec into a file called pod.yaml
```
kubectl run nginx --image=nginx --dry-run=client -o yaml > pod.yaml
```{{execute}}


Attach to Running Container
```
kubectl attach my-pod -i
```{{execute}}

Listen on port 5000 on the local machine and forward to port 6000 on my-pod
```
kubectl port-forward my-pod 5000:6000
```{{execute}}


Run command in existing pod (1 container case)
```
kubectl exec my-pod -- ls /                         
```{{execute}}


Interactive shell access to a running pod (1 container case) 
```
kubectl exec --stdin --tty my-pod -- /bin/sh        
```{{execute}}

Run command in existing pod (multi-container case)
```
kubectl exec my-pod -c my-container -- ls /
```{{execute}}


Show metrics for a given pod and its containers
```
kubectl top pod POD_NAME --containers
```{{execute}}

Show metrics for a given pod and sort it by 'cpu' or 'memory'
```
kubectl top pod POD_NAME --sort-by=cpu
```{{execute}}

# Interacting with Deployments and Services


dump Pod logs for a Deployment (single-container case)
```
kubectl logs deploy/my-deployment                         
```{{execute}}


dump Pod logs for a Deployment (multi-container case)
```
kubectl logs deploy/my-deployment -c my-container
```{{execute}}


listen on local port 5000 and forward to port 5000 on Service backend
```
kubectl port-forward svc/my-service 5000
```{{execute}}


listen on local port 5000 and forward to Service target port with name <my-service-port>
```
kubectl port-forward svc/my-service 5000:my-service-port
```{{execute}}


listen on local port 5000 and forward to port 6000 on a Pod created by <my-deployment>
```
kubectl port-forward deploy/my-deployment 5000:6000
```{{execute}}


run command in first Pod and first container in Deployment (single- or multi-container cases)
```
kubectl exec deploy/my-deployment -- ls
```{{execute}}

# Interacting with Nodes and cluster

Mark my-node as unschedulable
```
kubectl cordon my-node
```{{execute}}


Drain my-node in preparation for maintenance
```
kubectl drain my-node
```{{execute}}


Mark my-node as schedulable
```
kubectl uncordon my-node
```{{execute}}


Show metrics for a given node
```
kubectl top node my-node
```{{execute}}

Display addresses of the master and services
```
kubectl cluster-info
```{{execute}}


Dump current cluster state to stdout
```
kubectl cluster-info dump
```{{execute}}


Dump current cluster state to /path/to/cluster-state
```
kubectl cluster-info dump --output-directory=/path/to/cluster-state
```{{execute}}


# If a taint with that key and effect already exists, its value is replaced as specified.

```
kubectl taint nodes foo dedicated=special-user:NoSchedule
Resource types
```{{execute}}

List all supported resource types along with their shortnames, API group, whether they are namespaced, and Kind:

```
kubectl api-resources
```{{execute}}

Other operations for exploring API resources:

All namespaced resources
```
kubectl api-resources --namespaced=true
```{{execute}}


All non-namespaced resources
```
kubectl api-resources --namespaced=false
```{{execute}}


All resources with simple output (only the resource name)
```
kubectl api-resources -o name
```{{execute}}


All resources with expanded (aka "wide") output
```
kubectl api-resources -o wide
```{{execute}}

All resources that support the "list" and "get" request verbs
```
kubectl api-resources --verbs=list,get
```{{execute}}

All resources in the "extensions" API group
```
kubectl api-resources --api-group=extensions 
```{{execute}}


# Formatting output

All images running in a cluster
```
kubectl get pods -A -o=custom-columns='DATA:spec.containers[*].image'
```{{execute}}


All images running in namespace: default, grouped by Pod
```
kubectl get pods --namespace default --output=custom-columns="NAME:.metadata.name,IMAGE:.spec.containers[*].image"
```{{execute}}


All images excluding "k8s.gcr.io/coredns:1.6.2"
```
kubectl get pods -A -o=custom-columns='DATA:spec.containers[?(@.image!="k8s.gcr.io/coredns:1.6.2")].image'
```{{execute}}

All fields under metadata regardless of name
```
kubectl get pods -A -o=custom-columns='DATA:metadata.*'
```{{execute}}