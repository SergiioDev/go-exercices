package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	person := Person{
		name: "sergi",
		age:  5,
	}

	changeMe(&person)

	fmt.Println(person.age)
}

func changeMe(person *Person) {
	fmt.Println(person.age)
	person.age = 20
}
