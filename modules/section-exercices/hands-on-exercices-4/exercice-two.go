package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, value := range a {
		fmt.Printf("Position %v value %v\n", i, value)
	}
	fmt.Printf("Type of the array = %T\n", a)
}
