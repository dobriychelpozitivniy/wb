package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
func main() {
	n := 200
	v := 0

	c := Counter{count: int32(v)}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			c.Inc()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(c)
}

type Counter struct {
	count int32
}

func (c *Counter) Inc() {
	atomic.AddInt32(&c.count, 1)
}

func (c *Counter) Dec() {
	atomic.AddInt32(&c.count, -1)
}
