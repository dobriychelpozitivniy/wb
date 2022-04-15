package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
func main() {
	var arr []int = []int{11, 5, 10, 13, 2, 13, 10}

	sorted := QuickSort(arr, 0, len(arr)-1)

	fmt.Println(sorted)
}

func QuickSort(arr []int, minIndex int, maxIndex int) []int {
	if minIndex >= maxIndex {
		return arr
	}

	pivotIndex := GetPivotIndex(arr, minIndex, maxIndex)

	QuickSort(arr, minIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, maxIndex)

	return arr
}

func GetPivotIndex(arr []int, minIndex int, maxIndex int) int {
	pivot := minIndex - 1

	for i := minIndex; i <= maxIndex; i++ {
		if arr[i] < arr[maxIndex] {
			pivot++
			arr[pivot], arr[i] = arr[i], arr[pivot]
		}
	}

	pivot++
	arr[pivot], arr[maxIndex] = arr[maxIndex], arr[pivot]

	return pivot
}
