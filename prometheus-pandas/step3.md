# Anaconda Jupyter 




## Anaconda Jupyter

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
```


https://[[HOST_SUBDOMAIN]]-8888-[[KATACODA_HOST]].environments.katacoda.com/




# install python prometheus-pandas

https://pypi.org/project/prometheus-pandas/



```
pip install prometheus-pandas

```{{execute}}

