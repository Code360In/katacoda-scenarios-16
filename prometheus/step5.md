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

starting benchmark...


## sending one request

```
h2load "https://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D" -H "accept: application/json" -n 1
```{{execute T2}}


## sending 1k requests

```
h2load "https://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D" -H "accept: application/json" -n 1000 -c 10
```{{execute T2}}




## sending 10k requests

```
h2load "https://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D" -H "accept: application/json" -n 1 -c 10
```{{execute T2}}
