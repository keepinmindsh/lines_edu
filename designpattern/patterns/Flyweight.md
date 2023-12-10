~~# Pattern - Flyweight

## Basic Information

- 패턴 형식 : 목적 패턴 

## Gof's Description 

문서 편집기에서 문자단위 하나까지 객체로 구성할 때 얻을 수 있는 확장성과 서식화, 영향성의 최소화가 가능하지만 여기에서 문제가 되는 것은 비용입니다. 
그 이유는 얼마 되지 않는 문서일지라도 그 객체는 수천개의 문제 객체를 포함 할 수 있기 때문입니다. 이와 같은 문제를 플라이급(FlyWeight) 패턴을 이용해서 
객체를 공유하는 방법을 통해 해결하는 방법을 보여줍니다. 플라이급 객체는 공유가능한 객체로, 여러 비슷한 상황에서 사용될 수 있습니다.
그러나 각각의 상황에서는 독립적인 객체로 동작합니다. 이것은 공유될 수 없는 객체의 인스턴스와 구분이 안된다는 의미입니다. 
그러므로 플라이급 객체가 적용될 사항을 미리 가정하면서 소프트웨어를 개발할 수 없습니다. 즉 같은 것을 놓고 이런 상황에서는 이런 특징으로 정의하고, 
또 다른 상황에서는 다른 특징으로 정의할 수 없다는 것입니다. 플라이급 패턴에서 중요한 개념은 본질적 상태와 부가적 상태의 구분입니다. 
본질적 상태는 플라이급 객체에 저장되어야하며, 이것이 적용되는 상황과 상관없는 본질적 특성 정보들이 객체를 구성합니다.

## 언제 활용해야 하는가? 

플라이급 패턴은 언제 사용하는가에 따라서 그 효과가 달라집니다.
- 응용프로그램이 대량의 객체를 사용해야 할 때,
- 객체의 수가 너무 많아져 저장 비용이 너무 높아질 때,
- 대부분의 객체 상태를 부가적인 것으로 만들 수 있을 때,
- 부가적인 속성들을 제거한 후 객체들을 조사해보니 객체의 많은 묶음이 비교적 적은 수의 공유된 객체로 대체될 수 있을 때. 현재 서로 다른 객체로 간주한 이유는 이들 부가적인 속성 때문이었지 본질이 달라던 것은 아닐 때,
- 응용 프로그램이 객체의 정체성에 의존하지 않을 때. 플라이급 객체들은 공유될 수 있음을 의미하는데, 식별자가 있다는 것은 서로 다른 객체로 구별해야한다는 의미이므로 플라이급 객체를 사용할 수 없습니다.

![Flyweight](https://github.com/keepinmindsh/lines_edu/blob/main/assets/flyweight.png)


## 실제로 활용해보면, 

### Golang 에서는 

```go 
package main

import (
	"excel_sheet/app/usecase"
	"excel_sheet/domain"
	"excel_sheet/domain/position"
	"fmt"
)

func main() {

	var text domain.Text
	for i := 0; i < 5; i++ {
		text = usecase.GetText(domain.A)

		cell := usecase.NewCell(&position.Position{
			X: i,
			Y: i,
		}, &text)

		print(fmt.Sprintf("Cell Address : %v   ", &cell))

		cell.Draw()
	}
}
```

```go 
package usecase

import (
	"excel_sheet/domain"
	"excel_sheet/domain/position"
	"fmt"
)

type Cell struct {
	Position *position.Position
	Text     *domain.Text
}

func NewCell(position *position.Position, text *domain.Text) domain.Cell {
	c := &Cell{
		Position: position,
		Text:     text,
	}

	return c
}

func (c *Cell) Draw() {
	if c.Text != nil {

		print(fmt.Sprintf(" Text: Address - %v   ", &(*c.Text)))
		print(fmt.Sprintf(" Text: Pointer Address - %v   ", &c.Text))

		(*c.Text).Draw()
	}
}
```

```go 
package usecase

import "excel_sheet/domain"

type a struct{}

func (a a) Draw() {
	print("  A \r\n")
}

type b struct{}

func (b b) Draw() {
	print("  B \n")
}

type c struct{}

func (c c) Draw() {
	print("  C \n")
}

type d struct{}

func (d d) Draw() {
	print("  D \n")
}

type e struct{}

func (e e) Draw() {
	print("  E \n")
}

var TextA domain.Text
var TextB domain.Text
var TextC domain.Text
var TextD domain.Text
var TextE domain.Text

func GetText(txtType domain.TextType) domain.Text {
	switch txtType {
	case domain.A:
		if TextA != nil {
			return TextA
		} else {
			TextA = a{}
			return TextA
		}
	case domain.B:
		if TextB != nil {
			return TextB
		} else {
			TextB = b{}
			return TextB
		}
	case domain.C:
		if TextC != nil {
			return TextC
		} else {
			TextC = c{}
			return TextC
		}
	case domain.D:
		if TextD != nil {
			return TextD
		} else {
			TextD = d{}
			return TextD
		}
	case domain.E:
		if TextE != nil {
			return TextE
		} else {
			TextE = e{}
			return TextE
		}
	default:
		return nil
	}
}
```

```go 
package position

type Position struct {
	X int
	Y int
}
```

```go 
package domain

type (
	Cell interface {
		Draw()
	}

	Text interface {
		Draw()
	}
)

type TextType string

const (
	A TextType = "A"
	B TextType = "B"
	C TextType = "C"
	D TextType = "D"
	E TextType = "E"
)

func (t TextType) String() string {
	return string(t)
}
```

### Java 에서는 

```java 
package designpattern.gof_flyweight.sample01;

public class FlyweightMain {

    private static INoodle[] ramen = new Ramen[20];
    private static NoodleContext[] tables = new NoodleContext[20];
    private static int ordersCount = 0;
    private static NoodleFactory noodleFactory;

    public static void takeOrder(String flavorIn, int table) {
        ramen[ordersCount] = noodleFactory.getNoodleFlavor(flavorIn);
        tables[ordersCount] = new NoodleContext(table);
        ordersCount++;
    }

    public static void main(String args[]) {
        noodleFactory = new NoodleFactory();

        takeOrder("Zin Ramen", 2);
        takeOrder("Zin Ramen", 2);
        takeOrder("Ahn Sung Tang Men", 1);
        takeOrder("Ahn Sung Tang Men", 2);
        takeOrder("Ahn Sung Tang Men", 3);
        takeOrder("Ahn Sung Tang Men", 4);
        takeOrder("Zin Ramen", 4);
        takeOrder("Zin Ramen", 5);
        takeOrder("Ahn Sung Tang Men", 3);
        takeOrder("Zin Ramen", 3);

        for (int i = 0; i < ordersCount; i++) {
            ramen[i].serveNoodle(tables[i]);
        }

        System.out.println("\n Total Coffee objects made: "
                + noodleFactory.getTotalNoodleFlavorMade());
    }
}  

package DesignPattern.gof_flyweight.sample01;

public interface INoodle {
    public void serveNoodle(NoodleContext noodleContext);
}

package DesignPattern.gof_flyweight.sample01;

public class NoodleContext {
    private final int tableNumber;

    public NoodleContext(int tableNumber){
        this.tableNumber = tableNumber;
    }

    public int getTable(){
        return this.tableNumber;
    }
}

package DesignPattern.gof_flyweight.sample01;

import java.util.HashMap;

public class NoodleFactory {
    private HashMap <String, INoodle> flavors = new HashMap <String, INoodle>();

    public INoodle getNoodleFlavor(String flavorName){
        INoodle noodle = flavors.get(flavorName);

        if(noodle == null){
            noodle = new Ramen(flavorName);
            flavors.put(flavorName, noodle);
        }

        return noodle;
    }

    public int getTotalNoodleFlavorMade(){
        return flavors.size();
    }
}

package DesignPattern.gof_flyweight.sample01;

public class Ramen implements INoodle {

    private final String flavor;


    public Ramen(String flavor){
        this.flavor = flavor;
        System.out.println("Noodle is created!" + flavor);
    }

    public String getFlavor(){
        return this.flavor;
    }

    public void serveNoodle(NoodleContext noodleContext) {
        System.out.println("Serving" + flavor + " to table " + noodleContext.getTable());
    }

}
```