package main

import "fmt"

type Person struct {
	firstName                 string
	lastName                  string
	favouriteIceCreamFlavours []string
}

func main() {
	person1 := Person{
		firstName:                 "Jack",
		lastName:                  "Sparrow",
		favouriteIceCreamFlavours: []string{"Chocolate", "Cream"},
	}
	fmt.Printf("\tName: %v \n\tLast name: %v \n\tFavourite Ice Cream Flavours: ", person1.firstName, person1.lastName)
	for _, flavour := range person1.favouriteIceCreamFlavours {
		fmt.Printf("%v & ", flavour)
	}
}
