
# Deploy Vector


# create network

```
docker network create np 
```{{execute}}



## expose docker tcp  1234
```
docker run -d --name=dockersock -v /var/run/docker.sock:/var/run/docker.sock -p 127.0.0.1:1234:1234 bobrik/socat TCP-LISTEN:1234,fork UNIX-CONNECT:/var/run/docker.sock
```{{execute}}



```
 docker network connect np dockersock
 docker restart dockersock
```{{execute}}

verify

```
curl http://localhost:1234/images/json
```{{execute}}



# Deploy Vector

```
docker run \
  --name=vector \
  -d \
  -v /root/vector.toml:/etc/vector/vector.toml:ro \
  -p 8383:8383 \
  timberio/vector:0.17.0-debian
```{{execute}}


```
docker network connect np vector
docker restart vector
```{{execute}}


```
docker logs -f $(docker ps -aqf "name=vector")
```{{execute}}


