### Create local Docker Registry with authorization using Nginx
Refer to https://www.digitalocean.com/community/tutorials/how-to-set-up-a-private-docker-registry-on-ubuntu-14-04


### Docker Registry V2 HTTP API
* https://docs.docker.com/registry/spec/api/
* https://github.com/docker/docker/issues/9015


###

## Docker Registry V2 with auth using Nginx


####How to launch this Docker Registry
```
install docker-compose firstly:  pip install docker-compose 

[start] docker-compose up -d  
[stop ] docker-compose stop  
[remove] docker-compose rm  
```


####How to create/add/del users
```
[create] htpasswd -c registry.passwd ${username}
[  add ] htpasswd registry.passwd ${username} 
[  del ] htpasswd -D ${username}

```

####How to pull/push images to this registry
```
1.docker login --username=test --password=test --email=**@XXX.com ${registry_ip}:5000  
2.test curl -u test:test ****:5000/v2/_catalog 
3.[pull] docker pull ****:5000/swarm:latest  
  [push] docker tag  ** && docker push 
  
