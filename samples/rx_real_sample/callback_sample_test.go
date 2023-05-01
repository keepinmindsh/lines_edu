package rx_real_sample

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

var (
	e1 = errors.New("error 1")
	e2 = errors.New("error 2")
)

func Test_RealSample(t *testing.T) {
	observable := rxgo.Just(1, 2, e1, 3, 4, 5, e2, 6)().
		Map(times10).
		Filter(greaterThan30)

	done := observable.ForEach(onNext, onError, onComplete, rxgo.WithErrorStrategy(rxgo.ContinueOnError))
	<-done
}

func onNext(i interface{}) {
	fmt.Println("items ::", i)
}

func onError(err error) {
	fmt.Println("error :: ", err)
}

func onComplete() {
	fmt.Println("complete")
}

func times10(ctx context.Context, i interface{}) (interface{}, error) {
	return i.(int) * 10, nil
}

func greaterThan30(i interface{}) bool {
	return i.(int) > 30
}
