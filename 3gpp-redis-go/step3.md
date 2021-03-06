# LOAD
Now, we should generate traffic performing requests to the envoy proxy.

Install h2load to generate http2 traffic:

```
echo t2
```{{execute T2}}

`sudo apt install -y nghttp2-client`{{execute T2}}



# Start h2load

on terminal2

Run traffic with 4 clients and 500 requests:


```
h2load -v https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com -d "{\"reqNotifEvents\":[null],\"notifCondition\":{\"unmonitoredAttributes\":[\"string\"],\"monitoredAttributes\":[\"string\"]},\"plmnId\":{\"mnc\":\"string\",\"mcc\":\"string\"},\"nfStatusNotificationUri\":\"string\",\"reqNfFqdn\":\"string\",\"validityTime\":\"2021-03-06T01:48:35.399Z\"}" --h1 --header 'Content-Type: application/json' -n 500 -t 4 -c 4 -T 10
```{{execute T2}}


# Run RTOP

on terminal3
```
echo t3
```{{execute T3}}

`go get github.com/rapidloop/rtop `{{execute T3}}
`go build github.com/rapidloop/rtop `{{execute T3}}

Run
`rtop 127.0.0.1 `{{execute T3}}


ref:

https://nghttp2.org/documentation/h2load-howto.html