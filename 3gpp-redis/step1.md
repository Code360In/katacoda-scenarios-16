3GPP test new stuff


# Install [Redis](https://redis.io/)
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


Change to executable:
```
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


# Start test example:Nnrf

https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com/TS29510_Nnrf_NFManagement.html#/


Change API.root to be  https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com


![](3gpp-2.png)

ref:

https://github.com/OpenAPITools/openapi-generator
https://github.com/APIDevTools/swagger-parser
https://www.3gpp.org/DynaReport/29-series.htm
https://www.3gpp.org/ftp/Specs/archive/OpenAPI/Rel-15


todo:

https://www.3gpp.org/ftp/Specs/archive/OpenAPI/Rel-16


<pre class="file">
</pre>
