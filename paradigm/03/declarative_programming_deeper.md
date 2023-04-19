# More Deeper 

## 선언형 프로그래밍에서 중요한 것은 "무엇" 이다. "어떻게"에 초점을 맞추는 것이 아닌 "무엇"에 초점을 맞춘다. 

- 명령형 방식 : 서울에서 부산으로 가기 위해서는 고속 버스 터미널을 타고, 부산 사상 터미널에 내려야 한다. 

- 선언형 방식 : 내가 가야할 곳의 주소는 부산 광역시 해운대로 XXX라는 곳입니다. 

## 프로그램이 함수형 프로그래밍 언어, 논리형 프로그래밍 언어, 제한형 프로그래밍 언어로 작성된 경우를 "선언형"이라고 한다. 

### 구현을 내부에 숨겨서 캡슐화 시키는 것도 "선언형" 프로그래밍이라고 할 수 있다. 

명령형 방식과 선언형 방식의 차이는 아래의 예제코드에서 확인할 수 있다. 

- 명령형 방식 

```golang
package main 

func main(){
  list := makeList()
}

func makeList() []int {
  list := make([]int, 10)
  
  rowCount := len(list)
  
  for i := 0 i < rowCount; i ++ {
    list[i] = i 
  }
  
  return list
}
```

- 선언형 방식 

```golang 
package main

import "fmt"

func main() {
	d := declarative{}

	result := d.makeArray().loop().result()

	fmt.Printf("%v", result)
}

type declarative struct {
	list []int
}

func (d *declarative) makeArray() *declarative {
	d.list = make([]int, 10)
	return d
}

func (d *declarative) loop() *declarative {
	rowCount := len(d.list)

	for i := 0; i < rowCount; i++ {
		d.list[i] = i
	}

	return d
}

func (d *declarative) result() []int {
	return d.list
}

```

우리가 쉽게 이해 하기에는 SQL이 선언형 언어로 가장 명확하지만, 선언형 언어의 경우에도 결국 내부 구현을 캡슐화하여 숨기고 있을 뿐이다. 

그렇기 때문에 다양한 언어에서 함수를 정의하고 사용하는 방식에 따라서 선언적인 방식으로 코드를 작성할 수 있는 방법들을 제공하고 있다. 

선언형 프로그래밍 언어가 우리에게 시사하는 바는 우리가 다루는 도메인을 좀더 명확하게 이해할 수 있게 만들어 준다는 것이다. 


선언형 프로그래밍 언어로 파생된 다양한 개념/언어는, 

- SQL
  - []() 	
- Behavier Driven Development 
  - [https://github.com/jaypipes/gdt](https://github.com/jaypipes/gdt)
- Domain Specific Language 
  - [Method Chaining](https://github-history.netlify.app/keepinmindsh/lines_edu/blob/main/paradigm/03/declarative/method_chaning.go)  	
