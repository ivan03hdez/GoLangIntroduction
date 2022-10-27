package string_example

import (
	"fmt"
	strings "golang/introduction/pkg/packages/string"
)

func RunExample() {
	s := strings.ToUpperCase("  hola  ")
	fmt.Println(s)
	// s := strings.getSpaceChar() // this will error because the method is private
	s = strings.TrimSpaces(s)
	fmt.Println(s)
}
