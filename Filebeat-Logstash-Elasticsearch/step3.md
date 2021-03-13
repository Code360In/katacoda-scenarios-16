# LOAD
Now, we should generate traffic performing requests to the envoy proxy.

Install h2load to generate http2 traffic:

```
echo t2
```{{execute T2}}

`sudo apt install -y nghttp2-client`{{execute T2}}



# Start h2load

on terminal2

Run traffic with 10 clients and 100k requests:


100k
```
h2load -v "http://localhost:9000" -H "accept: application/json" -n 100000 -c 10
```{{execute T2}}

10k
```
h2load -v "http://localhost:9000" -H "accept: application/json" -n 10000 -c 10
```{{execute T2}}

1k
```
h2load -v "http://localhost:9000" -H "accept: application/json" -n 1000 -c 10
```{{execute T2}}


testpayload.json
```
cat << 'EOF' > /root/testpayload.json
{
    "query": "convert text into entities"
}
EOF
```{{execute}}

```
h2load -v http://localhost:9000 -d testpayload.json --h1 --header 'Content-Type: application/json' -n 500 -t 4 -c 4 -T 10
```{{execute}}


# Run RTOP to monitor Linux resources:

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


