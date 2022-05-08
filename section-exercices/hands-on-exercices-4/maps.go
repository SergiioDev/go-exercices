package main

import "fmt"

func main() {

	jamesBondCast := map[string]int{
		"James":          32,
		"Miss Monepenny": 27,
	}

	//if the key exists in the map the "exists" is true otherwise is false
	if value, exists := jamesBondCast["Tom"]; exists {
		fmt.Println(value)
	}

	//adding element to the map
	jamesBondCast["sergi"] = 33

	for key, value := range jamesBondCast {
		fmt.Println(key, value)
	}

	delete(jamesBondCast, "sergi")

	if _, exists := jamesBondCast["sergi"]; !exists {
		fmt.Println("sergi successfull deleted")
	}
}
