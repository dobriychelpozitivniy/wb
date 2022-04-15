package main

import (
	"fmt"
	"sort"
	"strings"
)

// Разработать программу, которая проверяет,
// что все символы в строке уникальные (true — если уникальные, false etc).
//  Функция проверки должна быть регистронезависимой.

func main() {
	s := "Thanks ц 😊"
	fmt.Println(checkUnique(s))
}

func checkUnique(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	arr := []rune(s)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
	}

	return true
}
