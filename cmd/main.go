package main

import (
	"fmt"
	"golang/introduction/pkg/goroutines"
	// string_example "golang/introduction/pkg/packages/example"
)

func main() {
	fmt.Println("Hello, world.")

	// Uncomment for testing it
	// string_example.RunExample()
	// structs.CreateNewFord()
	// functions.Fn()
	// loops.SimpleLoop()
	// loops.InfiniteLoop()
	// loops.LoopThroughSliceAndStructs()
	// loops.LoopThroughMap()
	goroutines.GoroutineWithBufferedChannel()
}
