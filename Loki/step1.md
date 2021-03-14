
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
curl -v -H "Content-Type: application/json" -XPOST -s "http://localhost:3100/loki/api/v1/push" --data-raw \
  '{"streams": [{ "stream": { "foo": "bar2" }, "values": [ [ "1570818238000000000", "fizzbuzz" ] ] }]}'
```{{execute}}
