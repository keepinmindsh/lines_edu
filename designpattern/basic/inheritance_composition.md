# 상속과 합성 

우리가 객체와 객체 사이의 연관관계를 가지는 방법은 **상속**과 **합성**을 통해서 이루어진다. 

**상속** 의 경우, 프로그래밍 언어에 따라서 지원되지 않은 개념일 수도 있다.
여기에서는 Java를 바탕으로 상속 및 합성의 개념을 설명하고, 다른 언어로 확장시켜 보겠다. 

## 상속의 개념

```java 
package com.lines.sample;

public class Inheritance {
    public static void main(String[] args) {
        Howard howard = new Howard();

        howard.DrinkDouble();

        howard.Drink();

        System.out.printf("Drunk Status: %s",howard.DrunkStatus());
    }
}


class DrinkingSkill {

    public int soju = 0;

    public void DrinkDouble(){
        soju += 2;
    }

    public int DrunkStatus(){
        return soju;
    }
}


class Howard extends DrinkingSkill {

    public void Drink(){
        super.soju += 1;
    }
}
```

## 합성의 개념 