

## Swarm Mode Overview
* https://docs.docker.com/engine/swarm/

```
Feature highlights
Cluster management integrated with Docker Engine: Use the Docker Engine CLI to create a swarm of Docker Engines where you can deploy application services. You don’t need additional orchestration software to create or manage a swarm.
Decentralized design: Instead of handling differentiation between node roles at deployment time, the Docker Engine handles any specialization at runtime. You can deploy both kinds of nodes, managers and workers, using the Docker Engine. This means you can build an entire swarm from a single disk image.
Declarative service model: Docker Engine uses a declarative approach to let you define the desired state of the various services in your application stack. For example, you might describe an application comprised of a web front end service with message queueing services and a database backend.
Scaling: For each service, you can declare the number of tasks you want to run. When you scale up or down, the swarm manager automatically adapts by adding or removing tasks to maintain the desired state.
Desired state reconciliation: The swarm manager node constantly monitors the cluster state and reconciles any differences between the actual state and your expressed desired state. For example, if you set up a service to run 10 replicas of a container, and a worker machine hosting two of those replicas crashes, the manager will create two new replicas to replace the replicas that crashed. The swarm manager assigns the new replicas to workers that are running and available.
Multi-host networking: You can specify an overlay network for your services. The swarm manager automatically assigns addresses to the containers on the overlay network when it initializes or updates the application.
Service discovery: Swarm manager nodes assign each service in the swarm a unique DNS name and load balances running containers. You can query every container running in the swarm through a DNS server embedded in the swarm.
Load balancing: You can expose the ports for services to an external load balancer. Internally, the swarm lets you specify how to distribute service containers between nodes.
Secure by default: Each node in the swarm enforces TLS mutual authentication and encryption to secure communications between itself and all other nodes. You have the option to use self-signed root certificates or certificates from a custom root CA.
Rolling updates: At rollout time you can apply service updates to nodes incrementally. The swarm manager lets you control the delay between service deployment to different sets of nodes. If anything goes wrong, you can roll-back a task to a previous version of the service.
```


## Swarm Mode Key Concepts
* https://docs.docker.com/engine/swarm/key-concepts/

* What is a swarm?
```
The cluster management and orchestration features embedded in the Docker Engine are built using SwarmKit. Docker engines participating in a cluster are running in swarm mode. You enable swarm mode for an engine by either initializing a swarm or joining an existing swarm.
A swarm is a cluster of Docker engines, or nodes, where you deploy services. The Docker Engine CLI and API include commands to manage swarm nodes (e.g., add or remove nodes), and deploy and orchestrate services across the swarm.
When you run Docker without using swarm mode, you execute container commands. When you run the Docker in swarm mode, you orchestrate services. You can run swarm services and standalone containers on the same Docker instances.
```

* What is a node?
```
A node is an instance of the Docker engine participating in the swarm. You can also think of this as a Docker node. You can run one or more nodes on a single physical computer or cloud server, but production swarm deployments typically include Docker nodes distributed across multiple physical and cloud machines.
To deploy your application to a swarm, you submit a service definition to a manager node. The manager node dispatches units of work called tasks to worker nodes.
Manager nodes also perform the orchestration and cluster management functions required to maintain the desired state of the swarm. Manager nodes elect a single leader to conduct orchestration tasks.
Worker nodes receive and execute tasks dispatched from manager nodes. By default manager nodes also run services as worker nodes, but you can configure them to run manager tasks exclusively and be manager-only nodes. An agent runs on each worker node and reports on the tasks assigned to it. The worker node notifies the manager node of the current state of its assigned tasks so that the manager can maintain the desired state of each worker.
```

* Services and tasks
```
A service is the definition of the tasks to execute on the worker nodes. It is the central structure of the swarm system and the primary root of user interaction with the swarm.
When you create a service, you specify which container image to use and which commands to execute inside running containers.
In the replicated services model, the swarm manager distributes a specific number of replica tasks among the nodes based upon the scale you set in the desired state.
For global services, the swarm runs one task for the service on every available node in the cluster.
A task carries a Docker container and the commands to run inside the container. It is the atomic scheduling unit of swarm. Manager nodes assign tasks to worker nodes according to the number of replicas set in the service scale. Once a task is assigned to a node, it cannot move to another node. It can only run on the assigned node or fail.
```

* Load balancing
```
The swarm manager uses ingress load balancing to expose the services you want to make available externally to the swarm. The swarm manager can automatically assign the service a PublishedPort or you can configure a PublishedPort for the service. You can specify any unused port. If you do not specify a port, the swarm manager assigns the service a port in the 30000-32767 range.
External components, such as cloud load balancers, can access the service on the PublishedPort of any node in the cluster whether or not the node is currently running the task for the service. All nodes in the swarm route ingress connections to a running task instance.
Swarm mode has an internal DNS component that automatically assigns each service in the swarm a DNS entry. The swarm manager uses internal load balancing to distribute requests among services within the cluster based upon the DNS name of the service.
```


## How Service Work
* https://docs.docker.com/engine/swarm/how-swarm-mode-works/services/
## How Node Work
* https://docs.docker.com/engine/swarm/how-swarm-mode-works/nodes/







## Get Started with Docker Swarm 

* docker swarm init --advertise-addr 172.16.0.2
Swarm initialized: current node (yloq2964nwvqnu39jkn6bq579) is now a manager.

To add a worker to this swarm, run the following command:
```
    docker swarm join \
    --token SWMTKN-1-5h9xzburikl2lkyj5gfp778rbqj4gpc0ri02bjv6cqshbzf68z-4pk0obugc6xg2yuxiffrm1ywo \
    172.16.0.2:2377
```
To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```
>>>docker node ls
ID                           HOSTNAME        STATUS  AVAILABILITY  MANAGER STATUS
56hsb1n1kbgcv3f49e04t55f3    VM_0_26_centos  Ready   Active        
yloq2964nwvqnu39jkn6bq579 *  VM_0_2_centos   Ready   Active        Leader

>>>docker info
Swarm: inactive

Swarm: active
 NodeID: yloq2964nwvqnu39jkn6bq579
 Is Manager: true
 ClusterID: sz2kkx68xfnelwb008su93y2a
 Managers: 1
 Nodes: 2
 Orchestration:
  Task History Retention Limit: 5
 Raft:
  Snapshot Interval: 10000
  Number of Old Snapshots to Retain: 0
  Heartbeat Tick: 1
  Election Tick: 3
 Dispatcher:
  Heartbeat Period: 5 seconds
 CA Configuration:
  Expiry Duration: 3 months
 Node Address: 172.16.0.2
 Manager Addresses:
  172.16.0.2:2377

Swarm: active
 NodeID: 56hsb1n1kbgcv3f49e04t55f3
 Is Manager: false
 Node Address: 172.16.0.26
 Manager Addresses:
  172.16.0.2:2377
```



>>>docker swarm join-token worker
To add a worker to this swarm, run the following command:

    docker swarm join \
    --token SWMTKN-1-5h9xzburikl2lkyj5gfp778rbqj4gpc0ri02bjv6cqshbzf68z-4pk0obugc6xg2yuxiffrm1ywo \
    172.16.0.2:2377

>>>docker swarm join-token manager
To add a manager to this swarm, run the following command:

    docker swarm join \
    --token SWMTKN-1-5h9xzburikl2lkyj5gfp778rbqj4gpc0ri02bjv6cqshbzf68z-4hsnl6tp32n3au36nb19qrmyo \
    172.16.0.2:2377



* Deploy a service to the swarm

[root@VM_0_2_centos ~]# docker service create --replicas 1 --name helloworld 172.16.0.2:5000/alpine:latest ping 192.168.1.174
ke9gw18jilvne3w7kuyrcyw8j

[root@VM_0_2_centos ~]# docker service ls 
ID            NAME        MODE        REPLICAS  IMAGE
ke9gw18jilvn  helloworld  replicated  1/1       alpine:latest

* Inspect a service on the swarm

[root@VM_0_2_centos ~]# docker service inspect --pretty helloworld  

ID:             ke9gw18jilvne3w7kuyrcyw8j
Name:           helloworld
Service Mode:   Replicated
 Replicas:      1
Placement:
UpdateConfig:
 Parallelism:   1
 On failure:    pause
 Max failure ratio: 0
ContainerSpec:
 Image:         alpine:latest@sha256:58e1a1bb75db1b5a24a462dd5e2915277ea06438c3f105138f97eb53149673c4
 Args:          ping 192.168.1.174 
Resources:
Endpoint Mode:  vip


[root@VM_0_2_centos ~]# docker service ps helloworld
ID            NAME          IMAGE          NODE           DESIRED STATE  CURRENT STATE          ERROR  PORTS
ni48bzefc4bg  helloworld.1  alpine:latest  VM_0_2_centos  Running        Running 4 minutes ago   


* Scale the service in the swarm

[root@VM_0_2_centos ~]# docker service scale helloworld=5
helloworld scaled to 5

[root@VM_0_2_centos ~]# docker service ps helloworld
ID            NAME          IMAGE                          NODE            DESIRED STATE  CURRENT STATE           ERROR  PORTS
k1sy7wcq4702  helloworld.1  172.16.0.2:5000/alpine:latest  VM_0_46_centos  Running        Running 34 seconds ago         
t4zb7bws5zxc  helloworld.2  172.16.0.2:5000/alpine:latest  VM_0_2_centos   Running        Running 6 seconds ago          
s1w0sq29drnq  helloworld.3  172.16.0.2:5000/alpine:latest  VM_0_26_centos  Running        Running 5 seconds ago          
llz8hximz6mb  helloworld.4  172.16.0.2:5000/alpine:latest  VM_0_26_centos  Running        Running 5 seconds ago          
wu78li2mrm2p  helloworld.5  172.16.0.2:5000/alpine:latest  VM_0_50_centos  Running        Running 5 seconds ago 

[root@VM_0_2_centos ~]# docker service inspect --pretty helloworld

ID:             rx4g03ewemppk2a3duy6aaons
Name:           helloworld
Service Mode:   Replicated
 Replicas:      5
Placement:
UpdateConfig:
 Parallelism:   1
 On failure:    pause
 Max failure ratio: 0
ContainerSpec:
 Image:         172.16.0.2:5000/alpine:latest@sha256:d0a670140e7d73562b401819f88fb68edca8aed45f4d2f835edb294e4a7a152a
 Args:          ping 192.168.1.174 
Resources:
Endpoint Mode:  vip


* Apply rolling updates to a service

[root@VM_0_2_centos ~]# docker service create --replicas 3 --name redis --update-delay 10s 172.16.0.2:5000/env/redis:3.0.5
sfhx6l23mcsmq418ax78p1jvv

[root@VM_0_2_centos ~]# docker service inspect --pretty redis


[root@VM_0_2_centos ~]# docker service ps redis
ID            NAME         IMAGE                            NODE            DESIRED STATE  CURRENT STATE            ERROR  PORTS
dzknvwuh2cm1  redis.1      172.16.0.2:5000/env/redis:3.0.7  VM_0_46_centos  Running        Running 9 minutes ago           
sipq05e991iz   \_ redis.1  172.16.0.2:5000/env/redis:3.0.5  VM_0_46_centos  Shutdown       Shutdown 9 minutes ago          
kkp2pq5txom3  redis.2      172.16.0.2:5000/env/redis:3.0.7  VM_0_50_centos  Running        Running 9 minutes ago           
wgjn2t1k6a0r   \_ redis.2  172.16.0.2:5000/env/redis:3.0.5  VM_0_50_centos  Shutdown       Shutdown 9 minutes ago          
lxxyu11q3cyu  redis.3      172.16.0.2:5000/env/redis:3.0.7  VM_0_2_centos   Running        Running 10 minutes ago          
sa8m4a340bj0   \_ redis.3  172.16.0.2:5000/env/redis:3.0.5  VM_0_2_centos   Shutdown       Shutdown 10 minutes ago        

* Drain a node on the swarm

[root@VM_0_2_centos ~]# docker node update --availability drain VM_0_50_centos
VM_0_50_centos

[root@VM_0_2_centos ~]# docker service ps redis
ID            NAME         IMAGE                            NODE            DESIRED STATE  CURRENT STATE            ERROR  PORTS
dzknvwuh2cm1  redis.1      172.16.0.2:5000/env/redis:3.0.7  VM_0_46_centos  Running        Running 11 minutes ago          
sipq05e991iz   \_ redis.1  172.16.0.2:5000/env/redis:3.0.5  VM_0_46_centos  Shutdown       Shutdown 11 minutes ago         
ujecuw2bmc22  redis.2      172.16.0.2:5000/env/redis:3.0.7  VM_0_26_centos  Running        Preparing 6 seconds ago         
kkp2pq5txom3   \_ redis.2  172.16.0.2:5000/env/redis:3.0.7  VM_0_50_centos  Shutdown       Shutdown 5 seconds ago          
wgjn2t1k6a0r   \_ redis.2  172.16.0.2:5000/env/redis:3.0.5  VM_0_50_centos  Shutdown       Shutdown 11 minutes ago         
lxxyu11q3cyu  redis.3      172.16.0.2:5000/env/redis:3.0.7  VM_0_2_centos   Running        Running 11 minutes ago          
sa8m4a340bj0   \_ redis.3  172.16.0.2:5000/env/redis:3.0.5  VM_0_2_centos   Shutdown       Shutdown 11 minutes ago  

[root@VM_0_2_centos ~]# docker node inspect --pretty VM_0_50_centos
ID:                     8s1b3c09ftug83ppqt4ol6wxx
Hostname:               VM_0_50_centos
Joined at:              2017-03-29 07:51:05.912493323 +0000 utc
Status:
 State:                 Ready
 Availability:          Drain
 Address:               172.16.0.50
Platform:
 Operating System:      linux
 Architecture:          x86_64
Resources:
 CPUs:                  8
 Memory:                7.796 GiB
Plugins:
  Network:              bridge, host, macvlan, null, overlay
  Volume:               local
Engine Version:         1.13.0
Engine Labels:
 - com.zhongan.bizcluster = anlink
 - com.zhongan.env = prd
 - com.zhongan.publicnetwork = false

* Use swarm mode routing mesh

[root@VM_0_2_centos ~]# docker service create --name nginx --publish 8080:80 --replicas 2 172.16.0.2:5000/env/nginx:latest
j4lvi4dwacrvzjtem1yr98teh

[root@VM_0_2_centos ~]# docker service ps nginx
ID            NAME     IMAGE                             NODE            DESIRED STATE  CURRENT STATE           ERROR  PORTS
e9nh1i688w9v  nginx.1  172.16.0.2:5000/env/nginx:latest  VM_0_50_centos  Running        Running 10 seconds ago         
pryg5wxhz1be  nginx.2  172.16.0.2:5000/env/nginx:latest  VM_0_46_centos  Running        Running 10 seconds ago

[root@VM_0_2_centos ~]# docker service inspect --format="{{json .Endpoint.Spec.Ports}}" nginx
[{"Protocol":"tcp","TargetPort":80,"PublishedPort":8080,"PublishMode":"ingress"}]



* Deploy services to a swarm


```
[root@VM_0_2_centos ~]# docker service create --mode global  --mount type=bind,source=/,destination=/rootfs,ro=1  --mount type=bind,source=/var/run,destination=/var/run  --mount type=bind,source=/sys,destination=/sys,ro=1  --mount type=bind,source=/var/lib/docker/,destination=/var/lib/docker,ro=1  --publish mode=host,target=8080,published=8080  --name=cadvisor   172.16.0.2:5000/google/cadvisor:latest
np08hqygmso4wyyrjed5d1749

[root@VM_0_2_centos ~]# docker service ls
ID            NAME      MODE    REPLICAS  IMAGE
np08hqygmso4  cadvisor  global  4/4       172.16.0.2:5000/google/cadvisor:latest
[root@VM_0_2_centos ~]# docker service ps cadvisor
ID            NAME                                IMAGE                                   NODE            DESIRED STATE  CURRENT STATE           ERROR  PORTS
qqg48dj09s3p  cadvisor.rodcc8yv6fku2hqiaw9sdyhye  172.16.0.2:5000/google/cadvisor:latest  VM_0_46_centos  Running        Running 10 seconds ago         *:8080->8080/tcp
54q9jg8pa02g  cadvisor.56hsb1n1kbgcv3f49e04t55f3  172.16.0.2:5000/google/cadvisor:latest  VM_0_26_centos  Running        Running 10 seconds ago         *:8080->8080/tcp
1g0lj1hmzwi0  cadvisor.yloq2964nwvqnu39jkn6bq579  172.16.0.2:5000/google/cadvisor:latest  VM_0_2_centos   Running        Running 8 seconds ago          *:8080->8080/tcp
0lglggcg2sx5  cadvisor.8s1b3c09ftug83ppqt4ol6wxx  172.16.0.2:5000/google/cadvisor:latest  VM_0_50_centos  Running        Running 10 seconds ago         *:8080->8080/tcp
```

