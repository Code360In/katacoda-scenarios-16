# Docker metrics


```
printf "{
  "metrics-addr" : "127.0.0.1:9323",
  "experimental" : true
}" >> /etc/docker/daemon.json
```{{execute T2}}


If you're interested in seeing the raw metrics, they can be viewed with 

```
curl localhost:9100/metrics
```{{execute T2}}