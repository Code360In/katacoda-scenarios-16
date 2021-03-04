

# Monitoring

[BCC](https://github.com/iovisor/bcc/blob/master/INSTALL.md#ubuntu---binary)

https://github.com/iovisor/bcc

install bpftrace
```
apt update && apt install -y bpftrace linux-headers-$(uname -r)
```{{execute T2}}


install tools/examples:
```
echo "deb [trusted=yes] https://repo.iovisor.org/apt/xenial xenial-nightly main" | sudo tee /etc/apt/sources.list.d/iovisor.list
sudo apt-get update
sudo apt-get install -y bcc-tools libbcc-examples linux-headers-$(uname -r)
```{{execute T2}}

```
cd /usr/share/bcc/tools
```{{execute T2}}



# ebpf_exporter

```
git clone https://github.com/cloudflare/ebpf_exporter.git

 go get -u -v github.com/cloudflare/ebpf_exporter/...
 
```{{execute T2}}

ref:
http://www.brendangregg.com/ebpf.html#bcc
https://kubernetes.io/blog/2017/12/using-ebpf-in-kubernetes/
https://github.com/cloudflare/ebpf_exporter.git