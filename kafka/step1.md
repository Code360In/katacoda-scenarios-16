# Deploy Kafka

https://docs.confluent.io/


Download or copy the contents of the [Confluent Platform all-in-one Docker Compose file](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html?utm_medium=sem&utm_source=google&utm_campaign=ch.sem_br.nonbrand_tp.prs_tgt.kafka_mt.mbm_rgn.namer_lng.eng_dv.all&utm_term=%2Bkafka%20%2Bdocker&creative=&device=c&placement=&gclid=Cj0KCQjwi7yCBhDJARIsAMWFScOdVPN4G23xe5z_MRhnQ78ZtyC1SADqb7zWLxvTUwMc1S8mvcmM-74aAtAKEALw_wcB), for example:

```
curl --silent --output docker-compose.yml \
  https://raw.githubusercontent.com/confluentinc/cp-all-in-one/6.1.0-post/cp-all-in-one/docker-compose.yml
```{{execute}}


Start Confluent Platform with the -d option to run in detached mode:

```
docker-compose up -d
```{{execute}}


To verify that the services are up and running, run the following command:

```
docker-compose ps
```{{execute}}


In this step, you create Kafka topics using Confluent Control Center. Confluent Control Center provides the functionality for building and monitoring production data pipelines and event streaming applications.

https://[[HOST_SUBDOMAIN]]-9021-[[KATACODA_HOST]].environments.katacoda.com/

Click the controlcenter.cluster tile.


# Client python

https://pypi.org/project/kafka-python/


``` 
pip install kafka-python

```{{execute}}

KafkaProducer

>>> from kafka import KafkaProducer
>>> producer = KafkaProducer(bootstrap_servers='localhost:1234')
>>> for _ in range(100):
...     producer.send('foobar', b'some_message_bytes')




KafkaConsumer

>>> from kafka import KafkaConsumer
>>> consumer = KafkaConsumer('my_favorite_topic')
>>> for msg in consumer:
...     print (msg)
