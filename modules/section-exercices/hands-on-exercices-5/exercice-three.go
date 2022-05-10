package main

import "fmt"

func main() {

	employee := struct {
		firstName string
		lastName  string
		job       string
	}{
		firstName: "Jack",
		lastName:  "Sparrow",
		job:       "Captain",
	}

	fmt.Println(employee)
}
