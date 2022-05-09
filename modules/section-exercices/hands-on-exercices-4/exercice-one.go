package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}

	for i, value := range a {
		fmt.Printf("Position %v value %v\n", i, value)
	}
	fmt.Printf("Type of the array = %T\n", a)

}
