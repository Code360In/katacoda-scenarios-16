
# LOAD to Promscale


Test, add one metric to timescaleDB:

```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"labels":{"__name__":"foo"},"samples":[[1577836800000, 100]]}' \
"http://localhost:9201/write"
```{{execute}}



Now, we should generate traffic .

Install h2load to generate http2 traffic:

```
sudo apt install -y nghttp2-client
```{{execute}}




### Start h2load

on terminal2

Run traffic with 10 clients and 100k requests:


```
curl -H "Content-Type: application/json" -XPOST -s "http://localhost:3100/loki/api/v1/push" --data-raw "{\"streams\": [{\"stream\": {\"job\": \"test\"}, \"values\": [[\"$(date +%s)000000000\", \"fizzbuzz\"]]}]}"

```{{execute}}


```
curl -H "Content-Type: application/json" -XPOST -s "http://localhost:9201/write" --data-raw "{"labels":{"__name__":"foo"},"samples":[[\"$(date +%s)000000000\", 100]]}"

```{{execute}}



newdate:
```
date --rfc-3339=ns | sed 's/ /T/'
```{{execute}}


testpayload.json

```
echo '{"labels":{"__name__":"foo"},"samples":[[1577836800000, 100]]}' > /root/testpayload.json
```{{execute}}


1
```
h2load -vvv http://localhost:9201/write -d /root/testpayload.json --h1 --header 'Content-Type: application/json' -n 1 -t 1 -c 1 -T 10

```{{execute}}


500
```
h2load -vvv http://localhost:9201/write -d /root/testpayload.json --h1 --header 'Content-Type: application/json' -n 500 -t 2 -c 4 -T 10

```{{execute}}






### Run RTOP to monitor Linux resources: 

on terminal3

```
echo t3
```{{execute T3}}


```go get github.com/rapidloop/rtop```{{execute T3}}


```go build github.com/rapidloop/rtop```{{execute T3}}

Run

```rtop 127.0.0.1 ```{{execute T3}}


ref:

https://nghttp2.org/documentation/h2load-howto.html



