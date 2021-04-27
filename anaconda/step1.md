# Anaconda Jupyter TimescaleDB




## Create network

```
docker network create --driver bridge pnet
```{{execute}}


## Anaconda Jupyter

https://anaconda.org/

```
docker run -i -t -d --name conda --network=pnet -p 8888:8888 continuumio/anaconda3 /bin/bash -c "/opt/conda/bin/conda install jupyter -y --quiet && mkdir /opt/notebooks && /opt/conda/bin/jupyter notebook --notebook-dir=/opt/notebooks --ip='*' --port=8888 --allow-root " 
```{{execute}}

##  timescaledb


Install and run TimescaleDB with Promscale extension:

```
docker run --name timescaledb -e POSTGRES_PASSWORD=secret -d -p 5432:5432 --network=pnet timescaledev/promscale-extension:latest-pg12 postgres -csynchronous_commit=off
```{{execute}}


## Create database name:db

```
docker exec -it timescaledb  createdb db -U postgres
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




##  Psycopg

https://www.psycopg.org/docs/install.html
https://www.postgresqltutorial.com/postgresql-python/connect/
https://anaconda.org/anaconda/psycopg2


<pre class="file">

!pip install psycopg2-binary



import psycopg2

# Connect to an existing database
>>> conn = psycopg2.connect(
    host="timescaledb",
    database="db",
    user="postgres",
    password="secret")

# Open a cursor to perform database operations
>>> cur = conn.cursor()

# Execute a command: this creates a new table
>>> cur.execute("CREATE TABLE test (id serial PRIMARY KEY, num integer, data varchar);")

# Pass data to fill a query placeholders and let Psycopg perform
# the correct conversion (no more SQL injections!)
>>> cur.execute("INSERT INTO test (num, data) VALUES (%s, %s)",
...      (100, "abc'def"))

# Query the database and obtain data as Python objects
>>> cur.execute("SELECT * FROM test;")
>>> cur.fetchone()
(1, 100, "abc'def")

# Make the changes to the database persistent
>>> conn.commit()

# Close communication with the database
>>> cur.close()
>>> conn.close()


</pre>

# seaborn
https://seaborn.pydata.org/installing.html

<pre class="file">

import seaborn as sns
df = sns.load_dataset("penguins")
sns.pairplot(df, hue="species")

</pre>


# pandas
<pre class="file">

import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

</pre>


#  networkx
https://networkx.org/documentation/networkx-2.3/auto_examples/drawing/plot_simple_path.html#sphx-glr-auto-examples-drawing-plot-simple-path-py

<pre class="file">

import matplotlib.pyplot as plt
import networkx as nx

G = nx.path_graph(8)
nx.draw(G)
plt.show()

</pre>

# bokeh

https://github.com/bokeh/bokeh/blob/2.3.0/examples/howto/notebook_comms/Basic%20Usage.ipynb

<pre class="file">

from bokeh.io import push_notebook, show, output_notebook
from bokeh.layouts import row
from bokeh.plotting import figure
output_notebook()
opts = dict(plot_width=250, plot_height=250, min_border=0)

p1 = figure(**opts)
r1 = p1.circle([1,2,3], [4,5,6], size=20)

p2 = figure(**opts)
r2 = p2.circle([1,2,3], [4,5,6], size=20)

# get a handle to update the shown cell with
t = show(row(p1, p2), notebook_handle=True)

# this will update the left plot circle color with an explicit handle
r1.glyph.fill_color = "white"
push_notebook(handle=t)

# and this will update the right plot circle color because it was in the last shown cell
r2.glyph.fill_color = "pink"
push_notebook()

p3 = figure(**opts)
r3 = p3.circle([1,2,3], [4,5,6], size=20)

# get a handle to update the shown cell with
t2 = show(p3, notebook_handle=True)
# show which cell t2 handles
t2

# this updates the immediately previous cell with an explicit handle
r3.glyph.fill_color = "orange"
push_notebook(handle=t2)


</pre>

ref:

https://docs.timescale.com/latest/tutorials/tutorial-forecasting
