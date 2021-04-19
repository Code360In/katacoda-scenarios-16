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

verify:


<pre class="file">


TLS Protocol: TLSv1.3
Cipher: TLS_AES_128_GCM_SHA256
Server Temp Key: X25519 253 bits
Application protocol: h2


or


TLS Protocol: TLSv1.3
Cipher: TLS_AES_128_GCM_SHA256
Server Temp Key: X25519 253 bits
No protocol negotiated. Fallback behaviour may be activated
Server does not support NPN/ALPN. Falling back to HTTP/1.1.
Application protocol: http/1.1


</pre>



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
h2load "https://localhost:9091/federate?match%5B%5D=%7B__name__%3D~%22.%2A%22%7D" -H "accept: application/json" -n 10000 -c 10
```{{execute T2}}



### test exporter port 8443

```
h2load "https://localhost:8443/metrics" -H "accept: application/json" -n 10000 -c 10
```{{execute T2}}



