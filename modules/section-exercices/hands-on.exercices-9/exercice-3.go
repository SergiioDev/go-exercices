package main

import "fmt"

type person struct {
	fist string
	last string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.fist, " speaking")
}

func saySomething(h human) {
	h.speak()
}

func main() {

	saySomething(&person{
		fist: "Sergi",
		last: "Aguilar",
	})

	// you need to pass a pointer
	//	saySomething(person{
	//		fist: "Sergi",
	//		last: "Aguilar",
	//	})

}
