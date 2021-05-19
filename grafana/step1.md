# Deploy Grafana

Now, let's run grafana with this command:



```
docker run -d  -u root\
  --net=host \
  --name=grafana \
  -p 3000:3000 \
  -e "GF_INSTALL_PLUGINS=grafana-piechart-panel,grafana-worldmap-panel,marcusolsson-json-datasource,magnesium-wordcloud-panel" \
  grafana/grafana:latest-ubuntu
```{{execute}}


`docker logs grafana`{{execute}}




And access to the dashboard using this url:

https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/


And use this default credentials:
`admin`{{copy}} \ `admin`{{copy}}
