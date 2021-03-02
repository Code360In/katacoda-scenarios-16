# Grafana Tempo

Let's define a datasource with [Tempo](https://grafana.com/oss/tempo/) data configured before:

The URL for Tempo should be:

`https://[[HOST_SUBDOMAIN]]-16686-[[KATACODA_HOST]].environments.katacoda.com`{{copy}}

For all the other fields you can use the default values


Test and save your datasource.


## Build a dashboard for Jaeger

In order to build a dashboard you can build one from scratch or you can use an existing one, for example:
https://grafana.com/grafana/dashboards/10001

Let's use this existing dashboard. Copy the ID, and use the option `Import`.

`10001`{{copy}}

Select **Prometheus** as the data source and Import.


