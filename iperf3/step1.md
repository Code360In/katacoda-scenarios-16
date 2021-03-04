# IPERF3 Install


[Packages](https://iperf.fr/iperf-download.php#more-recent) are manually installed via the dpkg command (Debian Package Management System). dpkg is the backend to commands like apt and aptitude, which in turn are the backend for GUI install apps like the Software Center and Synaptic.


<pre>[[HOST_IP]]</pre>

`[[HOST_IP]]`{{execute}}

<pre>[[HOST2_IP]]</pre>

`[[HOST2_IP]]`{{execute}}


- Ubuntu 64 bits / Debian 64 bits / Mint 64 bits (AMD64) :

1.
```
sudo apt remove iperf3 libiperf0
```{{execute}}

2.
```
sudo apt install libsctp1
```{{execute}}

3.
```
wget https://iperf.fr/download/ubuntu/libiperf0_3.9-1_amd64.deb
```{{execute}}

4.
```
wget https://iperf.fr/download/ubuntu/iperf3_3.9-1_amd64.deb
```{{execute}}

5.
```
sudo dpkg -i libiperf0_3.9-1_amd64.deb iperf3_3.9-1_amd64.deb
```{{execute}}

6.
```
rm libiperf0_3.9-1_amd64.deb iperf3_3.9-1_amd64.deb
```{{execute}}


# Run SCTP tests


host1=[[host1]]
host2=[[host2]]

```

```{{execute}}
