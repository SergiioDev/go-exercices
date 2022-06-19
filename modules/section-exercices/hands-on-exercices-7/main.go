package main

import "fmt"

func main() {

	a := 42
	fmt.Println(a)

	//Pointer
	address := &a
	fmt.Println(address)

	//Dereferencing the pointer to the a
	fmt.Println(*address)
	number := 19
	setNumber(&number)
	fmt.Println(number)
}

func setNumber(number *int) {
	*number = 20
}
