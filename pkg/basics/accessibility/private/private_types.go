package private

// Public: every package can access to those properties
// Private: only code placed at the same package can access to private properties

type animal struct {
	Name           string // Public
	ScientificName string // Public
	Species        string // Public
}

func NewAnimal() animal {
	return animal{Name: "peter", ScientificName: "cannis", Species: "dog"}
}
