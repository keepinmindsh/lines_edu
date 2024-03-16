# Strategy

## Basic Information

- 패턴 형식 : 행위패턴

## Gof's Description 

동일 계열의 알고리즘을 정의하고, 각 알고리즘을 캡슐화하며, 이들을 상호교환이 가능하도록 만듭니다.
알고리즘을 사용하는 클라이언트와 상관 없이 독립적으로 알고리즘을 다양하게 변경할 수 있습니다.

### 풀이 

- 업무 자체는 변경되지 않는데, 업무를 처리하는 담당자가 변경되는 경우
- 예약 시스템에서 Copy라는 기는이 있는데 이를 예약에서 쓸 경우와 투숙에서 쓸 경우
- 동사무소에서 신청서를 출력할 때, 신청서의 양식에 따라서 신청서의 내용이 달라질 경우 

**전략**이라는 것은 어떤 동일한 "행위"에 대해서 처리하는 "사용자"와 "방식"에 따라서 행위의 결과는 같더라도 결과는 동일하게 처리할 수 있습니다.    
여기에서 **전략**은 처리하고자 하는 "행위"를 위한 알고리즘, 로직이 되며, 전략의 변화로 인한 외부 영향을 미치지 않은 상태에서,  
Runtime 이나 로딩 시점에 객체를 선언해서 사용할 수 있습니다.

#### 결과적으로 

- 동일 계열의 관련 알고리즘 군이 생깁니다.
- 서브 클래싱을 사용하지 않는 대안입니다.
- 조건문을 없앨수 있습니다.
- 구현의 선택이 가능합니다.
- 사용자(프로그램)는 서로 다른 전략을 알아야 합니다.
- Strategy 객체와 Context 객체 사이에 의사소통 오버헤드가 있습니다.
- 객체 수가 증가합니다.

## 코드 예시

### Golang

```go 
package main

import (
	"chat/app/chat"
	"chat/app/message"
	"chat/domain"
)

func main() {
	runMessage(message.NewChannelMessage())

	runMessage(message.NewDirectMessage())
}

func runMessage(message domain.Message) {
	chatting := chat.NewChat()

	chatting.RunContext(message)

	chatting.MessageCreate()
}
```

```go 
package chat

import (
	"chat/domain"
	"sync"
)

type chat struct {
	Context domain.Message
}

func (c *chat) MessageCreate() {
	c.Context.Create()

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		c.Context.Notify()
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		c.Context.Notify()
		wg.Done()
	}()

	wg.Wait()
}

func (c *chat) RunContext(message domain.Message) {
	c.Context = message
}

func NewChat() domain.Chat {
	return &chat{}
}
```

```go 
package message

import "chat/domain"

type channelMessage struct{}

func (c channelMessage) Create() {
	//TODO implement me
	panic("implement me")
}

func (c channelMessage) Notify() {
	//TODO implement me
	panic("implement me")
}

func (c channelMessage) Send() {
	//TODO implement me
	panic("implement me")
}

func NewChannelMessage() domain.Message {
	return channelMessage{}
}
```

```go 
package message

import "chat/domain"

type directMessage struct{}

func (d directMessage) Create() {
	//TODO implement me
	panic("implement me")
}

func (d directMessage) Notify() {
	//TODO implement me
	panic("implement me")
}

func (d directMessage) Send() {
	//TODO implement me
	panic("implement me")
}

func NewDirectMessage() domain.Message {
	return directMessage{}
}
```