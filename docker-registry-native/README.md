## Local Docker Registry with Native Basic Auth

###https://docs.docker.com/registry/deploying/

###https://docs.docker.com/registry/spec/auth/token/




First create a password file with one entry for the user “testuser”, with password “testpassword”:
```
mkdir auth
docker run --entrypoint htpasswd registry:2 -Bbn testuser testpassword > auth/htpasswd
```


Make sure you stopped your registry from the previous step, then start it again:
```
docker run -d -p 5000:5000 --restart=always --name registry \
  -v `pwd`/auth:/auth \
  -e "REGISTRY_AUTH=htpasswd" \
  -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
  -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
  -v `pwd`/certs:/certs \
  -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt \
  -e REGISTRY_HTTP_TLS_KEY=/certs/domain.key \
  registry:2
```


You should now be able to:
```
docker login myregistrydomain.com:5000
```
