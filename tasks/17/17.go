package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	var arr []int = []int{2, 7, 9, 11, 54, 76, 87, 101, 215, 228, 654, 857, 931, 1011, 5302, 10222}

	fmt.Println(binarySearch(arr, 1011))
}

func binarySearch(arr []int, exist int) (int, error) {
	left := 0 // задаем левую и правую границы поиска
	right := len(arr) - 1
	search := -1 // найденный индекс элемента равен -1 (элемент не найден)

	for left <= right { // пока левая граница не "перескочит" правую
		mid := (left + right) / 2 // ищем середину отрезка
		if exist == arr[mid] {    // если ключевое поле равно искомому
			search = mid // мы нашли требуемый элемент,
			break
		}

		if exist < arr[mid] { // если искомое ключевое поле меньше найденной середины
			right = mid - 1
		} else { // иначе
			left = mid + 1 // смещаем левую границу, продолжим поиск в правой части
		}
	}

	if search == -1 {
		return search, fmt.Errorf("Element %d, not found in array.", exist)
	} else { // иначе выводим элемент, его ключ и значение
		return search, nil
	}
}
