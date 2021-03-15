
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
curl -i -H "Content-type: application/json" -X POST --data '{ "streams": [ { "labels": "{source=\"JSON\",job=\"simpleJsonJob\", host=\"SimpleHost\"}", "entries": [{ "ts": "2021-03-15T01:45:48.042084520+00:00", "line": "TEST!" }] } ] }' http://localhost:3100/api/prom/push
```{{execute}}


Query:

```
curl -G -s  "http://localhost:3100/api/prom/label" | jq
```{{execute}}




```
date --rfc-3339=ns | sed 's/ /T/'
```{{execute}}

expected:
`
2021-03-15T01:40:42.017393021+00:00
`


