
# Deploy Loki

https://grafana.com/docs/loki/latest/installation/docker/


```
wget https://raw.githubusercontent.com/grafana/loki/v2.2.0/cmd/loki/loki-local-config.yaml -O loki-config.yaml
```{{execute}}



```
docker run -v $(pwd):/mnt/config --name=loki -p 3100:3100 grafana/loki:2.2.0 -config.file=/mnt/config/loki-config.yaml
```{{execute}}



Verify:


```
docker logs loki
```{{execute}}
