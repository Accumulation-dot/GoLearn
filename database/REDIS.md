# Redis



## install
```
sudo apt-get install redis-server
```  

## config
```
redis 开启远程访问
# bind 127.0.0.1
protected-mode no

redis

maxmemory 256mb
maxmemory-policy allkeys-lru

重启redis 服务

sudo service redis-server restart

```