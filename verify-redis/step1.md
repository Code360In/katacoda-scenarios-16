# [Redis](https://redis.io/)

# Install 
```
apt install -y redis
```{{execute}}

# Verify
```
systemctl status redis
```{{execute}}

```
journal -u redis
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
