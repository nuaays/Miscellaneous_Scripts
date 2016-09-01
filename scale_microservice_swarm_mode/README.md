
#Scale a Microservice using Docker 1.12 Swarm Mode

```
npm init
npm install exporess uuid

docker swarm init --advertise-addr 10.253.6.128
docker node ls
docker service create -p 9000:9000 --name guid swarm-mode-guid
curl localhost:9000
docker service ls
docker service update --replicas=100 guid
docker service scale guid=10
service scale guid=2
docker service rm guid

```
