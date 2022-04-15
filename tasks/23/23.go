package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	var s []int = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(deleteElemSlice(s, 3))

	s = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(deleteElemSlice2(s, 4))
}

// Сохраняет порядок
func deleteElemSlice(s []int, n int) []int {
	s = append(s[:n], s[n+1:]...)

	return s
}

// не сохраняет порядок
func deleteElemSlice2(s []int, n int) []int {
	s[n] = s[len(s)-1]

	return s[:len(s)-1]
}
