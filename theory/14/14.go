package main

import "fmt"

// Т.к. функция append() создает новый слайс и массив.
// То мы меняем значения в уже новом слайсе созданном при вызове append()
func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
