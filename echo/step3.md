# LOAD
Now, we should generate traffic performing requests to the echo.

Install h2load to generate http2 traffic:

```
echo t2
```{{execute T2}}

`sudo apt install -y nghttp2-client`{{execute T2}}



# Start h2load

on terminal2

Run traffic with 10 clients and 100 requests:


```
h2load http://localhost:1323  --h1 -n 100 -c 10
```{{execute T2}}

results:

<pre class="file">

starting benchmark...
spawning thread #0: 100 total client(s). 10000 total requests
Application protocol: http/1.1
progress: 10% done
progress: 20% done
progress: 30% done
progress: 40% done
progress: 50% done
progress: 60% done
progress: 70% done
progress: 80% done
progress: 90% done
progress: 100% done

finished in 20.16s, 496.03 req/s, 62.97KB/s
requests: 10000 total, 10000 started, 10000 done, 10000 succeeded, 0 failed, 0 errored, 0 timeout
status codes: 10000 2xx, 0 3xx, 0 4xx, 0 5xx
traffic: 1.24MB (1300000) total, 839.84KB (860000) headers (space savings 0.00%), 126.95KB (130000) data
                     min         max         mean         sd        +/- sd
time for request:   200.07ms    219.19ms    201.06ms      1.48ms    91.96%
time for connect:     2.36ms     12.18ms      5.94ms      2.93ms    63.00%
time to 1st byte:   204.63ms    215.93ms    212.51ms      3.06ms    75.00%
req/s           :       4.96        4.98        4.97        0.00    50.00%
$ 

</pre>




ref:

https://nghttp2.org/documentation/h2load-howto.html


