package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
func main() {
	a := Action{
		Human: Human{
			sex: "women",
			age: 18,
		},
		action: "action",
	}

	fmt.Printf("%#v", a)
}

type Human struct {
	sex string
	age int
}

type Action struct {
	Human
	action string
}
