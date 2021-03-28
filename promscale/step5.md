# PostgreSQL Server Exporter

https://github.com/prometheus-community/postgres_exporter


Prometheus exporter for PostgreSQL server metrics.

```
docker run -d \
  --net=pnet -p 9187:9187 \
  --name=postgresexporter \
  -e DATA_SOURCE_NAME="postgresql://postgres:secret@timescaledb:5432/postgres?sslmode=disable" \
  quay.io/prometheuscommunity/postgres-exporter
```{{execute}}


Access Metrics:

https://[[HOST_SUBDOMAIN]]-9187-[[KATACODA_HOST]].environments.katacoda.com/metrics


Import Dashboards:

copy and import to grafana

`/root/grafana/grafana-dashboard-psql.json`{{open}}

original:
https://github.com/prometheus-community/postgres_exporter/tree/master/postgres_mixin



