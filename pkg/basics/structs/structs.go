package structs

import "fmt"

type Car struct {
	Doors  int
	Weight int
	Horses int
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
