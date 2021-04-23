# (M3DB)[https://m3db.io/docs/overview/components/]



<pre class="file">

Components
M3 Coordinator 
M3 Coordinator is a service that coordinates reads and writes between upstream systems, such as Prometheus, and M3DB. It is a bridge that users can deploy to access the benefits of M3DB such as long term storage and multi-DC setup with other monitoring systems, such as Prometheus. See this presentation for more on long term storage in Prometheus.

M3DB 
M3DB is a distributed time series database that provides scalable storage and a reverse index of time series. It is optimized as a cost effective and reliable realtime and long term retention metrics store and index. For more details, see the M3DB documentation.

M3 Query 
M3 Query is a distributed query engine for querying realtime and historical data stored in M3DB nodes, supporting several query languages. It is designed to support both low latency realtime queries and queries that can take longer to execute, aggregating over larger datasets for analytical use cases.

For example, if you are using the Prometheus remote write endpoint with M3 Coordinator, you can use M3 Query instead of the Prometheus remote read endpoint. By doing so, you get all the benefits of M3 Query’s engine such as block processing. As M3 Query provides a Prometheus compatible API, you can use 3rd party graphing and alerting solutions like Grafana.

M3 Aggregator 
M3 Aggregator is a dedicated metrics aggregator that provides stateful stream-based downsampling before storing metrics in M3DB nodes. It uses dynamic rules stored in etcd .

It uses leader election and aggregation window tracking, leveraging etcd to manage this state, to reliably emit at-least-once aggregations for downsampled metrics to long term storage. This provides cost effective and reliable downsampling & roll up of metrics.

M3 Coordinator can also perform this role, but M3 Aggregator shards and replicates metrics, whereas M3 Coordinator is not and requires care to deploy and run in a highly available way.

Similar to M3DB, M3 Aggregator supports clustering and replication by default. This means that metrics are correctly routed to the instance(s) responsible for aggregating each metric and you can configure multiple M3 Aggregator replicas so that there are no single points of failure for aggregation.

</pre>



![](http://1fykyq3mdn5r21tpna3wkdyi-wpengine.netdna-ssl.com/wp-content/uploads/2018/08/image4-1.png)

# Deploy M3DB


https://m3db.io/docs/quickstart/docker/


<pre class="file">

Start Docker Container 
By default the official M3 Docker image configures a single instance as one binary containing:

An M3DB storage instance for time series storage. It includes an embedded tag-based metrics index and an etcd server for storing the cluster topology and runtime configuration.
A M3 Coordinator instance for writing and querying tagged metrics, as well as managing cluster topology and runtime configuration.
The Docker container exposes three ports:

7201 to manage the cluster topology, you make most API calls to this endpoint
7203 for Prometheus to scrape the metrics produced by M3DB and M3 Coordinator

</pre>


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
http POST http://localhost:7203/metrics
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


You can check on the status by calling the `http://localhost:7201/api/v1/services/m3db/placement` endpoint:


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

You can also view the attributes of all namespaces by calling the `http://localhost:7201/api/v1/services/m3db/namespace` endpoint

```
curl http://localhost:7201/api/v1/services/m3db/namespace | jq .
```{{execute}}



# Writing and Querying Metrics 

<pre class="file">

Writing Metrics 
M3 supports ingesting statsd and Prometheus formatted metrics.

This quickstart focuses on Prometheus metrics which consist of a value, a timestamp, and tags to bring context and meaning to the metric.

You can write metrics using one of two endpoints:

`http://localhost:7201/api/v1/prom/remote/write` - Write a Prometheus remote write query to M3DB with a binary snappy compressed Prometheus WriteRequest protobuf message.
`http://localhost:7201/api/v1/json/write` - Write a JSON payload of metrics data. This endpoint is quick for testing purposes but is not as performant for production usage.



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


```
ls
cat /root/prometheus.yml
echo 
```{{execute}}





## Start Node Exporters

copy metrics-1.prom:
```
mkdir /root/node-1
cd /root/node-1
curl -LO https://m3db.io/docs/includes/quickstart/node-1/metrics-1.prom
cd 
```{{execute}}


Install:

```
cd /root/
curl -LO https://github.com/prometheus/node_exporter/releases/download/v1.1.2/node_exporter-1.1.2.linux-amd64.tar.gz
tar -xvf node_exporter-1.1.2.linux-amd64.tar.gz
mv node_exporter-1.1.2.linux-amd64/node_exporter .
chmod +x node_exporter
rm  -r -f node_exporter-1.1.2.linux-amd64.tar.gz node_exporter-1.1.2.linux-amd64
cd
```{{execute}}

Run 

```
echo T2
```{{execute T2}}


```
cd /root
#!/bin/bash
./node_exporter --collector.textfile.directory=/root/node-1/ --web.listen-address 127.0.0.1:8081
```{{execute T2}}

Verify:

```
curl localhost:8081/metrics
```{{execute T1}}


## Querying metrics

Return results in past 45 seconds 


```
curl -X "POST" -G "http://localhost:7201/api/v1/query_range" \
  -d "query=third_avenue" \
  -d "start=$(date "+%s" -d "45 seconds ago")" \
  -d "end=$( date +%s )" \
  -d "step=5s" | jq .  
```{{execute}}


Values above a certain number 


```
curl -X "POST" -G "http://localhost:7201/api/v1/query_range" \
  -d "query=third_avenue" \
  -d "start=$(date "+%s" -d "45 seconds ago")" \
  -d "end=$( date +%s )" \
  -d "step=5s" | jq .
```{{execute}}


# Writing and Querying Metrics

```
#!/bin/bash
curl -X POST http://localhost:7201/api/v1/json/write -d '{
  "tags": 
    {
      "__name__": "third_avenue",
      "city": "new_york",
      "checkout": "1"
    },
    "timestamp": '\"$(date "+%s")\"',
    "value": 3347.26
}'
```{{execute}}