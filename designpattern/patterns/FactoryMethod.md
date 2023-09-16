# Pattern - Factory Method

## Basic Information 

- 패턴형식 : 생성패턴 

## Gof's Description 

### 의도 

객체를 생성하기 위해서 인터페이스를 정의하지만, 어떤 클래스의 인스턴스를 생성할 지에 대한 결정은 서브 클래스가 내리도록 합니다.   

Factory 하나에 대한 객체의 종류에 따라 인스턴스를 반환할 수 있게 처리합니다.    

### 풀이 

Abstract Factory와 Factory Method는 주로 같이 활용되는데,   
Abstract Factory는 **객체의 생성**에 집중하고, 
Factory Method는 **객체의 구현**에 집중합니다.   

**객체의 구현**은 인터페이스를 통해 추상화된 행위에 대한 실질적은 행위를 정의하는 것인데,   

- 해외 여행을 위해서 숙박할 호텔을 예약한다고 했을 때, 
  - 아고다에서 얘약하는 방식
  - 부킹닷컴에서 예약하는 방식 
  - 익스피디아에서 예약하는 방식 

위의 각각의 방식은 각 사이트의 예약 방법에 따라 다르기 때문에 **숙박을 예약한다**는 실제 구현 방식은 
달라지게 된다. 

### 장점

- 어떤 클래스가 자신이 생성해야 하는 객체의 클래스를 예측할 수 없을 때,
- 생성할 객체를 기술하는 책임을 자신의 서브 클래스가 지정했으면 할 때,
- 객체 생성의 책임을 몇 개의 보조 서브 클래스 가운데 하나에게 위임하고, 어떤 서브 클래스가 위임자인지에 대한 정보를 국소화 시키고 싶을 때

> 팩토리 메소드 패턴을 사용하는 이유는 클래스간의 결합도를 낮추기 위한것입니다.  
> 결합도라는 것은 간단히 말해 클래스의 변경점이 생겼을 때 얼마나 다른 클래스에도 영향을 주는가입니다.  
> 팩토리 메소드 패턴을 사용하는 경우 직접 객체를 생성해 사용하는 것을 방지하고 서브 클래스에 위임함으로써 보다 효율적인 코드 제어를 할 수 있고 의존성을 제거합니다.   
> 결과적으로 결합도 또한 낮출 수 있습니다.   

### 단점 

- 하나틔 클래스 내에 추가적인 기능을 정의해서 사용해야할 경우, 변경의 영향도가 횡으로 확장한다. 
- 비즈니스 로직을 정의하기 위한 메소드를 정의할 때 책임의 범위를 정의하는 부분에서 고려야할 사항이 많다. 

## 실생활의 사례를 통한 패턴의 이해 

### 온라인 예약의 예시 ( With Go )

- 아고다 예약 
- 부킹닷컴 예약 
- 익스페디아 예약 

```go 
package domain

type ReservationType string

const (
	Agoda   ReservationType = "Agoda"
	Booking ReservationType = "Booking"
	Expedia ReservationType = "Expedia"
)

type (
	Reservation interface {
		SchedulePeriod()
		Book()
		PayMoney()
	}
)

```

```go 
package main

import (
	"factory_method/app/usecases"
	"factory_method/domain"
)

func main() {
	reservation := usecases.NewReservation(domain.Agoda)

	reservation.SchedulePeriod()

	reservation.Book()

	reservation.PayMoney()
}

```

```go 
package usecases

import (
	"factory_method/app/service/agoda"
	"factory_method/app/service/booking"
	"factory_method/app/service/expedia"
	"factory_method/domain"
)

func NewReservation(rsvnType domain.ReservationType) domain.Reservation {
	switch rsvnType {
	case domain.Agoda:
		return agoda.NewAgodaReservation()
	case domain.Expedia:
		return expedia.NewExpediaReservation()
	case domain.Booking:
		return booking.NewBookingReservation()
	}

	return nil
}

```


```go 
// agoda factory method implementation
package agoda

import (
	"factory_method/domain"
	"fmt"
)

type Reservation struct{}

func NewAgodaReservation() domain.Reservation {
	return &Reservation{}
}

func (r Reservation) SchedulePeriod() {
	fmt.Println("Agoda scheduled!")
}

func (r Reservation) Book() {
	fmt.Println("Agoda booked!")
}

func (r Reservation) PayMoney() {
	fmt.Println("Agoda payMoney!")
}

// booking factory method implementation
package booking

import (
	"factory_method/domain"
	"fmt"
)

type Reservation struct{}

func NewBookingReservation() domain.Reservation {
	return &Reservation{}
}

func (r Reservation) SchedulePeriod() {
	fmt.Println("Booking.com scheduled!")
}

func (r Reservation) Book() {
	fmt.Println("Booking.com booked!")
}

func (r Reservation) PayMoney() {
	fmt.Println("Booking.com paymoney!")
}

// expedia factory method implementation
package expedia

import (
	"factory_method/domain"
	"fmt"
)

type Reservation struct{}

func NewExpediaReservation() domain.Reservation {
	return &Reservation{}
}

func (r Reservation) SchedulePeriod() {
	fmt.Println("Expedia scheduled!")
}

func (r Reservation) Book() {
	fmt.Println("Expedia booked!")
}

func (r Reservation) PayMoney() {
	fmt.Println("Expedia paymoney!")
}
```

### 온라인 예약의 예시 ( With Kotlin )

```java
package domain

enum class ReservationType {
    Agoda, Booking, Expedia
}

interface Reservation {
    fun SchedulePeriod()
    fun Book()
    fun PayMoney()
}
```

```java 
import app.usecases.Factory
import domain.ReservationType

fun main(){
    val reservation = Factory.GetReservationWay(ReservationType.Agoda)

    reservation.SchedulePeriod()

    reservation.Book()

    reservation.PayMoney()
}
```

```java
package app.usecases

import app.service.agoda.Agoda
import app.service.booking.Booking
import app.service.expedia.Expedia
import domain.Reservation
import domain.ReservationType

class Factory {
    companion object {
        fun GetReservationWay(rsvnType: ReservationType): Reservation {
            when (rsvnType) {
                ReservationType.Agoda -> return Agoda()
                ReservationType.Booking -> return Booking()
                ReservationType.Expedia -> return Expedia()
                else -> throw java.lang.RuntimeException("Error!")
            }
        }
    }
}
```

```java
package app.service.agoda

import domain.Reservation

class Agoda : Reservation {
    override fun SchedulePeriod() {
        print("Agoda Scheduled!")
    }

    override fun Book() {
        print("Agoda Booked!")
    }

    override fun PayMoney() {
        print("Agoda Pay Money!")
    }
}

package app.service.booking

import domain.Reservation

class Booking : Reservation{
    override fun SchedulePeriod() {
        print("Booking.com Scheduled!")
    }

    override fun Book() {
        print("Booking.com Booked!")
    }

    override fun PayMoney() {
        print("Booking.com Pay Money")
    }
}

package app.service.expedia

import domain.Reservation

class Expedia : Reservation {
    override fun SchedulePeriod() {
        print("Expedia Scheduled!")
    }

    override fun Book() {
        print("Expedia Booked!")
    }

    override fun PayMoney() {
        print("Expedia Pay Money!")
    }
}
```

### 작성해보면 좋을 만한 예시 문제 

```
모든 유닛의 추상화된 행동 
- 이동 
- 공격
- 아드레날린 사용 


마린의 훈련 과정
 - 마린은 기본적으로 배워할 기능에 대해서 Barrack을 통해서 배울 수 있어야 한다.
 - 공격시 총을 사용한다. 
 - 이동시 1칸 씩 이동이 가능하다. 
 - 아드레날린 사용시 공격속도가 빨라진다. 

파이어뱃의 훈랸 과정 
  - 파이어뱃은 기본적으로 배워할 기능에 대해서 Barrack을 통해서 배울 수 있어야 한다.
  - 공격시 화염 방사기를 사용한다. 
  - 이동시 1칸 씩 이동이 가능하다. 
  - 아드레날린 사용시 공격속도가 빨라진다. 

위의 예시를 바탕으로 실제 코드 레벨에서 작성해볼 것! 
```