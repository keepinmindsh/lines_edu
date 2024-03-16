# State 

## Basic Information 

- 패턴 형식 : 행위패턴 

## Gof's Description 

객체의 내부 상태에 따라 스스로 행동을 변경할 수 있게 허가하는 패턴으로, 이렇게하면 객체는 마치 자신의 클래스를 바꾸는 것처럼 보입니다.

### 풀이 

- 객체의 행동이 상태에 따라서 달라질 수 있을때, 객체의 상태에 따라서 런타임에 행동이 변경되어야 한다. 
  - 하나의 객체가 가지는 행동(추상화)이 상태에 따라 변경된다는 것은 
    - 사람이 걷는 것을 
      - 아플때 
      - 건강할 때 
      - 발목을 삐었을 때 
      - 뼈가 부러졌을 때 
    - 사람의 내부 상태(건강상태)에 따라서 그 행동이 달라질 수 있습니다. 

#### 즉, 

객체의 내부 상태가 변경됨에 따라 Runtime의 행동이 변화가 일어날 수 있다.   
어떤 업무나 요구사항의 행위는 정해져 있지만 정해져있는 행위가 상태의 변경에 따라 행위가 변경되는 것.   

#### 예시) 

- 호텔의 객실에 손님을 "배정"하려할 할때, 객실의 "상태"에 따라서 "배정"이라는 행위의 결과가 달라진다.
- 카드를 이용해서 "결제"를 하려 할 때, 카드 내의 잔고 "상태"에 따라서 "결제" 행위의 결과다 달라진다. 

**위의 예시에서 중요한 부분은 상태를 관리하는 객체와 상태에 따라서 "행위"의 방식이 나뉘어지는 객체를 적절하게 분리하는 것입니다.**  

호텔의 객실에 대한 "배정"이라는 행위가 객체 생성의 기준이되고, 각 행위를 실행하고, 상태를 관리하는 "객실"는 상태에 따라 "배정"의 행위를 달리 실행하게 됩니다.   
**즉 객실은 행위를 수행할 수 있는 추상화된 로직이 존재 해야합니다.**    

## 코드 예시 

### Golang

```go 
package main

import (
	"human_walk/app/human"
	"human_walk/domain"
)

func main() {
	aPerson := human.NewHuman()

	aPerson.HealthStatus(domain.HEALTHY)
	aPerson.Walk()
	aPerson.HealthStatus(domain.SICK)
	aPerson.Walk()
	aPerson.HealthStatus(domain.BROKEN_ANKLE)
	aPerson.Walk()
	aPerson.HealthStatus(domain.BROKEN_BONES)
	aPerson.Walk()
}
```

```go 
package human

import (
	"human_walk/app/action"
	"human_walk/domain"
)

type human struct {
	HealthyStatus domain.HealthType

	SickWalk        domain.Walk
	HealthyWalk     domain.Walk
	BrokenAnkleWalk domain.Walk
	BrokenBonesWalk domain.Walk
}

func (h *human) HealthStatus(healthStatus domain.HealthType) {
	h.HealthyStatus = healthStatus
}

func (h *human) Walk() {
	switch h.HealthyStatus {
	case domain.SICK:
		h.SickWalk.Walk()
	case domain.HEALTHY:
		h.HealthyWalk.Walk()
	case domain.BROKEN_ANKLE:
		h.BrokenAnkleWalk.Walk()
	case domain.BROKEN_BONES:
		h.BrokenBonesWalk.Walk()
	}
}

func NewHuman() domain.Action {
	return &human{
		SickWalk:        action.NewSickWalk(),
		HealthyWalk:     action.NewHealthyWork(),
		BrokenAnkleWalk: action.NewBrokenAnkleWalk(),
		BrokenBonesWalk: action.NewBrokenBonesWalk(),
	}
}
```

```go 
package action

import "human_walk/domain"

type brokenAnkleWalk struct {
}

func (b brokenAnkleWalk) Walk() {
	//TODO implement me
	panic("implement me")
}

func NewBrokenAnkleWalk() domain.Walk {
	return brokenAnkleWalk{}
}


package action

import "human_walk/domain"

type brokenBonesWalk struct {
}

func (b brokenBonesWalk) Walk() {
  //TODO implement me
  panic("implement me")
}

func NewBrokenBonesWalk() domain.Walk {
  return brokenBonesWalk{}
}

package action

import "human_walk/domain"

type healthyWork struct {
}

func (h healthyWork) Walk() {
  //TODO implement me
  panic("implement me")
}

func NewHealthyWork() domain.Walk {
  return healthyWork{}
}

package action

import "human_walk/domain"

type sickWalk struct {
}

func (s sickWalk) Walk() {
  //TODO implement me
  panic("implement me")
}

func NewSickWalk() domain.Walk {
  return sickWalk{}
}
```