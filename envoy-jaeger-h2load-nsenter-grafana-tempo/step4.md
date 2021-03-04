# LOAD
Now, we should generate traffic performing requests to the envoy proxy.

Install h2load to generate http2 traffic:

`sudo apt install -y nghttp2-client`{{execute}}


# start TCPDUMP on container


```
cnid=`docker ps | grep front-envoy |awk 'NR==1{print $1}'`
pid=`docker inspect -f '{{.State.Pid}}' $cnid`
echo $pid
nsenter -t $pid --net tcpdump udp
```{{execute T2}}


#Start h2load
Run traffic with 100 clients and 1000 requests:

`h2load http://localhost:8000/trace/1 -c 100 -n 1000 `{{execute T1}}
