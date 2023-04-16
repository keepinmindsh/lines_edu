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

```Dockerfile 
FROM node:7
ADD app.js/app.js

ENTRYPOINT["node", "app.js"]
```




> [Minikube Start](https://minikube.sigs.k8s.io/docs/start/)   
> [Kubectl Context Switching](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/)
