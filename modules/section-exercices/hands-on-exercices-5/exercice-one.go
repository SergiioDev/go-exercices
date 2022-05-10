package main

import (
	"fmt"
)

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	Vehicle
	fourWheel bool
}

type Sedan struct {
	Vehicle
	luxury bool
}

func main() {
	greatTruck := Truck{
		Vehicle:   Vehicle{doors: 2, color: "blue"},
		fourWheel: true,
	}
	greatSedan := Sedan{
		Vehicle: Vehicle{doors: 4, color: "black"},
		luxury:  true,
	}
	fmt.Printf("Truck = %v \nSedan = %v\n", greatTruck, greatSedan)
}
