# Proxy 

## Basic Information

- 패턴형식 : 목적패턴 

## Gof's Description

다른 객체에 대한 접근을 제어하기 위한 대리자 또는 자리 채움자 역할을 하는 객체를 둡니다.

### 풀이

어떤 객체에 대한 접근을 제어하는 한 가지 이유는 실제로 그 객체를 사용할 수 있을 때까지 객체 생성과 초기화에 들어가는 비용 및 시간을 물지 않겠다는 것 입니다.    
그래픽 객체를 문서 안에 넣을 수 있는 문서 편집기의 예를 다시 생각해 봅시다. 래스터 이미지와 같은 그래픽 객체를 생성 하려면 비용이 많이 듭니다.   
그러나 문서를 읽어내는 것은 이런 그래픽 객체가 있든 없든 매우 빠르게 진행 되어야 합니다. 또한 문서가 읽히는 그 시점 에서 모든 내용을 다 읽어올 필요는 없습니다.   
이미지의 모든 내용이 한꺼번에 한 문서에 다 보일 필요는 없기 때문 입니다.

### 결과 

프록시 패턴은 어떤 객체에 접근할 때 추가적인 간접화 통로를 제공 합니다. 이렇게 추가된 간접화 통로는 프록시의 종류에 따라서 여러 가지 쓰임새가 있습니다.   

- 원격지 프록시는 객체가 다른 주소 공간에 존재 한다는 사실을 숨길 수 있습니다.
- 가상 프록시는 요구에 따라 객체를 생성하는 등 처리를 최적화 할 수 있습니다.
- 보호용 프록시 및 스마트 참조자는 객체가 접근할 때 마다 추가 관리를 책임 집니다. 객체를 생성할 것인지 삭제할 것 인지를 관리합니다.
- 기록 시점 복사 : 이 최적화는 요구가 들어올 때만 객체를 생성하는 개념과 관련 있는데, 덩치가 크고 복잡한 객체를 복사하려면 비용이 만만치 않습니다. 만약, 사본이 변경되지 않고 원본과 똑같다면, 굳이 이 비용을 물 필요가 없습니다. 프록시를 사용해서 복사 절차를 미룸으로써, 사본이 수정될 때만 실제 복사 비용을 물게 만드는 것 입니다.
- 프록시에서 중요한 부분 중의 하나는 흐름제어만 할 뿐 결과 값을 조장하거나 변경시키면 안됩니다.

## 실제로 사용할 수 있는 코드 예시 

### Java 

```java

package DesignPattern.gof_proxy.sample01;

public class OrderMain {

    public static void main(String[] args) throws Exception {

        OrderExecutor orderExecutor = new OrderExecutorProxy();

        orderExecutor.callOrder("커피요청");
    }
}
```

```java 
package DesignPattern.gof_proxy.sample01;

public interface OrderExecutor {
    public void callOrder(String requestName) throws Exception;
}

package DesignPattern.gof_proxy.sample01;

public class CoffeOrder implements OrderExecutor {
    public void callOrder(String requestName) throws Exception {
        System.out.println(requestName + " is waiting for receiving result.");
    }
}

package DesignPattern.gof_proxy.sample01;

public class OrderExecutorProxy implements OrderExecutor {

    private OrderExecutor orderExecutor;

    public OrderExecutorProxy(){
        orderExecutor = new CoffeOrder();
    }

    public void callOrder(String requestName) throws Exception {
        System.out.println("커피를 요청 하기 위한 사전 작업 진행 !");

        orderExecutor.callOrder(requestName);

        System.out.println("커피를 전달 받아 추가 작업 진행 !");
    }
}                                           

```