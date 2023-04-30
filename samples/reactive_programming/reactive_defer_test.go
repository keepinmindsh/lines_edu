package reactive_programming

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_Defer(t *testing.T) {

	observable := rxgo.Defer([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
	}})

	observe := observable.Observe()

	for i := 0; i < 3; i++ {
		fmt.Println(<-observe)
	}
}
