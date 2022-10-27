package hello_world

import (
	"fmt"
	"golang/introduction/pkg/basics/variables"
	"golang/introduction/pkg/basics/visibility/public"
)

func main() {
	fmt.Println("Hello, world.")
	p := public.Person{Name: "IVAN", LastName: "HERNANDEZ"}
	fmt.Println(p)

	fmt.Println(variables.SECOND)
}
