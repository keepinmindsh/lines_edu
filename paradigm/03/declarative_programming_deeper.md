# More Deeper 

## 선언형 프로그래밍에서 중요한 것은 "무엇" 이다. "어떻게"에 초점을 맞추는 것이 아닌 "무엇"에 초점을 맞춘다. 

- 명령형 방식 : 서울에서 부산으로 가기 위해서는 고속 버스 터미널을 타고, 부산 사상 터미널에 내려야 한다. 

- 선언형 방식 : 내가 가야할 곳의 주소는 부산 광역시 해운대로 XXX라는 곳입니다. 

## 프로그램이 함수형 프로그래밍 언어, 논리형 프로그래밍 언어, 제한형 프로그래밍 언어로 작성된 경우를 "선언형"이라고 한다. 

### 구현을 내부에 숨겨서 캡슐화 시키는 것도 "선언형" 프로그래밍이라고 할 수 있다. 

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

