

##Refer to https://docs.docker.com/swarm/install-manual/


搭建5个节点的DockerSwarm集群: Consul节点1个，Swarm Manager主/从各1个，Swarm Node两个 
https://docs.docker.com/swarm/install-manual/ 
Consul节点运行   
	docker run -d -p 8500:8500 --name=consul progrium/consul -server -bootstrap 
Swarm Primary Manager      :   
	docker run -d -p 4000:4000 swarm manage -H :4000 --replication --advertise <manager0_ip>:4000 consul://192.168.56.101:8500 
Swarm Secondary Manager : 
	docker run -d -p 4000:4000 swarm manage -H :4000 --replication --advertise <manager1_ip>:4000 consul://192.168.56.101:8500 
  
Swarm Node 1: 
	docker run -d  swarm join --advertise=<node1_ip>:2375 consul://192.168.56.101:8500 
Swarm Node 2: 
	docker run -d  swarm join --advertise=<node2_ip>:2375 consul://192.168.56.101:8500 
  
	
在Swarm主节点上运行 
	docker -H :4000 info 
	docker -H :4000 run hello-world 
	docker -H :4000 ps -a 

删除Exited的容器 
	docker -H :4000 ps -a | grep Exited | awk '{print $1}' | docker -H :4000 rm 




##High availability in Docker Swarm  
https://docs.docker.com/swarm/multi-manager-setup/
