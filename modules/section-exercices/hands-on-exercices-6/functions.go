package main

import (
	"fmt"
)

func main() {
	test := "sergi"
	valueReturned := Foo(test)
	fmt.Println(valueReturned)
	charactersReturned, booleanReturned := usersName()

	fmt.Println(charactersReturned, booleanReturned)

	// Defer is to delay the execution of a function or a statement until the nearby function returns
	defer foo(1, 2, 3)
	foo(3, 2, 1)
}

// Foo All parameters in Go are passed by VALUE!!!!!
func Foo(s string) string {
	fmt.Println(s)
	return "value returned"
}

func usersName() (string, bool) {
	return "This is the string returned", false
}

// This is a variadic parameter
func foo(numbers ...int) {
	fmt.Println(numbers)
}
