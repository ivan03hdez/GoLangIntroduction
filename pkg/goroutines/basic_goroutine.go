package goroutines

import "fmt"

// variables that come from the upper scope have to be passed as arguments to the goroutine, otherwise we cannot determine the value is ok.

func SimpleGoroutine() {
	go func(name string) {
		fmt.Println(name)
	}("IVAN")
}

func SimpleGoroutineWithChannel() {
	var ch chan int

	go func(x int) int {
		ch <- 10 * x
	}(10)

	var sum = <-ch
	fmt.Println("sum", sum)
}

func GoroutineWithBufferedChannel() {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(x int) {
			ch <- x
		}(i)
	}

	// select statement will block until one of its cases is executed
	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch:
			fmt.Println("received ", msg1)
		default:
			fmt.Println("no message received")
		}
	}
}
