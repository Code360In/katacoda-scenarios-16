# Build AWX


https://github.com/ansible/awx

https://github.com/ansible/awx-operator


Verify:

```
kubectl get nodes -o wide
```{{execute}}


This Kubernetes Operator is meant to be deployed in your Kubernetes cluster(s) and can manage one or more AWX instances in any namespace.

First you need to deploy AWX Operator into your cluster:

```
kubectl apply -f https://raw.githubusercontent.com/ansible/awx-operator/devel/deploy/awx-operator.yaml
```{{execute}}



Deploy new AWX instance:

```
cat <<EOF | kubectl apply -f -
apiVersion: awx.ansible.com/v1beta1
kind: AWX
metadata:
  name: awx
EOF
```{{execute}}


Results:

<pre class="file">
awx.awx.ansible.com/awx created
</pre>
