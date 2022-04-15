package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	array := [...]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, v := range array {
		wg.Add(1)
		go func(v int) {
			fmt.Println(v * v)
			wg.Done()
		}(v)
	}

	wg.Wait()
}
