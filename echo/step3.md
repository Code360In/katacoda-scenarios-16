# LOAD
Now, we should generate traffic performing requests to the echo.

Install h2load to generate http2 traffic:

```
echo t2
```{{execute T2}}

`sudo apt install -y nghttp2-client`{{execute T2}}



# Start h2load

on terminal2

Run traffic with 10 clients and 100k requests:


```
h2load -v "http://localhost:1323" -H "accept: application/json" -n 100000 -c 10
```{{execute T2}}


ref:

https://nghttp2.org/documentation/h2load-howto.html


