package go_pool

import (
	"fmt"
	"github.com/panjf2000/ants"
	"strconv"
	"testing"
)

func TestNewGoPool(t *testing.T) {
	p, _ := ants.NewPoolWithFunc(5, func(i interface{}) {
		data := i.(int)
		handlerTest(data)
	})

	ch := make(chan interface{}, 100)

	for i := 0; i <= 99; i++ {
		ch <- i
	}

	for {
		data := <-ch
		err := p.Invoke(data)
		if err != nil {

		}
		if len(ch) == 0 {
			break
		}
	}
}

func handlerTest(data int) {
	fmt.Println(strconv.Itoa(data) + "test")
}
