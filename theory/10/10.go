package main

import "fmt"

type Test struct {
	i int
}

// Т.к. в update передается копия указателя на p,
// то эта копия теперь указывает на b, и никак не затрагивает переданный аргумент
func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
