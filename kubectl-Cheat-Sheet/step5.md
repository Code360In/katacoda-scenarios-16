# Patching resources


Partially update a node
```
kubectl patch node k8s-node-1 -p '{"spec":{"unschedulable":true}}'
```{{execute}}

Update a container's image; spec.containers[*].name is required because it's a merge key
```
kubectl patch pod valid-pod -p '{"spec":{"containers":[{"name":"kubernetes-serve-hostname","image":"new image"}]}}'
```{{execute}}

Update a container's image using a json patch with positional arrays
```
kubectl patch pod valid-pod --type='json' -p='[{"op": "replace", "path": "/spec/containers/0/image", "value":"new image"}]'
```{{execute}}

Disable a deployment livenessProbe using a json patch with positional arrays
```
kubectl patch deployment valid-deployment  --type json   -p='[{"op": "remove", "path": "/spec/template/spec/containers/0/livenessProbe"}]'
```{{execute}}

Add a new element to a positional array
```
kubectl patch sa default --type='json' -p='[{"op": "add", "path": "/secrets/1", "value": {"name": "whatever" } }]'
```{{execute}}



# Editing resources

Edit the service named docker-registry

```
kubectl edit svc/docker-registry
```{{execute}}

Use an alternative editor
```
KUBE_EDITOR="nano" kubectl edit svc/docker-registry
```{{execute}}


# Scaling resources

Scale a replicaset named 'foo' to 3
```
kubectl scale --replicas=3 rs/foo
```{{execute}}


Scale a resource specified in "foo.yaml" to 3
```
kubectl scale --replicas=3 -f foo.yaml
```{{execute}}


If the deployment named mysql's current size is 2, scale mysql to 3
```
kubectl scale --current-replicas=2 --replicas=3 deployment/mysql  
```{{execute}}


Scale multiple replication controllers
```
kubectl scale --replicas=5 rc/foo rc/bar rc/baz
```{{execute}}



# Deleting resources 


Delete a pod using the type and name specified in pod.json
```
kubectl delete -f ./pod.json
```{{execute}}


Delete pods and services with same names "baz" and "foo"
```
kubectl delete pod,service baz foo
```{{execute}}


Delete pods and services with label name=myLabel
```
kubectl delete pods,services -l name=myLabel
```{{execute}}


Delete all pods and services in namespace my-ns,
```
kubectl -n my-ns delete pod,svc --all
```{{execute}}


Delete all pods matching the awk pattern1 or pattern2
```
kubectl get pods  -n mynamespace --no-headers=true | awk '/pattern1|pattern2/{print $1}' | xargs  kubectl delete -n mynamespace pod
```{{execute}}
