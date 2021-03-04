# [Robot](https://robotframework.org/)


# Build image

```
docker build -t robot:chrome .
```{{execute}}

# Deploy 

```
docker run --name=robot -d -p 80:80 robot:chrome
```{{execute}}

# Run

```
docker run exec robot --help
```{{execute}}
