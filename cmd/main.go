package main

import (
	"fmt"
	"golang/introduction/pkg/goroutines"
	// string_example "golang/introduction/pkg/packages/example"
)

func main() {
	fmt.Println("Hello, world.")
	// string_example.RunExample()
	// variables.PrintNumber()
	///////////////ACCESSIBILITY/////////////////////
	//var persona = public.Person{
	//Name: "IVAN",
	//LastName: "HERNANDEZ",
	//property age is not accessible from outer packages
	//}
	//fmt.Println("persona:", persona)
	//	var animal = private.NewAnimal()
	//	var animal = private.NewAnimal()
	//fmt.Println("animal:", animal)
	/////////////////////////////////////////////////
	// structs.CreateNewFord()
	// functions.Fn("")
	// x := 6
	// pointers.PointerFunc(&x)
	// conditionals.SwitchShowCase()
	// conditionals.IfShowcase()
	// loops.SimpleLoop()
	// loops.InfiniteLoop()
	// loops.LoopThroughSliceAndStructs()
	// loops.LoopThroughMap()
	// goroutines.SimpleGoroutine()
	// goroutines.SimpleGoroutineWithChannel()
	// goroutines.GoroutineWithBufferedChannel()
	goroutines.GoroutineWithBufferedChannelAndWaitingGroup()
}
