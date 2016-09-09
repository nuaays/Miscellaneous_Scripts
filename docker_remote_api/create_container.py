#!/usr/bin/python
# -*- coding: utf-8 -*-
import os, sys
reload(sys)
sys.setdefaultencoding('utf8')
import time, random, pprint
import json, requests, urllib2




USER_AGENTS = (
	"Mozilla/5.0 (Windows NT 6.3; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0",
    "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_7_0; en-US) AppleWebKit/534.21 (KHTML, like Gecko) Chrome/11.0.678.0 Safari/534.21",
    "Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
    "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:0.9.2) Gecko/20020508 Netscape6/6.1",
    "Mozilla/5.0 (X11;U; Linux i686; en-GB; rv:1.9.1) Gecko/20090624 Ubuntu/9.04 (jaunty) Firefox/3.5",
    'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:22.0) Gecko/20100101 Firefox/22.0',
    "Opera/9.80 (X11; U; Linux i686; en-US; rv:1.9.2.3) Presto/2.2.15 Version/10.10"
)


def getURLContent(url='http://nsso.zhonganonline.com/login/'):
    headers = {
               'Content-Type':'application/x-www-form-urlencoded',
               'Accept-Language': 'zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3',
               'Accept-Encoding': 'gzip, deflate', 
               'Accept' : 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
               'Connection' : 'keep-alive',
               #'Referer' : 'http://nsso.zhonganonline.com/login?service=keystone-manager&renew=false&target=http://keystone.zhonganonline.com/v1/dashboard'        
               }

    headers['User-Agent'] = USER_AGENTS[random.randint(0, len(USER_AGENTS)-1)]
    
    # data = {
    #          "username":"yangsen",
    #          "password":"asb#1234",
    #          "service":"keystone-manager",
    #          "target":"http://keystone.zhonganonline.com/v1/dashboard"
    #        }

    try:
        #r = requests.get(url, params={'ip': '8.8.8.8'}, headers=headers, timeout=1)
        s = requests.session()
        s.headers=headers
        
        #r = s.post("http://nsso.zhonganonline.com/login/", data=data)
        #r = s.get("http://keystone.zhonganonline.com/v1/dashboard", params=data, headers=headers, timeout=1)
        r = s.get(url, params=None, headers=headers, timeout=1)
        #print "XXXXX",r
        r.raise_for_status()
    except requests.RequestException as e:
        print e
        return None
    else:
        #print r.encoding
        return r.content


def post(url, datas=None):
    #payload = {'a':'杨','b':'hello'}
    # r = requests.post('http://httpbin.org/post', data=json.dumps(payload))
    #headers = {'content-type': 'application/json'}
    #r = requests.post(url, data=json.dumps(payload), headers=headers)
    
    response = requests.post(url, data=datas)  
    json = response.json()  
    return json  

#get images
DOCKER_IMAGES_RESTAPI="http://%s:2375/images/json"

#get containers
DOCKER_CONTAINERS_RESTAPI="http://%s:2375/containers/json?all=1"
#

def get_docker_images(host):
    url = DOCKER_IMAGES_RESTAPI % host
    #print getURLContent(url)
    res_json = json.loads(getURLContent(url))
    pprint.pprint(res_json)  

def get_docker_containers(host):
    url = DOCKER_CONTAINERS_RESTAPI % host
    #print getURLContent(url)
    res_json = json.loads(getURLContent(url))
    pprint.pprint(res_json)  


"""
tomcat:latest





##create
curl -X POST -H "Content-Type: application/json" -d '{
  "Hostname":"", "User":"", "Memory":0, "MemorySwap":0,
  "AttachStdin":false, "AttachStdout":true, "AttachStderr":true,
  "PortSpecs":null, "Tty":false, "OpenStdin":false,
  "StdinOnce":false, "Env":null, "Cmd":["date"],
  "Image":"ubuntu", "Tag":"latest", "Volumes":{"/tmp":{} },
  "WorkingDor":"", "DisableNetwork":false, "ExposedPorts":{"22/tcp":{} }
}'


curl -X POST -H "Content-Type: application/json;charset=utf-8" -d '{  "Hostname":"", "User":"", "Memory":0, "MemorySwap":0,  "AttachStdin":false, "AttachStdout":true, "AttachStderr":true,  "PortSpecs":null, "Tty":false, "OpenStdin":false,  "StdinOnce":false, "Env":null, "Cmd":["date"],  "Image":"ubuntu", "Tag":"latest", "Volumes":{"/tmp":{} },  "WorkingDor":"", "DisableNetwork":false, "ExposedPorts":{"22/tcp":{} }    }'
http://10.253.6.128:2375/containers/create

{"Id":"0a4956fb165f1c2d52520d3d210ad840f0e368ad704d296e9d2084782f829d73","Warnings":null}



curl -X POST -H "Content-Type: application/json;charset=utf-8" -d '{  "Hostname":"", "User":"", "Memory":0, "MemorySwap":0,  "AttachStdin":false, "AttachStdout":true, "AttachStderr":true,  "PortSpecs":null, "Tty":false, "OpenStdin":false,  "StdinOnce":false, "Env":null, "Cmd":["date"],  "Image":"ubuntu", "Tag":"latest", "Volumes":{"/tmp":{} }, "WorkingDor":"", "DisableNetwork":false, "ExposedPorts":{"22/tcp":{"HostPort":"11022"} } , "Binds":["/tmp:/tmp"],  "PortBindings":{ "22/tcp":[{"HostPort":"11022"}] },  "PublishAllPorts":false,  "Privileged":false }'









##start
curl -v -X POST "Content-Type: application/json" -d '{
  "Binds":["/tmp:/tmp"],
  "PortBindings":{ "22/tcp":[{"HostPort":"11022"}] },
  "PublishAllPorts":false,
  "Privileged":false
}'

curl -v -X POST "Content-Type: application/json" -d '{ "Binds":["/tmp:/tmp"],  "PortBindings":{ "22/tcp":[{"HostPort":"11022"}] },  "PublishAllPorts":false,  "Privileged":false }'
http://10.253.6.128:2375/containers/0a4956fb165f1c2d52520d3d210ad840f0e368ad704d296e9d2084782f829d73/start





##########
yangsen
##########
 curl -X POST -H "Content-Type: application/json;charset=utf-8" -d '{  "Hostname":"", "User":"", "Memory":0, "MemorySwap":0,  "AttachStdin":false, "AttachStdout":true, "AttachStderr":true,  "PortSpecs":null, "Tty":false, "OpenStdin":false,  "StdinOnce":false, "Env":null, "Cmd":["date"],  "Image":"ubuntu", "Tag":"latest",  "WorkingDor":"", "DisableNetwork":false, "ExposedPorts":{"22/tcp":{"HostPort":"11022"} } , "Binds":["/tmp:/tmp"],  "PortBindings":{ "22/tcp":[{"HostPort":"11022"}] },  "PublishAllPorts":false,  "Privileged":false }' http://10.253.6.128:2375/containers/create?name=senyang









"""



if __name__ == "__main__":
    #url="http://keystone.zhonganonline.com/v1/host/list/test"
    #print getURLContent(url="http://keystone.zhonganonline.com/v1/dashboard")
    #print getURLContent(url)
    #get_docker_images("10.253.6.128")
    get_docker_containers("10.253.6.128")
