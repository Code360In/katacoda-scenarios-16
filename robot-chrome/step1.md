# [Robot](https://robotframework.org/)



# Build image

wait to get:
```
ls
```{{execute}}
`
$ ls
docker-entrypoint.sh  Dockerfile  test.robot
`


```
docker build -t robot:01 .
```{{execute}}

wait for:
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
docker run --name=robot -d -p 8000:8000 robot:01
```{{execute}}

# Run:
```
docker exec robot robot --rpa --help
```{{execute}}


An example of robot file `test.robot`{{open}}


Run demo:

a) bash to container robot:
```
docker exec -it robot bash
```{{execute}}


b) run robot calling test.robot file.
```
robot /home/robotuser/test.robot
```{{execute}}

![](robot.png)

c) robot will connect to url https://www.katacoda.com/
`
( next step robot will start the kotacoda scenario and start robot, robot, robot....)`