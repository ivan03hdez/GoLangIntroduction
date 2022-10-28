package variables

import "fmt"

// VARIABLES EXIST IN THE SCOPE THEY ARE DECLARED

// local variables - Declared within a function, statement, etc.
// global variables - Declared in the upper scope of a package.

var VariableGlobal = 10

func PrintNumber() {
	var firstNumber int = 1
	var secondNumber = 2
	var (
		n     = 5
		hello = "HOLA"
	)
	fmt.Printf("%d %d %d %s", firstNumber, secondNumber, n, hello)

	number := 5 // sintatic sugar for var number int = 5
	fmt.Printf("%d", number)

	const NUMBER_1 int = 5
	const NUMBER_2 = 5
	fmt.Printf("%d %d", NUMBER_1, NUMBER_2)

	const (
		ZERO = 0
		FIRST
		SECOND
		THIRD
		FOURTH
		FIFTH
	)
	fmt.Printf("%d %d %d %d %d %d", ZERO, FIRST, SECOND, THIRD, FOURTH, FIFTH)

	const (
		_ZERO = iota
		_FIRST
		_SECOND
		_THIRD
		_FOURTH
		_FIFTH
	)
	fmt.Printf("%d %d %d %d %d %d", _ZERO, _FIRST, _SECOND, _THIRD, _FOURTH, _FIFTH)

	var one, two, three = 1, 2, 3
	fmt.Println(one, two, three)
}
