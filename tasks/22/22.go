package main

import (
	"math"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
func main() {
	a := big.NewFloat(math.Pow(30, 30))
	b := big.NewFloat(math.Pow(2, 35))

	// Сумма
	add := new(big.Float)
	add.Add(a, b)

	// Деление
	div := new(big.Float)
	div.Quo(a, b)

	// Вычитание
	sub := new(big.Float)
	sub.Sub(a, b)

	// Умножение
	mult := new(big.Float)
	mult.Mul(a, b)
}
