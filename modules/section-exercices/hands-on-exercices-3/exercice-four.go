package main

import "fmt"

func main() {
	fmt.Println(isFavouriteSportF1())
}

func isFavouriteSportF1() bool {
	switch favSport := "F1"; favSport {
	case "Diving":
		return false
	case "F1":
		return true
	case "GT3":
		return true
	default:
		return false
	}
}
