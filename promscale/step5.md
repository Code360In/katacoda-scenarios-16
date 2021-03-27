# Verification


# PostgreSQL Server Exporter

https://github.com/prometheus-community/postgres_exporter


Prometheus exporter for PostgreSQL server metrics.

```
docker run \
  --net=pnet \
  -e DATA_SOURCE_NAME="postgresql://postgres:secret@timescaledb:5432/postgres?sslmode=disable" \
  quay.io/prometheuscommunity/postgres-exporter
```{{execute}}


https://github.com/prometheus-community/postgres_exporter/tree/master/postgres_mixin

