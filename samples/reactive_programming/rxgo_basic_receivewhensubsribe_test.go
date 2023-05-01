package reactive_programming

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_ReceiveWhenSubscribe(t *testing.T) {

	observable := rxgo.Just("Hello World")()

	ch := observable.Observe()

	for item := range ch {
		message := item.V.(string)
		fmt.Println(message)
	}
}
