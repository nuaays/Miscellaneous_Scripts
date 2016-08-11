##Self-signed certificates
#####Refer to https://docs.docker.com/registry/insecure/

mkdir -p certs && openssl genrsa -out certs/domain.key 4096 && openssl req -x509 -new -nodes -sha256 -key certs/domain.key -days 10000 -out certs/domain.crt



#on ubuntu
cp certs/domain.crt /usr/local/share/ca-certificates/myregistrydomain.com.crt
update-ca-certificates
#on red hat or centos
cp certs/domain.crt /etc/pki/ca-trust/source/anchors/myregistrydomain.com.crt
update-ca-trust
update-ca-trust enable

service docker stop 
service docker start
