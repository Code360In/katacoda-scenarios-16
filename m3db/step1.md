# Deploy M3DB


https://m3db.io/docs/quickstart/docker/




```
docker run -d  -p 7201:7201 -p 7203:7203 --name m3db -v $(pwd)/m3db_data:/var/lib/m3db quay.io/m3db/m3dbnode:v1.0.0

```{{execute}}

Logs:
```
docker logs m3db
```{{execute}}

# Verify Metrics

```
curl http://localhost:7203/metrics
```{{execute}}


or

https://[[HOST_SUBDOMAIN]]-7203-[[KATACODA_HOST]].environments.katacoda.com/metrics




## httpie

https://httpie.io/docs

```
apt install httpie
```{{execute}}


```
http POST http://localhost:2379/metrics
```{{execute}}


## Create a Placement and Namespace



```
#!/bin/bash
curl -X POST http://localhost:7201/api/v1/database/create -d '{
  "type": "local",
  "namespaceName": "default",
  "retentionTime": "12h"
}' | jq .
```{{execute}}


<pre class="file">

</pre>


## You can check on the status by calling the http://localhost:7201/api/v1/services/m3db/placement endpoint:


```
curl http://localhost:7201/api/v1/services/m3db/placement | jq .
```{{execute}}


## Ready a Namespace 
Once a namespace has finished bootstrapping, you must mark it as ready before receiving traffic by using the http://localhost:7201/api/v1/services/m3db/namespace/ready.


```
#!/bin/bash
curl -X POST http://localhost:7201/api/v1/services/m3db/namespace/ready -d '{
  "name": "default"
}' | jq .
```{{execute}}


## View Details of a Namespace 
You can also view the attributes of all namespaces by calling the http://localhost:7201/api/v1/services/m3db/namespace endpoint

```
curl http://localhost:7201/api/v1/services/m3db/namespace | jq .
```{{execute}}



# Writing and Querying Metrics 

<pre class="file">

Writing Metrics 
M3 supports ingesting statsd and Prometheus formatted metrics.

This quickstart focuses on Prometheus metrics which consist of a value, a timestamp, and tags to bring context and meaning to the metric.

You can write metrics using one of two endpoints:

http://localhost:7201/api/v1/prom/remote/write - Write a Prometheus remote write query to M3DB with a binary snappy compressed Prometheus WriteRequest protobuf message.
http://localhost:7201/api/v1/json/write - Write a JSON payload of metrics data. This endpoint is quick for testing purposes but is not as performant for production usage.



</pre>

```
global:
  scrape_interval:     15s
  evaluation_interval: 15s
  external_labels:
    monitor: 'codelab-monitor'

remote_read:
  - url: "http://localhost:7201/api/v1/prom/remote/read"
    # To test reading even when local Prometheus has the data
    read_recent: true
remote_write:
  - url: "http://localhost:7201/api/v1/prom/remote/write"

scrape_configs:
  - job_name: 'prometheus'
    scrape_interval: 5s
    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'node'
    scrape_interval: 5s
    static_configs:
      - targets: ['localhost:8080', 'localhost:8081']
        labels:
          group: 'production'
      - targets: ['localhost:8082']
        labels:
          group: 'canary'

  - job_name: 'm3' 
    static_configs: 
      - targets: ['localhost:7203']
```{{copy}}



## Querying metrics

```
curl -X "POST" -G "http://localhost:7201/api/v1/query_range" \
  -d "query=third_avenue" \
  -d "start=$(date "+%s" -d "45 seconds ago")" \
  -d "end=$( date +%s )" \
  -d "step=5s" | jq .  
```{{execute}}
