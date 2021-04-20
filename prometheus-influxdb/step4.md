# Deploy influxDB + chronograf


```
docker run -d -p 8086:8086 --name influxdb \
      -v influxdb:/root/influxdb/data \
      -v influxdb2:/root/influxdb2/data \
      -v influxdb2-config:/etc/influxdb2 \
      -v $PWD/influxdb.conf:/root/influxdb/influxdb.conf \
      -e DOCKER_INFLUXDB_INIT_MODE=setup \
      -e DOCKER_INFLUXDB_INIT_USERNAME=admin \
      -e DOCKER_INFLUXDB_INIT_PASSWORD=system#1 \
      -e DOCKER_INFLUXDB_INIT_ORG=prometheus \
      -e DOCKER_INFLUXDB_INIT_BUCKET=prometheus \
      -e DOCKER_INFLUXDB_INIT_RETENTION=1w \
      -e DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=token \
      influxdb:1.8
```{{execute}}     

```
docker run -d -p 8086:8086 --name influxdb \
      -v influxdb:/var/lib/influxdb \
      -v $PWD/influxdb.conf:/etc/influxdb/influxdb.conf \
      influxdb:1.8
```{{execute}}  

```
docker ps
```{{execute}}


install client
```
apt install influxdb-client
```{{execute}}

```
influx
```{{execute}}



```
CREATE DATABASE "prometheus"
quit
```{{execute}}


access:

https://[[HOST_SUBDOMAIN]]-8086-[[KATACODA_HOST]].environments.katacoda.com/

