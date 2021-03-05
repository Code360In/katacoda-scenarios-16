

# Linux statistics


display complete report 

on terminal5
```
echo t5
```{{execute T5}}


# netstat on host


```
netstat -s
netstat -i 
```{{execute T5}}

# netstat on container


```
docker exec iperf3-server  netstat -s
docker exec iperf3-server  netstat -i 
```{{execute T5}}



ref:

`
`