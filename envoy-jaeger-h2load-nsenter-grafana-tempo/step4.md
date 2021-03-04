# LOAD
Now, we should generate traffic performing requests to the envoy proxy.

Install h2load to generate http2 traffic:

`sudo apt install -y nghttp2-client`{{execute}}


# start TCPDUMP on container

on terminal2
```
echo t2
```{{execute T2}}


```
cnid=`docker ps | grep front-envoy |awk 'NR==1{print $1}'`
pid=`docker inspect -f '{{.State.Pid}}' $cnid`
echo $pid
nsenter -t $pid --net tcpdump udp
```{{execute T2}}


#Start h2load

on terminal1

Run traffic with 100 clients and 1000 requests:

`h2load http://localhost:8000/trace/1 -c 100 -n 1000 `{{execute T1}}

# Run RTOP

on terminal3
```
echo t3
```{{execute T3}}

`go get github.com/rapidloop/rtop `{{execute T3}}
`go build github.com/rapidloop/rtop `{{execute T3}}

Run
`rtop 127.0.0.1 `{{execute T3}}
