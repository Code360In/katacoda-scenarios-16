# 3GPP test new stuff

# Start

Wait for file `/root/5GcoreOpenApiV1530 ` to be loaded:
```
pwd
ls
```{{execute}}


# Install [Redis](https://redis.io/)

Redis key-value database to store Json.



```
apt install -y redis
```{{execute}}

```
redis-cli info
```{{execute}}

```
redis-cli get keys
```{{execute}}


```
redis-cli SET mykey "Hello"
```{{execute}}


```
redis-cli GET mykey
```{{execute}}

# Start

Change to executable:
```
ls
chmod +x /root/5GcoreOpenApiV1530 
```{{execute}}

# Run
```
./5GcoreOpenApiV1530 
```{{execute}}


<pre class="file">
Webpages are embedded into the binnary.
</pre>


You can access the website using this link:

https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com

![](3gpp-1.png)


ref:

https://github.com/OpenAPITools/openapi-generator
https://github.com/APIDevTools/swagger-parser
https://www.3gpp.org/DynaReport/29-series.htm
https://www.3gpp.org/ftp/Specs/archive/OpenAPI/Rel-15
https://www.etsi.org/deliver/etsi_ts/123500_123599/123501/15.03.00_60/ts_123501v150300p.pdf
https://www.3gpp.org/news-events/1930-sys_architecture



todo:

https://www.3gpp.org/ftp/Specs/archive/OpenAPI/Rel-16
