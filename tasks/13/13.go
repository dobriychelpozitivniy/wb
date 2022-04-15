package main

import "fmt"

// Поменять местами два числа без создания временной переменной.
func main() {
	a := 10
	b := 12

	a, b = b, a

	fmt.Println(a, b)
}
