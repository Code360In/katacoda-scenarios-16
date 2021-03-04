

# IPERF3 SCTP Deploy and run tests.


[iperf3](https://iperf.fr/iperf-doc.php)

[docker iperf3](https://hub.docker.com/r/networkstatic/iperf3/)



# Iperf3 Server

Start a listener service on port 5201 and name the container "iperf3-server":
```
docker run  -it --rm --name=iperf3-server -p 5201:5201 networkstatic/iperf3 -s
```{{execute T1}}

That returns an iperf3 process bound to a soccer waiting for new connections:

<pre class="file">

-----------------------------------------------------------
Server listening on 5201
-----------------------------------------------------------

</pre>


# Iperf3 Client SIde
First, get the IP address of the new server container ypu just started:

```
echo t2
```{{execute T2}}

```
ip=`docker inspect --format "{{ .NetworkSettings.IPAddress }}" iperf3-server`
```{{execute T2}}

```
echo $ip
```{{execute T2}}

Next, initiate a client connection from another container to measure the bandwidth between the two endpoints.

Run a client container pointing at the server service IP address.

Note if you are new to Docker, the --rm flag will destroy the container after the test runs. I also left out explicitly naming the container on the client side since I don't need its IP address. I typically explicitly name containers for organization and to maintain a consistent pattern.

```
docker run  -it --rm networkstatic/iperf3 -c $ip --sctp
```{{execute T2}}


And the output is the following:

<pre class="file">

Connecting to host 172.17.0.163, port 5201
[  4] local 172.17.0.191 port 51148 connected to 172.17.0.163 port 5201
[ ID] Interval           Transfer     Bandwidth       Retr  Cwnd
[  4]   0.00-1.00   sec  4.16 GBytes  35.7 Gbits/sec    0    468 KBytes
[  4]   1.00-2.00   sec  4.10 GBytes  35.2 Gbits/sec    0    632 KBytes
[  4]   2.00-3.00   sec  4.28 GBytes  36.8 Gbits/sec    0   1.02 MBytes
[  4]   3.00-4.00   sec  4.25 GBytes  36.5 Gbits/sec    0   1.28 MBytes
[  4]   4.00-5.00   sec  4.20 GBytes  36.0 Gbits/sec    0   1.37 MBytes
[  4]   5.00-6.00   sec  4.23 GBytes  36.3 Gbits/sec    0   1.40 MBytes
[  4]   6.00-7.00   sec  4.17 GBytes  35.8 Gbits/sec    0   1.40 MBytes
[  4]   7.00-8.00   sec  4.14 GBytes  35.6 Gbits/sec    0   1.40 MBytes
[  4]   8.00-9.00   sec  4.29 GBytes  36.8 Gbits/sec    0   1.64 MBytes
[  4]   9.00-10.00  sec  4.15 GBytes  35.7 Gbits/sec    0   1.68 MBytes
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bandwidth       Retr
[  4]   0.00-10.00  sec  42.0 GBytes  36.1 Gbits/sec    0             sender
[  4]   0.00-10.00  sec  42.0 GBytes  36.0 Gbits/sec                  receiver

iperf Done.


</pre>

# Runt TCPDUMP on container

on terminal3
```
echo t3
```{{execute T3}}


```
cnid=`docker ps | grep iperf3-server |awk 'NR==1{print $1}'`
pid=`docker inspect -f '{{.State.Pid}}' $cnid`
echo $pid
nsenter -t $pid --net tcpdump sctp
```{{execute T3}}


# Run traffic again:

```
docker run  -it --rm networkstatic/iperf3 -c $ip --sctp
```{{execute T2}}


`
Let's change the world a little bit.
`