# Chain of Responsibility 

## Gof's Description

메시지를 보내는 객체와 이를 받아 처리하는 객체들 간의 결합도를 없애기 위한 패턴입니다.     
하나의 요청에 대한 처리가 반드시 한 객체에서만 되지 않고, 여러 객체에 그 처리 기회를 주려는 것 입니다.   

## 풀이 

책임 연쇄 패턴의 중요한 부분은 메세지를 보내는 요청자, 그리고 그 요청을 처리하는 프로세스 간의 **관심사의 분리**가 얼마나 효과적으로 되는가 입니다.   

### 특정 요청에 대해서 처리하는 객체가 종속되어 있을 경우, 특정 요청 객체의 범위가 늘어날 경우, 처리하는 객체 내에서 추가적으로 로직이 작성되면, SRP 원칙을 위배하게 됩니다.
### 들어오는 요청 객체와 처리하는 객체가 분리되면, 처리하는 객체는 어떤 객체가 요청이 들어올 지 확신할 수 없습니다. 

만약 우리가 Web Application Server를 개발한다고 할 때, 통신 Protocol이 초기에는 QueryString에 제공되는 값만을 처리할 수 있다고 해봅시다.   

이후 서비스가 확장하면서 QueryString 외에도 Json, MultiPart도 지원해야한다고 할 때, Socket을 통해 들어오는 데이터에 대해서 데이터 타입에 따라서
처리할 수 있는 방식을 달리해야 합니다. 이때 책임 연쇄 패턴을 기반하여 코드를 작성해볼 수 있습니다. 

### **책임 연쇄 패턴에서의 처리자에 대한 추가, 삭제로 인한 요청자의 변경 사항에 발생해서는 안된다.**

처리자를 확장할 수 있는 구조여야 하고, 

```go 
package main

import (
	"fmt"
	"webserver/app/executor"
	"webserver/app/filter"
	"webserver/app/interceptor"
	"webserver/app/parser"
	"webserver/app/servlet"
	"webserver/app/webserver"
	"webserver/domain"
)

func main() {
	// Parser의 확장이 가능하다. 
	serverChain := webserver.NewWebServer(executor.NewDataParser(
		[]domain.Parser{
			parser.NewJsonParser(),
			parser.NewStringParser(),
		}))

	// servlet 내의 각 단계의 proxy 패턴 기반의 확장이 가능하다.
	newServlet := servlet.MustNewServlet(serverChain, []domain.Executor{
		filter.NewFilter(),
		interceptor.NewHttpInterceptor(),
	})

	newServlet.Process(domain.Request{
		URL:         "https://lines/hello.do",
		ContentType: "Json",
		Data:        "Im a Json",
	})

	fmt.Println("Finished 1 Thread")

	newServlet.Process(domain.Request{
		URL:         "https://lines/hello",
		ContentType: "MultiPart",
		Data:        "Im a MultiPart",
	})

	fmt.Println("Finished 2 Thread")
}
```

### 책임 연쇄 패턴의 처리자는 본인이 어떤 요청을 처리할지/할수 없을지 모르기에 들어오는 요청을 대상으로 자동으로 체크 및 처리 된다. 

- 파서의 들어오는 요청에 대한 각각의 처리 

```go 
package executor

import "webserver/domain"

type DataParser struct {
	parserChain []domain.Parser
}

func (d DataParser) Do(request domain.Request) interface{} {
	for _, parser := range d.parserChain {
		if _, ok := parser.Parse(); ok {
			return domain.Response{Data: "Hi"}
		}
	}

	return nil
}

func NewDataParser(parserChain []domain.Parser) domain.HttpExecutor {
	return DataParser{
		parserChain: parserChain,
	}
}

```

- 서블릿 내에서 filter와 interceptor를 처리한다. main에서 실제 정의한 책임에 대해서 자동으로 모두 체크할 수 있게 된다. 

```go 
package servlet

import "webserver/domain"

type Servlet struct {
	processor domain.Processor
	executors []domain.Executor
}

func MustNewServlet(executor domain.Processor, executors []domain.Executor) Servlet {
	return Servlet{processor: executor, executors: executors}
}

func (s *Servlet) Process(request domain.Request) {
	proxyLength := len(s.executors)

	for i := 0; i < proxyLength; i++ {
		s.executors[i].Pre(request)
	}

	response := s.processor.Do(request)

	for i := proxyLength; i > 0; i-- {
		data, ok := response.(domain.Response)
		if ok {
			s.executors[i-1].Post(data)
		}
	}
}
```

## 어떻게 활용 해야 하는가? 

- 하나 이상의 객체가 요청을 처리해야 하고, 그 요청 처리자 중 어떤 것이 선행자 인지 모를 때, 처리자가 자동으로 확정 되어야 한다.
- 메세지를 받을 객체를 명시 하지 않은 채 여러 객체 중 하나에게 처리를 요청 하고 싶을 때
- 요청을 처리할 수 있는 객체 집합이 동적으로 정의 되어야 할 때

## 실제 활용 했을 때 얻는 이점 

**다른 객체가 어떻게 요청을 처리하는지 몰라도 됩니다.**     
단지 요청을 보내는 객체는 이 메시지가 적절하게 처리될 것이라는 것만 확신하면 됩니다.   
**객체에 책임을 할당하는 데 유연성을 높일 수 있습니다.**     
**객체의 책임을 여러 객체에게 분산시킬 수 있으므로 런타임에 객체 연결 고리를 변경하거나 추가하여 책임을 변경하거나 확장할 수 있습니다.**   

## 다만,

어떤 객체가 이 처리에 대한 수신을 담당한다는 것을 명시하지 않으므로 요청이 처리된다는 보장은 없습니다.

### 또한 처리자가 너무 많거나 범위가 특정되기 어려울 경우, 반드시 시간 복작도를 고려한 알고리즘도 검토되어야 합니다.    