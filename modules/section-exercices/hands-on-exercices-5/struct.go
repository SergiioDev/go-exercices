package main

import "fmt"

type Peerson struct {
	firstName string
	lastName  string
	birthYear int
}

type worker struct {
	Peerson
	company string
}

func main() {
	person1 := Peerson{firstName: "Lightning", lastName: "Mcqueen", birthYear: 2001}
	fmt.Println(person1)

	firstWorker := worker{
		Peerson: Peerson{firstName: "Mcqueen", lastName: "Lightning", birthYear: 2001},
		company: "ss",
	}
	fmt.Println(firstWorker.lastName)

	// Anonymous struct
	contact := struct {
		name    string
		number  string
		company string
		address string
	}{
		name:    "Mater",
		number:  "1111",
		company: "Google",
		address: "Los Angeles",
	}

	fmt.Println(contact)
}
