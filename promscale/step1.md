# Promscale

https://hub.docker.com/r/timescale/promscale/tags?page=1&ordering=last_updated
https://github.com/timescale/promscale

https://github.com/timescale/promscale/blob/master/docs/docker.md


https://prometheus.io/docs/prometheus/latest/installation/


docker network create --driver bridge promscale-timescaledb

docker run --name timescaledb -e POSTGRES_PASSWORD=secret -d -p 5432:5432 --network promscale-timescaledb timescaledev/promscale-extension:latest-pg12 postgres -csynchronous_commit=off

docker run --name promscale -d -p 9201:9201 --network promscale-timescaledb timescale/promscale:0.2 -db-password=secret -db-port=5432 -db-name=postgres -db-host=timescaledb -db-ssl-mode=allow

docker run -d \
  -p 3000:3000 \
  --name=grafana \
  --net=promscale-timescaledb \
  --hostname=grafana \
  -e "GF_INSTALL_PLUGINS=grafana-worldmap-panel,marcusolsson-json-datasource,grafana-polystat-panel" \
  grafana/grafana:latest-ubuntu


