package conditionals

func switchFunc() int {
	switch x := 5; x {
	case 5:
		return x
	default:
		return 0
	}

	y := 5
	switch {
	case y == 5:
		return y
	default:
		return 0
	}
}

func ifFunc() {
	if x := 10; x < 5 {
		println("x vale menos de 5")
	} else if x > 5 {
		println("x vale más de 5")
	} else {
		println("x vale 5")
	}
}
