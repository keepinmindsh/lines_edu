package rx_real_sample

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
)

func Test_ColdSample(t *testing.T) {
	observable := rxgo.Defer([]rxgo.Producer{
		func(_ context.Context, next chan<- rxgo.Item) {
			for i := 0; i < 3; i++ {
				next <- rxgo.Of(i)
			}
		},
	})

	for item := range observable.Observe() {
		fmt.Println("Cold observable 1 ::", item.V)
	}

	for item := range observable.Observe() {
		fmt.Println("Cold observable 2 ::", item.V)
	}
}

func Test_ColdSampleWithCreate(t *testing.T) {
	observable := rxgo.Defer([]rxgo.Producer{
		func(_ context.Context, next chan<- rxgo.Item) {
			for i := 0; i < 30; i++ {
				next <- rxgo.Of(i)
			}
		},
	})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for item := range observable.Observe() {
			fmt.Println("Cold observable 1 ::", item.V)
		}
		observable.DoOnCompleted(func() {
			wg.Done()
		})
	}()

	wg.Add(1)
	go func() {
		for item := range observable.Observe() {
			fmt.Println("Cold observable 2 ::", item.V)
		}
		observable.DoOnCompleted(func() {
			wg.Done()
		})
	}()

	wg.Wait()
}
