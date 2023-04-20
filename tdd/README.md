
# Tdd Basic 

### [Why TDD](https://github.com/keepinmindsh/lines_edu/blob/main/tdd/basic/README.md)

# 요구사항 

- java/golang의 순수 함수를 제외하고 어떤 라이브러리도 사용해서는 안된다. ( gin, mux 등등 사용 불가 ) 
- http://localhost:8080/index.html 로 접속했을 때 webapp 디렉토리의 index.html 파일을 읽어 클라이언트에 응답한다. 
  - [index.html](https://github.com/keepinmindsh/lines_edu/blob/main/tdd/codesample/index.html) 


## HTTP Header 

GET /index.html  HTTP/1.1  
Host: localhost:8080  
Connection: keep-alive   
Accept: */*     


