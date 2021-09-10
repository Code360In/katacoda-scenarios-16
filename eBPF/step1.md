# eBPF

https://github.com/iovisor/bcc/blob/master/INSTALL.md#ubuntu---binary




```
sudo apt-get install bpfcc-tools linux-headers-$(uname -r)
```{{execute}}



# pixie-labs

https://github.com/pixie-labs/pixie-demos


```
git clone https://github.com/pixie-labs/pixie-demos.git
```{{execute}}


```
cd /root/pixie-demos/simple-gotracing
```{{execute}}

```
 cd app
 go build
 ./app 
```{{execute}}
Starting server on: :9090


terminal2


```
echo terminal2
```{{execute T2}}



```
 curl localhost:9090/e
```{{execute T2}}

e = 2.7183

```
cd /root/pixie-demos/simple-gotracing/
ls
```{{execute T2}}


go get -u github.com/iovisor/gobpf/bcc

https://github.com/iovisor/bcc/blob/master/INSTALL.md#ubuntu---binary


echo "deb [trusted=yes] https://repo.iovisor.org/apt/xenial xenial-nightly main" | sudo tee /etc/apt/sources.list.d/iovisor.list
sudo apt-get update
sudo apt-get install bcc-tools libbcc-examples linux-headers-$(uname -r)

