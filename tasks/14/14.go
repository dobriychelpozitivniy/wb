package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип переменной:
//  int, string, bool, channel из переменной типа interface{}.
func main() {
	var v interface{} = make(chan int)

	fmt.Println(getType(v))
}

func getType(v interface{}) string {
	return fmt.Sprintf("%T\n", v)
}
