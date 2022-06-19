package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 9, 1, 15}
	xs := []string{"zulu", "bravo", "alpha"}

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}
