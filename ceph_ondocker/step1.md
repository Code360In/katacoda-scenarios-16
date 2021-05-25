# Deploy Ceph


https://docs.ceph.com/en/latest/cephadm/install/


update linux
```
apt-get update
```{{execute}}


cephadm

```
apt install -y cephadm
```{{execute}}


Get ipaddr ens3
```
ip_ens3=`ifconfig ens3 | awk '/inet / {print $2}'`
```{{execute}}


bootstrap

```
cephadm bootstrap --mon-ip  $ip_ens3
```{{execute}}


Verify:

<pre class="file">
Ceph Dashboard is now available at:

             URL: https://localhost:8443/
            User: admin
        Password: cx7ynhiwh4

You can access the Ceph CLI with:

        sudo /usr/sbin/cephadm shell --fsid 5a4f5066-bd5e-11eb-a42b-d58d1675a004 -c /etc/ceph/ceph.conf -k /etc/ceph/ceph.client.admin.keyring

Please consider enabling telemetry to help improve Ceph:

        ceph telemetry on

For more information see:

        https://docs.ceph.com/docs/master/mgr/telemetry/

</pre>



Enable Telemetry
```
apt install -y ceph-common
```{{execute}}


```
ceph telemetry on --license sharing-1-0
```{{execute}}

```
docker ps
```{{execute}}


ceph status 
```
ceph status  
```{{execute}}


```
ceph osd create
```{{execute}}

```
ceph fs volume create fstest
```{{execute}}

https://docs.ceph.com/en/latest/cephfs/createfs/


```
ceph fs flag set enable_multiple true --yes-i-really-mean-it
```{{execute}}

```
ceph osd pool create cephfs_data
ceph osd pool create cephfs_metadata
```{{execute}}


```
ceph fs new cephfs cephfs_metadata cephfs_data
```{{execute}}

```
ceph mds stat
```{{execute}}

```
ceph health detail
```{{execute}}

```
ceph osd tree
```{{execute}}
