# Verify Logstash and Elasticsearh Performance


https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html


### Deploy Elasticsearch
```
docker run -d  --name=elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.11.2
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

Copy logstash.conf to /root/


```
input {
  beats {
    port => 5044
  }
}

output {
  elasticsearch {
    hosts => ["http://localhost:9200"]
    index => "%{[@metadata][beat]}-%{[@metadata][version]}" 
  }
}
```{{copy}}


Deploy logstash
```
docker run -d  --name=logstash --rm -it -v /root/logstash.conf:/usr/share/logstash/config/logstash.conf docker.elastic.co/logstash/logstash:7.11.1
```{{execute}}


Verify:
```
docker ps -a
```{{execute}}
