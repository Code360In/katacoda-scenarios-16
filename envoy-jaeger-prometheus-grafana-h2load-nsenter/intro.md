[Envoy](https://www.envoyproxy.io/) is easily integrated with open source projects in order to have metrics and distributed tracing.

The following scenario demonstrates how to:

- Expose Envoy's statistics in [Prometheus](https://prometheus.io/).  

- Use prometheus datasource to build dashboards in [Grafana](https://grafana.com/).

- Configure envoy to send traces to [Jaeger](https://www.jaegertracing.io/).

- Configure Metrics for Jaeger and Grafana dashboard.

- Use h2load to generate load to envoy.

- Use nsenter tcpdump to verify UDP msg to Jaeger



https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/jaeger_native_tracing.html?highlight=jaeger