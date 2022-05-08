package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	birthYear int
}

type worker struct {
	Person
	company string
}

func main() {
	person1 := Person{firstName: "Sergi", lastName: "Aguilar", birthYear: 2001}
	fmt.Println(person1)

	firstWorker := worker{
		Person:  Person{firstName: "Aguikar", lastName: "Ruiz", birthYear: 2001},
		company: "ss",
	}
	fmt.Println(firstWorker.lastName)

}
