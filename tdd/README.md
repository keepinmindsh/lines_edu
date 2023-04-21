# Good Code ( feat. Clean Code ) 

- 깨끗한 코드는 ‘보기에 즐거운’ 코드다.
  - 깨끗한 코드는 한가지에 집중한다.
- 코드는 추측이 아니라 사실에 기반해야한다. 
  - 반드시 필요한 내용만 담아야 한다.
- 깨끗한 코드란 다른 사람이 고치기 쉽다고 단언한다.
  - 테스트케이스가 없는 코드는 깨끗한 코드가 아니다.
- 깨끗한 코드는 주의 깊게 작성한 코드다. 
  - 누군가 시간을 들여 깔끔하고 단정하게 정리한 코드다. 
  - 세세한 사항까지 꼼꼼하게 신경쓴 코드다.
- 모든 테스트를 통과한다. 중복이 없다. 
  - 시스템 내 모든 설계 아이디어를 표현한다. 
  - 클래스, 메서드, 함수등을 최대한 줄인다. 
  - 중복 줄이기, 표현력 높이기, 초반부터 간단한 추상화 고려하기.
- 깨끗한 코드는 읽으면서 놀랄 일이 없어야 한다. 
  - 코드를 독해하느라 머리를 쥐어짤 필요가 없어야 한다. 
  - 읽으면서 짐작한대로 돌아가는 코드가 깨끗한 코드다.


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


