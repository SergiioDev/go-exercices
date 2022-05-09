package main

import "fmt"

func main() {
	personsFavouriteThings := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Martinis", "Women"},
		"no_dr":      {"Being evil", "Ice cream", "Sunsets"},
	}

	for key, favouriteThings := range personsFavouriteThings {
		fmt.Printf("Character %v favourite things: \n ", key)
		for pos, value := range favouriteThings {
			fmt.Printf("\tPosition %v value: %v\n", pos, value)
		}

	}
}
