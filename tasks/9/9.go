package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	first := make(chan int)
	second := make(chan int)
	wg := sync.WaitGroup{}

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for v := range first {
			second <- v * 2
		}
	}()

	go func() {
		for v := range second {
			fmt.Println(v)
			wg.Done()
		}
	}()

	for _, v := range arr {
		wg.Add(1)
		first <- v
	}

	wg.Wait()
}
