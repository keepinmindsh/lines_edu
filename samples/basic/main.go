package main

import "fmt"

func main() {
	NewDeclare().MakeArray().Loop()
}

func imperative_programming() {
	list := make([]int, 10)

	for _, value := range list {
		fmt.Println(value)
	}
}

func declarative_programming() {

}

type Declarative struct {
	List []int
}

func NewDeclare() *Declarative {
	return &Declarative{}
}

func (d *Declarative) MakeArray() *Declarative {
	d.List = make([]int, 10)
	return d
}

func (d *Declarative) Loop() {
	for _, value := range d.List {
		fmt.Println(value)
	}
}
