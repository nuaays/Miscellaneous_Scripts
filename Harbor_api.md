

## Harbor API
* 200 OK
* 201 Created
* 401 Unauthorized
* 409 Conflict


* Login in
```
curl -v -X POST -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" "http://10.253.20.65/login?principal=admin&password=zcloud"
```

* list current user
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3" http://10.253.20.65/api/users/current
```

* list all users info
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=23a788127c1362630b5de12bfba22b81" http://10.253.20.65/api/users
```

* search a user
```
curl -v -X GET http://10.253.20.65/api/users?username=nihao
```

* add an user
```
curl -v -X POST -H "Content-Type: application/json;charset=utf-8" --data '{"username":"xixihaha1","email":"nuaays1@qq.com","password":"Test1234","realname":"xixihaha1","comment":"For Test1"}' http://10.253.20.65/api/users
curl -v -X POST -H "Content-Type: application/json;charset=utf-8" --data '{"username":"nihao","email":"nuaays2@qq.com","password":"Test1234","realname":"管理员","comment":"管理员你好"}' http://10.253.20.65/api/users
```

* change a user to admin
```
curl -v -X PUT -J "Content-Type: application/json;charset=utf-8" --data '{"has_admin_role":1}' http://10.253.20.65/api/users/5/sysadmin
```



* create a project
```
curl -v -X POST -H "Content-Type: application/json;charset=utf-8" -H "Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3" --data '{"project_name":"myproject11111","public":1}' "http://10.253.20.65/api/projects"
```

* delete project
```
curl -v -X DELETE http://10.253.20.65/api/projects/4
```


* list all projects basic info
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3" "http://10.253.20.65/api/projects?is_public=0"
```

* list all projects basic info which is public
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3" "http://10.253.20.65/api/projects?is_public=1"
```


* list related projects info with project keyword 
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3" "http://10.253.20.65/api/projects?project_name=myproject11111"
```

* list the specified project info with project id
```
curl -v -X GET http://10.253.20.65/api/projects/3/
```


* list project 1's all members
```
curl -v -X GET -H 'Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3' "http://10.253.20.65/api/projects/1/members"
```

* list project 1's members with name keyword
```
curl -v -X GET -H 'Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3' "http://10.253.20.65/api/projects/1/members?username=xixihaha"
```

* add person into some project as member with project id
```
Admin: "roles":[1]
Developer : "roles":[2]
Guest: "roles":[3]
curl -v -X POST -H "Content-Type: application/json;charset=utf-8" -H 'Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3' --data '{"roles":[2],"username":"xixihaha"}' http://10.253.20.65/api/projects/1/members/
```


* delete some member from project with project id
```
curl -v -XDELETE -H "Cookie: language=en-US; beegosessionID=078d5fd527e0961622ccbd483abe7ef3" http://10.253.20.65/api/projects/1/members/5
```



* query project' repos with project id
```
curl -v -X GET "http://10.253.20.65/api/repositories?project_id=1"
```


* query repo
```
http://10.253.20.65/api/repositories/tags?repo_name=library%2Fcentos

```


* query tags of the specified image in some repo
```
curl -v -X GET "http://10.253.20.65/api/repositories/tags?repo_name=library%2Fcentos"

```

* query manifest info of some image with tag
```
curl -v -X GET http://10.253.20.65/api/repositories/manifests?repo_name=test/alpine&tag=3.4&version=v1
curl -v -X GET http://10.253.20.65/api/repositories/manifests?repo_name=test%2Falpine&tag=3.4&version=v1
```


* delete images with tag
```
curl -v -X DELETE http://10.253.20.65/api/repositories?repo_name=test%2Fregistry&tag=2.5.0
```




* system information
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=23a788127c1362630b5de12bfba22b81" http://10.253.20.65/api/systeminfo/volumes

```


* repo top 10
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=23a788127c1362630b5de12bfba22b81" h"ttp://10.53.20.65/api/repositories/top?count=10"
```

* query registry log
```
curl -v -X GET -H "Cookie: language=en-US; beegosessionID=23a788127c1362630b5de12bfba22b81" http://10.253.20.65/api/logs?lines=100
```

* query project log
```
curl -v -X POST -H "Cookie: language=en-US; beegosessionID=23a788127c1362630b5de12bfba22b81" -H "Content-Type: application/json;charset=utf-8" --data '{"begin_timestamp":1480521600,"end_timestamp":1480953599,"keywords":"create/pull/push/delete","project_id":1,"username":""}' "http://10.253.20.65/api/projects/1/logs/filter?page=3&page_size=500"
```

* query statistics
```
curl -v -X POST -H "Cookie: language=en-US; beegosessionID=23a788127c1362630b5de12bfba22b81" http://10.253.20.65/api/statistics
```
