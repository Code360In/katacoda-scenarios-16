# Verify Logstash and Elasticsearh Performance

https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html


### Deploy Elasticsearch
```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.11.2
```{{execute}}



```
curl -X GET "localhost:9200/_cat/nodes?v=true&pretty"
```{{execute}}


### Deploy Logstash

```
docker run --rm -it -v ~/settings/logstash.yml:/usr/share/logstash/config/logstash.yml docker.elastic.co/logstash/logstash:7.11.1
```{{execute}}


### Deploy Grafana

