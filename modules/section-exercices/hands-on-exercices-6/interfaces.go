package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

//Polymorphism
func (p person) speak() {
	fmt.Println("I'm a person my name is", p.firstName)
}

type secretAgent struct {
	person
	agency string
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

//Polymorphism
//This function only can be used with types secretAgent
func (agent secretAgent) speak() {
	fmt.Println("I'm", agent.firstName, "working with", agent.agency)
}

type human interface {
	speak()
}

func main() {
	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		"CIA",
	}
	sa1.speak()

	p1 := person{
		"Doctor",
		"No",
	}
	bar(p1)
	bar(sa1)
}
