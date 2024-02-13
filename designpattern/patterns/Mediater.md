# Mediator 

## Gof's Description 

한 집합에 속해 있는 객체의 상호 작용을 캡슐화 하는 객체를 정의 합니다.
객체 들이 직접 서로를 참조하지 않도록 하여 객체 사이의 소결합을 촉진시키며, 개발자가 객체의 상호 작용을 독립적으로 다양화 시킬 수 있게 만듭니다.

## 사용하게 되는 경우 

- 여러 객체가 잘 정의된 형태이기는 하지만 복잡한 상호작용을 가질 때. 객체 간의 의존성이 구조화되지 않으며, 잘 이해하기 어려울 때
- 한 객체가 다른 객체를 너무 많이 참조하고, 너무 많은 의사소통을 수행해서 그 객체를 재사용하기 힘들 때
- 여러 클래스에 분산된 행동들이 상속 없이 상황에 맞게 수정되어야 할 때

## 풀이 

우리는 흔히 Mediator(중재자) 패턴을 비행기와 관제탑 사이의 관계에 빗대어 많이 이야기하기 됩니다.

- 비행기 
  - 비행기는 자신이 이동하면서 착륙하기 위한 순서를 알아야 합니다. 
  - 다른 비행기가 현재 착륙 중인지, 착륙장에 비행기가 이동 중인지, 자신의 뒤에 비행기가 따라오고 있는지 각각의 상황을 개별 비행기를 통해 확인하기 어렵습니다. 
  - 그렇다고 비행기 내에 관제탑 역할을 하는 기기를 넣을 수 없습니다. 각 착륙 장소마다의 방식도 다를 수 있기 때문입니다. 
- 관제탑
  - 관제탑은 해당 공항으로 이동 중인 모든 비행기의 정보를 실시간으로 수신합니다. 
  - 수신된 정보를 바탕으로 현재 활주로의 상태를 고려하여 비행기를 순차적으로 착륙할 수 있도록 정보를 비행기에 제공합니다. 
- 공항 활주로 
  - 활주로는 한번에 하나의 비행기만 착륙할 수 있습니다. 

위의 상황에서 관제탑은 각 비행기(객체) 간의 상호 작용을 중재하여 의사소통을 수행할 수 있습니다.  
각 비행기는 관제탑에만 현재 상태를 확인 요청하고, 비행기(객체) 자체의 역할을 수행할 수 있습니다.  

## 실제 코드로 구현해보면, 


- 중재를 위한 상태 객체 정의  

```go 
package strip

import (
  "right_airplan/domain"
)

type strip struct {
  AirPlanName string
}

func (s *strip) SetAirPlanName(airPlaneName string) {
  s.AirPlanName = airPlaneName
}

func (s *strip) GetAirPlanName() string {
  return s.AirPlanName
}

func (s *strip) LandingProcedure() {
  s.AirPlanName = "RUNNING"
}

func (s *strip) Complete() {
  s.AirPlanName = "COMPLETE"
}

func (s *strip) GetRoadStatus() string {
  return s.AirPlanName
}

func NewStrip() domain.Strip {
  return &strip{
    AirPlanName: "BEFORE",
  }
}
```
- 항공기 객제 정의 

```go 
package plane

import (
	"fmt"
	"right_airplan/domain"
	"strconv"
	"time"
)

type Plane struct {
	Name      string
	Landed    bool
	IsLanding bool
	Strip     *domain.Strip
}

func (p *Plane) GetName() string {
	return p.Name
}

func (p *Plane) LandingEnable() bool {
	return p.IsLanding
}

func (p *Plane) SetLandingEnable(isLanding bool) {
	p.IsLanding = isLanding
}

func (p *Plane) Landing() {
	p.IsLanding = true
	fmt.Println(p.Name + "이 착륙 중 입니다. ")
	var appendString string
	for i := 10; i > 0; i-- {
		time.Sleep(time.Second * 1)
		appendString += strconv.Itoa(i) + " 미터 상공 > "
	}
	fmt.Println(p.Name + " : " + appendString + "착륙 완료")
}

func (p *Plane) Stop() {
	(*p.Strip).Complete()
	p.IsLanding = false
	p.Landed = true
	fmt.Println(p.Name + "가 착륙했습니다.")
}

func (p *Plane) PlaneLanded() bool {
	return p.Landed
}

func NewPlane(planeName string, strip *domain.Strip) *Plane {
	return &Plane{Name: planeName, Strip: strip, Landed: false, IsLanding: false}
}

```

- Mediator 역할을 하는 TestDo 핸들링 (Singleton 과 함께 활용)

```go 
package airport

import (
  "fmt"
  "right_airplan/app/plane"
  "right_airplan/app/strip"
  "right_airplan/domain"
  "sync"
  "testing"
  "time"
)

var lock = &sync.Mutex{}

func TestDo(t *testing.T) {
  airStrip := strip.NewStrip()

  airplanes := []*plane.Plane{
    plane.NewPlane("대한항공", &airStrip),
    plane.NewPlane("아시아나 항공", &airStrip),
    plane.NewPlane("제주 항공", &airStrip),
    plane.NewPlane("티웨이 항공", &airStrip),
    plane.NewPlane("동방항공", &airStrip),
  }

  lock := &sync.Mutex{}

  var wg sync.WaitGroup

  for {
    for i := 0; i < len(airplanes); i++ {
      if !airplanes[i].PlaneLanded() {
        wg.Add(1)

        go func(item domain.AirPlane) {
          if airStrip.GetRoadStatus() == "BEFORE" || airStrip.GetRoadStatus() == "COMPLETE" {
            if !item.PlaneLanded() {
              if airStrip.GetRoadStatus() != "RUNNING" {
                lock.Lock()
                defer lock.Unlock()
                if airStrip.GetRoadStatus() != "RUNNING" {
                  airStrip.LandingProcedure()

                  item.Landing()

                  item.Stop()
                }
              }
            }
          }

          wg.Done()
        }(airplanes[i])
      }
    }

    var waitingAirplane string
    for _, airplane := range airplanes {
      if !airplane.LandingEnable() && !airplane.PlaneLanded() {
        waitingAirplane += "[" + airplane.GetName() + "]"
      }
    }
    fmt.Println(waitingAirplane + " 착륙 대기중")

    var isFinished bool
    for _, airplane := range airplanes {
      if !airplane.PlaneLanded() {
        isFinished = false
        break
      }

      isFinished = true
    }

    if isFinished {
      break
    }
    time.Sleep(time.Second * 1)
  }

  wg.Wait()
}
```

### 실행 결과

```shell 
=== RUN   TestDo
[대한항공][아시아나 항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
아시아나 항공이 착륙 중 입니다. 
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
아시아나 항공 : 10 미터 상공 > 9 미터 상공 > 8 미터 상공 > 7 미터 상공 > 6 미터 상공 > 5 미터 상공 > 4 미터 상공 > 3 미터 상공 > 2 미터 상공 > 1 미터 상공 > 착륙 완료
아시아나 항공가 착륙했습니다.
동방항공이 착륙 중 입니다. 
[대한항공][제주 항공][티웨이 항공][동방항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
동방항공 : 10 미터 상공 > 9 미터 상공 > 8 미터 상공 > 7 미터 상공 > 6 미터 상공 > 5 미터 상공 > 4 미터 상공 > 3 미터 상공 > 2 미터 상공 > 1 미터 상공 > 착륙 완료
동방항공가 착륙했습니다.
[대한항공][제주 항공][티웨이 항공] 착륙 대기중
티웨이 항공이 착륙 중 입니다. 
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
[대한항공][제주 항공] 착륙 대기중
티웨이 항공 : 10 미터 상공 > 9 미터 상공 > 8 미터 상공 > 7 미터 상공 > 6 미터 상공 > 5 미터 상공 > 4 미터 상공 > 3 미터 상공 > 2 미터 상공 > 1 미터 상공 > 착륙 완료
티웨이 항공가 착륙했습니다.
[대한항공][제주 항공] 착륙 대기중
대한항공이 착륙 중 입니다. 
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
대한항공 : 10 미터 상공 > 9 미터 상공 > 8 미터 상공 > 7 미터 상공 > 6 미터 상공 > 5 미터 상공 > 4 미터 상공 > 3 미터 상공 > 2 미터 상공 > 1 미터 상공 > 착륙 완료
대한항공가 착륙했습니다.
[제주 항공] 착륙 대기중
[제주 항공] 착륙 대기중
제주 항공이 착륙 중 입니다. 
 착륙 대기중
 착륙 대기중
 착륙 대기중
 착륙 대기중
 착륙 대기중
 착륙 대기중
 착륙 대기중
 착륙 대기중
 착륙 대기중
 착륙 대기중
제주 항공 : 10 미터 상공 > 9 미터 상공 > 8 미터 상공 > 7 미터 상공 > 6 미터 상공 > 5 미터 상공 > 4 미터 상공 > 3 미터 상공 > 2 미터 상공 > 1 미터 상공 > 착륙 완료
제주 항공가 착륙했습니다.
 착륙 대기중
--- PASS: TestDo (54.06s)
```