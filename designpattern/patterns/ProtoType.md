# Pattern - Builder 

## Basic Information

- 패턴형식 : 생성패턴 

## Gof's Description

객체 생성과 관련된 패턴들은 일부 영역이 겹치는 부분들이 있는데, 프로트 타입 패턴은 주로 다른 객체 생성 패턴들과 많이 같이 사용될 수 있습니다. 추상 팩토리 패턴이 프로토타입들의 집합을 가지고 있다가, 복제(clone)한 프로덕트 객체를 반환할 수 있습니다. 일반화 관계로 표현을 할 때, 파생 클래스의 개수가 과도히 많아지고 각 클래스의 메서드가 수행하는 알고리즘에서 차이가 없고 생성 시에 개체의 속성에 차이만 있다면 원형 패턴을 사용하는 것이 효과적입니다.

### 풀이

- 예를 들어, DB로 부터 데이터를 가져올 때, 해당 데이터의 유형이 실시간으로 갱신이 필요한 데이터가 아니고, 주기적으로 갱신처리를 해도 된다면 초기에 가져온 데이터를 새로운 객체에 복사하여 데이터 수정 작업을 하는 것이 더 효율적이다. 
- 객체의 속성만 차이가 있는 경우, 객체를 복제하는 시점에 속성 값에 대해서만 변경하여 객체를 사용하는 것이 훨씬 효율적일 수 있다. 

### Deep Copy & Shallow Copy 에 대한 이해 

- Deep Copy : 깊은 복사 

하나의 객체를 복사할 때 메모리 주소까지 분리해서 복사하는 경우를 **깊은 복사** 라고 한다.     
깊은 복사를 이용한 객체의 경우, 완전히 독립적인 새로운 객체이므로 기존의 객체에 영향을 줄 수 없다.       

- Shallow Copy : 얕은 복사  

하나의 객체를 복사할 때 동일한 메모리 주소를 사용하는 경우를 **얕은 복사** 라고 한다.    
얕은 복사를 이용한 객체의 경우, 기존의 객체와 연결 되어 있기 때문에 사용 방식에 따라 기존의 객체에 영향을 줄 수 있다.   

- **ProtoType 패턴 에서의 복사 방식** 

프로토 타입 에서의 복사 방식은 동일한 알고리즘, 프로세스, 행위에 대해서는 **얕은 복사**를 수행하되, 각 객체의 상태, 변경에 대한 부분은 **깊은 복사**를 수행 한다. 

### 장점 

- Object 생성이 높은 비용으로 수 많은 요청을 하는 경우, 
- 비슷한 Object를 지속적으로 생성해내야 하는 경우,  

## 실생활에서의 사용 예시 

스타크래프트에서 미네랄, 가스를 캐기 위해서 사용되는 유닛은 프로브가 있습니다. 프로브를 생성하는 방식에 대해서 프로토타입을 적용해봅니다. 

### Go 에서는

```go 
package unit

import (
	"fmt"
	"starcraft/domain"
)

type Probe struct {
	MineralCapacity    int
	AccumulatedMineral int
}

func (p *Probe) Harvest() {
	fmt.Println("자원을 캡니다.")

	if p.MineralCapacity >= p.AccumulatedMineral {
		p.AccumulatedMineral += 1
	} else {
		fmt.Println("이미 캘수 있는 모든 미네랄을 다 캐었기 때문에 더이상 캘 수 없습니다.")
	}
}

func (p *Probe) Attack() {
	fmt.Println("공격을 합니다.")
}

func (p *Probe) Building() {
	fmt.Println("건물을 짓습니다.")
}

func (p *Probe) GerMineralCapacity() int {
	return p.MineralCapacity
}

func NewCloneUnit() domain.Unit {
	cloneObj := &Probe{
		MineralCapacity:    50,
		AccumulatedMineral: 0,
	}

	return cloneObj
}

```

```go 
package building

import (
	"starcraft/app/unit"
	"starcraft/domain"
)

type Nexus struct{}

func (n Nexus) CreateUnit() domain.Unit {
	return unit.NewCloneUnit()
}

func NewNexus() domain.Building {
	return &Nexus{}
}
```

```go 
package main

import "starcraft/app/unit"

func main() {
	clonedUnit := unit.NewCloneUnit()

	clonedUnit.Building()

	clonedUnit.Attack()

	clonedUnit.Harvest()
}

```

### Go에서 위의 코드를 좀더 효율적으로 분리한다면, 

action은 모든 프로브의 공통적으로 사용되어야 할 객체이므로, 해당 객체는 신규 프로브 생성시 마다 생성하지 않는다. 

```go 
package action

import (
	"effective_startcraft/domain"
	"fmt"
)

type Action struct {
	MineralCapacity int
}

func (p *Action) Harvest(accumulatedMineral int) int {
	fmt.Println("자원을 캡니다.")

	if p.MineralCapacity >= accumulatedMineral {
		accumulatedMineral += 1
	} else {
		fmt.Println("이미 캘수 있는 모든 미네랄을 다 캐었기 때문에 더이상 캘 수 없습니다.")
	}

	return accumulatedMineral
}

func (p *Action) Attack() {
	fmt.Println("공격을 합니다.")
}

func (p *Action) Building() {
	fmt.Println("건물을 짓습니다.")
}

func NewAction() domain.UnitAction {
	return &Action{
		MineralCapacity: 50,
	}
}
```

```go 
package unit

import (
	"effective_startcraft/domain"
)

type Probe struct {
	AccumulatedMineral int
	domain.UnitAction
}

func (p *Probe) GerMineralCapacity() int {
	return p.AccumulatedMineral
}

func (p *Probe) SetMineralCapacity(accumulatedMineral int) {
	p.AccumulatedMineral = accumulatedMineral
}

func NewCloneUnit(action domain.UnitAction) domain.Unit {
	cloneObj := &Probe{
		0,
		action,
	}

	return cloneObj
}
```


```go 
package domain

type (
	Unit interface {
		GerMineralCapacity() int
		SetMineralCapacity(accumulatedMineral int)
		UnitAction
	}

	UnitAction interface {
		Harvest(accumulatedMineral int) int
		Attack()
		Building()
	}
)
```

### Java 에서는 

```java 
package lines.model;

public interface Building {
    Unit CreateUnit();
}

package lines.model;

public interface Unit {
    int GerMineralCapacity();
    void SetMineralCapacity(int accumulatedMineral);

    void Harvest();
    void Attack();
    void Building();
}

package lines.model;

public interface UnitAction {
    void Harvest();
    void Attack();
    void Building();
}
```

```java 
package application;

import lines.model.Unit;
import lines.model.UnitAction;
import lines.service.domain.MarineAction;
import lines.service.factory.UnitFactory;

public class Application {
    public static void main(String[] args) {
        UnitAction unitAction = new MarineAction();

        for (int i = 0; i < 60; i++) {
            Unit unit = UnitFactory.CreateUnit(unitAction);

            unit.Attack();

            unit.Building();

            unit.Harvest();
        }
    }
}
```

```java 
package lines.service.usecase;

import lines.model.Unit;
import lines.model.UnitAction;

public class Marine implements Unit {

    private final UnitAction unitAction;
    private int accumulatedMineral;

    public Marine(UnitAction unitAction){
        this.unitAction = unitAction;
    }

    @Override
    public int GerMineralCapacity() {
        return accumulatedMineral;
    }

    @Override
    public void SetMineralCapacity(int accumulatedMineral) {
        this.accumulatedMineral = accumulatedMineral;
    }

    @Override
    public void Harvest() {
        this.unitAction.Harvest();
    }

    @Override
    public void Attack() {
        this.unitAction.Attack();
    }

    @Override
    public void Building() {
        this.unitAction.Building();
    }
}
```

```java 
package lines.service.factory;

import lines.model.Unit;
import lines.model.UnitAction;
import lines.service.usecase.Marine;

public class UnitFactory {

    public static Unit CreateUnit(UnitAction action) {
        return new Marine(action);
    }
}
```

```java 
package lines.service.domain;

import lines.model.UnitAction;

public class MarineAction implements UnitAction {
    @Override
    public void Harvest() {
        System.out.println("수확합니다.");
    }

    @Override
    public void Attack() {
        System.out.println("공격합니다.");
    }

    @Override
    public void Building() {
        System.out.println("건설합니다.");
    }
}
```

### Kotlin 에서는 

### Dart 에서는 