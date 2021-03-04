# [Robot](https://robotframework.org/)


# Build image

```
docker build -t robot:01 .
```{{execute}}

Wait for:
 <pre class="file">
Successfully built cfd4e019797c
Successfully tagged robot:chrome
 </pre>

# Verify:
```
docker ps -a
```{{execute}}

# Deploy:
```
docker run --name=robot -d -p 80:80 robot:01
```{{execute}}

# Run:
```
docker exec robot robot --rpa --help
```{{execute}}

Clone demo:
```
docker exec robot git clone https://github.com/robotframework/WebDemo.git
```{{execute}}


Run demo:
```
docker exec robot robot /WebDemo/login_tests/ valid_login.robot
```{{execute}}