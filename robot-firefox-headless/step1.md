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



Start build:
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
docker logs grafana
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

![](robot_exec.png)

c) robot will connect to url https://www.katacoda.com/
`
( next step robot will start the kotacoda scenario and start robot, robot, robot....)`