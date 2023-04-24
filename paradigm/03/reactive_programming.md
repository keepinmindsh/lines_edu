# 웹 개발 패러다임의 거대한 변화 "Reactive"

***

 막힘 없이 흘러다니는 데이터(이벤트)를 통해 사용자에게 자연스러운 응답을 주고, 규모 탄력적으로 리소스를 사용하며 실패에 있어서 유용하게 대처한다.

# Reactive 패러다임 

***

"막힘없이 흘러다니는 데이터(이벤트)를 통해 사용자에게 자연스러운 응답을 주고, 규모 탄력적으로 리소스를 사용하며 실패에 있어서 유연하게 대처한다."  
"모든 지점에서 블럭되지 않게 하자."  

# Reactive 

***

- "외부의 어떤 이벤트나 데이터가 발생하면 거기에 대응하면 방식으로 동작하는 프로그램을 만드는 것"  
- "데이터 플로우와 상태 변경을 전파한다는 생각에 근간을 둔 프로그래밍 패러다임"  
- "막힘 없이 흘러다니는 데이터(이벤트)를 통해 사용자에게 자연스러운 응답을 주고, 규모 탄력적으로 리소스를 사용하며 실패에 있어서 유연하게 대처한다."   

![](https://keepinmindsh.github.io/lines/assets/img/Reactive.png)

#### Responsive : 사용자에 대한 반응  

  시스템이 적시에 응답합니다. 응답성은 사용성과 기능성의 기반인데, 그것보다 더 응답성은 문제에 대해서 빠르게 감지하는 것과 효율성을 다루는 것에 초점을 둡니다. 반응성(응답성)이 좋은 시스템은 속도와 일정한 응답성을 제공하고 , 신뢰할 수 있는 상향 경계를 수립하므로써, 일정한 품질의 서비스를 제공하는 것에 있습니다. 이 일관성있는 행동은 에러 처리를 간편화 하고, 사용자 신뢰를 구축하고, 앞선 상호작용을 장려합니다.

#### Scalable(Elastic) : 부하에 대한 반응

  시스템은 다양한 작업하에 응답성을 유지합니다. Reactive System은 서비스에 제공되기 위한 입력을 할당한 자원을 증가시키거나 감소시키는 것으로써 입력 량의 변화에 응답할 수 있다. 이것은 경합 포인트나 병목현상을 가지지 않게 설계되었으며, 공유하고, 컴포넌트를 복제하고, 입력을 분산할 수 있도록 하는 결과를 의도합니다. Reactive System은 응답성 뿐만 아니라 예측 가능성을 지원하는데 이는 실시간 성능 측정을 제공하여 알고리즘을 조정할 수 있습니다. 상용 하드웨어 및 소프트웨어 플랫폼에서 비용 효율적인 방식으로 탄력성을 얻을 수 있습니다.

#### Resillent : 실패 상황에 대한 반응

  시스템은 장애가 발생하더라도 응답성을 유지합니다. 이는 고 가용성, Mission Critical 시스템에 적용됩니다. 탄력성은 복제, 유지, 격리 및 위임에 의해 획득할 수 있습니다. 장애는 각각의 컴포넌트에 포함되어 있기 때문에 각각으로 부터 컴포넌트가 고립하되는 것으로 시스템 전체에 영향일 미치지 않고, 시스템의 일부가 장애 및 복구되는 것을 보장할 수 있습니다. 각각의 컴포넌트의 복수는 새로운 컴포넌트에게 위임되고, 고가용성은 필요에 따라 복제 따라서 보장됩니다. 고객의 기능은 장애를 처리함으로써 부담을 받지 않아도 됩니다.

#### Event-Driven( Message-Driven ) : 이벤트에 대한 반응

  Reactive 시스템은 Location Transparency, Isolation, Loose Coupling을 보장하는 컴포넌트들 사이의 경계를 관리하기 위해서 비동기적인 Message 전달 (Asynchronouse message-passing)에 의존합니다. 이 경계는 메세지로서 장애를 위임하기 위한 의도를 제공합니다. 명시적인 Message 전달을 이용하면 부하관리, 탄력성, 흐름제어 및 시스템에서의 메세지 큐 모니터링, 필요에 따라 Back Pressure를 적용하는 것을 가능하게 합니다. 통신수단으로의 Location Transparent Message는 동일한 구조와 의미의 단일 호스트 또는 클러스터와 동작하기 위한 장애의 관리를 가능하게 합니다. Non-Blocking Communication은 수신자로 하여금 활성 상태에서만 자원을 소모할 수 있게 하여 시스템의 오버헤드를 줄일 수 있습니다.


# Functional Reactive Programming 

***
  
  Functional Reactive Programming 은 위에서 살펴본 Reactive Programming을 Functional Programming의 원리를 통해 구현하는 것을 말합니다. 쉽게 말하자면, 비동기적인 데이터 처리를 간단한 함수를 통해 수행하는 프로그래밍을 말합니다. RP에 FP에서 제공하는 함수를 활용하는 것입니다.   

  여기에서 Functional Programming에 대해서 먼저 알아보고 가야하는데, Functional Programming이란 어떤 문제를 해결하는데 있어서 그 과정이나 공식에 매몰되기보다는 그냥 이미 만들어진 함수를 활용하는 방식을 말합니다. 다만, 무조건 활용하는 것이 아니라 함수 자체가 '숨겨진 input과 output'이 없도록 설계되어야 하는 것이 전제 조건입니다.

```java


// side-cause와 side-effect가 존재하는 함수 
// 함수를 콜할 때 implementation detail을 살펴보지 못한다면 무엇이 어떻게 변할지 알 수 없음                                     
func add() {
  number = 5
  letter = "S"
  title = title + " \ (number) " + letter 
}                                         

// 숨겨진 input/ouput이 없는 함수
func add(numberOnt: Int, numberTwo: Int) -> Int {
  return numberOne + numberTwo
}

```

 결국 Functional Programming이라는 것은, 결과에 집중하는 실용적인 함수를 정의하고 활용하되, 이러한 함수안에 숨겨진 input과 output이 최대한 없을 수 있도록 선언하는 프로그래밍 패러다임입니다.  

# Reactive Streams 

***

  Volume이 정해져 있지 않는 실시간의 데이터 스트림을 처리하는 것은 비동기 시스템에서 필요합니다. 
그리고 대부분의 이슈는 해당 스트팀의 최종 목적이가 과부하가 발생하지 않도록 데이터 자원을 조심스럽게 통제하기 위한 것입니다.
협력하는 다수의 네트워크 호스트 또는 다중 CPU 코에 대해서 컴퓨팅 자원의 병렬 사용을 활성하기 위해서 비동기적으로 필요합니다.
Reactive Stream의 주요 목표는 스트림 데이터의 변화를 비동기 영역에서 통제하는 것입니다. 다른말로 하면 , 내부 압력은 해당 모델의 완벽한 부분인데, 
연결을 맺기위한 스레드들 사이에서 중재하기 위한 큐들을 허락하기 위해서입니다. 
Reactive Stream은 서로 다른 API Component 사이의 스트림 데이터를 중재하기 위한 것에 대해서만 오로지 고려하고 있습니다.

- Reactive Stream : <https://www.reactive-streams.org/>  
- Reactive Menifesto : <https://www.reactivemanifesto.org>  

 우리가 앞으로 Reactive Programming을 다루기 위해 사용할 Project Reactor 에 대해서 이해하기 전에 반드시 알아야할 내용이 있습니다.
그것은 현재 Rective Programming을 지원하고 있는 Reactor, RxJava, RxSwift, RxJS, ... 등등의 근간이 되는 *Reactive Stream* 이라는 것입니다. 

Reactive Stream은 위의 링크(Reactive Stream 의 링크)에 영어로 설명되어 있는 것을 간단하게 정의해보면, 

 **비동기적인 stream 프로세싱을 논-블록킹 방식의 배압(Back Pressure)를 이용해 표준을 제공합니다.**

  - 위의 굵은 글씨의 내용이 Reactive Stream에서 가장 중요한 내용이며 이를 구현하기 위한 기본적인 개념을 하나씩 알아보겠습니다. 

Reactive Stream에 대한 GitHub : <https://github.com/reactive-streams/reactive-streams-jvm/tree/v1.0.3>

#### Reactive Stream 의 수학적 개념 

 - 가져오다 ( Pull )

```java

// List Type 은 Iterable 의 sub type 이다.
List<Integer> list = Arrays.asList(1, 2, 3, 4, 5);

// 이것은 Iterable 이기 때문에 여기에서 List 를 사용할 수 있는 것임
for (Integer i : list) { 
    System.out.println(i);
}

// Iterable 로 받아도 문제가 없음
Iterable<Integer> iter = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);

// 이터레이터를 통해서 반복적으로 사용할 수 있게 하려면 1.5 이전에는 아래와 같이 사용해야 했다. 
for (Iterator<Integer> it = iter.iterator(); it.hasNext(); ) {
    System.out.println(((Iterator) it).next());
}

// Lamda 의 관심은 함수 하나
// 1부터 10까지를 계속해서 호출해서 쓸 수 있음
Iterable<Integer> iter1 = () ->
  new Iterator<Integer>() {  

      int i = 0;
      final static int MAX = 10;

      @Override
      public boolean hasNext() {
          return i < MAX;
      }

      @Override
      public Integer next() {
          return ++i;
      }
  };

// 1.5 이상은 아래와 같이 사용 가능하다. 
for (Integer i : iter1) {
    System.out.println(i);
}

```

- 넣다 ( Push )

```java

public class Ob {

  public static void main(String[] args) {        
    // notifyObservers 호출이되면 update 가 호출이 된다.
    Observer ob = new Observer() {
        @Override
        public void update(Observable o, Object arg) {
            System.out.println(arg);

            // Thread는 풀에 있는 스레드를 받게 처리되는 것은 ExecutorService 의 Thread를 메인스레드에서 분리해서 사용할 수 있다.
            System.out.println(Thread.currentThread().getName() + " " + arg);
        }
    };

    IntObservable observable = new IntObservable();

    observable.addObserver(ob);

    // Event Driven에서 주로 사용하는 방식인 감시자 패턴
    // 메인 스레드에서 동작하게 하고 있음
    observable.run();
  }

  static class IntObservable extends Observable implements Runnable {

    @Override
    public void run() {
        for (int i = 0; i < 10; i++) {
            setChanged();
            // 아래의 두개는 Duality 라고 이야기 할 수 있다.
            notifyObservers(i);    // push  method(DATA)
            // int i = it.next()   // pull  DATA method(void)
        }
    }
  }
}

```
###### Iterable <------> Observable ( 쌍대성 ) - Duality : 비슷한 구조를 가진다.

- Iterable 은 Pulling 방식
- Observable 은 Push 방식

Observable 와 Iterable 이 쌍대성을 가지는 구조  
처리하려는 기능은 같지만 처리 과정에 있어서 서로 상반되는 부분을 duality라고 하면 좀 더 이해가 쉬울 수 있다.  
{: .notice--info}

쌍대성(雙對性; duality)은 수학과 물리학에서 자주 등장하는 표현이다. 보통 어떤 수학적 구조의 쌍대(雙對; dual)란 그 구조를 ‘뒤집어서’ 구성한 것을 말하는데, 엄밀한 정의는 세부 분야와 대상에 따라 각각 다르다. 쌍대의 쌍대는 자기 자신이므로 어떤 대상과 그 쌍대는 서로 일종의 한 ‘켤레’를 이룬다고 할 수 있으며, 이를 쌍대관계(雙對關係)라고 한다.
{: .notice--info}


#### Reactive Stream 의 API Component 

<https://github.com/reactive-streams/reactive-streams-jvm/blob/v1.0.0/README.md#specification>

Reactive Stream에서 제공하는 기본 4가지 컴포넌트를 알아보고자합니다. 아래의 컴포넌트들은 초기에 Observer 패턴을 기반으로하여 Observer/Observable에서 부족한 부분을 보완한 컴포넌트입니다. 

- Publisher
- Subscriber
- Subscription
- Processor

***

- Subscriber
  - Subscriber는 Observer 입니다.

```java

// 아래의 4개의 메소드는 반드시 implement 되어야 합니다.                                   
public interface Subscriber {
  public void onSubscribe(Subscription s);
  public void onNext(T t);
  public void onError(Throwable t);
  public void onComplete();
}   

```

![](https://keepinmindsh.github.io/lines/assets/img/subscriber_process.png){: .align-center}

  - onSubscribe : 최초 호출되는 메소드, Subscriber를 사용할 때는 무조건 처음에 호출해야합니다.
  - onNext : 기존의 Observer에서 update와 같은 역할을 합니다. 데이터를 받을 때 사용합니다.
  - onComplete : 완료 되었을 때,
  - onError : 에러가 발생했을 때,

- Publisher 

  - Publisher는 Observable 입니다. Subscriber는 Publisher의 subscribe를 통해 등록합니다. 

```java

public interface Publisher {
  public void subscribe(Subscriber<? super T> s);
}

```

- Subscription

```java

public interface Subscription {
  public void request(long n);
  public void cancel();
}

```

request는 long 타입의 파라미터를 받고 있는데 Subscriber가 이 메소드를 통해 요청을 하게 됩니다. 만약 5개의 데이터를 필요하다고 가정했을 때, reqeust 에 5를 넣어서 호출하면 Subscription은 5개를 호출하게 됩니다. 즉, 10개의 데이터가 있을 때, reqeust가 5를 받아 처리한다면 5개 -> 5개 를 보내줄 수 있게 처리합니다.
이는 publisher를 통해서 들어오는 데이터 스트림을 request를 이용해 subscriber에서 처리하는데 적절한 범위로 처리될 수 있게 제어를 할 수 있습니다. 이를 Reactive Stream에서 가장 중요한 Back-Pressure를 제어할 수 있는 방법입니다. 

#### Reactive Basic Flow : Publisher - Subscriber - Subscription

![](https://keepinmindsh.github.io/lines/assets/img/reactive_basic_flow.png)

- Publisher에 Subscriber가 구독(등록)되면, Publisher가 실행(subscribe)될 때 Publisher 통해서 데이터(스트림) 또는 시퀀스를 Subscriber로 전달하게됩니다. 
- 이때 Publisher는 Subscriber에 정의된 OnSubscribe()를 호출하고, Subscriber는 request(n)를 호출하여 몇개의 데이터를 보내달라고 Publisher에게 Subscription을 통해서 요청하게 됩니다. 
- Subscrition을 통해 정의된 요청 갯수에 의해서 request 메소드 내에서 Subscriber의 onNext, onError, OnComplete를 제어할 수 있습니다.
- Subscriber가 동작하던 도중에 장애/에러 발생으로 인하여 처리를 중단해야할 때 subscription 객체를 이용해서 cancel을 호출 하고 Flag를 관리한다면, 해당 Flow 전체를 중단할 수 있습니다. 

![](https://keepinmindsh.github.io/lines/assets/img/reactive_interaction.png)
