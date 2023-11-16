# Pattern - Bridge

## Basic Information 

- 패턴 형식 : 목적 패턴 

## Gof's Description 

구현에서 추상을 분리하여, 이들이 독립적으로 다양성을 가질 수 있도록 합니다.

## 필요한 경우 

하나의 추상적 개념이 여러가지 구현으로 구체화될 수 있을 때, 대부분은 상속을 통해서 이 문제를 해결합니다. 
추상 클래스로 추상적 개념에 대한 인터페이스를 정의하고, 구체적인 서브 클래스에서 서로 다른 방식으로 이들 인터페이스를 구현합니다. 
그러나 이 방법만으로는 충분한 융통성을 얻을 수 없습니다. 상속은 구현과 추상적 개념을 영구적으로 종속시키기 때문에, 추상적 개념과 구현을 분리해서 재사용하거나 수정, 확장하기가 쉽지 않습니다.

![Bridge Pattern](https://github.com/keepinmindsh/lines_edu/blob/main/assets/bridge_pattern.png)

- 추상적 개념과 이에 대한 구현 사이의 지속적인 종속 관계를 피하고 싶을 때, 이를테면, 런타임에 구현 방법을 선택하거나 구현 내용을 변경하고 싶을 때가 여기에 해당합니다
- 추상적 개념과 구현 모두가 독립적으로 서브클래싱을 통해 확장되어야 할 때, 이때, 가교(Bridge) 패턴은 개발자가 구현을 또 다른 추상적 개념과 연결할 수 있게 할 뿐 아니라, 각각을 독립적으로 확장가능하게 합니다.
- 추상적 개념에 대한 구현 내용을 변경하는 것이 다른 관련 프로그램에 아무런 영향을 주지 않아야 할 때, 즉, 추상적 개념에 해당하는 클래스를 사용하는 코드들은 구현 클래스가 변경되었다고 해서 다시 컴파일 되지 않아야 합니다.

## 실생활에서 활용해보면, 

### Go 에서는 

객체의 생성시점에 정해진 역할에 의해서 프로세스가 동작한다.   

**개발자**는 어떤 API로 개발을 할 수 있기 때문에, 필요한 상황에 따라서 API를 받아서 사용할 수 있게 된다.   

```go 
package designpattern

import "testing"

func Test_BridgeCall(t *testing.T) {
	developer := NewHowardDeveloper(NewHanabankAPI())

	developer.DevelopWithdrawProcess()
}
```

```go 
package designpattern

type (
	BankAPI interface {
		Withdraw()
	}

	Developer interface {
		DevelopWithdrawProcess()
	}
)


type KookminbankAPI struct{}

func (k KookminbankAPI) Withdraw() {}

func NewKookminbankAPI() BankAPI {
	return &KookminbankAPI{}
}

type HanabankAPI struct{}

func (h HanabankAPI) Withdraw() {}

func NewHanabankAPI() BankAPI {
	return &HanabankAPI{}
}

type Howard struct {
	BankAPI
}

func (h Howard) DevelopWithdrawProcess() {

}

func NewHowardDeveloper(bankAPI BankAPI) Developer {
	return &Howard{bankAPI}
}
```

### Java 에서는 

아래의 코드가 만약 패턴이 적용되지 않은 상태에서 모든 코드를 작성해야 했다면, 

- Not Good 

해당 코드는 서비스의 요구사항이 증가함에 따라 SRP 원칙이 깨질수 밖에 없게 된다. 
개발자가 개발할 수 있는 API 의 경우 수에 따라서 그만큼 처리해야할 로직이 한 곳에 모이게 되는 구조가 된다. 

```java 

class Sample {
    public static void WithDrawAPIDevelop(apiType APIType){
        Developer developer = new Developer(); 

        switch (apiType) {
            case HANA: 
                developer.withdrawFroHana()

                // do something...

                // anything will be added later
                break;
            case KOOKMIN:
                developer.withdrawForKookmin()

                // do something...

                // anything will be added later
                break;
        }
    }
}
```

- Better 

브릿지 패턴과 전략 패턴은 유사한 모습을 가지는 것 같지만, 
브릿지 패턴의 경우, 객체의 생성 시점(Compile)에 역할이 정해지며, 이는 객체의 행동시점(Runtime)에 따라 달라지는 전략 패턴과는 차이가 있다. 

```java
// Client에서 Bridge(가교) 패턴을 구현하는 예시입니다.  
class BridgePattern {
    public static void main(String[] args)
    {
        Developer[] bankDevelopers = new Developer[2];

        bankDevelopers[0] = new APIDeveloper(new KookminBankAPI1());
        bankDevelopers[1] = new APIDeveloper(new HanaBankAPI2());

        for (Developer developer : bankDevelopers)
        {
            developer.withdraw();
        }
    }
}
```

```java
interface BankAPI
{
    public void withdraw();
}

// 개발자가 사용할수 있는 국민은행 API 입니다. 
class KookminBankAPI1 implements BankAPI
{
    public void withdraw()
    {
        System.out.printf("API1.withdraw");
    }
}

// 개발자가 사용할 수 있는 하나은행 API 입니다. 
class HanaBankAPI2 implements BankAPI
{
    public void withdraw()
    {
        System.out.printf("API2.withdraw");
    }
}

interface Developer
{
    public void withdraw();   
}

// 개발자는 인출이라는 하나의 프로세스에 대해서 집중해서 처리합니다. 
class APIDeveloper implements Developer
{
    private BankAPI bankAPI;
    public APIDeveloper(BankAPI bankAPI)
    {
        this.bankAPI = bankAPI;
    }
    
    public void withdraw()
    {
      bankAPI.withdraw();
    }
}
```



## Terms 

- 서브클래싱 
> Super Class(슈퍼 클래스)에 구현된 코드와 내부 표현 구조를 Sub Class(하위 클래스)가 이어받습니다.클래스 상속으로 불리기도 하며 하위 클래스에서 슈퍼 클래스에 구현된 코드의 재사용이 가능합니다.


## References 

> 출처: https://epicdevsold.tistory.com/177 [프로그래밍로그:티스토리]  