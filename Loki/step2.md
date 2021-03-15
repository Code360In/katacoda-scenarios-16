
# Deploy Promtail


```
wget https://raw.githubusercontent.com/grafana/loki/v2.2.0/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml
```{{execute}}


Change loki:3100 to localhost:3100 

```
sed -i 's/loki:3100/localhost:3100/' /root/promtail-config.yaml
```{{execute}}

Open file:

```
/root/promtail-config.yaml
```{{open}}


Add push config for promtail:

```
- job_name: push1
  loki_push_api:
    server:
      http_listen_port: 3500
      grpc_listen_port: 3600
    labels:
      pushserver: push1
```{{copy}}



Deploy promtail:

```
docker run -d --user=root -v /root/promtail-config.yaml:/usr/share/promtail-config.yaml --net=host -p 9080:9080 -p 3500:3500 --name=promtail -v /var/log:/var/log grafana/promtail:2.2.0 -config.file=/usr/share/promtail-config.yaml
```{{execute}}



Verify:

```
curl localhost:9080/metrics
```{{execute}}


```
docker logs promtail
```{{execute}}




