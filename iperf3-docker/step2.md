

# MTU


Find MTU size
```
docker exec iperf3-server ip l  | grep eth0 | awk '{print $4,$5}'
```{{execute T2}}

# PING $ip


sending a packet larger than the current MTU setting, but is still accepted:
```
ping -M do -s 8972 $ip
```{{execute T2}}


Sending one too large:
```
ping -M do -s 118972 $ip
```{{execute T2}}

# Change MTU to 4000 bytes
```
docker exec iperf3-server ip link set eth0 mtu 4000
```{{execute T2}}


https://tools.ietf.org/html/rfc4459
https://tools.ietf.org/id/draft-ietf-intarea-frag-fragile-17.html#rfc.section.2.1

