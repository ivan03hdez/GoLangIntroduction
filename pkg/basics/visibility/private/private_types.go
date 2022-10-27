package private

import "fmt"

// Public: every package can access to those properties
// Private: only code placed at the same package can access to private properties

type animal struct {
	Name           string // Public
	ScientificName string // Public
	Species        string // Public
}

func main() {
	var a = animal{ScientificName: "", Species: "dog"}
	fmt.Println(a)
}
