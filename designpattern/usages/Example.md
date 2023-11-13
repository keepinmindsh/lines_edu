# 실제의 케이스에 대한 정의 

우리는 Java 진영에서 흔히 WebApplication Server/Web Server를 구현해보고자 한다.   
해당 구현에서 같이 논의해볼 내용은 MVC 패턴를 활용한 Servlet과 정적 파일을 로딩할 수 있는 WebServer를 만들어보자. 

## 기본적인 요구 사항 

- http://localhost:8080/hello 호출 시 

Json 모델로 응답을 { "Value" : "Hello World"} 로 반환 할 수 있어야 한다.  
요청시의 Content-Type이 Application/Json이 아닐 경우, 에러를 반환한다.   
에러 반환 내용은 "Content-Type is mismatched" 가 Json으로 표시되어야 한다.   

- http://localhost:8080/hello.do 호출 시 

Chrome 화면 내에 Hello World가 표시되어야 한다. hello.do 호출 시 내부적으로 특정경로의 html을 반환해줘야 한다.   

```
Path : /webapps/view/hello.html   
Content :     
```


- http://localhost:8080/hello.

