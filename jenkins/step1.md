# Jenkins
The Jenkins Continuous Integration and Delivery server.

This is a fully functional Jenkins server, based on the Long Term Support release http://jenkins.io/.

For weekly releases check out [jenkinsci/jenkins](https://hub.docker.com/_/jenkins)


### Deploy
```
docker run --name=jenkins -d -p 8080:8080 -p 50000:50000 jenkins:lts
```{{execute}}


### Verify
```
docker ps -a
```{{execute}}


### Logs
```
docker logs jenkins
```{{execute}}


access by:
https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com