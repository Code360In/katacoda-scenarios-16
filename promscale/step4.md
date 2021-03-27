# Verification



psql:

```
docker exec -it timescaledb psql -U postgres
```{{execute}}

to show all databases: 

```
\l
```{{execute}}


show all tables:

```
\dt
```{{execute}}


Grafana:

```
select * from metric;
```{{copy}}


```
SELECT *
FROM pg_catalog.pg_tables
WHERE schemaname != 'pg_catalog' AND 
    schemaname != 'information_schema';
```{{copy}}



```
http_request_duration_milliseconds_sum
```{{copy}}


[PromQL Compliance Test Results: Promscale (2020-12-01)](https://promlabs.com/promql-compliance-test-results/2020-12-01/promscale/)

[PromQL Compliance Test Results: Overview](https://promlabs.com/promql-compliance-tests/)


[Prometheus HA](https://github.com/timescale/promscale/blob/master/docs/high-avaliability/prometheus-HA.md)
![](https://github.com/timescale/promscale/blob/master/docs/assets/promscale-HA-arch.png)



https://docs.timescale.com/latest/tutorials/tutorial-howto-monitor-django-prometheus

