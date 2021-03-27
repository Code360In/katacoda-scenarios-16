# Promscale

https://hub.docker.com/r/timescale/promscale/tags?page=1&ordering=last_updated
https://github.com/timescale/promscale

https://github.com/timescale/promscale/blob/master/docs/docker.md


https://prometheus.io/docs/prometheus/latest/installation/

https://docs.timescale.com/latest/tutorials/getting-started-with-promscale


```
docker network create --driver bridge promscale-timescaledb
```{{execute}}

```
docker run --name timescaledb -e POSTGRES_PASSWORD=secret -d -p 5432:5432 --network host timescaledev/promscale-extension:latest-pg12 postgres -csynchronous_commit=off
```{{execute}}

```
docker run --name promscale -d -p 9201:9201 --network host timescale/promscale:0.2 -db-password=secret -db-port=5432 -db-name=postgres -db-host=timescaledb -db-ssl-mode=allow
```{{execute}}



