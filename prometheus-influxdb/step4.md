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
      influxdb:2.0
```{{execute}}     

```
docker ps
```{{execute}}


install client
```
apt install influxdb-client
```{{execute}}



access:

https://[[HOST_SUBDOMAIN]]-8086-[[KATACODA_HOST]].environments.katacoda.com/


get token


https://[[HOST_SUBDOMAIN]]-8086-[[KATACODA_HOST]].environments.katacoda.com/load-data/tokens


curl --header "Authorization: Token $(bin/$(uname -s | tr '[:upper:]' '[:lower:]')/influx auth list --json | jq -r '.[0].token')" \
--data-raw "m v=2 $(date +%s)" \
"http://localhost:8086/api/v2/write?org=InfluxData&bucket=telegraf&precision=s"


curl --header "Authorization: Token $(bin/$(uname -s | tr '[:upper:]' '[:lower:]')/influx auth list --json | jq -r '.[0].token')" \
--data-raw "m v=2 $(date +%s)" \
"http://localhost:8086/api/v2/write?org=InfluxData&bucket=telegraf&precision=s"


curl -XPOST localhost:8086/api/v2/query -sS \
  -H 'Accept:application/csv' \
  -H 'Content-type:application/vnd.flux' \
  -d 'from(bucket:"telegraf")
        |> range(start:-5m)
        |> filter(fn:(r) => r._measurement == "cpu")'
