# Pattern - Abstract Factory 

## Basic Information 

- 패턴형식 : 생성패턴 

## Gof's Description 

### 의도 

구체적인 클래스를 지정하지 않고 관련성을 갖는 객체의 집합을 생성하거나 서로 독립적인 객체들의 집합을 생성할 수 있는 인터페이스를 제공하는 패턴이다.

### 풀이  

**추상화된 인터페이스 정의를 통해서 객체를 생성하는 방식**  

여기에서 중요한 부분은 **추상화**된 객체를 정의한다는 것은 처리 절차, 객체의 유형이 유사한 업무, 개념이 존재한다는 것을 의미 한다.  
예를 들어, 예약을 할 경우, 예약 절차는 모두 동일하지만, 온라인 에약, 오프라인 예약, 전화 예약 등이 있다고 할 때 각 예약의 추상화된 인터페이스를 정의하여 
Abstract Factory 패턴을 이용해서 인터페이스를 통한 객체를 반환한다는 것을 의미 한다.   

### 장점 

- 객체가 생성되거나 구성,표현되는 방식과 무관하게 시스템을 독립적으로 만들고자 할 때 이는 Abstract Factory가 생성을 담당하는데, 실제 생성되는 객체에 대한 구현 및 표현은 Class로 분리되어 관리하기 때문이다.
- 여러 제품군 중 하나를 선택해서 시스템을 설정해야하고 한번 구성한 제품을 다른 것으로 대체 할 수 있을 때 Application이 요청에 따라 필요한 객체를 반환하는데, 해당 - 객체는 하나의 인터페이스에 의해 추상화된 제품의 구현을 이야기한다.
- 관련된 제품 객체들이 함께 사용되도록 설계되었고, 이 부분에 대한 제약이 외부에도 지켜지도록 하고 싶을 때
- 제품에 대한 클래스 라이브러리를 제공하고, 그들의 구현이 아닌 인터페이스를 노출시키고 싶을 때

### 단점 

- 새로운 종류의 제품을 제공하기 어려움 
- 새로운 종류의 제품을 만들기 위해 기존 추상 팩토리를 확장하기가 쉽지 않음, 생성되는 제품은 추상 팩토리가 생성할 수 있는 제품 집합에만 고정되어 있기 때문이다.

## 실생활의 사례를 통한 패턴의 이해 

### 자동차의 예시 

자동차에 달려있는 타이어는 모두 동일하게 앞,뒤로 구르는 역할을 한다. 

```go 
type Tire interface {
	Forward()
	Backward()
	Stop()
	Start()
}
```


```go 
package tire

import (
	"design_pattern/oop/app/car/service/moving"
	tierUcase "design_pattern/oop/app/car/usecase/tire"
	domainMaps "design_pattern/oop/domain/maps"
	steeringDomain "design_pattern/oop/domain/steering"

	"design_pattern/oop/domain/tire"
)

type TireName string

const (
	KUMHO TireName = "Kumho"
	NEXEN TireName = "Nexen"
)

// Abstract Factory 패턴이 적용된 함수 
func NewTire(tireName TireName, moving *moving.Moving, steering *steeringDomain.Steering, mapValidator domainMaps.MapValidate) tire.Tire {
	switch tireName {
	case KUMHO:
		return tierUcase.NewKumhoTire(moving, steering, mapValidator)
	case NEXEN:
		return tierUcase.NewNexenTire(moving, steering, mapValidator)
	}

	return nil
}

```

```go 
package main 


func main(){
	tire := tire.NewTire(tire.NEXEN, moving, &steering, validater)
}

```

- tire.Tire 인터페이스를 통해서 금호타이어, 넥센 타이어를 받을 수 있다. ( PolyMorphism )
- 인터페이스를 Client로 반환하기 때문에 내부 객체를 캡슐화할 수 있다. ( Capsulation )
- 인터페이스를 반환하기 때문에 Client에서 **사용** 만 가능하며, **변경**할 수 없는 구조로 제약을 적용할 수 있다. ( Open & Close )
- 객체의 생성 방식에 대해서는 main(client)는 신경쓰지 않는다. 오로지 사용할 뿐이다. ( Loose Coupling )


### Java 에서는 


```java 
package DesignPattern.gof_abstractFactory;
  
public interface Soldier {
	String getSoldier();

	String attack();
}  
```

```java 
package DesignPattern.gof_abstractFactory;
  
public class Marine implements Soldier {

	public String getSoldier() {
		return null;
	}

	public String attack() {
		return null;
	}
}
```

```java
package DesignPattern.gof_abstractFactory;
  
public interface TrainingFactory {

	/**
	* @param soldierType
	* @return
	*/
	Soldier create(String soldierType);
} 
```

```java
package DesignPattern.gof_abstractFactory;

public class infantryTraingCenter implements TrainingFactory {
	// infantryTraingCenter는 TrainingFactory의 구현을 담당하면서, 추상 팩토리에서 객체를 생성하는 역할을 맞는다..
	public Soldier create(String soldierType) {
		switch (soldierType){
			case "marine":
				return new Marine();
			case "fire":
				return new Firebat();
		}
		// Null 을 쓰는 것은 좋지 않으나 패턴 설명을 위해 사용함.
		return null;
	}
} 
```

아래의 코드는 실제 Client에서 사용되는 코드의 예시입니다.  

```java 
package DesignPattern.gof_abstractFactory;
  
public class Military {

	public static void main(String[] args) {

		TrainingFactory infantryFactory = TrainingProvider.getFactory("infantry");

		Soldier marine = infantryFactory.create("marine");

		marine.attack();
		marine.getSoldier();

		Soldier firbat = infantryFactory.create("fire");

		firbat.getSoldier();
		firbat.attack();
	}
} 
```

### 작성해보면 좋을 만한 예시 문제 

```
- 요구사항 정의
스타 크래프트에서는 다양한 유닛이 존재합니다 탱크와 골리앗을 생산할 수 있는 Factory 마린, 메딕, 파이어뱃을 생산할 수 있는 Barrack 레이스, 사이언스 베슬, 베틀 크루저를 생산할 수 있는 Starport

- 요구사항의 공통화
모든 건물은 유닛을 생산한다.
모든 시스템 건물은 추가적으로 지을수 있어야 한다. 하지만 각 건물의 기능이 추가적으로 변경되지 않는다.
모든 유닛은 공격이 주된 목적을 가진다.
```


