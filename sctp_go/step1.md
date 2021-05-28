# SCTP

```
git clone https://github.com/ishidawataru/sctp.git
```{{execute}}


```
cd /root/sctp/ecample
```{{execute}}

```
go build
```{{execute}}



Get ipaddr ens3
```
ip_ens3=`ifconfig ens3 | awk '/inet / {print $2}'`
```{{execute}}



start SCTP server
```
./example -server -port 1000 -ip $ip_ens3
```{{execute}}


start SCTP client 
on terminal 2
```
cd /root/sctp/ecample
./example -port 1000 -ip $ip_ens3
```{{execute T2}}