
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
curl -H "Content-Type: application/json" -XPOST -s "http://localhost:3100/loki/api/v1/push" --data-raw "{\"streams\": [{\"stream\": {\"job\": \"test\"}, \"values\": [[\"$(date +%s)000000000\", \"fizzbuzz\"]]}]}"

```{{execute}}


Query:

```
curl -G -s  "http://localhost:3100/api/prom/label" | jq

curl "http://localhost:3100/loki/api/v1/query_range" --data-urlencode 'query={job="test"}' --data-urlencode 'step=300' | jq .data.result

```{{execute}}




```
date --rfc-3339=ns | sed 's/ /T/'
```{{execute}}

expected:
`
2021-03-15T01:40:42.017393021+00:00
`



And access to the loki using this url:

https://[[HOST_SUBDOMAIN]]-3100-[[KATACODA_HOST]].environments.katacoda.com/


Debug:


https://[[HOST_SUBDOMAIN]]-3100-[[KATACODA_HOST]].environments.katacoda.com/debug/pprof/

