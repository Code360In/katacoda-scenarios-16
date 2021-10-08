


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

