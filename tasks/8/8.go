package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var i int64 = 12

	strconv.FormatInt(i, 2)

	newI, err := Bit(i, 1, 2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newI)
}

func Bit(v int64, bit int, index int) (int64, error) {
	bits := strings.Split(strconv.FormatInt(v, 2), "")

	fmt.Println(bits)

	if index > len(bits)-1 {
		return 0, fmt.Errorf("Error index - %s out of range %s", index, len(bits)-1)
	}

	if bit > 2 || bit < 0 {
		return 0, fmt.Errorf("Bit value may be only 0 or 1")
	}

	bits[index] = strconv.Itoa(bit)

	fmt.Println(bits)

	result, err := strconv.ParseInt(strings.Join(bits, ""), 2, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
