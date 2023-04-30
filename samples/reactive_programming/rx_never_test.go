package reactive_programming

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_Never(t *testing.T) {

	observable := rxgo.Never()

	_ = observable.Observe()

	//item := <-observe

	//fmt.Println(item.V)

}
