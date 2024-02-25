# Observer 

## Basic Information 

- 패턴형식 : 행위패턴

## Gof's Description 

객체 사이에 일 대 다의 의존 관계를 정의해 두어, 어떤 객체의 상태가 변할 때 그 객체에 의존성을 가진 다른 객체들이 그 변화를 통지받고 자동으로 갱신될 수 있게 만듭니다.

## 풀이 

Gof에서 이야기하는 것은 다수의 객체들이 하나의 데이터를 바라보는 의존성을 가질 때, 
즉 데이터 변경이 여러 객체로 통지되어야하는 구조일 때, 객체간의 결합도를 높이지 않으면서
객체간의 일관성을 유지하는 것입니다.    

### 예시 

- 스프레드 시트의 데이터 값
- 스프레드 시트에 의존하는 표
- 스프레드 시트에 의존하는 막대 그래프 
- 스프레드 시트에 의존하는 원형 그래프 

스프레드 시트에 데이터 값이 변경될 때, 표, 막대 그래프, 원형 그래프의 값도 변경되어야 합니다.    

여기에서 스프레드 시트의  데이터를 **Subject(주체)** 라고 하고, 표, 막대 그래프, 원형 그래프를 **Observers(감시자)** 라고 합니다.  

- 주체는 데이터의 상태 변경을 감시자에게 통지해야 합니다. 
- 감시자는 주체의 상태 변경을 통지 받아 객체가 수행해야할 일을 처리합니다. 

위와 같은 종류의 상호 작용을 **게시-구독** 관계라고 합니다.  

- 주체는 상태 변경에 대한 통보를 하는 것이므로 누가 감시자인지 모른채 통보를 발송합니다. 
- 불특정 다수의 감시자가 이 통보를 수신하기 위해서 구독을 신청하는 것입니다. 

### 활용 가능한 경우 

- 어떤 추상 개념이 두 가지 양상을 갖고 다른 하나가 종속적일 때, 각 양상을 별도의 객체로 캡슐화하여 이들 각각을 재사용할 수 있습니다. 
- 한 객체에 가해진 변경으로 다른 객체를 변경해야 하고, 프로그래머들은 얼마나 많은 객체들이 변경되어야 하는지 몰라도 될때, 
- 어떤 객체가 다른 객체에 자신의 변화를 통보할 수 있는데, 그 변화에 관심이 있어하는 객체들이 누구인지에 대한 가정 없이도 그러한 통보가 될 때, 

### 중요 포인트 

- Subject, Object 클래스 간에는 추상적인 결합도 많이 존재합니다. 
- 브로드 캐스트 방식의 교류를 가능하게 합니다. 
- 예측하지 못한 정보를 갱신합니다. 

## 코드예시 

### Golang 

- 감시자 ( 상태 변경을 객체에 갱신한다. )

```go
package domain

type MarketData struct {
	Name    string
	Pricing int
}

type Client interface {
	Update(marketData MarketData)
}
```

- 주체 ( 상태 변경을 감시자에 통지 한다. )

```go 
package domain

type Stock interface {
	Register(client Client)
	RunStockMarket()
}
```

- 주체와 감시자 사이의 관계 

```go 
package main

import (
	"stock/app/client"
	"stock/app/stock"
	"stock/domain"
)

func main() {
	clients := []domain.Client{
		client.NewStockClient("Client1"),
		client.NewStockClient("Client2"),
		client.NewStockClient("Client3"),
		client.NewStockClient("Client4"),
		client.NewStockClient("Client5"),
	}

	server := stock.NewStockServer()

	for _, item := range clients {
		server.Register(item)
	}

	server.RunStockMarket()
}
```

- 감시자의 내부 구현

```go 
package client

import (
	"fmt"
	"stock/domain"
	"strconv"
)

type StockClient struct {
	name string
}

func (s StockClient) Update(data domain.MarketData) {
	fmt.Print(s.name + " 주식 정보 갱신")
	fmt.Print("주식명 : " + data.Name)
	fmt.Print("주식 가격 : " + strconv.Itoa(data.Pricing) + "원")
	fmt.Println(" ")
}

func NewStockClient(name string) domain.Client {
	return &StockClient{
		name: name,
	}
}
```

- 주체의 내부 구현 

```go 
package stock

import (
	"fmt"
	"math/rand"
	"stock/domain"
	"sync"
	"time"
)

type StockServer struct {
	stockClients []domain.Client
	data         domain.MarketData
}

func (s *StockServer) Register(client domain.Client) {
	s.stockClients = append(s.stockClients, client)
}

func (s *StockServer) notifyPricing() {
	clientCount := len(s.stockClients)

	if clientCount > 0 {
		for i := 0; i < clientCount; i++ {
			client := s.stockClients[i]

			client.Update(s.data)
		}
	}
}

func (s *StockServer) runChangePricing() {
	go func() {
		for {
			time.Sleep(time.Second * 4)

			s.data = domain.MarketData{
				Name:    "Corp",
				Pricing: rand.Int(),
			}
		}
	}()

}

func (s *StockServer) RunStockMarket() {
	var wg sync.WaitGroup

	wg.Add(1)
	s.runChangePricing()

	wg.Add(1)
	go func() {
		for {
			time.Sleep(time.Second * 5)

			fmt.Println("[주식 정보 갱신]")

			s.notifyPricing()
		}
	}()

	wg.Wait()
}

func NewStockServer() domain.Stock {
	return &StockServer{}
}
```

### Java 

```java 
package designpattern.gof_observer.sample01;

import designpattern.gof_observer.sample01.publisher.NewsMachine;
import designpattern.gof_observer.sample01.subscriber.AnnualSubscriber;
import designpattern.gof_observer.sample01.subscriber.EventSubscriber;

public class MainClass {
    public static void main(String[] args) {
        NewsMachine newsMachine = new NewsMachine();
        AnnualSubscriber annualSubscriber = new AnnualSubscriber(newsMachine);
        EventSubscriber eventSubscriber = new EventSubscriber(newsMachine);
        newsMachine.setNewsInfo("오늘 한파", "전국 영하 18도 입니다.");
        newsMachine.setNewsInfo("벛꽃 축제합니다", "다같이 벚꽃보러~");
    }
}
```

```java 
package designpattern.gof_observer.sample01.publisher;

import designpattern.gof_observer.sample01.subscriber.Observer;

import java.util.ArrayList;

public class NewsMachine implements Publisher {

    private ArrayList<Observer> observers;
    private String title;
    private String news;

    public NewsMachine() {
        observers = new ArrayList<>();
    }

    @Override
    public void add(Observer observer) {
        observers.add(observer);
    }

    @Override
    public void delete(Observer observer) {
        int index = observers.indexOf(observer);
        observers.remove(index);
    }

    @Override
    public void notifyObserver() {
        for (Observer observer : observers) {
            observer.update(title, news);
        }
    }

    public void setNewsInfo(String title, String news) {
        this.title = title;
        this.news = news;
        notifyObserver();
    }

    public String getTitle() {
        return title;
    }

    public String getNews() {
        return news;
    }
}

package designpattern.gof_observer.sample01.publisher;

import designpattern.gof_observer.sample01.subscriber.Observer;

public interface Publisher {
    public void add(Observer observer);

    public void delete(Observer observer);

    public void notifyObserver();
}

package designpattern.gof_observer.sample01.subscriber;

import designpattern.gof_observer.sample01.publisher.Publisher;

public class AnnualSubscriber implements Observer {
    private String newsString;
    private Publisher publisher;

    public AnnualSubscriber(Publisher publisher) {
        this.publisher = publisher;
        publisher.add(this);
    }

    @Override
    public void update(String title, String news) {
        this.newsString = title + " \n -------- \n " + news;
        display();
    }

    private void display() {
        System.out.println("\n\n오늘의 뉴스\n============================\n\n" + newsString);
    }
}

package designpattern.gof_observer.sample01.subscriber;

import designpattern.gof_observer.sample01.publisher.Publisher;

public class EventSubscriber implements Observer {
    private String newsString;
    private Publisher publisher;

    public EventSubscriber(Publisher publisher) {
        this.publisher = publisher;
        publisher.add(this);
    }

    @Override
    public void update(String title, String news) {
        newsString = title + "\n------------------------------------\n" + news;
        display();
    }

    public void display() {
        System.out.println("\n\n=== 이벤트 유저 ===");
        System.out.println("\n\n" + newsString);
    }

}

package designpattern.gof_observer.sample01.subscriber;

public interface Observer {
    public void update(String title, String news);
}
```

## Observer 패턴은, 

- Reactive 패러다임의 기본 개념이다.  
  - [Reactive Manifesto](https://reactivemanifesto.org/)
  - [Reactive Streams](https://www.reactive-streams.org/) 

### Reactive Manifesto

Reactive 시스템이 구성해야할 기본적인 규칙을 정의한다. Reactive Mesto 아래에 많은 Reactive 패러다임에 근간한 방법론이 제시되었다.

#### “외부의 어떤 이벤트나 데이터가 발생하면 거기에 대응하면 방식으로 동작하는 프로그램을 만드는 것”
#### “데이터 플로우와 상태 변경을 전파한다는 생각에 근간을 둔 프로그래밍 패러다임”
#### “막힘 없이 흘러다니는 데이터(이벤트)를 통해 사용자에게 자연스러운 응답을 주고, 규모 탄력적으로 리소스를 사용하며 실패에 있어서 유연하게 대처한다.”

- RxJava
- RxJS
- RxGo
- RxSwift 
- RxPy
- RxScala

### Responsive : 사용자에 대한 반응 

시스템이 적시에 응답합니다. 응답성은 사용성과 기능성의 기반인데, 그것보다 더 응답성은 문제에 대해서 빠르게 감지하는 것과 효율성을 다루는 것에 초점을 둡니다. 반응성(응답성)이 좋은 시스템은 속도와 일정한 응답성을 제공하고 , 신뢰할 수 있는 상향 경계를 수립하므로써, 일정한 품질의 서비스를 제공하는 것에 있습니다. 이 일관성있는 행동은 에러 처리를 간편화 하고, 사용자 신뢰를 구축하고, 앞선 상호작용을 장려합니다.

### Scalable(Elastic) : 부하에 대한 반응

시스템은 다양한 작업하에 응답성을 유지합니다. Reactive System은 서비스에 제공되기 위한 입력을 할당한 자원을 증가시키거나 감소시키는 것으로써 입력 량의 변화에 응답할 수 있다. 이것은 경합 포인트나 병목현상을 가지지 않게 설계되었으며, 공유하고, 컴포넌트를 복제하고, 입력을 분산할 수 있도록 하는 결과를 의도합니다. Reactive System은 응답성 뿐만 아니라 예측 가능성을 지원하는데 이는 실시간 성능 측정을 제공하여 알고리즘을 조정할 수 있습니다. 상용 하드웨어 및 소프트웨어 플랫폼에서 비용 효율적인 방식으로 탄력성을 얻을 수 있습니다.

### Resillent : 실패 상황에 대한 반응

시스템은 장애가 발생하더라도 응답성을 유지합니다. 이는 고 가용성, Mission Critical 시스템에 적용됩니다. 탄력성은 복제, 유지, 격리 및 위임에 의해 획득할 수 있습니다. 장애는 각각의 컴포넌트에 포함되어 있기 때문에 각각으로 부터 컴포넌트가 고립하되는 것으로 시스템 전체에 영향일 미치지 않고, 시스템의 일부가 장애 및 복구되는 것을 보장할 수 있습니다. 각각의 컴포넌트의 복수는 새로운 컴포넌트에게 위임되고, 고가용성은 필요에 따라 복제 따라서 보장됩니다. 고객의 기능은 장애를 처리함으로써 부담을 받지 않아도 됩니다.

### Event-Driven( Message-Driven ) : 이벤트에 대한 반응

Reactive 시스템은 Location Transparency, Isolation, Loose Coupling을 보장하는 컴포넌트들 사이의 경계를 관리하기 위해서 비동기적인 Message 전달 (Asynchronouse message-passing)에 의존합니다. 이 경계는 메세지로서 장애를 위임하기 위한 의도를 제공합니다. 명시적인 Message 전달을 이용하면 부하관리, 탄력성, 흐름제어 및 시스템에서의 메세지 큐 모니터링, 필요에 따라 Back Pressure를 적용하는 것을 가능하게 합니다. 통신수단으로의 Location Transparent Message는 동일한 구조와 의미의 단일 호스트 또는 클러스터와 동작하기 위한 장애의 관리를 가능하게 합니다. Non-Blocking Communication은 수신자로 하여금 활성 상태에서만 자원을 소모할 수 있게 하여 시스템의 오버헤드를 줄일 수 있습니다.

### Reactive 패러다임의 전체는 수학적 근거를 바탕으로 요청 처리의 주체를 변경한다. 

- Iterable <——> Observable ( 쌍대성 ) - Duality : 비슷한 구조를 가진다.
  - Iterable 은 Pulling 방식
  - Observable 은 Push 방식

> 쌍대성(雙對性; duality)은 수학과 물리학에서 자주 등장하는 표현이다. 보통 어떤 수학적 구조의 쌍대(雙對; dual)란 그 구조를 ‘뒤집어서’ 구성한 것을 말하는데,
> 엄밀한 정의는 세부 분야와 대상에 따라 각각 다르다. 쌍대의 쌍대는 자기 자신이므로 어떤 대상과 그 쌍대는 서로 일종의 한 ‘켤레’를 이룬다고 할 수 있으며, 이를 쌍대관계(雙對關係)라고 한다.

### 비동기적인 stream 프로세싱을 논-블록킹 방식의 배압(Back Pressure)를 이용해 표준을 제공합니다.

### Reactive Programming 은 Microservice 내에서 요청 처리에 효과적으로 사용이 가능합니다. 