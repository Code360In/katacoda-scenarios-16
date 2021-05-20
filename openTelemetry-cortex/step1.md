

Clone:
``` 
git clone https://github.com/open-telemetry/opentelemetry-go-contrib.git
cd opentelemetry-go-contrib/
```{{execute}}



Run the docker container with the following command:
``` 
cd exporters/metric/cortex/example
docker-compose up -d
```{{execute}}



Connect to Grafana


https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com


Get ipadr from Cortex:
``` 
docker inspect example_cortex_1 | grep IPAddress
```{{execute}}




```
$ docker inspect example_cortex_1  | grep IPAddress
            "SecondaryIPAddresses": null,
            "IPAddress": "",
                    "IPAddress": "172.19.0.2",
$ 
```




Add a new Prometheus Data Source

Use http://172.19.0.2:9009/api/prom/ as the URL

Optionally, set the scrape interval to 3s to make updates appear quickly

Click Save & Test



Connect to Cortex:

https://[[HOST_SUBDOMAIN]]-9009-[[KATACODA_HOST]].environments.katacoda.com