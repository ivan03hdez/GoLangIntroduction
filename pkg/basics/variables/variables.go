package variables

// VARIABLES EXIST IN THE SCOPE THEY ARE DECLARED

// local variables - Declared within a function, statement, etc.
// global variables - Declared in the upper scope of a package.

func printNumber() {
	var firstNumber int = 5
	var secondNumber = 5
	var (
		n       = 5
		string_ = "HOLA"
	)

	number := 5 // sintatic sugar for var number int = 5

	const NUMBER_1 int = 5
	const NUMBER_2 = 5
	const (
		ZERO = 0
		FIRST
		SECOND
		THIRD
		FOURTH
		FIFTH
	)

	var one, two, three = 1, 2, 3
}
