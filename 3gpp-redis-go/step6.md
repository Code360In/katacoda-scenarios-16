# Task

Now, let's run grafana with this command:

`docker run --name=grafana -d -p 3000:3000 grafana/grafana:latest-ubuntu`{{execute T4}}

And access to the dashboard using this url:

https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/

And use this default credentials:
`admin`{{copy}} \ `admin`{{copy}}

The initial step will ask you to change the password, you can do it if you want or you can skip this step.


# Task

Let's define a datasource with [Tempo](https://grafana.com/oss/tempo/):

The URL for Tempo should be:

`https://[[HOST_SUBDOMAIN]]-16686-[[KATACODA_HOST]].environments.katacoda.com`{{copy}}

For all the other fields you can use the default values


Test and save your datasource.

