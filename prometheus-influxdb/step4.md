# Deploy influxDB 


```
docker run -d -p 8086:8086 --name influxdb \
      -v influxdb:/root/influxdb/data \
      -v influxdb2:/root/influxdb2/data \
      -v influxdb2-config:/etc/influxdb2 \
      -v $PWD/influxdb2.0.conf:/root/influxdb/influxdb.conf \
      -e DOCKER_INFLUXDB_INIT_MODE=setup \
      -e DOCKER_INFLUXDB_INIT_USERNAME=admin \
      -e DOCKER_INFLUXDB_INIT_PASSWORD=system#1 \
      -e DOCKER_INFLUXDB_INIT_ORG=prometheus \
      -e DOCKER_INFLUXDB_INIT_BUCKET=prometheus \
      -e DOCKER_INFLUXDB_INIT_RETENTION=1w \
      -e DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=token \
      influxdb:2.0.4
```{{execute}}     

```
docker run -d -p 8086:8086 --name influxdb \
      -v influxdb:/var/lib/influxdb \
      -v $PWD/influxdb.conf:/etc/influxdb/influxdb.conf \
      -e DOCKER_INFLUXDB_INIT_ORG=prometheus \
      -e DOCKER_INFLUXDB_INIT_BUCKET=prometheus \
      influxdb:1.8
```{{execute}}  

```
docker ps
```{{execute}}


influxDB 1.8

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

influx org create -n prometheus


influxDB 2.0

```
docker exec -it influxdb bash
```{{execute}}


find id
```
influx auth list
```{{execute}}


inactive
```
influx auth inactive --id 076878bcf6ecf000
```{{execute}}





access:

https://[[HOST_SUBDOMAIN]]-8086-[[KATACODA_HOST]].environments.katacoda.com/

http://localhost:8086/v2/api/prometheus/prometheus/prometheus

influxQL:

```
curl -G 'http://localhost:8086/query?pretty=true' --data-urlencode "db=prometheus" --data-urlencode "q=SELECT \"value\" FROM \"up\""
```{{execute}}



Flux:

```
curl -XPOST localhost:8086/api/v2/query -sS \
  -H 'Accept:application/csv' \
  -H 'Content-type:application/vnd.flux' \
  -d 'from(bucket:"prometheus")
        |> range(start:-5m)
        |> filter(fn:(r) => r._measurement == "up")'  
```{{execute}}

```
```{{execute}}
