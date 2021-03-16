
# Deploy Kafka

https://docs.confluent.io/


create network

```
docker network create es
```{{execute}}

--network=es


Download or copy the contents of the [Confluent Platform all-in-one Docker Compose file](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html?utm_medium=sem&utm_source=google&utm_campaign=ch.sem_br.nonbrand_tp.prs_tgt.kafka_mt.mbm_rgn.namer_lng.eng_dv.all&utm_term=%2Bkafka%20%2Bdocker&creative=&device=c&placement=&gclid=Cj0KCQjwi7yCBhDJARIsAMWFScOdVPN4G23xe5z_MRhnQ78ZtyC1SADqb7zWLxvTUwMc1S8mvcmM-74aAtAKEALw_wcB), for example:

```
curl --silent --output docker-compose.yml \
  https://raw.githubusercontent.com/confluentinc/cp-all-in-one/6.1.0-post/cp-all-in-one/docker-compose.yml
```{{execute}}


Start Confluent Platform with the -d option to run in detached mode:

```
docker-compose up -d
```{{execute}}

results:

<pre class="file">

$ docker-compose ps
     Name                    Command               State                       Ports                     
---------------------------------------------------------------------------------------------------------
broker            /etc/confluent/docker/run        Up      0.0.0.0:9092->9092/tcp, 0.0.0.0:9101->9101/tcp
connect           /etc/confluent/docker/run        Up      0.0.0.0:8083->8083/tcp, 9092/tcp              
control-center    /etc/confluent/docker/run        Up      0.0.0.0:9021->9021/tcp                        
ksql-datagen      bash -c echo Waiting for K ...   Up                                                    
ksqldb-cli        /bin/sh                          Up                                                    
ksqldb-server     /etc/confluent/docker/run        Up      0.0.0.0:8088->8088/tcp                        
rest-proxy        /etc/confluent/docker/run        Up      0.0.0.0:8082->8082/tcp                        
schema-registry   /etc/confluent/docker/run        Up      0.0.0.0:8081->8081/tcp                        
zookeeper         /etc/confluent/docker/run        Up      0.0.0.0:2181->2181/tcp, 2888/tcp, 3888/tcp 


</pre>



To verify that the services are up and running, run the following command:

```
docker-compose ps
```{{execute}}


In this step, you create Kafka topics using Confluent Control Center. Confluent Control Center provides the functionality for building and monitoring production data pipelines and event streaming applications.

https://[[HOST_SUBDOMAIN]]-9021-[[KATACODA_HOST]].environments.katacoda.com/

Click the controlcenter.cluster tile.

# Go client

https://github.com/confluentinc/confluent-kafka-go.git

```
git clone https://github.com/confluentinc/confluent-kafka-go.git
```{{execute}}


```
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)
mkdir -p $GOPATH/src/github.com/user
go get -u gopkg.in/confluentinc/confluent-kafka-go.v1/kafka
```{{execute}}


build
```
cd /root/confluent-kafka-go/examples/admin_create_topic
```{{execute}}

```
go build admin_create_topic.go 
```{{execute}}

```
./admin_create_topic 
```{{execute}}

<pre class="file">

Usage: ./admin_create_topic <broker> <topic> <partition-count> <replication-factor>

</pre>



create topic:

```
./admin_create_topic localhost:9092 gugu 1 1

```{{execute}}


Producer:

Usage: ./producer_example <broker> <topic>


```
cd /root/confluent-kafka-go/examples/producer_example
```{{execute}}


```
go build producer_example.go 
```{{execute}}

```
./producer_example localhost:9092 1 
```{{execute}}


<pre class="file">

Created Producer rdkafka#producer-1
Delivered message to topic 1 [0] at offset 0

</pre>



Consumer:

Usage: ./consumer_example <broker> <group> <topics..>

```
cd /root/confluent-kafka-go/examples/consumer_example
```{{execute}}


```
go build consumer_example.go 
```{{execute}}


```
./consumer_example localhost:9092  gugu 1
```{{execute}}



results:

<pre class="file">

terminal 1

$ ./consumer_example localhost:9092  gugu 1
Created Consumer rdkafka#consumer-1
% Message on 1[0]@0:
Hello Go!
% Headers: [myTestHeader="header values are binary"]
Ignored OffsetsCommitted (<nil>, [1[0]@1])
% Message on 1[0]@1:
Hello Go!
% Headers: [myTestHeader="header values are binary"]
Ignored OffsetsCommitted (<nil>, [1[0]@2])
% Message on 1[0]@2:
Hello Go!
% Headers: [myTestHeader="header values are binary"]
Ignored OffsetsCommitted (<nil>, [1[0]@3])


terminal 2

$ ./producer_example localhost:9092 1
Created Producer rdkafka#producer-1
Delivered message to topic 1 [0] at offset 1
$ ./producer_example localhost:9092 1
Created Producer rdkafka#producer-1
Delivered message to topic 1 [0] at offset 2
$ 
</pre>



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