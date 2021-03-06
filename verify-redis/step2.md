
# Start test


```
/usr/bin/python3 -m pip install --upgrade pip
pip3 install redis
```{{execute}}


```
git clone https://github.com/valdemarpavesi/verifyREDIS
```{{execute}}


<pre class="file">

Usage: verifyREDIS.py [options] --create --delete --verify [--keyname] [--host] [--port] 

Options:
  -h, --help         show this help message and exit
  --create           set keys
  --delete           del keys
  --verify           verify REDIS database SET DEL VERIFY
  --keyname=KEYNAME  key name prefix
  --host=HOST        REDIS ipaddr
  --port=PORT        REDIS tcp port

</pre>
 


```
python3 /root/verifyREDIS/verifyREDIS.py
```{{execute}}

```
python3 /root/verifyREDIS/verifyREDIS.py --help
```{{execute}}

# CREATE
```
python3 /root/verifyREDIS/verifyREDIS.py --create
```{{execute}}

# DELETE
```
python3 /root/verifyREDIS/verifyREDIS.py --delete
```{{execute}}


# VERIFY
```
python3 /root/verifyREDIS/verifyREDIS.py --verify
```{{execute}}

<pre class="file">
</pre>


<pre class="file">
# end.
</pre>

