# 실제의 케이스에 대한 정의 

우리는 Java 진영에서 흔히 WebApplication Server/Web Server를 구현해보고자 한다.   
해당 구현에서 같이 논의해볼 내용은 MVC 패턴를 활용한 Servlet과 정적 파일을 로딩할 수 있는 WebServer를 만들어보자. 

## 기본적인 요구 사항 

- http://localhost:8080/hello 호출 시

```
HTTP Method : POST   
Request Body   
{  
    "value" : "hello world"  
}
```

Json 모델로 응답을 { "Value" : "Hello World"} 로 반환 할 수 있어야 한다.  
요청시의 Content-Type이 Application/Json이 아닐 경우, 에러를 반환한다.   
에러 반환 내용은 "Content-Type is mismatched" 가 Json으로 표시되어야 한다.   

- http://localhost:8080/hello.do 호출 시 

Chrome 화면 내에 Hello World가 표시되어야 한다. hello.do 호출 시 내부적으로 특정경로의 html을 반환해줘야 한다.   

```
Path : /webapps/view/hello.html   
Content :     

<html>
    <head>
    
    </head>
    <body>
        Hello World! 
    </body>
</html>

```

- http://localhost:8080/hello.png 호출시 

hello.png 파일이 다운로드 되어야 한다. 

## 향후 우리가 만들어 나갈 구조 

![Spring Dispatcher Servlet](https://github.com/keepinmindsh/lines_edu/blob/main/assets/spring-dispatcher-servlet.png)

**Servlet** : 클라이언트의 요청을 처리하고, 그 결과를 반환하는 Servlet 클래스의 구현 규칙을 지킨 자바 웹 프로그래밍 기술

## 요구사항 확장 ( 1차 )

- HTTP Method 인 Get/Post/Put/Delete 를 받을 수 있도록 구현한다.
  - Server가 기동될 때 URL과 실행될 메소드를 등록할 수 있는 Route Handler를 구현한다. 
- HTTP Request Header 를 분석할 수 있는 Parser 를 구현한다. 
- Content-Type 에 따라서 요청의 파라미터 를 Parsing 하는 로직을 구현한다. 