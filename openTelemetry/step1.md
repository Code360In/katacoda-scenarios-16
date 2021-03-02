# Deploys a load generator, agent and gateway as well as Jaeger, Zipkin and Prometheus back-ends. 

Go to Demo:
```
cd /root/opentelemetry-collector/examples/demo/
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