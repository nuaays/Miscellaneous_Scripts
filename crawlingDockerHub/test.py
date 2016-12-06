#!/usr/bin/python
# -*- coding: utf-8 -*-
__author__ = 'yangsen'

import os, sys
reload(sys)
sys.setdefaultencoding('utf8')
import re, random, pprint
import json, requests, urllib2
from bs4 import BeautifulSoup



USER_AGENTS = (
    "Mozilla/5.0 (Windows NT 6.3; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0",
    "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_7_0; en-US) AppleWebKit/534.21 (KHTML, like Gecko) Chrome/11.0.678.0 Safari/534.21",
    "Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
    "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:0.9.2) Gecko/20020508 Netscape6/6.1",
    "Mozilla/5.0 (X11;U; Linux i686; en-GB; rv:1.9.1) Gecko/20090624 Ubuntu/9.04 (jaunty) Firefox/3.5",
    'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:22.0) Gecko/20100101 Firefox/22.0',
    "Opera/9.80 (X11; U; Linux i686; en-US; rv:1.9.2.3) Presto/2.2.15 Version/10.10"
)


def getURLContent(url='https://hub.docker.com/explore/?page=1'):
    headers = {
        'Content-Type':'application/x-www-form-urlencoded',
        'Accept-Language': 'zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3',
        'Accept-Encoding': 'gzip, deflate',
        'Accept' : 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
        'Connection' : 'keep-alive',
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
        r = s.get(url, params=None, headers=headers, timeout=30)
        r.raise_for_status()
    except requests.RequestException as e:
        print e
        return None
    else:
        #print r.encoding
        return r.content




def getPage(page=1):
    html_doc = getURLContent("https://hub.docker.com/explore/?page={}".format(page))
    repos_info = html_doc.split('"DashboardReposStore":{"repos":')[1].split(',"count":')[0]
    #print repos_info
    repos = json.loads(repos_info)
    for repo in repos:
        repo["pic"] = "https://hub.docker.com/public/images/official/{}.png".format(repo["name"])
        repo["url"] = "https://hub.docker.com/_/{}/".format(repo["name"])
        repo["tags"] = getRepoSupportedTags(repoName=repo["name"])
        repo["tagsDetail"] = getRepoDetailedInfo(repoName=repo["name"], tagslist=repo["tags"])

        pprint.pprint(repo)



def getRepoTags(repoName="rethinkdb"):
    html_doc = getURLContent("https://hub.docker.com/r/library/{}/tags/".format(repoName))
    tags_info = "{" + html_doc.split('"is_private":false,"repository_type":null,')[1].split(';</script>')[0].replace('},"scans":{},"status":{}','')
    # pprint.pprint(tags_info)
    tags = json.loads(tags_info)
    # print tags["description"]
    # print tags["pull_count"]
    # print tags["last_updated"]
    print tags["tags"]["library"][repoName]["result"]
    # for tag in tags["tags"]["library"][repoName]["result"]:
    #     print tag, tags["tags"]["library"][repoName]["tags"][tag]["latest_scan_status"]

    # for tag in tags["tags"]["library"][repoName]["tags"]:
    #     print tag, tags["tags"]["library"][repoName]["tags"][tag]
    #
    # print "="*20
    # for tag in sorted(tags["tags"]["library"][repoName]["tags"].items(), lambda x, y: cmp(x[1]["last_updated"], y[1]["last_updated"]), reverse=True):
    #     print tag[0], tag[1]["full_size"]/1000.0/1000.0, tag[1]
    print [ tag for tag in sorted(tags["tags"]["library"][repoName]["tags"].items(), lambda x, y: cmp(x[1]["last_updated"], y[1]["last_updated"]), reverse=True) ]
    #print tags["tags"]["library"][repoName]["tags"]["2.3.5"]

    #every tag page
    #https://hub.docker.com/r/library/rethinkdb/tags/2.3.4/



def getRepoDetailedInfo(repoName="rethinkdb", tagslist=[]):
    detailedInfo = {}
    html_doc = getURLContent("https://hub.docker.com/r/library/{}/tags/".format(repoName))
    tags_info = "{" + html_doc.split('"is_private":false,"repository_type":null,')[1].split(';</script>')[0].replace('},"scans":{},"status":{}','')
    tags = json.loads(tags_info)
    for tag in tagslist:
        detailedInfo[tag] = tags["tags"]["library"][repoName]["tags"][tag]
    return detailedInfo

def getRepoSupportedTags(repoName="node"):
    #https://hub.docker.com/r/library/mongo/
    tags_supported = []
    html_doc = getURLContent("https://hub.docker.com/_/{}/".format(repoName))
    soup = BeautifulSoup(html_doc, "html.parser")
    for i in soup.select("li > a"):
        if 'Dockerfile' in str(i) and "windows" not in str(i):
            tags_supported.extend(i.get_text().split(" (")[0].split(", "))
    return tags_supported


if __name__ == "__main__":
    #getRepoTags(repoName="rethinkdb")
    #getRepoTags(repoName="node")
    #getRepoSupportedTags(repoName="node")

    getPage(page=1)





    """
    "DashboardReposStore":{"repos":[{"user":"library","name":"nginx","namespace":"library","repository_type":null,"status":1,"description":"Official build of Nginx.","is_private":false,"is_automated":false,"can_edit":false,"star_count":4705,"pull_count":535908586,"last_updated":"2016-11-28T18:20:22.673554Z"},{"user":"library","name":"redis","namespace":"library","repository_type":null,"status":1,"description":"Redis is an open source key-value store that functions as a data structure server.","is_private":false,"is_automated":false,"can_edit":false,"star_count":3027,"pull_count":168231459,"last_updated":"2016-11-09T01:08:09.206096Z"},{"user":"library","name":"busybox","namespace":"library","repository_type":null,"status":1,"description":"Busybox base image.","is_private":false,"is_automated":false,"can_edit":false,"star_count":862,"pull_count":160787750,"last_updated":"2016-10-07T21:07:05.797711Z"},{"user":"library","name":"ubuntu","namespace":"library","repository_type":null,"status":1,"description":"Ubuntu is a Debian-based Linux operating system based on free software.","is_private":false,"is_automated":false,"can_edit":false,"star_count":5117,"pull_count":141788454,"last_updated":"2016-11-16T21:04:35.301712Z"},{"user":"library","name":"registry","namespace":"library","repository_type":null,"status":1,"description":"Containerized docker registry","is_private":false,"is_automated":false,"can_edit":false,"star_count":1194,"pull_count":58620290,"last_updated":"2016-11-14T17:36:51.825175Z"},{"user":"library","name":"alpine","namespace":"library","repository_type":null,"status":1,"description":"A minimal Docker image based on Alpine Linux with a complete package index and only 5 MB in size!","is_private":false,"is_automated":false,"can_edit":false,"star_count":1619,"pull_count":53883943,"last_updated":"2016-10-18T20:35:09.725485Z"},{"user":"library","name":"mongo","namespace":"library","repository_type":null,"status":1,"description":"MongoDB document databases provide high availability and easy scalability.","is_private":false,"is_automated":false,"can_edit":false,"star_count":2623,"pull_count":47437759,"last_updated":"2016-11-23T23:27:16.874125Z"},{"user":"library","name":"mysql","namespace":"library","repository_type":null,"status":1,"description":"MySQL is a widely used, open-source relational database management system (RDBMS).","is_private":false,"is_automated":false,"can_edit":false,"star_count":3467,"pull_count":44655463,"last_updated":"2016-11-23T23:33:28.767879Z"},{"user":"library","name":"swarm","namespace":"library","repository_type":null,"status":1,"description":"Swarm: a Docker-native clustering system.","is_private":false,"is_automated":false,"can_edit":false,"star_count":550,"pull_count":41846698,"last_updated":"2016-08-19T16:36:40.582717Z"},{"user":"library","name":"hello-world","namespace":"library","repository_type":null,"status":1,"description":"Hello World! (an example of minimal Dockerization)","is_private":false,"is_automated":false,"can_edit":false,"star_count":207,"pull_count":33642651,"last_updated":"2016-11-21T20:24:32.669143Z"},{"user":"library","name":"elasticsearch","namespace":"library","repository_type":null,"status":1,"description":"Elasticsearch is a powerful open source search and analytics engine that makes data easy to explore.","is_private":false,"is_automated":false,"can_edit":false,"star_count":1804,"pull_count":29317070,"last_updated":"2016-11-23T23:54:34.043345Z"},{"user":"library","name":"postgres","namespace":"library","repository_type":null,"status":1,"description":"The PostgreSQL object-relational database system provides reliability and data integrity.","is_private":false,"is_automated":false,"can_edit":false,"star_count":2899,"pull_count":26849935,"last_updated":"2016-11-23T23:43:38.903881Z"},{"user":"library","name":"node","namespace":"library","repository_type":null,"status":1,"description":"Node.js is a JavaScript-based platform for server-side and networking applications.","is_private":false,"is_automated":false,"can_edit":false,"star_count":3165,"pull_count":25622164,"last_updated":"2016-11-28T22:32:28.977742Z"},{"user":"library","name":"httpd","namespace":"library","repository_type":null,"status":1,"description":"The Apache HTTP Server Project","is_private":false,"is_automated":false,"can_edit":false,"star_count":800,"pull_count":23790738,"last_updated":"2016-11-29T00:37:23.415411Z"},{"user":"library","name":"wordpress","namespace":"library","repository_type":null,"status":1,"description":"The WordPress rich content management system can utilize plugins, widgets, and themes.","is_private":false,"is_automated":false,"can_edit":false,"star_count":1401,"pull_count":18199335,"last_updated":"2016-11-18T01:05:48.251430Z"}],"count":125
    """