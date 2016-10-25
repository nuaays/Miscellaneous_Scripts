


## dokcer-py
* https://my.oschina.net/guol/blog/271416
* https://github.com/docker/docker-py



## Docker Remote API
* http://www.open-open.com/lib/view/open1419921028828.html 
* https://my.oschina.net/guol/blog/271416




网上的例子
https://segmentfault.com/a/1190000002711475
	列出所有容器
	GET /containers/json
	创建新容器
	POST /containers/create
	>curl -v -X POST -H "Content-Type: application/json" -d '{}' http://localhost:2375/containers/create?name=xxxxxx
	检查容器
	GET /containers/(id)/json
	获取容器内部进程列表
	GET /containers/(id)/top
	获取容器日志log(stdout & stderr)
	GET /containers/(id)/logs
	导出容器
	GET /containers/(id)/export
	>curl -o testcontainer.tar.gz http://localhost:2375/containers/2sr4r4r32r2rww/export
	启动|停止|重启|杀死容器
	POST /containers/(id)/start|stop|restart?t=5
	POST /containers/(id)/kill
	
	
	
https://segmentfault.com/a/1190000002711508
https://my.oschina.net/guol/blog/271416

	POST /images/create?name=xxx
	POST /images/create?fromImage=base&tag=latest 获取名叫 base 的镜像。
	POST /images/create?fromSrc=url 从 url 导入镜像。
	curl -v -X POST "http://localhost:5555/images/create?fromImage=base&tag=latest"
	推送镜像
	Curl -v -X POST http://localhost:5000/images/xxxx/push
	为镜像打标签
	POST /images/(name)/tag
	查看镜像历史
	GET /images/(name)/history
	构建镜像
	POST /build
	>tar zcf Dockerfile.tar.gz Dockerfile
	>curl -v -X POST -H "Content-Type: application/tar" --data-binary "@Dockerfile.tar.gz" http://localhost:2375/build?t=samplerep
	若是远程Dockerfile URI
	curl -X POST "127.0.0.1:4243/build?t=asd&remote=http%3A%2F%2Flocalhost%2FDockerfile"
	获取容器日志
	>curl -s -XGET 'http://0.0.0.0:2375/containers/3454a5bf585f/logs?stdout=1&follow=1&timestamps=1'
	检查容器的改变
	>curl -s -XGET 'http://0.0.0.0:2375/containers/3454a5bf585f/changes'
	启动容器
	> curl -s -XGET 'http://0.0.0.0:2375/containers/3454a5bf585f/start'
	停止容器
	> curl -s -XGET 'http://0.0.0.0:2375/containers/3454a5bf585f/stop'
	重启容器
	> curl -s -XGET 'http://0.0.0.0:2375/containers/3454a5bf585f/restart'
	杀死容器
	> curl -s -XGET 'http://0.0.0.0:2375/containers/3454a5bf585f/kill?signal=SIGINT'
	附着到容器
	> curl -s -XGET 'http://0.0.0.0:2375/containers/3454a5bf585f/attach?logs=1&stream=1&stdout=1'
	删除一个容器
	>curl -s -XDELETE 'http://0.0.0.0:2375/containers/3454a5bf585f'
	从容器中Copy文件
	POST /containers/${containerid}/copy
	
	列出image
	curl -s -XGET 'http://0.0.0.0:2375/images/json?all=1'
	检查docker daemon是否存活
	curl -s -XGET 'http://0.0.0.0:2375/_ping'
	检查docker版本
	curl -s -XGET 'http://0.0.0.0:2375/version'
	显示系统信息
	curl -s -XGET 'http://0.0.0.0:2375/info'
	检查认证
	POST /auth
	删除本地一个镜像
	curl -s -XDELETE 'http://0.0.0.0:2375/images/telsa:soapui'



## Create and start a container
* curl -v -X POST -H "Content-Type: application/json" -d '{"Image": " registry:2.",}' http://localhost:2376/containers/create?name=registry
* curl -v -X POST -H "Content-Type: application/json" -d '{"PortBindings": { "5000/tcp": [{ "HostPort": "5000" }] },"RestartPolicy": { "Name": "always",},}' http://localhost:2376/containers/registry/start?name=registry
