package main

import "fmt"

// В какой последовательности будут выведены элементы map[int]int?

// По порядку т.к. map выводится отсортированными ключами по возрастанию
func main() {
	var m map[int]int = map[int]int{}

	m[0] = 1
	m[1] = 124
	m[2] = 281

	for i, v := range m {
		fmt.Println(i, v)
	}
}
