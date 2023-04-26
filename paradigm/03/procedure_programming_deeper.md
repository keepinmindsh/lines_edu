# More Deeper 

절차적 프로그래밍은 논리적으로 작성된 코드가 순차적으로 처리되는 것이 중요하다.  
절차적인 프로그래밍은 하나의 함수내에서 정의된 논리적인 흐름이다. 

## 장점 

- 컴퓨터의 처리 구조와 유사 -> 빠른 실행 속도 
- 같은 코드를 복사하지 않고 다른 위치에서 호출하여 사용 가능 
- 모듈 구성이 용이하며, 구조적인 프로그래밍이 가능 

## 단점 

- 프로그램 분석의 어려움 
- 유지보수의 어려움 
- 코드의 수정이 어려움 

# 절차지향 vs 객체지향 

예시 : 서울에서 부산으로 가는 방법에 대해서 정의하라!   

서울에서 부산으로 가는 방법이  
- 서울 고속 버스 터미널로 간다. 
- 버스를 탄다
- 부산에 도착한다. 
라는 하나의 방법에 대해서 

### 절차적인 방식으로 작성되었을 경우 

아래와 같이 함수 내에서 순차적으로 코드가 작성되고 호출된다. 

```go 
package main 

func GoFromSeoulToBusanByBusTerminal(){
	goToBusTerminal()

	getOnTheBus()

	arriveToBusan()
}

func goToBusTerminal(){
	
}

func getOnTheBus(){
	
}

func arriveToBusan(){
	
}

func main(){
	GoFromSeoulToBusanByBusTerminal()
}

```

### 객체지향적인 방식으로 작성되었을 경우

```go
package main

type TheWayToGoBuSan interface {
	UseTransportationInSeoul()
	Moving()
	ArriveToDestination()
}

type UseBus struct{}

func NewTheWayBus() TheWayToGoBuSan {
	return &UseBus{}
}

func (ub *UseBus) UseTransportationInSeoul() {

}

func (ub *UseBus) Moving() {

}

func (ub *UseBus) ArriveToDestination() {

}

func main() {
	howToGoBusan := NewTheWayBus()

	howToGoBusan.UseTransportationInSeoul()

	howToGoBusan.Moving()

	howToGoBusan.ArriveToDestination()
}
```

### 만약 서울 어디에서 출발해야하는지를 정한다면? 
### 만약 고속버스가 아닌 비행기를 이용해야 한다면? 
### 만약 걸어가야 한다면? 

