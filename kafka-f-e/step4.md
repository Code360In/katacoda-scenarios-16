

# Deploy Filebeat input kafka

https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-overview.html

https://www.elastic.co/guide/en/beats/filebeat/current/running-on-docker.html

https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-kafka.html


Config file /root/filebeat.yml

```
cat << 'EOF' > /root/filebeat-input-kafka.yml
###################### Filebeat Configuration Example #########################

# This file is an example configuration file highlighting only the most common
# options. The filebeat.reference.yml file from the same directory contains all the
# supported options with more comments. You can use it as a reference.
#
# You can find the full configuration reference here:
# https://www.elastic.co/guide/en/beats/filebeat/index.html

# For more available modules and options, please see the filebeat.reference.yml sample
# configuration file.

# ============================== Filebeat inputs ===============================

filebeat.inputs:

# Each - is an input. Most options can be set at the input level, so
# you can use different inputs for various configurations.
# Below are the input specific configurations.

#- type: log

  # Change to true to enable this input configuration.
  #enabled: true

  # Paths that should be crawled and fetched. Glob based paths.
  #  paths:
  #    - /var/log/*.log

  # Exclude lines. A list of regular expressions to match. It drops the lines that are
  # matching any regular expression from the list.
  #exclude_lines: ['^DBG']

  # Include lines. A list of regular expressions to match. It exports the lines that are
  # matching any regular expression from the list.
  #include_lines: ['^ERR', '^WARN']

  # Exclude files. A list of regular expressions to match. Filebeat drops the files that
  # are matching any regular expression from the list. By default, no files are dropped.
  #exclude_files: ['/var/log/yum.log']

  # Optional additional fields. These fields can be freely picked
  # to add additional information to the crawled log files for filtering
  #fields:
  #  level: debug
  #  review: 1

  ### Multiline options

  # Multiline can be used for log messages spanning multiple lines. This is common
  # for Java Stack Traces or C-Line Continuation

  # The regexp Pattern that has to be matched. The example pattern matches all lines starting with [
  #multiline.pattern: ^\[

  # Defines if the pattern set under pattern should be negated or not. Default is false.
  #multiline.negate: false

  # Match can be set to "after" or "before". It is used to define if lines should be append to a pattern
  # that was (not) matched before or after or as long as a pattern is not matched based on negate.
  # Note: After is the equivalent to previous and before is the equivalent to to next in Logstash
  #multiline.match: after

# filestream is an experimental input. It is going to replace log input in the future.
- type: filestream

  # Change to true to enable this input configuration.
  enabled: false

  # Paths that should be crawled and fetched. Glob based paths.
  paths:
    - /var/log/*.log

  # Exclude lines. A list of regular expressions to match. It drops the lines that are
  # matching any regular expression from the list.
  #exclude_lines: ['^DBG']

  # Include lines. A list of regular expressions to match. It exports the lines that are
  # matching any regular expression from the list.
  #include_lines: ['^ERR', '^WARN']

  # Exclude files. A list of regular expressions to match. Filebeat drops the files that
  # are matching any regular expression from the list. By default, no files are dropped.
  #prospector.scanner.exclude_files: ['.gz$']

  # Optional additional fields. These fields can be freely picked
  # to add additional information to the crawled log files for filtering
  #fields:
  #  level: debug
  #  review: 1


- type: kafka
  hosts:
    - localhost:9092
  topics: ["my-topic"]
  group_id: "filebeat"



# ============================== Filebeat modules ==============================

filebeat.config.modules:
  # Glob pattern for configuration loading
  path: ${path.config}/modules.d/*.yml
  enabled: true

  # Set to true to enable config reloading
  reload.enabled: false

  # Period on which files under path should be checked for changes
  #reload.period: 10s

# ======================= Elasticsearch template setting =======================

setup.template.settings:
  index.number_of_shards: 1
  #index.codec: best_compression
  #_source.enabled: false


# ================================== General ===================================

# The name of the shipper that publishes the network data. It can be used to group
# all the transactions sent by a single shipper in the web interface.
#name:

# The tags of the shipper are included in their own field with each
# transaction published.
#tags: ["service-X", "web-tier"]

# Optional fields that you can specify to add additional information to the
# output.
#fields:
#  env: staging

# ================================= Dashboards =================================
# These settings control loading the sample dashboards to the Kibana index. Loading
# the dashboards is disabled by default and can be enabled either by setting the
# options here or by using the `setup` command.
#setup.dashboards.enabled: false

# The URL from where to download the dashboards archive. By default this URL
# has a value which is computed based on the Beat name and version. For released
# versions, this URL points to the dashboard archive on the artifacts.elastic.co
# website.
#setup.dashboards.url:

# =================================== Kibana ===================================

# Starting with Beats version 6.0.0, the dashboards are loaded via the Kibana API.
# This requires a Kibana endpoint configuration.
#setup.kibana:

  # Kibana Host
  # Scheme and port can be left out and will be set to the default (http and 5601)
  # In case you specify and additional path, the scheme is required: http://localhost:5601/path
  # IPv6 addresses should always be defined as: https://[2001:db8::1]:5601
  #host: "localhost:5601"

  # Kibana Space ID
  # ID of the Kibana Space into which the dashboards should be loaded. By default,
  # the Default Space will be used.
  #space.id:

# =============================== Elastic Cloud ================================

# These settings simplify using Filebeat with the Elastic Cloud (https://cloud.elastic.co/).

# The cloud.id setting overwrites the `output.elasticsearch.hosts` and
# `setup.kibana.host` options.
# You can find the `cloud.id` in the Elastic Cloud web UI.
#cloud.id:

# The cloud.auth setting overwrites the `output.elasticsearch.username` and
# `output.elasticsearch.password` settings. The format is `<user>:<pass>`.
#cloud.auth:

# ================================== Outputs ===================================

# Configure what output to use when sending the data collected by the beat.

# ---------------------------- Elasticsearch Output ----------------------------
output.elasticsearch:
  # Array of hosts to connect to.
  hosts: ["localhost:9200"]

  # Protocol - either `http` (default) or `https`.
  #protocol: "https"

  # Authentication credentials - either API key or username/password.
  #api_key: "id:api_key"
  #username: "elastic"
  #password: "changeme"

# ------------------------------ Logstash Output -------------------------------
#output.logstash:
  # The Logstash hosts
  #hosts: ["localhost:5044"]

  # Optional SSL. By default is off.
  # List of root certificates for HTTPS server verifications
  #ssl.certificate_authorities: ["/etc/pki/root/ca.pem"]

  # Certificate for SSL client authentication
  #ssl.certificate: "/etc/pki/client/cert.pem"

  # Client Certificate Key
  #ssl.key: "/etc/pki/client/cert.key"

# ================================= Processors =================================
processors:
  - add_host_metadata:
      when.not.contains.tags: forwarded
  - add_cloud_metadata: ~
  - add_docker_metadata: ~
  - add_kubernetes_metadata: ~

# ================================== Logging ===================================

# Sets log level. The default log level is info.
# Available log levels are: error, warning, info, debug
#logging.level: debug

# At debug level, you can selectively enable logging only for some components.
# To enable all selectors use ["*"]. Examples of other selectors are "beat",
# "publisher", "service".
#logging.selectors: ["*"]

# ============================= X-Pack Monitoring ==============================
# Filebeat can export internal metrics to a central Elasticsearch monitoring
# cluster.  This requires xpack monitoring to be enabled in Elasticsearch.  The
# reporting is disabled by default.

# Set to true to enable the monitoring reporter.
#monitoring.enabled: false

# Sets the UUID of the Elasticsearch cluster under which monitoring data for this
# Filebeat instance will appear in the Stack Monitoring UI. If output.elasticsearch
# is enabled, the UUID is derived from the Elasticsearch cluster referenced by output.elasticsearch.
#monitoring.cluster_uuid:

# Uncomment to send the metrics to Elasticsearch. Most settings from the
# Elasticsearch output are accepted here as well.
# Note that the settings should point to your Elasticsearch *monitoring* cluster.
# Any setting that is not set is automatically inherited from the Elasticsearch
# output configuration, so if you have the Elasticsearch output configured such
# that it is pointing to your Elasticsearch monitoring cluster, you can simply
# uncomment the following line.
#monitoring.elasticsearch:

# ============================== Instrumentation ===============================

# Instrumentation support for the filebeat.
#instrumentation:
    # Set to true to enable instrumentation of filebeat.
    #enabled: false

    # Environment in which filebeat is running on (eg: staging, production, etc.)
    #environment: ""

    # APM Server hosts to report instrumentation results to.
    #hosts:
    #  - http://localhost:8200

    # API Key for the APM Server(s).
    # If api_key is set then secret_token will be ignored.
    #api_key:

    # Secret token for the APM Server(s).
    #secret_token:


# ================================= Migration ==================================

# This allows to enable 6.7 migration aliases
#migration.6_to_7.enabled: true
EOF
```{{execute}}


Open file filebeat.yml

```
/root/filebeat-input-kafka.yml
```{{open}}



Deploy filebeat:

```
docker run -d --net=host --name=filebeat-i \
--user=root \
-v /root/filebeat-input-kafka.yml:/usr/share/filebeat/filebeat.yml \
-v /var/log/:/var/log:ro \
docker.elastic.co/beats/filebeat:7.11.2
```{{execute}}


Verify:
```
docker logs filebeat-i
```{{execute}}

Results:

`
2021-03-13T02:43:54.400Z        INFO    instance/beat.go:468    filebeat start running.
`

and


`
2021-03-13T16:30:17.024Z        INFO    log/harvester.go:302    Harvester started for file: /var/log/dpkg.log
2021-03-13T16:30:17.025Z        INFO    log/harvester.go:302    Harvester started for file: /var/log/kern.log
2021-03-13T16:30:17.025Z        INFO    log/harvester.go:302    Harvester started for file: /var/log/ubuntu-advantage.log
2021-03-13T16:30:17.026Z        INFO    log/harvester.go:302    Harvester started for file: /var/log/bootstrap.log
2021-03-13T16:30:17.026Z        INFO    log/harvester.go:302    Harvester started for file: /var/log/alternatives.log
2021-03-13T16:30:17.027Z        INFO    log/harvester.go:302    Harvester started for file: /var/log/boot.log
2021-03-13T16:30:17.027Z        INFO    log/harvester.go:302    Harvester started for file: /var/log/auth.log
`



### tcpdump
```
cnid=`docker ps | grep filebeat-i |awk 'NR==1{print $1}'`
pid=`docker inspect -f '{{.State.Pid}}' $cnid`
echo $pid
nsenter -t $pid --net tcpdump tcp
```{{execute}}



### netstat
```
cnid=`docker ps | grep filebeat-i |awk 'NR==1{print $1}'`
pid=`docker inspect -f '{{.State.Pid}}' $cnid`
echo $pid
nsenter -t $pid netstat -s
nsenter -t $pid netstat -at
```{{execute}}

count TCP ports open:
```
nsenter -t $pid netstat -at | wc -l
```{{execute}}


Display all indeces:
```
curl -X GET "localhost:9200/_cat/indices/*?v&s=index&pretty"
```{{execute}}

results:

<pre class="file">

yellow open   filebeat-7.11.2-2021.03.13-000001 VQ_UsZkMQ_q36RWXjU3Pdw   1   1      10483            0      1.4mb          1.4mb

</pre>




### Generate logs:

terminal 1

```
logger comment to be added to log
```{{execute}}


terminal 2

```
tail -f /var/log/syslog
```{{execute T2}}



Show all open ports:
```
netstat -at -n
```{{execute}}


<pre class="file">

$ netstat -at -n
Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State      
tcp        0      0 127.0.0.53:53           0.0.0.0:*               LISTEN     
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN     
tcp        0      0 127.0.0.1:9000          0.0.0.0:*               LISTEN     
tcp        0      0 0.0.0.0:111             0.0.0.0:*               LISTEN     
tcp        0      0 172.17.0.18:22          172.17.0.119:59604      ESTABLISHED
tcp        0      0 172.17.0.18:22          172.17.0.119:59360      ESTABLISHED
tcp6       0      0 :::8083                 :::*                    LISTEN     
tcp6       0      0 :::22                   :::*                    LISTEN     
tcp6       0      0 :::8088                 :::*                    LISTEN     
tcp6       0      0 :::9021                 :::*                    LISTEN     
tcp6       0      0 :::9092                 :::*                    LISTEN     
tcp6       0      0 :::2181                 :::*                    LISTEN     
tcp6       0      0 :::9101                 :::*                    LISTEN     
tcp6       0      0 :::111                  :::*                    LISTEN     
tcp6       0      0 :::8081                 :::*                    LISTEN     
tcp6       0      0 :::8082                 :::*                    LISTEN     
$ 

</pre>