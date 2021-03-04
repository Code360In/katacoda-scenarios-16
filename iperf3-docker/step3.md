

# Monitoring

[BCC](https://github.com/iovisor/bcc/blob/master/INSTALL.md#ubuntu---binary)

https://github.com/iovisor/bcc

install bpftrace

```
echo t5
```{{execute T5}}

```
apt update && apt install -y bpftrace linux-headers-$(uname -r)
```{{execute T5}


install tools/examples:
```
echo "deb [trusted=yes] https://repo.iovisor.org/apt/xenial xenial-nightly main" | sudo tee /etc/apt/sources.list.d/iovisor.list
sudo apt-get update
sudo apt-get install -y bcc-tools libbcc-examples linux-headers-$(uname -r)
```{{execute T5}}

```
cd /usr/share/bcc/tools
```{{execute T5}}



# ebpf Exporter

To build a package from latest sources:
```
go get -u -v github.com/cloudflare/ebpf_exporter/...
```{{execute T5}}

To run with bio config:
```
~/go/bin/ebpf_exporter --config.file=src/github.com/cloudflare/ebpf_exporter/examples/bio.yaml
```{{execute T5}}


ref:

`
http://www.brendangregg.com/ebpf.html#bcc
https://kubernetes.io/blog/2017/12/using-ebpf-in-kubernetes/
https://github.com/cloudflare/ebpf_exporter.git
`