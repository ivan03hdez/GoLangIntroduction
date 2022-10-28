package goroutines

import (
	"fmt"
	"sync"
)

// variables that come from the upper scope have to be passed as arguments to the goroutine, otherwise we cannot determine the value is ok.

func SimpleGoroutine() {
	fmt.Println("executing a goroutine")
	go func(name string) {
		fmt.Println(name)
	}("IVAN")
}

//channels
// ch <- VARIABLE   --> we send a value over the channel if possible
// var VARIABLE = <- ch ---> we receive the value thrown at the channel
// channels block when someone tries to send and it is full and when trying to receive from the channel and it is empty

func SimpleGoroutineWithChannel() {
	ch := make(chan int, 1)

	go func(x int) {
		ch <- 10 * x
	}(10)

	sum := <-ch
	fmt.Println("sum", sum)
}

func GoroutineWithBufferedChannel() {
	numberOfGoroutines := 10
	var ch = make(chan int, numberOfGoroutines)

	for i := 0; i < numberOfGoroutines; i++ {
		go func(x int) {
			ch <- x
		}(i)
	}

	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Println("n: ", n)
	}
}

// waiting group --> used to wait for all the goroutines launched here to finish

func GoroutineWithBufferedChannelAndWaitingGroup() {
	var ch = make(chan int, 2)
	var wg sync.WaitGroup

	numberOfGoRoutines := 10
	for i := 0; i < numberOfGoRoutines; i++ {
		wg.Add(1)
		ch <- struct{} // If there are no more messages in the channel, it will block the thread until they fit.
		go func(x int) {
			defer wg.Done()
			<-ch
		}(i)
	}
	wg.Wait()

	// select statement will block until one of its cases is executed
	// this is useful when waiting to get messages for more than one channel
	for i := 0; i < numberOfGoRoutines; i++ {
		select {
		case msg1 := <-ch:
			fmt.Println("received ", msg1)
		default: // The default case is always able to proceed and runs if all other cases are blocked.
			fmt.Println("no message received")
		}
	}
}
