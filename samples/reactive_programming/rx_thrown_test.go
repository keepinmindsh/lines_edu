package reactive_programming

import (
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"log"
	"testing"
)

func Test_Thrown(t *testing.T) {
	observable := rxgo.Thrown(errors.New("Foo"))

	ch := observable.Observe()

	item := <-ch

	if item.E != nil {
		log.Fatal(item.Error())
	}

	fmt.Print(item.V)
}
