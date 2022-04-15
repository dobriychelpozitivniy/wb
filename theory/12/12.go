package main

import "fmt"

// n объявленная внутри if имеет собственную зону видимости и переопределяет внешнюю n
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
