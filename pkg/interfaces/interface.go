package interfaces

import "fmt"

type persona interface {
	GetAge() int
	FullName() string
	IsMinor() (bool, error)
}

type user struct {
	name     string
	lastname string
	age      int
}

func (u user) GetAge() int {
	return u.age
}

func (u user) FullName() string {
	return fmt.Sprintf("%s %s", u.name, u.lastname)
}

func (u user) IsMinor() (bool, error) {
	if u.age < 0 {
		return true, fmt.Errorf("Age is not valid")
	}
	return u.age > 18, nil
}

func displayValue(i interface{}) {
	fmt.Println(i)
}

func RunExample() {
	var u persona

	u = user{
		name:     "John",
		lastname: "Snow",
		age:      30,
	}

	fmt.Println(u.IsMinor())

	displayValue("Hello World!")
	displayValue(20)
	displayValue(false)
}
