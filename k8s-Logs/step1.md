#### Launch Cluster

`launch.sh`{{execute}}

This will create a two node Kubernetes cluster using WeaveNet for networking.

#### Health Check

`
kubectl cluster-info
`{{execute}}


# Kubernetes logging




```
cat << 'EOF' > example.go

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

