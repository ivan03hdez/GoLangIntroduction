package public

// Public: every package can access to those properties
// Private: only code placed at the same package can access to private properties

type Person struct {
	Name     string // Public
	LastName string // Public
	age      int    // Private
}
