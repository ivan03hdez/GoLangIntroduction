package structs

import "fmt"

type Car struct {
	Doors  int `json:"car_doors"`
	Weight int `json:"car_weight"`
	Horses int `json:"car_horses"`
}

type Ford struct {
	Car        // creates a property Car that has a car Object
	FordNumber int
}

func CreateNewFord() {
	var ford = Ford{
		Car: Car{
			Doors:  3,
			Weight: 1000,
			Horses: 150,
		},
		FordNumber: 5,
	}

	fmt.Println("ford:", ford)
}
