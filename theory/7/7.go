package main

import "fmt"

// В какой последовательности будут выведены элементы map[int]int?

// В рандомном порядке
func main() {
	var m map[int]int = map[int]int{}

	m[0] = 1
	m[1] = 124
	m[2] = 281

	for i, v := range m {
		fmt.Println(i, v)
	}
}
