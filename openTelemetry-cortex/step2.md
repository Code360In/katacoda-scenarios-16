
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



# Intall Cortex tools

https://github.com/grafana/cortex-tools/releases


# download the binary (adapt os and arch as needed)

``` 
curl -fSL -o "/usr/local/bin/cortextool" "https://github.com/grafana/cortex-tools/releases/download/v0.10.1/cortextool_0.10.1_linux_x86_64"
```{{execute}}



# make it executable
``` 
chmod a+x "/usr/local/bin/cortextool"
```{{execute}}


# have fun :)

``` 
cortextool --help
```{{execute}}


# Generate load

``` 
cortextool loadgen --write-url=http://127.0.0.1:9009 --active-series=136805  --series-name="node_cpu_seconds_total"  --parallelism=100  --batch-size=1000
```{{execute}}


# Verify
https://[[HOST_SUBDOMAIN]]-9009-[[KATACODA_HOST]].environments.katacoda.com/distributor/all_user_stats

