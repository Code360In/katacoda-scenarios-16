# LOAD

Now, we should generate traffic performing requests to the envoy proxy.

Install h2load to generate http2 traffic:

```
sudo apt install -y nghttp2-client
```{{execute}}




### Start h2load

on terminal2

Run traffic with 10 clients and 100k requests:




testpayload.json
```
cat << 'EOF' > /root/testpayload.json
{ "streams": [ { "labels": "{source=\"JSON\",job=\"simpleJsonJob\", host=\"SimpleHost\"}", "entries": [{ "ts": "olddate", "line": "TEST!" }] } ] }
EOF
```{{execute}}

newdate:
```
date --rfc-3339=ns | sed 's/ /T/'
```{{execute}}


```
sed -i 's/olddate/newdate/' /root/testpayload.json
```



500
```
h2load -v http://localhost:3100/loki/api/v1/push -d testpayload.json --h1 --header 'Content-Type: application/json' -n 500 -t 4 -c 4 -T 10
```{{execute}}



100k
```
h2load -v http://localhost:3100/loki/api/v1/push -d testpayload.json --h1 --header 'Content-Type: application/json' -n 100000 -t 4 -c 4 -T 10
```{{execute}}

10k
```
h2load -v http://localhost:3100/loki/api/v1/push -d testpayload.json --h1 --header 'Content-Type: application/json' -n 10000 -t 4 -c 4 -T 10
```{{execute}}

1k
```
h2load -v http://localhost:3100/loki/api/v1/push -d testpayload.json --h1 --header 'Content-Type: application/json' -n 1000 -t 4 -c 4 -T 10
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


