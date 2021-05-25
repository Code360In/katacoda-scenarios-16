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
ifconfig
```{{execute}}


bootstrap

```
cephadm bootstrap --mon-ip  172.17.0.26
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
