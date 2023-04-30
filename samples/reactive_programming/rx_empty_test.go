package reactive_programming

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test(t *testing.T) {
	observable := rxgo.Empty()

	ch := observable.Observe()

	item := <-ch

	fmt.Println(item.V)
}
