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

Display Docker file content `Dockerfile`{{open}}


```
docker build -t robot:01 .
```{{execute}}

wait for:
 <pre class="file">
Successfully built xxxxx
Successfully tagged robot:01
 </pre>

# Verify:
```
docker ps -a
```{{execute}}

# Deploy:
```
docker run --name=robot -d -p 8000:8000 robot:01
```{{execute}}

Verify
```
docker logs robot
```{{execute}}


# Run:
```
docker exec robot robot --rpa --help
```{{execute}}


An example of robot file `test.robot`{{open}}


Run demo:

a) create dir on container:
```
docker exec robot mkdir /home/robotuser
```{{execute}}

b) Copy robot file to container:
```
docker cp /root/test.robot robot:/home/robotuser/
```{{execute}}

c) bash to container
```
docker exec -it robot bash
```{{execute}}


d) run robot calling test.robot file.
```
robot /home/robotuser/test.robot
```{{execute}}

![](robot_exec.png)


e) robot will connect to url https://www.katacoda.com/
`
( next step robot will start the kotacoda scenario and start robot, robot, robot....)`