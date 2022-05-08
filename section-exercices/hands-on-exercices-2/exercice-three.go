package main

import "fmt"

func main() {
	//Untyped constant
	const playerMinimumAge = 18

	//Typed constant
	const playerMaxAge int = 100

	fmt.Println(playerMinimumAge, playerMaxAge)
}