# Container의 이해 / Kubernetes의 적용 

### Container 구성 

#### NodeJS 서비스 생성 

- app.js로 아래 파일로 생성하기 

```javascript 
const http = require('http')
const os = require('os')

console.log("Kubia server starting...")

const handler =function(request, response) {
    console.log("Received request from " + request.connection.remoteAddress);
    response.writeHead(200)
    response.end("You've hit " + os.hostname() + "\n")
};

const www = http.createServer(handler);
www.listen(8080)
```

#### Dockerfile 생성하기 

- Dockerfile 생성하기 

```Dockerfile 
FROM node:7
ADD app.js/app.js

ENTRYPOINT["node", "app.js"]
```

#### Docker의 image를 build 하기 

- Shell에서 Dockerfile을 Image 로 빌드하기 

```shell 
$ docker build . -t sample_apps -f ./Dockerfile
```

- Shell에서 Image를 Container로 기동하기 

```shell 
$ docker run --name sample_apps_container -p 8080:8080 -d sample_apps
```

- Container 내부 정보 자세히 살피기 

```shell 
$ docker inspect sample_apps_container
```


> [Minikube Start](https://minikube.sigs.k8s.io/docs/start/)   
> [Kubectl Context Switching](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/)
