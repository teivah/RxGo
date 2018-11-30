package rxgo

import (
	"fmt"
	"github.com/reactivex/rxgo/handlers"
	"testing"
)

func TestConnect(t *testing.T) {
	just := Just(1, 2, 3).Publish()

	just.Subscribe(handlers.NextFunc(func(i interface{}) {
		fmt.Printf("%v\n", i)
	}))

	just.Subscribe(handlers.NextFunc(func(i interface{}) {
		fmt.Printf("%v\n", i)
	}))

	just.Connect()

	select {}
}
