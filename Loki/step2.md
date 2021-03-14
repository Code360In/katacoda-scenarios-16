
# Deploy Promtail


```
wget https://raw.githubusercontent.com/grafana/loki/v2.2.0/cmd/promtail/promtail-docker-config.yaml -O promtail-config.yaml
```{{execute}}


Change loki:3100 to be localhost:3100 

```
sed -i 's/loki:3100/localhost:3100/' /root/promtail-config.yaml
```{{execute}}


Deploy promtail:

```
docker run -d --user=root -v /root/promtail-config.yaml:/usr/share/promtail-config.yaml --net=host -p 9080:9080 --name=promtail -v /var/log:/var/log grafana/promtail:2.2.0 -config.file=/usr/share/promtail-config.yaml
```{{execute}}



Verify:

```
curl localhost:9080/metrics
```{{execute}}


```
docker logs promtail
```{{execute}}



