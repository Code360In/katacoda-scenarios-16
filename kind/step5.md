
# Deploy Ha-Proxy


``` 
helm repo add haproxytech https://haproxytech.github.io/helm-charts
helm repo update
```{{execute}}


create namespace
``` 
kubectl create namespace ingress-controller
```{{execute}}


Run helm ha-proxy:
``` 
helm install haproxy -n ingress-controller haproxytech/kubernetes-ingress \
   --set controller.ingressClass=haproxy \
   --set controller.kind=DaemonSet \
   --set controller.logging.level=debug
```{{execute}}
   

Create Ingress for Prometheus:

``` 
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: mon-kube-prometheus-stack-prometheus
  namespace: mon
  annotations:
    haproxy.org/ingress.class: "haproxy"


spec:
  tls:
    - hosts:
        - host01

  rules:
  - host: host01
    http:
      paths:
      - path: /prometheus
        pathType: Prefix
        backend:
          service:
            name: mon-kube-prometheus-stack-prometheus
            port:
              number: 9090
EOF
```{{execute}}


Verify:
``` 
kubectl get ingress -A
```{{execute}}