# UML 

- [UML](https://keepinmindsh.github.io/1lines/document/architect_051.html)

## 항목 정의 

- 접근 제어자 

|항목|기호|
|---|---|
|public|+|
|private|-|
|protected|#|
|package|~|

- 클래스 표현 

|항목|설명|
|---|---|
|가장 윗 부분|클래스 이름|
|중간 부분|속성 |
|마지막 부분|연산|

## UML 관계 정의 

- 연관관계 ( Association )

![Association](https://github.com/keepinmindsh/lines_edu/blob/main/assets/association.jpg)

클래스 들이 개념 상 서로 연결되었음을 나타낸다. 보통은 한 클래스가 다른 클래스에서 제공하는 기능을 사용하는 상황일 때 표시한다.

- 일반화 관계 ( Generalization )

![Generalization](https://github.com/keepinmindsh/lines_edu/blob/main/assets/generalization.png)

객체지향 개념에서 상속 관계라고 한다. 한 클래스가 다른 클래스를 포함하는 상위 개념일 때 이를 IS-A 관계라고 하며 UML에서는 일반화 관계를 모델링한다.

- 집약 관계 ( Aggregation )

![Aggregation](https://github.com/keepinmindsh/lines_edu/blob/main/assets/aggregation.png)

클래스들의 사이의 전체 또는 부분 같은 관계를 나타낸다. 전체 객체의 라이프 타입과 부분 객체의 라이프 타임은 독립적이다. 즉, 전체 객체가 없어지면 부분 객체도 없어진다.

- 합성 관계 ( Composition )

![Composition](https://github.com/keepinmindsh/lines_edu/blob/main/assets/composition.png)

클래스 들 사이의 전체 또는 부분 같은 관계를 나타낸다. 전체 객체의 라이프 타입과 부분 객체의 라이프 타임은 의존적이다. 즉, 전체 객체가 없어지면 부분 객체도 없어진다.

- 의존 관게 ( Dependency )

![dependency](https://github.com/keepinmindsh/lines_edu/blob/main/assets/dependency.png)

연관 관계와 같이 한 클래스가 다른 클래스에서 제공하는 기능을 사용할 때를 나타낸다. 차이점은 두 클래스의 관계가 한 메서드를 실행하는 동안 같은, 매우 짧은 시간만 유지 된다는 점이다.

- 실체화 관계 ( Realization )

![realization](https://github.com/keepinmindsh/lines_edu/blob/main/assets/realization.png)

책임들의 집합인 인터페이스와 이 책임들을 실제로 실현한 클래스들 사이의 관계를 나타낸다.