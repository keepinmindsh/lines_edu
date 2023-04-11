# 관심있는 프로그래밍 패러다임에 대해서 실제 코드로 구현해볼까요? 

## 반응형 프로그래밍 ( Reactive Programming ) 

> [ReactiveX](https://reactivex.io/)   
> [RxGo](https://github.com/ReactiveX/RxGo)


## 시작해보자 

```shell 
$ go get -u github.com/reactivex/rxgo/v2
```

### 기본적인 사용 방법

```go 
package sample 

func TestSaveMessage(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()
	ch := observable.Observe()
	item := <-ch
	fmt.Println(item.V)
}

```

### Callback을 이용한 호출 방식 


```go
package sample 

func TestSaveMessageWithObservable(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()

	observable.ForEach(func(v interface{}) {
		// Next Func
		fmt.Printf("received: %v\n", v)
	}, func(err error) {
		// Err Func
		fmt.Printf("error: %e\n", err)
	}, func() {
		// Complete Func
		fmt.Println("observable is closed")
	})

	ch := observable.Observe()

	// 채널을 통해서 하나의 아이템을 소모하고, 해당 아이템의 값을 item.V를 통해서 출력한다.
	item := <-ch
	fmt.Printf("final result : %v", item.V)
}
```
