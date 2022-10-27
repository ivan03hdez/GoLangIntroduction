package main

import (
	"fmt"
	// "golang/introduction/pkg/basics/structs"

	// "golang/introduction/pkg/basics/loops"
	"golang/introduction/pkg/goroutines"
	// string_example "golang/introduction/pkg/packages/example"
)

func main() {
	fmt.Println("Hello, world.")

	// Uncomment for testing it
	// string_example.RunExample()
	// structs.CreateNewFord()
	// loops.SimpleLoop()
	// loops.InfiniteLoop()
	// loops.LoopThroughSliceAndStructs()
	// loops.LoopThroughMap()
	goroutines.GoroutineWithBufferedChannel()
}
