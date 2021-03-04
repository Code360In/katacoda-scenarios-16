<pre class="file">
Before routing a request to the appropriate service Envoy or the application, Envoy will take care of generating the appropriate spans for tracing (parent/child context spans). At a high-level, each span records the latency of upstream API calls as well as information needed to correlate the span with other related spans (e.g., the trace ID).

One of the most important benefits of tracing from Envoy is that it will take care of propagating the traces to the Jaeger service cluster. However, in order to fully take advantage of tracing, the application has to propagate trace headers that Envoy generates, while making calls to other services. In the sandbox we have provided, the simple flask app (see trace function in front-proxy/service.py) acting as service1 propagates the trace headers while making an outbound call to service2.
</pre>
