package main

import "fmt"

func main() {
	usersFavouriteThings := make([][]string, 50)

	usersFavouriteThings = [][]string{{"melkey", "prime"}, {"trashDev", "bashbunny"}}

	for i, value := range usersFavouriteThings {
		fmt.Printf("Line number %v \n", i)
		for i2, columnValue := range value {
			fmt.Printf("\tColumn position %v Value: %v \n ", i2, columnValue)
		}
	}

}
