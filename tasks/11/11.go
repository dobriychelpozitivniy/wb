package main

import (
	"fmt"
)

// Реализовать пересечение двух неупорядоченных множеств.
func main() {
	var arr1 []int = []int{4, 2, 1, 53, 375, 33, 11, 58, 63}
	var arr2 []int = []int{5, 3, 8, 1, 2, 375, 11}

	fmt.Println(Intersection(arr1, arr2))
}

func Intersection(arr1 []int, arr2 []int) []int {
	var result []int = []int{}
	var exist map[int]struct{} = make(map[int]struct{})

	for _, v := range arr1 {
		exist[v] = struct{}{}
	}

	for _, v := range arr2 {
		if _, ok := exist[v]; ok {
			result = append(result, v)
		}
	}

	return result
}
