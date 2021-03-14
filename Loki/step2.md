
# Deploy Promtail


```
wget https://raw.githubusercontent.com/grafana/loki/v2.2.0/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml
```{{execute}}



```
docker run -v $(pwd):/mnt/config -v /var/log:/var/log grafana/promtail:2.2.0 -config.file=/mnt/config/promtail-config.yaml
```{{execute}}
