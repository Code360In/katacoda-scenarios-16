# Heartbeat / Uptime



following DEB instructions:

https://[[HOST_SUBDOMAIN]]-5601-[[KATACODA_HOST]].environments.katacoda.com/app/home#/tutorial/uptimeMonitors



### Download and install Heartbeat

First time using Heartbeat? See the Quick Start.


```
curl -L -O https://artifacts.elastic.co/downloads/beats/heartbeat/heartbeat-7.11.1-amd64.deb
sudo dpkg -i heartbeat-7.11.1-amd64.deb
```{{execute}}


Modify /etc/heartbeat/heartbeat.yml to set the connection information:

<pre class="file">

output.elasticsearch:
  hosts: ["<es_url>"]
  username: "elastic"
  password: "<password>"
setup.kibana:
  host: "<kibana_url>"

</pre>


Where <password> is the password of the elastic user, <es_url> is the URL of Elasticsearch, and <kibana_url> is the URL of Kibana.



### Edit the configuration - Add monitors

Edit the heartbeat.monitors setting in the heartbeat.yml file.


<pre class="file">

heartbeat.monitors:
- type: http
  urls: ["<http://localhost:9200>"]
  schedule: "@every 10s"

</pre>


Where <http://localhost:9200> is your monitored URL, For more details on how to configure Monitors in Heartbeat, read the Heartbeat configuration docs.



Open file heartbeat.yml
```
/etc/heartbeat/heartbeat.yml
```{{open}}


### Start Heartbeat

The setup command loads the Kibana index pattern.


```
sudo heartbeat setup
sudo service heartbeat-elastic start
```{{execute}}



And access to the dashboard using this url:

https://[[HOST_SUBDOMAIN]]-5601-[[KATACODA_HOST]].environments.katacoda.com/app/uptime

