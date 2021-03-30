# Updating resources



Rolling update "www" containers of "frontend" deployment, updating the image

```
kubectl set image deployment/frontend www=image:v2               
```{{execute}}

Check the history of deployments including the revision 
```
kubectl rollout history deployment/frontend
```{{execute}}

Rollback to the previous deployment
```
kubectl rollout undo deployment/frontend
```{{execute}}


Rollback to a specific revision
```
kubectl rollout undo deployment/frontend --to-revision=2
```{{execute}}


Watch rolling update status of "frontend" deployment until completion
```
kubectl rollout status -w deployment/frontend
```{{execute}}

Rolling restart of the "frontend" deployment
```
kubectl rollout restart deployment/frontend
```{{execute}}


Replace a pod based on the JSON passed into std
```
cat pod.json | kubectl replace -f - 
```{{execute}}


 Force replace, delete and then re-create the resource. Will cause a service outage.
```
kubectl replace --force -f ./pod.json
```{{execute}}

Create a service for a replicated nginx, which serves on port 80 and connects to the containers on port 8000
```
kubectl expose rc nginx --port=80 --target-port=8000
```{{execute}}

Update a single-container pod's image version (tag) to v4
```
kubectl get pod mypod -o yaml | sed 's/\(image: myimage\):.*$/\1:v4/' | kubectl replace -f -
```{{execute}}


Add a Label
```
kubectl label pods my-pod new-label=awesome
```{{execute}}

Add an annotation
```
kubectl annotate pods my-pod icon-url=http://goo.gl/XXBTWq
```{{execute}}

Auto scale a deployment "foo"
```
kubectl autoscale deployment foo --min=2 --max=10
```{{execute}}
