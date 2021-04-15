# Deploy Telegraf



```
docker run -d --name=telegraf \
      --net=host \
      -p 8125:8125/udp \
      -v $PWD/telegraf.conf:/etc/telegraf/telegraf.conf:ro \
      telegraf:alpine
```{{execute}}


federate
```
http://localhost:9090/federate?match[]={__name__=~"[a-z].*"}

http://localhost:9090/federate?match[]={__name__=~"node.*"}&match[]={__name__=~"mysql.*"}

http://localhost:9090/federate?match[]={__name__=~".*"}

```

