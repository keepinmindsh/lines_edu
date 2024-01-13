# Command 

## Basic Information

- 패턴형식 : 명령패턴 

## Gof's Description

요청 자체를 캡슐화 하는 것 입니다. 이를 통해 요청이 서로 다른 사용자를 매개 변수로 만들고, 요청을 대기 시키거나 로깅 하며, 되돌릴 수 있는 연산을 지원 합니다.  

## 동기

항상 그렇지는 않지만, 요청 받은 연산이 무엇 이며, 이를 처리할 객체가 누구 인지에 대한 아무런 정보 없이 임의의 객체에 메세지를 보내야 할 때가 간간이 있습니다.     
예를 들어, 사용자 인터페이스 툴킷은 버튼, 메뉴 같은 객체를 포함하는데, 이는 사용자의 메세지를 처리하게 됩니다.     
그러나 사용자 인터페이스 툴킷은 버튼과 메뉴에서 요청을 처리할 수 없습니다. 툴킷을 사용하는 응용프로그램만이 어떤 객체를 통해서 어떤 일이 되어야 하는지 알기 때문입니다.          
사용자 인터페이스 툴킷 설계자의 입장에서는 어떤 객체가 이 요청을 처리할 지를 알아낼 방법이 없습니다.

## 좀 더 쉽게 풀이 하기

어떤 특정한 작업이 **요청**할 주체가 **누구**인지 명확하게 알 수 없을 때 사용한다. 

![Pacade Pattern](https://github.com/keepinmindsh/lines_edu/blob/main/assets/command_pattern.png)

즉, 처음에는 채널을 통해서 메세지를 보내는 채널 메세지 생성 함수가 있었는데, 다이렉트 메세지를 통해서도 메세지 생성이 필요하게 되었다. 이럴 경우, 

### Channel Message 생성만 존재 했을때, 

```go 
package main

import "fmt"

func main() {
	ChannelMessageCreate()
}

func ChannelMessageCreate() {
	fmt.Println("채널 메세지를 생성합니다.")
}
```

### 만약 DirectMessageCreate 가 생성되는 경우, 

```go
package main

func main() {

}

type Caller struct {
	message Message
}

func NewCaller(msg Message) *Caller {
	return &Caller{
		message: msg,
	}
}

func (c Caller) ChannelCall() {
	// 메세지를 전송한다는 명령은 동일하다. 
	c.message.MessageCreate()
}

func (c Caller) DirectMessageCall() {
	// 메세지를 전송한다는 명령은 동일하다.
	c.message.MessageCreate()
}

type Message interface {
	MessageCreate()
}

type messageCreate struct {
}

func NewMessage() Message {
	return &messageCreate{}
}

// 명령을 수행하는 MessageCreate는 호출 이후의 상태 관리를 자체적으로 수행해야 한다.
func (m messageCreate) MessageCreate() {

}
```

### 요청을 수행하는 Command 는 

- Command 내에서 메세지 생성을 위한 연산 및 상태에 대해서 관리 및 처리되어야 한다. 
  - 예시) 메세지를 생성하다가 실패할 경우, 실패에 대한 RollBack은 Command 객체가 내부에서 처리되어야 한다. 
    - Method 호출, RPC Call, API 콜 등등 

## 정리해보면, 

- Command 는 연산을 호출하는 객체와 연산 수행 방법을 구현하는 객체를 분리합니다.
  - 여기에서 분리의 개념은 단순히 객체로만의 분리가 아닌 
    - Rest API 
    - RPC 
    - Event 기반의 Pub/Sub 
    - etc
- Command 는 [일급 클래스](#일급-클래스란) 입니다. 다른 객체와 같은 방식으로 조작되고 확장할 수 있습니다.
- 명령을 여러 개를 복합해서 복합 명령을 만들 수 있습니다. 앞에서 Macro Command 클래스를 예로 들었지만, 복합체 패턴을 이용하여 여러 명령어를 구성할 수도 있습니다.
- 새로운 Command 객체를 추가하기 쉽습니다. 기존 클래스를 변경할 필요 없이 단지 새로운 명령어에 대응하는 클래스만 정의하면 됩니다.

## 실제 사용하는 코드 예시 

### Java 

```java 
/*
     문제를 해결하기 위해서는 구체적인 기능을 직접 구현하는 대신 실행될 기능을 캡슐화해야 한다.
     즉, Button 클래스의 pressed 메서드에서 구체적인 기능(램프 켜기, 알람 동작 등)을 직접 구현하는 대신 버튼을 눌렀을 때 
     실행될 기능을 Button 클래스 외부에서 제공받아 캡슐화해 pressed 메서드에서 호출한다.
*/
package DesignPattern.gof_command.sample02;

public class Client {
    public static void main(String[] args) {
        Lamp lamp = new Lamp();
        Command lampOnCommand = new LampOnCommand(lamp);
        Alarm alarm = new Alarm();
        Command alarmOnCommand = new AlarmStartCommand(alarm);

        Button button1 = new Button(lampOnCommand);

        button1.pressed();

        Button button2 = new Button(alarmOnCommand);
        button2.pressed();
        button2.SetCommand(lampOnCommand);
        button2.pressed();
    }
}  
```

```java 
package DesignPattern.gof_command.sample02;

public class Button {

    private Command theCommand;

    public Button(Command command ){
        SetCommand(command);
    }

    public void SetCommand(Command command){
        this.theCommand = command;
    }

    public void pressed(){
        theCommand.execute();
    }
}   

package DesignPattern.gof_command.sample02;

public interface Command {

    public abstract void execute();
}

package DesignPattern.gof_command.sample02;

public class Alarm {
    public void start(){
        System.out.println("Alarming");
    }
}

package DesignPattern.gof_command.sample02;

public class AlarmStartCommand implements Command {

    private Alarm theAlarm;

    public AlarmStartCommand(Alarm alarm){
        this.theAlarm = alarm;
    }

    @Override
    public void execute() {
        theAlarm.start();
    }
}

package DesignPattern.gof_command.sample02;

public class Lamp {
    public void turnOn(){
        System.out.println("Lamp On");
    }
}

package DesignPattern.gof_command.sample02;

public class LampOnCommand implements Command {

    private Lamp theLamp;

    public LampOnCommand(Lamp lamp){
        this.theLamp = lamp;
    }

    @Override
    public void execute() {
        theLamp.turnOn();
    }
}
```

### Go 

```go 
package main

import (
	"message/apps/usecase"
	"message/domain"
)

func main() {
	messageCreate := usecase.NewMessageCreate()

	publicCaller := usecase.MustNewCaller(messageCreate, domain.PUBLIC_CHANNEL)
	publicCaller.Send()

	privateCaller := usecase.MustNewCaller(messageCreate, domain.PRIVATE_CHANNEL)
	privateCaller.Send()
}
```

```go 
package usecase

import "message/domain"

func MustNewCaller(message domain.Message, callType domain.CallType) domain.Caller {
	switch callType {
	case domain.PRIVATE_CHANNEL:
		return newPrivateChannel(message)
	case domain.PUBLIC_CHANNEL:
		return newPublicChannel(message)
	default:
		panic("error, check your callType")
	}
}

type message struct{}

func NewMessageCreate() domain.Message {
	return &message{}
}
```

```go


```

# Tips

## 일급 클래스란?

다른 객체들에 일반적으로 적용 가능한 연산을 모두 지원하는 객체를 가리킨다, 보통 함수에 인자로 넘기기, 수정하기, 변수에 대인하기와 같은 연산을 지원할 때 일급 객체라고 한다. 

- 객체는 변수나 매개변수에 할당할 수 있다. 
- 객체는 다른 객체와 동등한 지위를 가진다. 
- 객체는 반환 값으로 사용할 수 있다. 
- 객체는 필요한 경우 메서드에서 생성할 수 있다. 
- 