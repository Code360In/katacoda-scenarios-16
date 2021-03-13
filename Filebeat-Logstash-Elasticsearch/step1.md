
### Deploy Elasticsearch

https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html


```
docker run -d --net=host --name=elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.11.2
```{{execute}}


Verify:
```
curl -X GET "localhost:9200/_cat/nodes?v=true&pretty"
```{{execute}}


Results:

<pre class="file">

ip         heap.percent ram.percent cpu load_1m load_5m load_15m node.role  master name
172.18.0.2           62          95  11    0.69    0.40     0.17 cdhilmrstw *      37afd33e70f0

</pre>


```
curl localhost:9200
```{{execute}}

Results:

<pre class="file">

{
  "name" : "host01",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "miGT76pkRB6izijC4CYHUg",
  "version" : {
    "number" : "7.11.2",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "3e5a16cfec50876d20ea77b075070932c6464c7d",
    "build_date" : "2021-03-06T05:54:38.141101Z",
    "build_snapshot" : false,
    "lucene_version" : "8.7.0",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}


</pre>


```
docker logs elasticsearch
```{{execute}}

Results:

<pre class="file">

{"type": "server", "timestamp": "2021-03-13T18:27:30,522Z", "level": "INFO", "component": "o.e.x.s.s.SecurityStatusChangeListener", "cluster.name": "docker-cluster", "node.name": "host01", "message": "Active license is now [BASIC]; Security is disabled", "cluster.uuid": "miGT76pkRB6izijC4CYHUg", "node.id": "pgMNigQWSCy1JK1lNR7caQ"  }

</pre>
