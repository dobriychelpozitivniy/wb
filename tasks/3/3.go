package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
func main() {
	fmt.Println(summ(2, 4, 6, 8, 10))

}

func summ(args ...int64) int64 {
	var wg sync.WaitGroup
	var result int64 = 0

	summ := make(chan int64, 1)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case v := <-summ:
				result += v
				wg.Done()
			case <-done:
				return
			}

		}
	}()

	for _, v := range args {
		wg.Add(1)
		go func(v int64) {
			summ <- v * v
		}(v)
	}

	wg.Wait()
	done <- true

	return result
}
