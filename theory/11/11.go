package main

import (
	"fmt"
	"sync"
)

// Программа выведет дедлок т.к. мы передаем копию sync.WaitGroup
// из-за чего не меняется значение wg и мы вечно ждём wg.Wait()
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
