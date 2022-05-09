package main

import "fmt"

func main() {

	streamers := [][]string{{"melkey", "prime"}, {"trashDev", "bashbunny"}}

	for pos, line := range streamers {
		fmt.Printf("Line: %v \n", pos)
		for columns := range line {
			fmt.Printf("\t Column %v value: %v  \n ", columns, streamers[pos][columns])
		}
	}

}
