
# Deploy Logstash

Copy logstash.conf to /root/

```
cat << 'EOF' > /root/logstash.conf
input {
  beats {
    port => 5044
    ssl => false
    ssl_verify_mode => "none"
  }
}

output {
  elasticsearch {
    hosts => ["http://localhost:9200"]
    index => "%{[@metadata][beat]}-%{[@metadata][version]}" 
  }
}
EOF
```{{execute}}


logstash.yml
```
cat << 'EOF' > /root/logstash.yml
http.host: "0.0.0.0"
xpack.monitoring.enabled: true
xpack.monitoring.elasticsearch.hosts : "http://localhost:9200"
EOF
```{{execute}}


Verify:

```
pwd;ls
```{{execute}}



Deploy logstash
```
docker run -d -it --net=host --name=logstash -p 5044:5044 -v /root/logstash.conf:/usr/share/logstash/pipeline/logstash.conf -v /root/logstash.yml:/usr/share/logstash/config/logstash.yml docker.elastic.co/logstash/logstash:7.11.1
```{{execute}}


Verify:
```
docker ps -a
```{{execute}}


Config file:
```
docker exec logstash more /usr/share/logstash/pipeline/logstash.conf
```{{execute}}


```
docker logs logstash
```{{execute}}

Results:
<pre class="file">

[2021-03-13T02:22:50,015][INFO ][logstash.outputs.elasticsearch][main] Installing elasticsearch template to _template/logstash
[2021-03-13T02:22:50,964][INFO ][logstash.javapipeline    ][main] Pipeline Java execution initialization time {"seconds"=>1.01}
[2021-03-13T02:22:50,983][INFO ][logstash.inputs.beats    ][main] Starting input listener {:address=>"0.0.0.0:5044"}
[2021-03-13T02:22:51,001][INFO ][logstash.javapipeline    ][main] Pipeline started {"pipeline.id"=>"main"}
[2021-03-13T02:22:51,305][INFO ][logstash.agent           ] Successfully started Logstash API endpoint {:port=>9600}

</pre>


Verify:
```
curl localhost:9600
```{{execute}}

```
curl -XGET 'localhost:9600/?pretty'
```{{execute}}

```
curl -XGET 'localhost:9600/_node/plugins?pretty'
```{{execute}}

```
curl -XGET 'localhost:9600/_node/hot_threads?pretty'
```{{execute}}

```
curl -XGET 'localhost:9600/_node/pipelines?pretty'
curl -XGET 'localhost:9600/_node/os?pretty'
curl -XGET 'localhost:9600/_node/jvm?pretty'
```{{execute}}



Test from logstash to elasticsearch:
```
docker exec -it logstash curl localhost:9200
```{{execute}}


