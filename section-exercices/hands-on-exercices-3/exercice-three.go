package main

import "fmt"

func main() {
	for i := 0; i < 2022; i++ {
		moduleResult := i % 2
		fmt.Printf("The result of dividing %v between 2 is %v\n", i, moduleResult)

	}
}
