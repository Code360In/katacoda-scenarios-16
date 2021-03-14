
# Deploy Loki

https://grafana.com/docs/loki/latest/installation/docker/

https://grafana.com/docs/loki/latest/api/


```
wget https://raw.githubusercontent.com/grafana/loki/v2.2.0/cmd/loki/loki-local-config.yaml -O loki-config.yaml
```{{execute}}



```
docker run -d -v /root/loki-config.yaml:/usr/share/loki-config.yaml --net=host --name=loki -p 3100:3100 grafana/loki:2.2.0 -config.file=/usr/share/loki-config.yaml
```{{execute}}



Verify:


```
docker logs loki
```{{execute}}


```
curl localhost:3100/metrics
```{{execute}}



sending data to loki:

```
curl -H "Content-Type: application/json" -XPOST -s "http://localhost:3100/api/prom/push" --data-raw \
  '{"streams": [{ "labels": "{foo=\"bar\"}", "entries": [{ "ts": "2018-12-18T08:28:06.801064-04:00", "line": "fizzbuzz" }] }]}'
```{{execute}}
