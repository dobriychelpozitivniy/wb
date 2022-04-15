package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(unite(s))
}

func unite(args []string) []string {
	s := make([]string, 0, len(args))
	have := make(map[string]struct{})

	for _, v := range args {
		_, ok := have[v]
		if !ok {
			have[v] = struct{}{}
			s = append(s, v)
		}
	}

	return s
}
