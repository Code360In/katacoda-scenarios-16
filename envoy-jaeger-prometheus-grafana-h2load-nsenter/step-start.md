An initial envoy configuration file has been created at
`envoy.yaml`{{open}}

In this file, it is defined that the server will run with a listener using all network interfaces in port 10000.

Start the envoy proxy with the defined configuration using this command:



To build this sandbox example:

```
pwd;ls
cd envoy/examples/jaeger-native-tracing
```{{execute}}

```
docker-compose pull
```{{execute}}

```
docker-compose up --build -d
```{{execute}}

```
docker-compose ps
```{{execute}}

You should get an answer similar to:
<pre class="file">
 Name                              Command                State                     Ports
jaeger-native-tracing_front-envoy_1   /start-front.sh                Up      10000/tcp, 0.0.0.0:8000->8000/tcp
jaeger-native-tracing_jaeger_1        /go/bin/all-in-one-linux - ... Up      14250/tcp, 14268/tcp, 0.0.0.0:16686->16686/tcp, 5775/udp, 5778/tcp, 6831/udp, 6832/udp
jaeger-native-tracing_service1_1      /start-service.sh              Up      10000/tcp
jaeger-native-tracing_service2_1      /start-service.sh              Up      10000/tcp

</pre>


You can now send a request to service1 via the front-envoy as follows:

```curl -v localhost:8000/trace/1```{{execute}}




Envoy is answering the request and balancing between the two nodes with a `ROUND_ROBIN` strategy according to our configuration.

Also, you can test this via your local browser with the URL https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/
