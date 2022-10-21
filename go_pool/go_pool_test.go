package go_pool

import (
	"fmt"
	"github.com/panjf2000/ants"
	"strconv"
	"testing"
)

func TestGoPool(t *testing.T) {
	p, _ := ants.NewPoolWithFunc(5, func(i interface{}) {
		data := i.(int)
		handlerTest(data)
	})
	// 关闭协程池
	defer p.Release()
	ch := make(chan interface{}, 100)

	for i := 0; i <= 99; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		err := p.Invoke(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func handlerTest(data int) {
	fmt.Println(strconv.Itoa(data) + "test")
}
