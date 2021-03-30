
# Viewing, finding resources 


Get commands with basic output

List all services in the namespace
```
kubectl get services
```{{execute}}

List all pods in all namespaces
```
kubectl get pods --all-namespaces
```{{execute}}


List all pods in the current namespace, with more details
```
kubectl get pods -o wide
```{{execute}}

List a particular deployment
```
kubectl get deployment my-dep
```{{execute}}


List all pods in the namespace
```
kubectl get pods
```{{execute}}

Get a pod's YAML
```
kubectl get pod my-pod -o yaml
```{{execute}}

Describe commands with verbose output
```
kubectl describe nodes my-node
```{{execute}}

```
kubectl describe pods my-pod
```{{execute}}

List Services Sorted by Name
```
kubectl get services --sort-by=.metadata.name
```{{execute}}

List pods Sorted by Restart Count
```
kubectl get pods --sort-by='.status.containerStatuses[0].restartCount'
```{{execute}}

List PersistentVolumes sorted by capacity
```
kubectl get pv --sort-by=.spec.capacity.storage
```{{execute}}

Get the version label of all pods with label app=cassandra
```
kubectl get pods --selector=app=cassandra -o \
  jsonpath='{.items[*].metadata.labels.version}'
```{{execute}}

Retrieve the value of a key with dots, e.g. 'ca.crt'
```
kubectl get configmap myconfig \
  -o jsonpath='{.data.ca\.crt}'
```{{execute}}

Get all worker nodes (use a selector to exclude results that have a label
named 'node-role.kubernetes.io/master')
```
kubectl get node --selector='!node-role.kubernetes.io/master'
```{{execute}}

Get all running pods in the namespace
```
kubectl get pods --field-selector=status.phase=Running
```{{execute}}

Get ExternalIPs of all nodes
```
kubectl get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="ExternalIP")].address}'
```{{execute}}

List Names of Pods that belong to Particular RC
"jq" command useful for transformations that are too complex for jsonpath, it can be found at https://stedolan.github.io/jq/
```
sel=${$(kubectl get rc my-rc --output=json | jq -j '.spec.selector | to_entries | .[] | "\(.key)=\(.value),"')%?}
echo $(kubectl get pods --selector=$sel --output=jsonpath={.items..metadata.name})
```{{execute}}

Show labels for all pods (or any other Kubernetes object that supports labelling)
```
kubectl get pods --show-labels
```{{execute}}

Check which nodes are ready
```
JSONPATH='{range .items[*]}{@.metadata.name}:{range @.status.conditions[*]}{@.type}={@.status};{end}{end}' \
 && kubectl get nodes -o jsonpath="$JSONPATH" | grep "Ready=True"
```{{execute}}

Output decoded secrets without external tools
```
kubectl get secret my-secret -o go-template='{{range $k,$v := .data}}{{"### "}}{{$k}}{{"\n"}}{{$v|base64decode}}{{"\n\n"}}{{end}}'
```{{execute}}

List all Secrets currently in use by a pod
```
kubectl get pods -o json | jq '.items[].spec.containers[].env[]?.valueFrom.secretKeyRef.name' | grep -v null | sort | uniq
```{{execute}}

List all containerIDs of initContainer of all pods
Helpful when cleaning up stopped containers, while avoiding removal of initContainers.
```
kubectl get pods --all-namespaces -o jsonpath='{range .items[*].status.initContainerStatuses[*]}{.containerID}{"\n"}{end}' | cut -d/ -f3
```{{execute}}

List Events sorted by timestamp
```
kubectl get events --sort-by=.metadata.creationTimestamp
```{{execute}}

Compares the current state of the cluster against the state that the cluster would be in if the manifest was applied.
```
kubectl diff -f ./my-manifest.yaml
```{{execute}}

Produce a period-delimited tree of all keys returned for nodes
Helpful when locating a key within a complex nested JSON structure
```
kubectl get nodes -o json | jq -c 'path(..)|[.[]|tostring]|join(".")'
```{{execute}}

Produce a period-delimited tree of all keys returned for pods, etc
```
kubectl get pods -o json | jq -c 'path(..)|[.[]|tostring]|join(".")'
```{{execute}}
