package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//  Приведите корректный пример реализации.
func main() {
	someFunc()
}

var justString string

func someFunc() {
	// Уборщик мусора не сможет очистить память,
	// т.к. второй срез будет ссылаться на первый
	// в итоге оперативная память переполнится
	v := createHugeString(100000)
	b := make([]rune, len(v[:100]))

	copy(b, []rune(v[:100]))
	justString = string(b)
	fmt.Println(justString)
}

func createHugeString(count int) string {
	i := 0
	s := strings.Builder{}
	r := byte(72)

	for i != count {
		s.WriteByte(r)
		i++
	}

	return s.String()
}
