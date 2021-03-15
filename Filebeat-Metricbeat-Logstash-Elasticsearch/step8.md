# Deploy Metricbeat

https://www.elastic.co/guide/en/beats/metricbeat/current/running-on-docker.html

https://github.com/elastic/beats/blob/master/metricbeat/metricbeat.yml



Running Metricbeat with the setup command will create the index pattern and load visualizations , dashboards, and machine learning jobs. Run this command:

```
docker run -d --net=host --name=metricbeat \
--user=root \
-v /root/metricbeat.yml:/usr/share/metricbeat/metricbeat.yml \
docker.elastic.co/beats/metricbeat:7.11.2
```{{execute}}


```
docker logs metricbeat
```{{execute}}


```
docker ps -a
```{{execute}}
