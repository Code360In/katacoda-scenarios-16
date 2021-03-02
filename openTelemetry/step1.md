# Deploys a load generator, agent and gateway as well as Jaeger, Zipkin and Prometheus back-ends. 


Get code:

```
git clone https://github.com/open-telemetry/opentelemetry-collector.git
```{{execute}}


Go to Demo:
```
cd opentelemetry-collector/examples/demo;
```{{execute}}

Build all images and create container:
```
docker-compose up -d
```{{execute}}




Verify:
```
docker ps -a
```{{execute}}


<pre class="file">
.
</pre>