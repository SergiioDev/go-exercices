package main

import "fmt"

func main() {
	personsFavouriteThings := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Martinis", "Women"},
		"no_dr":      {"Being evil", "Ice cream", "Sunsets"},
	}
	personsFavouriteThings["npc"] = []string{"talk to others", "walk"}

	for key, favouriteThings := range personsFavouriteThings {
		fmt.Printf("Character %v favourite things: %v \n ", key, favouriteThings)
	}

}
