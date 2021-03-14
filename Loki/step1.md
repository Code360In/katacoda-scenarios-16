
# Deploy Loki

https://grafana.com/docs/loki/latest/installation/docker/


```
wget https://raw.githubusercontent.com/grafana/loki/v2.2.0/cmd/loki/loki-local-config.yaml -O loki-config.yaml
```{{execute}}



```
docker run -d -v /root/loki-config.yaml:/usr/share/loki-config.yaml --name=loki -p 3100:3100 grafana/loki:2.2.0 -config.file=/usr/share/loki-config.yaml
```{{execute}}



Verify:


```
docker logs loki
```{{execute}}
