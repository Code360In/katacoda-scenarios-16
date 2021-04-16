
# Deploy Cortex

following:
https://www.infracloud.io/blogs/cortex-prometheus-ha/

Install multiple Prometheus and Grafana to visualize the data.


``` 
git clone https://github.com/kanuahs/cortex-demo.git
cd cortex-demo
```{{execute}}



``` 
docker-compose -f docker-demo/docker-compose.yaml up
```{{execute}}


test , go to grafana explore:
```
up{cluster="one"}
prometheus_tsdb_head_samples_appended_total{cluster="one"}
```



clean-up
```
docker-compose -f docker-demo/docker-compose.yaml down
```{{execute}}
