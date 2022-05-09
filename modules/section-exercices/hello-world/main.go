package main

import "fmt"

type user int

var b user = 23
var a int

func main() {
	a = int(b)

	fmt.Printf("%T\n", a)
	fmt.Println(a)
}
