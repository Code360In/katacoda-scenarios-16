

# MTU


Find MTU size
```
docker exec iperf3-server ip l  | grep eth0 | awk '{print $4,$5}'
```{{execute T2}}


```
ping -h
```{{execute T2}}



# PING to $ip
`
-M define mtu discovery, can be one of <do|dont|want>
-s <size>          use <size> as number of data bytes to be sent
`

sending a packet larger than the current MTU setting:
```
ping -M do -s 1501 -c 10 $ip
```{{execute T2}}


Sending one too large:
```
ping -M do -s 118972 -c 10 $ip
```{{execute T2}}

<pre class="file">

Expected Error: packet size 118972 is too large. Maximum is 65507

</pre>

ok:
```
ping -M do -s 1472 -c 10 $ip
```{{execute T2}}

fail:
```
ping -M do -s 1474 -c 10 $ip
```{{execute T2}}


# TCPDUMP docker0 interface

on terminal4
```
echo t4
```{{execute T4}}

```
tcpdump -i docker0 icmp
```{{execute T4}}



# Change MTU 

...

```
```{{execute T2}}


# IP Fragmentation
https://tools.ietf.org/html/rfc4459
https://tools.ietf.org/id/draft-ietf-intarea-frag-fragile-17.html#rfc.section.2.1

