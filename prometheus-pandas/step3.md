# Anaconda Jupyter 


https://anaconda.org/

```
docker run -i -t -d --net=host --name conda  -p 8888:8888 continuumio/anaconda3 /bin/bash -c "/opt/conda/bin/conda install jupyter -y --quiet && mkdir /opt/notebooks && /opt/conda/bin/jupyter notebook --notebook-dir=/opt/notebooks --ip='*' --port=8888 --allow-root " 
```{{execute}}



Connect to Jupyter:


```
curl localhost:8888/test
```{{execute}}

get token.

```
docker logs conda
```{{execute}}


https://[[HOST_SUBDOMAIN]]-8888-[[KATACODA_HOST]].environments.katacoda.com/




# prometheus-pandas

https://pypi.org/project/prometheus-pandas/



new jupyter notebook python3

new cell , install 
```
!pip install prometheus-pandas

```{{copy}}


```
p = query.Prometheus('http://localhost:9090')
```{{copy}}

```
p.query('process_cpu_seconds_total{}', '2021-05-16T00:00:00Z')
```{{copy}}



copy notebook from:

https://raw.githubusercontent.com/dcoles/prometheus-pandas/master/Prometheus.ipynb

