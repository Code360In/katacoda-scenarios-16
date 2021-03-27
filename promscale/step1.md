# Promscale

https://hub.docker.com/r/timescale/promscale/tags?page=1&ordering=last_updated
https://github.com/timescale/promscale

https://github.com/timescale/promscale/blob/master/docs/docker.md


https://prometheus.io/docs/prometheus/latest/installation/

https://docs.timescale.com/latest/tutorials/getting-started-with-promscale


Create network
```
docker network create --driver bridge pnet
```{{execute}}


Install and run TimescaleDB with Promscale extension:
```
docker run --name timescaledb -e POSTGRES_PASSWORD=secret -d -p 5432:5432 --network=pnet timescaledev/promscale-extension:latest-pg12 postgres -csynchronous_commit=off
```{{execute}}


run Promscale:
```
docker run --name promscale -d -p 9201:9201 --network=pnet timescale/promscale:0.2 -db-password=secret -db-port=5432 -db-name=postgres -db-host=timescaledb -db-ssl-mode=allow
```{{execute}}


Verify:
```
docker ps -a
```{{execute}}

```
docker logs timescaledb
```{{execute}}

```
docker logs promscale
```{{execute}}

```
curl localhost:9201/metrics
```{{execute}}
