# Deploys a load generator, agent and gateway as well as Jaeger, Zipkin and Prometheus back-ends. 


Get code:

```
git clone https://github.com/open-telemetry/opentelemetry-collector.git
```{{execute}}


go to demo dir:
```
cd /root/opentelemetry-collector/examples/demo;
```{{execute}}

Build all images and Deploy containers:
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