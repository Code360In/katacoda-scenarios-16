Now, let's run grafana with this command:

`docker run --name=grafana -d -p 3000:3000 grafana/grafana`{{execute T1}}

And access to the dashboard using this url:

https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/

And use this default credentials:
`admin`{{copy}} \ `admin`{{copy}}

The initial step will ask you to change the password, you can do it if you want or you can skip this step.

The first step to create a dashboard is to have a datasource. Let's define a datasource with phometheus data configured before:

The URL for phometheus should be:

`http://172.18.0.5:9090`{{copy}}

For all the other fields you can use the default values

![](/envoyproxy/scenarios/implementing-metrics-tracing/assets/prometheus-data-source.png)

Test and save your datasource.

## Build a dashboard

In order to build a dashboard you can build one from scratch or you can use an existing one, for example:
https://grafana.com/dashboards/6693

Let's use this existing dashboard. Copy the ID, and use the option `Import`.

`6693`{{copy}}

Select **Prometheus** as the data source and Import.

![](/envoyproxy/scenarios/implementing-metrics-tracing/assets/import.png)

[View Dashboard for the targetCluster](https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/d/000000003/envoy-proxy?refresh=5s&orgId=1&var-cluster=targetCluster&var-hosts=All)


