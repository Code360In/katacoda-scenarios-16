# Anaconda Jupyter TimescaleDB




## Create network

```
docker network create --driver bridge pnet
```{{execute}}


## Anaconda Jupyter

https://anaconda.org/

```
docker run -i -t -d --name conta --network=pnet -p 8888:8888 continuumio/anaconda3 /bin/bash -c "/opt/conda/bin/conda install jupyter -y --quiet && mkdir /opt/notebooks && /opt/conda/bin/jupyter notebook --notebook-dir=/opt/notebooks --ip='*' --port=8888 --allow-root " 
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
```


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


<pre class="file">

import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

</pre>


ref:

https://docs.timescale.com/latest/tutorials/tutorial-forecasting
