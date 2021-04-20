# Deploy Grafana

Now, let's run grafana with this command:



```
docker run -d  -u root\
  --name=grafana \
  -p 3000:3000 \
  -v $PWD/certs:/certs \
  --net=host \
  -e "GF_SERVER_PROTOCOL=h2" \
  -e "GF_SERVER_CERT_FILE=/certs/prometheus.crt" \
  -e "GF_SERVER_CERT_KEY=/certs/prometheus.key" \
  -e "GF_INSTALL_PLUGINS=grafana-piechart-panel,grafana-worldmap-panel,marcusolsson-json-datasource,magnesium-wordcloud-panel" \
  grafana/grafana:latest-ubuntu
  ```{{execute}}


`docker logs grafana`{{execute}}

results:
`
"HTTP Server Listen" logger=http.server address=[::]:3000 protocol=http subUrl= socket=
`


And access to the dashboard using this url:
https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/

And use this default credentials:
`admin`{{copy}} \ `admin`{{copy}}


The initial step will ask you to change the password, you can do it if you want or you can skip this step.

The first step to create a dashboard is to have a datasource. Let's define a datasource with Prometheus data configured before:

The URL for phometheus should be:
`http://localhost:9090`{{copy}}  

or

`https://[[HOST_SUBDOMAIN]]-9090-[[KATACODA_HOST]].environments.katacoda.com`{{copy}}

For all the other fields you can use the default values



Test and save your datasource.



### Verify Grafana own Metrics:

https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/metrics



