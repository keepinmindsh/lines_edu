package reactive_programming

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()

	ch := observable.Observe()

	item := <-ch

	fmt.Println(item.V)
}

func TestHelloWorldWithError(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()

	ch := observable.Observe()

	item := <-ch

	if item.Error() {
		t.Error(item.E)
	}

	fmt.Println(item.V)
}

func TestHelloWorldWithForeach(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()

	<-observable.ForEach(func(v interface{}) {
		fmt.Printf("received: %v\n", v)
	}, func(err error) {
		fmt.Printf("error: %e\n", err)
	}, func() {
		fmt.Println("observable is closed")
	})

	ch := observable.Observe()

	item := <-ch

	if item.Error() {
		t.Error(item.E)
	}

	fmt.Println(item.V)
}
