package conditionals

import (
	"fmt"
	"math/rand"
)

func SwitchShowCase() {
	switch x := rand.Intn(11); x {
	case 2:
		fmt.Println("Number generated has been 2")
	case 5, 10:
		fmt.Println("Number generated has been 5 or 10")
	default:
		fmt.Println("Number generated not in the list")
	}

	switch y := rand.Intn(11); {
	case y > 5, y == 5:
		fmt.Println("Number generated greater or equal than 5")
	default:
		fmt.Println("Number generated less than 5")
	}
}

func IfShowcase() {
	if x := 10; x < 5 {
		fmt.Println("x vale menos de 5")
	} else if x > 5 {
		fmt.Println("x vale más de 5")
	} else {
		fmt.Println("x vale 5")
	}
}
