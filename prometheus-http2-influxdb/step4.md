# Deploy influxDB + chronograf

```
cat <<EOF >> docker-compose.yml
version: '3'
services:
  influxdb:
    image: influxdb:latest
    volumes:
      # Mount for influxdb data directory
      - ./influxdb/data:/var/lib/influxdb
      # Mount for influxdb configuration
      - ./influxdb/config/:/etc/influxdb/
    ports:
      # The API for InfluxDB is served on port 8086
      - "8086:8086"
      - "8082:8082"

  chronograf:
    image: chronograf:latest
    volumes:
      # Mount for chronograf database
      - ./chronograf/data/:/var/lib/chronograf/
    ports:
      # The WebUI for Chronograf is served on port 8888
      - "8888:8888"
    depends_on:
      - influxdb
EOF
```{{execute}}


```
docker-compose up -d
```{{execute}}


docker ps
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
