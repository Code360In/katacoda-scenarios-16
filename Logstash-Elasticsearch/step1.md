# Verify Logstash and Elasticsearh Performance


https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html



### Deploy Elasticsearch
```
docker run -d  -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.11.2
```{{execute}}


Verify:
```
curl -X GET "localhost:9200/_cat/nodes?v=true&pretty"
```{{execute}}


Results:
`
ip         heap.percent ram.percent cpu load_1m load_5m load_15m node.role  master name
172.18.0.2           62          95  11    0.69    0.40     0.17 cdhilmrstw *      37afd33e70f0
`



### Deploy Logstash

```
docker run -d  --rm -it -v /root/logstash.yml:/usr/share/logstash/config/logstash.yml docker.elastic.co/logstash/logstash:7.11.1
```{{execute}}

```
docker run -d  --rm -it docker.elastic.co/logstash/logstash:7.11.1
```{{execute}}
