# Verification


# PostgreSQL Server Exporter

https://github.com/prometheus-community/postgres_exporter


Prometheus exporter for PostgreSQL server metrics.

```
docker run \
  --net=pnet -p 9187:9187 \
  --name=postgresexporter \
  -e DATA_SOURCE_NAME="postgresql://postgres:secret@timescaledb:5432/postgres?sslmode=disable" \
  quay.io/prometheuscommunity/postgres-exporter
```{{execute}}


Import Dashboards:
https://github.com/prometheus-community/postgres_exporter/tree/master/postgres_mixin


More examples:


https://docs.timescale.com/latest/tutorials/tutorial-howto-monitor-django-prometheus
