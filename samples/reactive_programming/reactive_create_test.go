package reactive_programming

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func TestCreate(t *testing.T) {
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
	}})

	ch := observable.Observe()

	item1 := <-ch
	item2 := <-ch
	item3 := <-ch

	fmt.Println(item1.V)
	fmt.Println(item2.V)
	fmt.Println(item3.V)
}
