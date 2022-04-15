package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	words := "snow dog sun"

	s := strings.Split(words, " ")

	reverse(s)

	fmt.Println(s)
}

func reverse[T any](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}
