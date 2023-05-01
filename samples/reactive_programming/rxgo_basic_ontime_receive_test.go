package reactive_programming

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_RxGo_OneTimeReceive(t *testing.T) {

	observable := rxgo.Just("Hello World!")()

	ch := observable.Observe()

	item := <-ch
	message := item.V.(string)
	fmt.Println(message)

}
