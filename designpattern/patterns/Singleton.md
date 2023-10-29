# Pattern - Singleton

## Basic Information

- 패턴 형식 : 생성 패턴 

## Gof's Description

오직 한 개의 클래스 인스턴스 만을 갖도록 보장하고, 이에 대한 전역적인 접근점을 제공합니다.

### 풀이

- 기본적인 전제 
    - 어떤 클래스는 정확히 하나의 인스턴스만을 갖도록 하는 것이 좋습니다. 
    - 시스템에 많은 프린터가 있다 하더라도, 프린터 스풀은 오직 하나여야 합니다. 
    - 파일 시스템도, 윈도우 관리자도 오직 하나여야 합니다.
    - 한 회사에서는 하나의 회계 시스템만 운영될 것입니다 

### 필요할 경우 

- 클래스의 인스턴스가 오직 하나여야 함을 보장하고, 잘 정의된 접근 점으로 모든 사용자가 접근할 수 있도록 할 때,
- 유일한 인스턴스가 서브클래싱으로 확장되어야 하며, 사용자는 코드의 수정없이 확장된 서브 클래스의 인스턴스를 사용할 수 있어야 할 때,  

## 실생활에서의 사용 예시 

우리가 알고 있는 프린터기는 다수 컴퓨터의 출력 요청을 전달받아 프린터를 출력합니다. 각 다수의 컴퓨터에게 하나의 프린터 객체만 존재해야 함을 의미합니다. 

### Go 에서는 

#### 호출 테스트

```go 
package test

import (
  "fmt"
  "sync"
  "testing"
)

func Test_Service_Singleton(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {

			fmt.Printf("Address : %v \r\n", NewGetPrinterSingleton_Way3())

			wg.Done()
		}()
	}

	wg.Wait()
}
```

#### 신규 객체를 반환하는 방식  

```go 
package singleton 

func NewGetPrinterNotSingleton() *domain.Printer {
	return NewPrinter(domain.SAMSUNG_PRINTER)
}
```

#### 싱글턴 패턴에 맞춘 객체 생성 방식 1 ( 잘못된 케이스 )

```go 
package singleton

import (
  "fmt"
  "printer/domain"
  "sync"
)

var SingletonPrinter *domain.Printer

func NewGetPrinterNotSingleton_Way1() *domain.Printer {
	if SingletonPrinter == nil {
		SingletonPrinter = NewPrinter(domain.SAMSUNG_PRINTER)
		return SingletonPrinter
	} else {
		return SingletonPrinter
	}
}
```

#### 싱글턴 패턴에 맞춘 객체 생성 방식 2 ( 올바른 케이스 ) 

```go 
package singleton

import (
  "fmt"
  "printer/domain"
  "sync"
)

var once sync.Once

func NewGetPrintSingleton_Way2() *domain.Printer {
	once.Do(func() {
		SingletonPrinter = NewPrinter(domain.SAMSUNG_PRINTER)
	})

	return SingletonPrinter
}
```

#### 싱글턴 패턴에 맞춘 객체 생성 방식 3 ( 올바른 케이스 )

```go 
package service

import (
  "fmt"
  "printer/domain"
  "sync"
)

var SingletonPrinter *domain.Printer

var lock = &sync.Mutex{}

func NewGetPrinterSingleton_Way3() *domain.Printer {
	if SingletonPrinter == nil { // Golang#1
		lock.Lock()
		defer lock.Unlock()
		if SingletonPrinter == nil { // Golang#2
			SingletonPrinter = NewPrinter(domain.SAMSUNG_PRINTER) // Golang#3
			fmt.Println("Create new printer")
			return SingletonPrinter
		} else {
			fmt.Println("Already printer was created")
			return SingletonPrinter
		}
	} else {
		fmt.Println("Already printer was created")
		return SingletonPrinter
	}
}
```

### Java 에서는 

```java 
package com.moong.realiticdesignpattern;
    
public class Singleton {
      private static Singleton instance = null;
      public static Singleton instance(){
          if( instance == null ){
              instance = new Singleton();
          }
          return instance;
      }
      //... 
}  
```

![Generalization](https://github.com/keepinmindsh/lines_edu/blob/main/assets/thread_problem.png)

- 위의 Singleton 예제의 문제점
  - 스레드 1이 instance()를 호출하고 4번째 줄을 검사하고 있다. 그런데 다음 줄로 넘어가기 전에 클록 틱에 의해 선점되었다.
  - 스레드 2가 instance()를 호출하고 메소드 전체를 실행한다. 인스턴스가 생성되었다.
  - 스레드 1이 잠에서 깨어나 인스턴스가 아직 존재하지 않는다고 생각하고(이 스레드는 전에 중지되기 전에 null 테스트를 마쳤다
  - Singleton의 두 번째 인스턴스를 생성한다


#### 또다른 대안은 

```java 
public class Singleton2 {
  private volatile static Singleton2 single;
  public static Singleton2 getInstance(){
      if (single == null) {
          synchronized(Singleton2.class) {
              if (single == null) {
                  single = new Singleton2();
              }
          }
      }
      return single;
  }
  private Singleton2(){
  }
}   
```

#### static 초기화를 사용하면, 

static 초기화를 사용하면 오버헤드를 격지 않으면서도 동기화 문제를 해결할 수 있다.

```java  
class Singleton3 {
  private static Singleton3 instance = new Singleton3();
  public instance() { return instance; }
  //…
} 

```

#### 초기화 대상에 대해서 파악이 어려울 경우, 늦은 초기화가 필요할 경우 

```java 
class Connection {
  private static URL server;
  public static void pointAt( URL serverUrl){
    server = serverUrl;
  }

  private Connection(){
      // ...
      
      // ...
  }

  private static Connection instance;
  public synchronized static Connection getInstance(){
      if(instance == null) {
          instance = new Connection();
      }
      return instance;
  }
}   
```

### Kotlin 에서는

```kotlin
package com.printer

import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    repeat(50_000) {
        launch {
            val printerInstance = SingletonPrinter.getPrinterInstance()

            print("Memory Address is $printerInstance \r\n")
        }
        launch {
            val printerInstance = SingletonPrinter.getPrinterInstance()

            print("Memory Address is $printerInstance \r\n")
        }
        launch {
            val printerInstance = SingletonPrinter.getPrinterInstance()

            print("Memory Address is $printerInstance \n")
        }
        launch {
            val printInstance = SingletonPrinter2.getPrintInstance()

            print("Memory Address is $printInstance \n")
        }
        launch {
            val printInstance = SingletonPrinter2.getPrintInstance()

            print("Memory Address is $printInstance \n")
        }
        launch {
            val printInstance = SingletonPrinter2.getPrintInstance()

            print("Memory Address is $printInstance \n")
        }
    }
}
```

```kotlin 
class SamsungPrinter : Printer {
    override fun PutPaperIn() {
        print("Put Paper In!")
    }

    override fun Print() {
        print("Print!")
    }
}

interface Printer {
    fun PutPaperIn()
    fun Print()
}
```

#### object를 사용하는 방식 

클래스 외부에서 선언하는 방식으로 별도의 객체 생성 없이 바로 호출하여 사용할 수 있다. Object는 접근시점에 하나만 생성되므로 별도의 객체 생성 없이 바로 호출할 수 있다. 

```kotlin 
package com.printer

object SingletonPrinter {
    private val printer = SamsungPrinter()

    fun getPrinterInstance() : Printer {
        return printer
    }
}
```

#### companion object를 사용하는 방식 

companion object로 생성한 파라미터가 있는 싱글톤 클래스는 객체를 생성할 수 있다. 여러 개의 객체를 생성할 수 있지만, 클래스의 메모리 주소 값은 동일하다. 

```kotlin 
package com.printer

class SingletonPrinter2 private constructor() {
    companion object {
        private var printer: Printer? = null

        fun getPrintInstance(): Printer {
            return printer ?: synchronized(this) {
                printer ?: SamsungPrinter()
            }
        }
    }
}
```

## 고민해볼까? 

### Double-Checked Locking ( 사용하지 말자! )

위의 예시 코드들 중에서 Golang#3 부분은 코드상으로는 한줄로 표현되어있지만, 실제로는 새로운 객체를 위한 메모리를 할당하고  할당된 메모리 주소를 변수에 저장하고  만들어진 객체를 초기화하는 여러 단계를 거친다.
조건이 하나일 경우에는 여러개의 스레드가 Golang#1 의 첫번째 조건을 다같이 통과 한 것조차 막을 수 없다.

조건이 두개인 경우에는, 여러개의 스레드가 Golang#1 의 첫번째 조건을 다같이 통과 한 것을 mutex 로 조절할 수 있다.   
적어도 Golang#2 영역의 critical section 에는 하나의 스레드가 접근하는 것을 보장한다.   
하지만, 메모리를 할당하고, 변수에 메모리의 주소를 저장한 것이 실제 메모리에 반영이 되었는지는 보장할 수가 없다.  
따라서, 처음으로 lock 을 획득한 스레드가 변수 초기화를 하는 것에 대한 보장은 되지만, 다음 스레드가 진입했을때, 해당 객체가 null 이 아니라고 해서 초기화 작업을 끝냈을 것이라는 보장이 없는 것이다.    

> [Double Check Locking](https://herdin.github.io/2020/12/25/about-double-check-locking)

### Memory Visibility ( 메모리 가시성 ) 와 Memory Barrier ( 메모리 장벽 )

- [메모리 가시성과 메모리 장벽](https://keepinmindsh.github.io/lines/parrallel/parrallel-01/#%EA%B0%80%EC%8B%9C%EC%84%B1)

### 단일 연산과 복합 연산의 개념 

- [단일연산과 복합연산, 그리고 원자성](https://keepinmindsh.github.io/lines/parrallel/parrallel-01/#%EB%8B%A8%EC%9D%BC-%EC%97%B0%EC%82%B0)

### 동시성 제어 

- [동시성 제어](https://velog.io/@ha0kim/%EB%8F%99%EC%8B%9C%EC%84%B1-%EC%A0%9C%EC%96%B4)