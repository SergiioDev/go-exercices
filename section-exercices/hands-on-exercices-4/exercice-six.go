package main

import "fmt"

func main() {
	states := make([]string, 50)
	states = []string{"Alabama", "Alaska", "Arizona", "Arkansas"}

	fmt.Println(cap(states))
	fmt.Println(len(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

}
