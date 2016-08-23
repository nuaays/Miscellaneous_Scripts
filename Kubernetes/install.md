sudo mkdir /kubernetes && sudo chmod -R 777 /kubernetes 


>:/boot# cat /proc/cmdline  
BOOT_IMAGE=/boot/vmlinuz-3.13.0-65-generic root=UUID=c70a26ec-1dda-455c-acfd-792015b2bb6f ro console=tty1 console=ttyS0 

echo "CONFIG_MEMCG_SWAP_ENABLED=y" >> /boot/config-3.13.0-65-generic 
echo 'GRUB_CMDLINE_LINUX="cgroup_enable=memory swapaccount=1"' >> /etc/default/grub 

sudo update-grub





#install etcd
```
curl -L https://github.com/coreos/etcd/releases/download/v3.0.6/etcd-v3.0.6-linux-amd64.tar.gz -o etcd-v3.0.6-linux-amd64.tar.gz  
tar -zxvf etcd-v3.0.6-linux-amd64.tar.gz 
cd etcd-v3.0.6-linux-amd64 && cp etcd* /usr/local/bin 
```

#start etcd
```
sudo docker run --net=host -d gcr.io/google_containers/etcd:2.0.12 /usr/local/bin/etcd --addr=127.0.0.1:4001 --bind-addr=0.0.0.0:4001 --data-dir=/var/etcd/data 
```

#start kubernetes master
###kubernetes master is started using the kubelet , which also starts the other master components:
# apiserver, scheduler, controller, and pause , 
```
sudo docker run --volume=/:/rootfs:ro --volume=/sys:/sys:ro --volume=/dev:/dev --volume=/var/lib/docker/:/var/lib/docker:ro --volume=/var/lib/kubelet/:/var/lib/kubelet:rw  --volume=/var/run:/var/run:rw --net=host --pid=host --privileged=true -d  gcr.io/google_containers/hyperkube:v1.0.1 /hyperkube kubelet --containerized --hostname-override="127.0.0.1" --address="0.0.0.0" --api-servers=http://localhost:8080 --config=/etc/kubernetes/manifests  

####Apiserver
The apiserver takes API requests, processes them, and stores the result in etcd if required and returns the result
####Scheduler
The scheduler monitors the API for unscheduled pods and schedules them on a node to run and also notifies the about the same to the API
####Controller
The controller manages the replication level of the pods, starting new pods in a scale up event and stopping some of the pods in a scale down
####Pause
The pause keeps the port mappings of all the containers in the pod or the network endpoint of the pod

```

#start Service Porxy
```
sudo docker run -d --net=host --privileged gcr.io/google_containers/hyperkube:v1.0.1 /hyperkube proxy -- master=http://127.0.0.1:8080 --v=2

```






