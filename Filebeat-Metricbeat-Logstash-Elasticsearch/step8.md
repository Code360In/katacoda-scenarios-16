# Deploy Metricbeat

https://www.elastic.co/guide/en/beats/metricbeat/current/running-on-docker.html


Running Metricbeat with the setup command will create the index pattern and load visualizations , dashboards, and machine learning jobs. Run this command:

```
docker run -d --net=host --name=metricbeat \
--user=root \
docker.elastic.co/beats/metricbeat:7.11.2 \
setup -E setup.kibana.host=localhost:5601 \
-E output.elasticsearch.hosts=["localhost:9200"]
```{{execute}}


```
docker logs metricbeat
```{{execute}}


```
docker ps -a
```{{execute}}
