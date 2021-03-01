
Install Kubernetes dashboard [dashboard](https://helm.sh/docs/intro/install/):

[Helm package](https://artifacthub.io/packages/helm/k8s-dashboard/kubernetes-dashboard)

Create a namespace:
``` 
kubectl create namespace dashboard
```{{execute}}


Add kubernetes-dashboard repository
``` 
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
```{{execute}}

Deploy a Helm Release named "dashboard" using the kubernetes-dashboard chart
``` 
helm install dashboard kubernetes-dashboard/kubernetes-dashboard -n dashboard
```{{execute}}



Create a new account that can be used to authenticate to the dashboard:
``` 
kubectl create serviceaccount me -n kube-system
```{{execute}}

 
Make that new account have the right permissions:
``` 
kubectl create clusterrolebinding me -n kube-system --clusterrole=cluster-admin --serviceaccount=kube-system:me
```{{execute}}
 

Show the token needed to log in to the dashboard:
``` 
kubectl get secret $(kubectl get serviceaccount me -n kube-system -o jsonpath="{.secrets[0].name}") -o jsonpath="{.data.token}" -n kube-system | base64 --decode; echo
```{{execute}}


Show all values from helm package:
``` 
helm get -a
```{{execute}}


Acess the kubernetes dashboard, go to the dashboard:
https://[[HOST_SUBDOMAIN]]-8443-[[KATACODA_HOST]].environments.katacoda.com