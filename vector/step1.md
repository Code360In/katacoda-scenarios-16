
# Deploy Vector


# create network

```
docker network create np 
```{{execute}}



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


