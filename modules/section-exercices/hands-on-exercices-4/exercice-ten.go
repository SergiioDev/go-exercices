package main

import "fmt"

func main() {
	personsFavouriteThings := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Martinis", "Women"},
		"no_dr":      {"Being evil", "Ice cream", "Sunsets"},
	}
	delete(personsFavouriteThings, "no_dr")

	if _, exists := personsFavouriteThings["no_dr"]; !exists {
		fmt.Println("Record deleted successfully")
	}

	for key, favouriteThings := range personsFavouriteThings {
		fmt.Printf("Character %v favourite things: %v \n ", key, favouriteThings)
	}
}
