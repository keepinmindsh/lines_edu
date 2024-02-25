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