package main

import "fmt"

const (
	firstYear = iota + 2010
	secondYear
	ThirdYear
	FourthYear
)

func main() {
	fmt.Println(firstYear, secondYear, ThirdYear, FourthYear)
}
