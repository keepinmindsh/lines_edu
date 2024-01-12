# Template Method 

## Basic Information

- 패턴형식 : 행위패턴 

## Gof's Description

객체의 연산이 알고리즘의 뼈대 만들 정의하고 각 단계에서 수행할 구체적 처리는 서브 클래스 쪽에서 미룹니다.   
알고리즘의 구조 자체는 그대로 놔둔 채 알고리즘 각 단계의 처리를 서브 클래스에서 처리할 수 있게 합니다.

## 동기

Application 클래스와 Draw 클래스를 제공하는 응용 프로그램 프레임워크를 생각해 봅시다.   
Application 클래스는 파일이나 특정한 외부 형식으로 저장된 문서를 열 수 있고,   
Document 객체를 파일에서 읽은 문서에 정보를 나타냅니다.   
특정한 요구에 따라서 프레임워크 정의한 Application 클래스와 Document 클래스를 상속한 서브 클래스를 정의하여 새로운 응용 프로그램을 구축할 수 있을 것 입니다.     
예를 들어, 그림 그리기 응용 프로그램은 DrawApplication 클래스와 DrawDocument 클래스를 정의할 수 있을 것이고,      
스프레드 시트 응용프로그램은 SpreadsheetApplication 클래스와 SpreadsheetDocument 클래스를 정의할 수 있습니다.    
이와 같은 패턴을 구현할 때 사용하는 것이 탬플릿 메서드라고 합니다.   

