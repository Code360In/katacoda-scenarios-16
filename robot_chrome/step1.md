# [Robot](https://robotframework.org/)


# Build image

```
docker build -t robot:chrome .
```{{execute}}

Wait for:
 <pre class="file">
Successfully built cfd4e019797c
Successfully tagged robot:chrome
 </pre>

Verify:
```
docker ps -a
```{{execute}}

# Deploy 
```
docker run --name=robot -d -p 80:80 robot:chrome
```{{execute}}

# Run
```
docker run exec robot robot --rpa --help
```{{execute}}
