#Installing  Kubernetes 


###Kubernetes is an open source container cluster manager. The main components of Kubernetes are the 
following:
 1 . etcd 
 2 . Kubernetes  master 
 3 . Service  proxy 
 4 . kubelet 

etcd is a simple, secure, fast and reliable distributed key-value store.   

Kubernetes master exposes the Kubernetes API using which containers are run on nodes to handle tasks.  

kubelet is an agent that runs on each node to monitor the containers running on the node, restarting 
them if required to keep the replication level.  

A service proxy runs on each node to provide the Kubernetes service interface for clients. A service is an 
abstraction for the logical set of pods represented by the service, and a service selector is used to select the 
pods represented by the service. The service proxy routes the client traffic to a matching pod. Labels are used 
to match a service with a pod.

