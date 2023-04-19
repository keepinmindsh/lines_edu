
# 요구사항 

- golang의 순수 함수를 제외하고 어떤 라이브러리도 사용해서는 안된다. ( gin, mux 등등 사용 불가 ) 
- http://localhost:8080/index.html 로 접속했을 때 webapp 디렉토리의 index.html 파일을 읽어 클라이언트에 응답한다. 
  - Web Browser를 통해서 해당 URL로 접속해서 아래의 이미지를 표시해보자. 


## HTTP Header 

GET /index.html  HTTP/1.1  
Host: localhost:8080  
Connection: keep-alive   
Accept: */*     


