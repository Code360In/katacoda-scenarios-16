# OpenTracing

Distributed tracing allows developers to obtain visualizations of call flows in large service oriented architectures. It can be invaluable in understanding serialization, parallelism, and sources of latency. For this example we will use Jaeger, an open source, end-to-end distributed tracing.



Ensure that Jaeger is configured to accept Zipkin requests via the *COLLECTOR_ZIPKIN_HTTP_PORT* Environment Variable.

One important configuration for our example is telling to the connection manager that generates the
 `x-request-id` header if it does not exist.


```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 9411:9411 -p 5775:5775/udp -p 16686:16686 -p 6831:6831/udp -p 6832:6832/udp\
  --net=host  -u root \
  jaegertracing/all-in-one:latest
```{{execute}}


Logs:

```
docker logs grafana
```{{execute}}


access:


https://[[HOST_SUBDOMAIN]]-16686-[[KATACODA_HOST]].environments.katacoda.com/


https://m3db.io/docs/operational_guide/monitoring/

