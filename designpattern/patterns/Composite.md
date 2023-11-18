# Pattern - Composite

## Basic Information 

- 패턴형식 : 목적패턴 

## Gof's Description 

부분과 전체의 계층을 표현하기 위해 객체들을 모아 트리 구조로 구성합니다. 사용자로 하여금 개별 객체와 복합 객체를 모두 동일하게 다룰 수 있도록 하는 패턴입니다.

## 필요한 경우 

우리가 사진(Picture)를 모니터 내에 표현해야한다고 했을 때, 사진내에 표현되고 있는 Element를 어떤 방식으로 관리하면 좋을 지에 대해서 고민한다면, 

- 구성 요소 
    - 선
    - 사각형 
    - 글자 

위의 구성 요소에서 선, 사각형, 글자로 구성된 사진을 표현하기 위해서 처리할 수 있는 방법을 그려보면, 

### 첫번째 케이스 

그래픽 편집기의 경우 Line, Rectangle, Text, Picture 등의 다양한 구성요소가 존재하고, 이를 하나의 컨테이너에 담아서 처리할 수 있습니다. 하지만 각 구성요소 별로 크기,색깔, 위치 등의 기본속성이 다를 것이고, 이를 그래픽 편집기에서 쓰기위해서는 각 구성요소 별로의 속성 값에 따라서 분기 처리해서 사용해야하고 이는 개발자의 기억에 의존하는 문제가 발생할 수 있습니다.

![Composite Wrong](https://github.com/keepinmindsh/lines_edu/blob/main/assets/composite_wrong.png)

### 두번째 케이스 

![Composite Right](https://github.com/keepinmindsh/lines_edu/blob/main/assets/composite_graphic.png)

복합 객체의 경우, 하위와 상위 계층이 존재할 수 있습니다. 이 때문에 재귀적인 특성으로 이를 구현할 수 있으며, 상위 요소에서 하위요소의 각 객체의 연산을 호출하여 사용할 수 있게 구성합니다. 상위 계층이라는 것은 하위 계층을 묶어 하나로 구성하는 객체(클래스)라고 생각해도 좋습니다. 이 때문에 그림판이나 파워 포인트 등을 쉬운 예제로 많이 들수 있습니다.

- 부분 - 객체의 객체 계통을 표현하고 싶을 때,
    - 사용자가 객체의 합성으로 생긴 복합 객체와 개개의 객체 사이의 차이를 알지 않고도 자기 일을 할 수 있도록 만들고 싶을 때, 사용자는 복합 구조의 모든 객체를 똑같이 취급하게 됩니다.
    - 복합체 패턴의 주요 목표 중 하나는 사용자가 어떤 Leaf나 Composite 클래스가 존재하는지 모르도록 하고자 할 때,

## 중요 포인트 

복합체 패턴의 가장 중요한 요소는 기본 클래스와 이들의 컨테이너를 모두 표현할 수 있는 하나의 추상화 클래스를 정의하는 것입니다. 그래픽 응용 프로그램 예에서 추상 클래스로 그래픽 편집기를 정의하였을 때, 그래픽 편집기는 그림을 그리기위한 기본 연산인 Draw() 뿐만이 아니라, 이런 기본 클래스를 포함하고 관리하는데 필요한 Add(), Remove(), GetChild() 등의 연산도 정의되어 있습니다.

## 실행활 활용 예제 

### Go 에서는 

```go 
package controller

import (
	"composite/app/service/element"
	"composite/app/service/graphic"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) DrawPicture() {
	picture := graphic.NewGraphic()

	text := element.NewText()
	text.Add(element.NewLine())
	text.Add(element.NewRectangle())

	picture.Add(text)
	picture.Add(element.NewLine())
	
	rectangle := element.NewRectangle()
	rectangle.Add(element.NewLine())
	rectangle.Add(element.NewLine())
	rectangle.Add(element.NewLine())

	picture.Add(rectangle)
	picture.Add(element.NewRectangle())

	picture.Draw()
}
```

```go 
package graphic

import "composite/domain"

type Graphic struct {
	elements []domain.Graphic
}

func NewGraphic() domain.Graphic {
	return &Graphic{}
}

func (g *Graphic) Draw() {
	for _, element := range g.elements {
		element.Draw()
	}
}

func (g *Graphic) Add(element domain.Graphic) {
	g.elements = append(g.elements, element)
}

func (g *Graphic) Remove(element domain.Graphic) {

}
```

```go 
package element

import (
	"composite/domain"
	"fmt"
)

type Line struct {
	elements []domain.Graphic
}

func NewLine() domain.Graphic {
	return &Line{}
}

func (g *Line) Draw() {
	for _, element := range g.elements {
		element.Draw()
	}

	fmt.Println("선을 그립니다.")
}

func (g *Line) Add(element domain.Graphic) {
	g.elements = append(g.elements, element)
}

func (g *Line) Remove(element domain.Graphic) {

}
```

```go 
package element

import (
	"composite/domain"
	"fmt"
)

type Rectangle struct {
	elements []domain.Graphic
}

func NewRectangle() domain.Graphic {
	return &Rectangle{}
}

func (g *Rectangle) Draw() {
	for _, element := range g.elements {
		element.Draw()
	}

	fmt.Println("사각형을 그립니다.")
}

func (g *Rectangle) Add(element domain.Graphic) {
	g.elements = append(g.elements, element)
}

func (g *Rectangle) Remove(element domain.Graphic) {

}
```

```go 
package element

import (
	"composite/domain"
	"fmt"
)

type Text struct {
	elements []domain.Graphic
}

func NewText() domain.Graphic {
	return &Text{}
}

func (g *Text) Draw() {
	for _, element := range g.elements {
		element.Draw()
	}

	fmt.Println("글자를 작성한다.")
}

func (g *Text) Add(element domain.Graphic) {
	g.elements = append(g.elements, element)
}

func (g *Text) Remove(element domain.Graphic) {

}
```

### Java 에서는 

```java 
package DesignPattern.gof_composite.sample002;

public class Client {

    public static void main(String[] args) {

        // 각각의 세부 객체 - 각각의 역할에 따라 다르게 처리됨
        Component leaf1 = new Leaf1();
        Component leaf2 = new Leaf2();
        Component leaf3 = new Leaf3();
        Component leaf4 = new Leaf4();

        Component composite = new Composite();

        // 복합체에 각각의 고유 프로세스 객체를 담는다.
        composite.Add(leaf1);
        composite.Add(leaf2);
        composite.Add(leaf3);
        composite.Add(leaf4);

        // 복합체의 실행
        composite.Operation();
    }
}
```

```java
package DesignPattern.gof_composite.sample002;

import java.util.ArrayList;
import java.util.List;

public class Component {

    public List<Component> children= new ArrayList<>();

    public void Operation(){
        for(Component component : children){
            component.Operation();
        }
    }

    public void Add(Component component){
        children.add(component);
    }

    public void Remove(Component component){
        children.remove(component);
    }

    public void GetChild(int index){
        children.get(index);
    }

}  

package DesignPattern.gof_composite.sample002;

import java.util.ArrayList;
import java.util.List;

public class Composite extends Component {
    public List<Component> children= new ArrayList<>();

    public void Operation(){
        for(Component component : children){
            component.Operation();
        }
    }

    public void Add(Component component){
        children.add(component);
    }

    public void Remove(Component component){
        children.remove(component);
    }

    public void GetChild(int index){
        children.get(index);
    }
}

package DesignPattern.gof_composite.sample002;

public class Leaf1 extends Component {

    @Override
    public void Operation(){
        System.out.println("실질적인 프로세스 처리 : Leaf1");
    }
}

package DesignPattern.gof_composite.sample002;

public class Leaf2 extends Component {

    @Override
    public void Operation(){
        System.out.println("실질적인 프로세스 처리 : Leaf2");
    }
}

package DesignPattern.gof_composite.sample002;

public class Leaf3 extends Component {

    @Override
    public void Operation(){
        System.out.println("실질적인 프로세스 처리 : Leaf3");
    }
}

package DesignPattern.gof_composite.sample002;

public class Leaf4 extends Component {

    @Override
    public void Operation(){
        System.out.println("실질적인 프로세스 처리 : Leaf4");
    }
}
```