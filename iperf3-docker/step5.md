
# Deploy Grafana

Now, let's run grafana with this command:

`docker run --name=grafana --net=host -d -p 3000:3000 grafana/grafana:master-ubuntu`{{execute T2}}

And access to the dashboard using this url:

https://[[HOST_SUBDOMAIN]]-3000-[[KATACODA_HOST]].environments.katacoda.com/

And use this default credentials:
`admin`{{copy}} \ `admin`{{copy}}

The initial step will ask you to change the password, you can do it if you want or you can skip this step.

The first step to create a dashboard is to have a datasource. Let's define a datasource with phometheus data configured before:

The URL for phometheus should be:

`http://localhost:9090`{{copy}}

`http://localhost:9090`{{Execute T2}}


For all the other fields you can use the default values

![](import.png)

Test and save your datasource.
