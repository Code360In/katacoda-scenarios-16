# Deploy Kibana

```
cat << 'EOF' > /root/kibana.yml
# Default Kibana configuration for docker target
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://localhost:9200" ]
xpack.monitoring.ui.container.elasticsearch.enabled: true
EOF
```{{execute}}


```
docker run -d  --net=host \
--name kibana  -p 5601:5601 \
-v /root/kibana.yml:/usr/share/kibana/config/kibana.yml \
kibana:7.11.1
```{{execute}}


```
docker logs kibana
```{{execute}}


And access to the kibana using this url:
https://[[HOST_SUBDOMAIN]]-5601-[[KATACODA_HOST]].environments.katacoda.com/app/home


