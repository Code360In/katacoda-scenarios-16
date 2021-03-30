#### Launch Cluster

`launch.sh`{{execute}}

This will create a two node Kubernetes cluster using WeaveNet for networking.

#### Health Check

`
kubectl cluster-info
`{{execute}}


# Kubernetes logging




```
cat << 'EOF' > example.yaml

apiVersion: v1
kind: Pod
metadata:
  name: example
spec:
  containers:
  - name: example
    image: busybox
    args: [/bin/sh, -c, 'while true; do echo $(date); sleep 1; done']
EOF

```{{execute}}

```
kubectl apply -f example.yaml
```{{execute}}



```
kubectl get pods -o wide
```{{execute}}



```
kubectl logs example
```{{execute}}



Node1:
```
cd /var/log/containers
```{{execute HOST2}}


```
more /var/log/pods/default_example_*/example/0.log
```{{execute HOST2}}




# k8s events

```
kubectl get events -n default
```{{execute}}


https://kubernetes.io/docs/reference/kubectl/cheatsheet/


<pre class="file">

kubectl logs -f # stream logs
kubectl logs --since=1h # return logs newer than a relative duration
kubectl logs --since-time=2020-08-13T10:46:00.000000000Z # return logs after a specific date (RFC3339)
kubectl logs --previous # print the logs for the previous instance of the container
kubectl logs -c # print the logs of this container
kubectl logs -l #  print logs from all containers in pods defined by label
kubectl get events --sort-by='.metadata.creationTimestamp' # print all events in chronological order
kubectl describe pod  # print pod details like status or recent events

</pre>

