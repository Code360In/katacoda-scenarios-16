
``` 
cd /root/opentelemetry-go-contrib/exporters/metric/cortex/example
sed -i 's/host.docker.internal/localhost/' config.yml
```{{execute}}


``` 
go mod verify
go build
```{{execute}}



run:


``` 
 ./example 
```{{execute}}


