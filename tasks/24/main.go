package main

import (
	"fmt"

	"24/point"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

func main() {
	p := point.NewPoint(2.5, 3.5)
	p2 := point.NewPoint(4.8, 10)

	fmt.Println(point.Distance(*p, *p2))
}
