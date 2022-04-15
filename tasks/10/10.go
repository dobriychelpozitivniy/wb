package main

import (
	"fmt"
	"strconv"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

type Unite map[string][]float64

func main() {
	s := unite(-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.5, 10, 9, 5, -1, 1, -5, -9.9, 0)

	fmt.Printf("%#v", s)
}

func unite(args ...float64) Unite {
	var u Unite = make(Unite)
	var step string
	var vInt int

	for _, v := range args {
		vInt = int(v)

		if vInt < 10 && vInt >= 0 {
			step = "0"
		} else if vInt < 0 && vInt > -10 {
			step = "-0"
		} else {
			step = strconv.Itoa(vInt - (vInt % 10))
		}

		s := u[step]
		s = append(s, v)
		u[step] = s
	}

	return u
}
