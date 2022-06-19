package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type user struct {
	First string
	Age   int
}

type UsersByAge []user

func (u UsersByAge) Len() int {
	return len(u)
}
func (u UsersByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u UsersByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

type UsersByName []user

func (u UsersByName) Len() int {
	return len(u)
}
func (u UsersByName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u UsersByName) Less(i, j int) bool {
	return u[i].First < u[j].First
}

func main() {

	users := []user{
		{
			First: "Aames",
			Age:   32,
		},
		{
			First: "Zoneypenny",
			Age:   27,
		},
	}

	bs, err := json.Marshal(users)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Users marshaled: %v\n", string(bs))

	var usersConverted []user

	err = json.Unmarshal(bs, &usersConverted)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Users unmarshaled: %v\n", usersConverted)

	sort.Sort(UsersByAge(users))

	fmt.Printf("Users by age: %v\n", users)

	sort.Sort(UsersByName(users))
	fmt.Printf("Users by Name: %v\n", users)
}
