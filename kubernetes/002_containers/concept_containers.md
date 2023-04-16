# Container의 이해 / Kubernetes의 적용 

### Container 구성 

```script 
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


> [Minikube Start](https://minikube.sigs.k8s.io/docs/start/)   
> [Kubectl Context Switching](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/)
