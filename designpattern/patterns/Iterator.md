# Iterator 

## Gof's Description

내부 표현부를 노출하지 않고 어떤 집합 객체에 속한 원소들을 순차적으로 접근할 수 있는 방법을 제공 합니다.

## 풀이 

리스트 등의 집합 객체들은 내부 표현 구조를 노출하지 않고도 자신의 원소를 접근할 수 있는 방법을 제공하는 게 좋습니다. 여기에 덧붙여서, 이미 정의한 방법과 다른 방법으로 이들  
원소들을 순회하고자 할 수 도 있습니다. 그러나 다른 순회 방법이 바뀌었다고 List 클래스의 인터페이스를 부풀리고 싶지는 않을 것입니다. 또한 동일한 리스트에 대해서 하나 이상의 
순회 방법을 정의하고 싶을 때도 있습니다.  

### 여기에서 Iterator는 

next() 메소드를 이용하여 데이터를 순회할 수 있는 Object를 말합니다.     

#### iterate는 반복한다는 의미를 가지고 있고, **프로그래밍**에서 반복기는 개발자가 컨테이너, 특히 리스트를 순회할 수 있게 해주는 객체다.  

### Iterator의 장점은, 

- 요소를 제어하는 기능 
- next() 및 prev()를 이용해서 앞뒤로 이동할 수 있는 기능 
- haxNext()를 써서 더많은 요소가 있는지를 확인하는 기능 

### 중요한 점은, 

- 우리가 순회하고자 하는 데이터의 형식, 자료구조를 Iterator 개념을 기반으로 만들 수 있는 것을 의미합니다.
  - Primitive 객체가 제공해주는 순수 기능에 대해서 알고리즘을 적용하여 좀더 빠르고, 쉽게 아이템의 위치, 이동, 조회를 할 수 있는 자료구조 인터페이스를 정의
  - 정의한 자료구조 인터페이스에 대해서 내가 원하는 자료구조를 Iterator 기반으로 사용하는 것

## 실제 Golang에서 Iterator 기반으로 처리해보면, 

```go 
package main

import (
	"bookshelf/app/shelf"
	"fmt"
)

func main() {
	bookList := []string{
		"Book1",
		"Book2",
		"Book3",
		"Book4",
		"Book5",
		"Book6",
		"Book7",
		"Book8",
		"Book9",
		"Book10",
	}

	bookShelf := shelf.NewBookShelf(bookList)

	for bookShelf.HasNext() {
		item := bookShelf.Next()

		fmt.Println(item)
	}
}
```

위에서 언급한 Iterator 기본 인터페이스를 정의하면,   

- Next() 메소드를 이용하여 아이템을 호출 
- HasNext()를 통해서 다음 아이템을 가져올 지를 정의

```go 
package domain

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
```


```go 
package shelf

import "bookshelf/domain"

type Shelf struct {
	index  int
	shelfs []string
}

func (s *Shelf) HasNext() bool {
	if s.index < len(s.shelfs) {
		return true
	}
	return false
}

func (s *Shelf) Next() interface{} {
	if s.HasNext() {
		user := s.shelfs[s.index]
		s.index++
		return user
	}
	return nil
}

func NewBookShelf(shelfs []string) domain.Iterator {
	return &Shelf{
		shelfs: shelfs,
	}
}
```