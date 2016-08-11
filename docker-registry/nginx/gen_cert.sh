
#https://www.digitalocean.com/community/tutorials/how-to-set-up-a-private-docker-registry-on-ubuntu-14-04

#create user and passwd
htpasswd -c registry.passwd XXXXXXX
#add other user
htpasswd registry.passwd YYYYYY



#Signing Your Own Certificate
#Generate a new root key:
openssl genrsa -out devdockerCA.key 2048
#Generate a root certificate
openssl req -x509 -new -nodes -key devdockerCA.key -days 10000 -out devdockerCA.crt
#Generate a key for your server
openssl genrsa -out domain.key 2048
#make a certificate signing request, "Common Name" should be the domain or IP of your server
openssl req -new -key domain.key -out dev-docker-registry.com.csr
#sign the certificate request
openssl x509 -req -in dev-docker-registry.com.csr -CA devdockerCA.crt -CAkey devdockerCA.key -CAcreateserial -out domain.crt -days 10000



#On local client machine
sudo mkdir /usr/local/share/ca-certificates/docker-dev-cert
sudo cp devdockerCA.crt /usr/local/share/ca-certificates/docker-dev-cert
sudo update-ca-certificates
sudo service docker restart

#
