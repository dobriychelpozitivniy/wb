package main

import "fmt"

// Т.к. функция append() создает новый слайс и массив.
// То это не отражается на изначальном слайсе

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	a := make([]int8, 0, 10)
	a = append(a, 1, 2, 3, 4, 5)
	someAction(a, 6)
	fmt.Println(a)
}
