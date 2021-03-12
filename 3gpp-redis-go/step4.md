# OpenMetrics 

You can access the website using this link:

https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com/metrics




<pre class="file">
promhttp_metric_handler_requests_total{code="503"} 0
# HELP request_duration_seconds Time (in seconds) spent serving HTTP requests.
# TYPE request_duration_seconds histogram
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="0.005"} 0
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="0.01"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="0.025"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="0.05"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="0.1"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="0.25"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="0.5"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="1"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="2.5"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="5"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="10"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="200",ws="false",le="+Inf"} 1
request_duration_seconds_sum{method="GET",route="other",status_code="200",ws="false"} 0.007334103
request_duration_seconds_count{method="GET",route="other",status_code="200",ws="false"} 1
request_duration_seconds_bucket{method="GET",route="other",status_code="204",ws="false",le="0.005"} 129742
request_duration_seconds_bucket{method="GET",route="other",status_code="204",ws="false",le="0.01"} 129986
request_duration_seconds_bucket{method="GET",route="other",status_code="204",ws="false",le="0.025"} 130000
</pre>


https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile


`
example:
request_duration_seconds_bucket{method="GET",route="other",status_code="204",ws="false",le="0.01"} 2000

A 99th percentile latency of 10 ms means that every 1 in 100 requests experience 10 ms of delay.

`

