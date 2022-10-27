package hello_world

import (
	"fmt"
	"golang/introduction/pkg/basics/loops"
	"golang/introduction/pkg/basics/variables"
)

func main() {
	fmt.Println("Hello, world.")

	fmt.Println(variables.SECOND)
	loops.InfiniteLoop()
}
