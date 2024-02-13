# Memento 

## Gof's Description 

캡슐화를 위배하지 않은 채 어떤 객체의 내부 상태를 잡아내고 실체화 시켜둠으로써, 이후 해당 객체가 그 상태로 되돌아 올 수 있도록 합니다.

## 사용하게 되는 경우 

때에 따라서는 객체의 내부 상태를 기록해 둘 필요가 있습니다.   
예를 들어, 체크 포인트를 구현할 때나 오류를 복구하거나 연산 수행 결과를 취소하는 매커니즘을 구현하려면 내부 상태 기록이 절대적으로 필요합니다.   
객체를 이전 상태로 복구하려면 복구할 시점의 상태 정보가 있어야 합니다.   
그러나 객체는 자체적으로 상태의 일부나 전부를 캡슐화 하여 상태를 외부에 공개하지 않아야 하기 때문에, 다른 객체는 상태에 접근하지 못해야 합니다.

- 어떤 객체의 상태에 대한 스냅샷(몇 개의 일부)을 저장한 후 나중에 이 상태로 복구 해야 할 때,
- 상태를 얻는 데 필요한 직접적인 인터페이스를 두면 그 객체의 구현 세부 사항이 드러날 수 밖에 없고, 이 때문에 객체의 캡슐화가 깨질 때

## 해당 패턴이 가지는 결과 

- 캡슐화된 경계를 유지할 수 있습니다.
- Originator 클래스를 단순화 할 수 있습니다.
- 메멘토의 사용으로 더 많은 비용이 들어갈 수도 있습니다.
- 제한 범위 인터페이스와 광범위 인터페이스를 정의 해야 합니다.
- 메멘토를 관리하는 데 필요한 비용이 숨어 있습니다.

## 메멘토 패턴을 활용하려면, 

명령 패턴을 실행 취소가 가능한 연산의 상태가 저장할 때 메멘토 패턴을 사용할 수 있습니다. 메멘토 패턴은 반복자 패턴에서의 반복 과정 상태를 관리할 수 있습니다.

## [SQL] Oracle - SavePoint 

```
1. Commit 
2. Insert 
3. Insert 
-- SavePoint 1 
4. Update 
5. Update 
6. Delete 
-- SavePoint 2
7. Insert 
8. Update 
Commit; 8 까지의 모든 프로세스가 확정 됨 
RollBack; 8 까지의 모든 프로세스가 취소됨 
RollBack to SavePoint 1; 4 ~ 8까지의 모든 프로세스 취소 
```

## 실사용 예제를 살펴보면, 

```go 
package main

import (
	"gamesave/app/play"
	"gamesave/app/save"
	"gamesave/domain"
)

const Latest int = 0

func main() {
	gamePlay := play.NewGamePlay(save.NewSave())

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 20,
		},
		Status: domain.CharaterStatus{
			Status: "good",
		},
		Address: domain.StoryPoint{
			Address: "story-1-1",
		},
	})

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 22,
		},
		Status: domain.CharaterStatus{
			Status: "good",
		},
		Address: domain.StoryPoint{
			Address: "story-1-1",
		},
	})

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 25,
		},
		Status: domain.CharaterStatus{
			Status: "bad",
		},
		Address: domain.StoryPoint{
			Address: "story-1-2",
		},
	})

	gamePlay.Save()

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 35,
		},
		Status: domain.CharaterStatus{
			Status: "good",
		},
		Address: domain.StoryPoint{
			Address: "story-1-8",
		},
	})

	gamePlay.Save()

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 35,
		},
		Status: domain.CharaterStatus{
			Status: "dead",
		},
		Address: domain.StoryPoint{
			Address: "story-1-20",
		},
	})

	gamePlay.Play(gamePlay.Load(Latest))
}
```

- Game Play 중 게임 저장 

```go 
package play

import (
	"gamesave/domain"
)

type GamePlay struct {
	GameData   domain.GameSave
	SaveModule domain.Save
}

func (g *GamePlay) Save() {
	g.SaveModule.Save(g.GameData)
}

func (g *GamePlay) Load(savePoint int) domain.GameSave {
	return g.SaveModule.Load(savePoint)
}

func (g *GamePlay) Play(data domain.GameSave) {
	g.GameData = data
}

func NewGamePlay(save domain.Save) domain.Game {
	return &GamePlay{
		SaveModule: save,
	}
}
```

- 저장을 위한 용도 

```go 
package save

import (
	"gamesave/domain"
)

type Save struct {
	SaveData []domain.GameSave
}

func (s *Save) Save(data domain.GameSave) {
	s.SaveData = append(s.SaveData, data)
}

func (s *Save) Load(savePoint int) domain.GameSave {
	if s.SaveData[savePoint] != (domain.GameSave{}) {
		return s.SaveData[savePoint]
	} else {
		return domain.GameSave{}
	}
}

func NewSave() domain.Save {
	return &Save{
		SaveData: []domain.GameSave{},
	}
}
```