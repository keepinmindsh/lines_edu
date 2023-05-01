package rx_real_sample

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
)

func Test_HotSample(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	observable := rxgo.FromChannel(ch)

	for item := range observable.Observe() {
		fmt.Println("Hot observable 1 ::", item.V)
	}

	for item := range observable.Observe() {
		fmt.Println("Hot observable 1 ::", item.V)
	}
}

func Test_HotSampleWithCreate(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for item := range observable.Observe() {
			fmt.Println("Hot observable 1 ::", item.V)
		}
		observable.DoOnCompleted(func() {
			wg.Done()
		})
	}()

	wg.Add(1)
	go func() {
		for item := range observable.Observe() {
			fmt.Println("Hot observable 2 ::", item.V)
		}
		observable.DoOnCompleted(func() {
			wg.Done()
		})
	}()

	wg.Wait()
}
