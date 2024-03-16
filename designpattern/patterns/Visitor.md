# Visitor 

## Basic Information

- 패턴 형식 : 행위패턴

## Gof's Description

객체 구조를 이루는 원소에 대해 수행할 연산을 표현합니다. 연산을 적용할 원소의 클래스를 변경하지 않고도 새로운 연산을 정의할 수 있게 합니다.

### 풀이

- 다른 인터페이스를 가진 클래스가 객체 구조에 포함되어 있으며, 구체 클래스에 따라 달라진 연산을 이들 클래스의 객체에 대해 수정하고자 할 때
- 각각 특징이 있고, 관련되지 않은 많은 연산이 한 객체 구조에 속해있는 객체들에 대해 수행될 필요가 있으며, 연산으로 클래스들을 “더럽히고” 싶지 않을 때, Visitor 클래스는 관련된 모든 연산을 하나의 클래스 안에다 정의해 놓음으로써 관련된 연산이 함께 있을 수 있게 해줍니다. 어떤 객체 구조가 많은 응용 프로그램으로 공유될 때, Visitor 클래스를 사용하면 이 객체 구조가 필요한 응용 프로그램에만 연산을 둘 수 있습니다
- 객체 구조를 정의한 클래스는 거의 변하지 않지만, 전체 구조에 걸쳐 새로운 연산을 추가하고 싶을 때. 객체 구조를 변경하려면 모든 방문자에 대한 인터페이스를 재정의해야 하는데, 이 작업에 잠재된 비용이 클 수 있습니다. 객체 구조가 자주 변경 될 때는 해당 연산을 클래스에 정의하는 편이 더 낫습니다.

Visitor 패턴은 응용 프로그램의 로직/연산을 한 곳에 모아서 처리하는 것이 더욱 효과적인 경우에 사용할 수 있습니다. 
즉, 여러 객체가 동일하가 하나의 연산을 수행해야할 때, 해당 연산의 처리할 때, IO 등을 제어해야할 때 사용하기 효율적일 수 있습니다.   

## 코드 예시

### Golang

```go 
package main

import (
	"printer/app/action"
	"printer/app/client"
	"printer/app/input"
	"printer/domain"
)

func main() {
	list := make([]domain.Input, 4)

	list = append(list, input.NewLG())
	list = append(list, input.NewSamsung())
	list = append(list, input.NewEpson())
	list = append(list, input.NewPDF())

	printer := client.NewPrinter(list)

	printer.Add(action.NewCheckPrint())
	printer.Add(action.NewInputPrint())
	printer.Add(action.NewNextPrint())

	printer.Print()
}
```

```go 
package client

import (
	"printer/domain"
)

type Printer struct {
	PrintInput []domain.Input
}

var printAction []domain.Printer

func (p *Printer) Add(printerAction domain.Printer) {
	printAction = append(printAction, printerAction)
}

func (p *Printer) Print() {
	for _, input := range p.PrintInput {
		for _, printer := range printAction {
			input.Accept(printer)
		}
	}
}

func NewPrinter(inputList []domain.Input) Printer {
	return Printer{
		PrintInput: inputList,
	}
}
```

```go
package action

import "printer/domain"

type checkPrint struct{}

func (c checkPrint) ExecuteForEpson() {
	//TODO implement me
	panic("implement me")
}

func (c checkPrint) ExecuteForPDF() {
	//TODO implement me
	panic("implement me")
}

func (c checkPrint) ExecuteForSamsung() {
	//TODO implement me
	panic("implement me")
}

func (c checkPrint) ExecuteForLG() {
	//TODO implement me
	panic("implement me")
}

func NewCheckPrint() domain.Printer {
	return checkPrint{}
}

package action

import "printer/domain"

type inputPrint struct{}

func (i inputPrint) ExecuteForEpson() {
	//TODO implement me
	panic("implement me")
}

func (i inputPrint) ExecuteForPDF() {
	//TODO implement me
	panic("implement me")
}

func (i inputPrint) ExecuteForSamsung() {
	//TODO implement me
	panic("implement me")
}

func (i inputPrint) ExecuteForLG() {
	//TODO implement me
	panic("implement me")
}

func NewInputPrint() domain.Printer {
	return inputPrint{}
}

package action

import "printer/domain"

type nextPrint struct{}

func (n nextPrint) ExecuteForEpson() {
	//TODO implement me
	panic("implement me")
}

func (n nextPrint) ExecuteForPDF() {
	//TODO implement me
	panic("implement me")
}

func (n nextPrint) ExecuteForSamsung() {
	//TODO implement me
	panic("implement me")
}

func (n nextPrint) ExecuteForLG() {
	//TODO implement me
	panic("implement me")
}

func NewNextPrint() domain.Printer {
	return &nextPrint{}
}
```

```go
package input

import "printer/domain"

type epson struct {
}

func (e *epson) Accept(printer domain.Printer) {
	printer.ExecuteForEpson()
}

func NewEpson() domain.Input {
	return &epson{}
}

package input

import "printer/domain"

type lg struct {
}

func (l *lg) Accept(printer domain.Printer) {
	printer.ExecuteForLG()
}

func NewLG() domain.Input {
	return &lg{}
}

package input

import "printer/domain"

type pdf struct {
}

func (p pdf) Accept(printer domain.Printer) {
	printer.ExecuteForPDF()
}

func NewPDF() domain.Input {
	return &pdf{}
}

package input

import "printer/domain"

type samsung struct {
}

func (s *samsung) Accept(printer domain.Printer) {
	printer.ExecuteForSamsung()
}

func NewSamsung() domain.Input {
	return &samsung{}
}
```