
# Deploy Vector

```
docker run \
  -d \
  -v ~/vector.toml:/etc/vector/vector.toml:ro \
  -p 8383:8383 \
  timberio/vector:0.17.0-debian
```{{execute}}


```
docker logs -f $(docker ps -aqf "name=vector")
```{{execute}}
