package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   32,
	}

	people := []Person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Unmarshall \n %v\n", string(bs))

	peoples := []Person{}

	err = json.Unmarshal(bs, &peoples)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(peoples)

}
