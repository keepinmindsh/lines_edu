# Pattern - Adapter 

## Basic Information 

- 패턴형식 : 목적패턴 

## Gof's Description 

클래스의 인터페이스를 사용자가 기대하는 인터페이스 형태로 적응(변환)시킵니다. 서로 일치하지 않는 인터페이스를 갖는 클래스들을 함께 동작시킵니다. 

## 필요한 경우 

- 기존 클래스를 사용하고 싶은데 인터페이스가 맞지 않을 때,
- 아직 예측하지 못한 클래스나 실제 관련되지 않은 클래스들이 기존 클래스를 재사용하고자 하지만, 이미 정의된 재사용 가능한 클래스가 지금 요청하는 인터페이스를 꼭 정의하고 있지 않을 때, 
- 이미 만든 것을 재사용하고자 하나 이 재사용 가능한 라이브러리를 수정할 수 없을 때, 

### 풀이

![adapter_pattern_01](https://github.com/keepinmindsh/lines_edu/blob/main/assets/adapter_pattern01.png)

- Target : 사용자가 사용할 응용 분야에 종속적인 인터페이스를 정의하는 클래스 
- Client : Target 인터페이스를 만족하는 객체와 동작할 대상을 처리
- Adaptee : 인터페이스의 적응이 필요한 기존 인터페이스를 정의하는 클래스
- Adapter : Target 인터페이스에 Adaptee의 인터페이스를 적응시키는 클래스

![adapter_pattern_02](https://github.com/keepinmindsh/lines_edu/blob/main/assets/adapter_pattern02.png)

- Adapter 클래스는 하나만 존재해도 수많은 Adaptee 클래스 등과 동작할 수 있다. 
- Adapter 객체가 포함하는 Adaptee에 대한 참조자는 Adaptee의 인스턴스를 관리할 수도 있고, Adaptee 클래스를 상속받는 다른 서브 클래스들의 인스턴스도 관리할 수 있다. 그러므로 하나의 Adapter 클래스로 모든 Adaptee 클래스와 이를 상속 받는 서브클래스 모두를 이용할 수 있게 된다.

> [UML 이해하기](https://github.com/keepinmindsh/lines_edu/blob/main/designpattern/basic/uml.md)

## 실생활에서의 사용 예시 

### Golang 에서는

```go 
package domain

type (
	V110ElectricAdapter interface {
		UseElectric()
	}
	V220ElectricAdapter interface {
		UseElectric()
	}
)
```

```fo 
package usecases

import (
	"adapter/domain"
	"fmt"
)

type V110 struct {
}

func (v V110) UseElectric() {
	fmt.Println("110 볼트의 전기 제품을 사용할 수 있습니다.")
}

func NewV110() domain.V110ElectricAdapter {
	return &V110{}
}
```


```go 
package usecases

import (
	"adapter/domain"
	"fmt"
)

type V220 struct {
}

func (v V220) UseElectric() {
	fmt.Println("220 볼트의 전기 제품을 사용할 수 있습니다.")
}

func NewV220() domain.V220ElectricAdapter {
	return &V220{}
}
```

```go 
package usecases

import (
	"adapter/domain"
	"fmt"
)

type VoltageAdapter struct {
	Adapter domain.V220ElectricAdapter
}

func NewVoltageAdapter(v220 domain.V220ElectricAdapter) domain.V110ElectricAdapter {
	return &VoltageAdapter{Adapter: v220}
}

func (v VoltageAdapter) UseElectric() {

	fmt.Println("V110 볼트의 전압을 V220의 기기를 사용할 수 있게 전압을 조정합니다. ")

	v.Adapter.UseElectric()
}
```


```go 
package main

import "adapter/app/usecases"

func main() {

	v110 := usecases.NewV110()
	v110.UseElectric()

	v220 := usecases.NewV220()
	v220.UseElectric()

	adapter := usecases.NewVoltageAdapter(v220)
	adapter.UseElectric()
}
```

### Java 에서는 

```java 
package DesignPattern.gof_adapter.sample01;

public class Electric {

    public static void main(String[] args) {

        V220 v220Product = new KoreaElecticProduct();

        V110 v110Adatper = new VoltageAdater(v220Product);

        v110Adatper.useElectric();
    }
}  
```

```java 
// 전압 V110                               
package DesignPattern.gof_adapter.sample01;

public interface V110 {
    public void useElectric();
}

// 전압 V110 용 베트남 제품   
package DesignPattern.gof_adapter.sample01;

public class VietnamElecticProduct implements V110 {

    public void useElectric() {
        System.out.println("베트남의 110 볼트 전기 제품을 사용합니다.");
    }
}

// 전압 V220          
package DesignPattern.gof_adapter.sample01;

public interface V220 {
    public void useElectric();
}

// 전압 V220 용 한국 제품
package DesignPattern.gof_adapter.sample01;

public class KoreaElecticProduct implements V220 {

    public void useElectric() {
        System.out.println("한국의 220 볼트 전기 제품을 사용합니다.");
    }
}

// 전압 V110에서 V220을 사용할 수 있게 해주는 Adapter
package DesignPattern.gof_adapter.sample01;

public class VoltageAdater implements V110{

    private final V220 v220;

    public VoltageAdater(V220 v220){
        this.v220 = v220;
    }

    public void useElectric() {

        System.out.println("V110 볼트의 전압을 V220의 기기를 사용할 수 있게 전압을 조정합니다.");

        v220.useElectric();
    }
}
   
```


### Kotlin 에서는 

```kotlin 
package bong.lines.patterns.adapter.inf

interface V110 {
    fun useElectric()
}

package bong.lines.patterns.adapter.inf

interface V220 {
    fun useElectric()
}
```

```kotlin 
package bong.lines.patterns.adapter.impl

import bong.lines.patterns.adapter.inf.V220

class KoreanElectricProduct : V220 {
    override fun useElectric() {
        print("한국의 220 볼트 전기 제품을 사용합니다.")
    }
}

package bong.lines.patterns.adapter.impl

import bong.lines.patterns.adapter.inf.V110

class VietnamElectricProduct : V110 {
    override fun useElectric() {
        print("베트남의 110 볼트 전기 제품을 사용합니다.")
    }
}

package bong.lines.patterns.adapter.impl

import bong.lines.patterns.adapter.inf.V110
import bong.lines.patterns.adapter.inf.V220

class VoltageAdapter(val v220: V220) : V110 {
    override fun useElectric() {
        v220.useElectric()
    }
}
```

```kotlin
package bong.lines.patterns.adapter

import bong.lines.patterns.adapter.impl.KoreanElectricProduct
import bong.lines.patterns.adapter.impl.VoltageAdapter
import bong.lines.patterns.adapter.inf.V110
import bong.lines.patterns.adapter.inf.V220

fun main() {
    var v220: V220 = KoreanElectricProduct()

    var voltageAdapter : V110 = VoltageAdapter(v220)

    voltageAdapter.useElectric()
}

```
