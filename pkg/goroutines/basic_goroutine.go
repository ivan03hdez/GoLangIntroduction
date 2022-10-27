package goroutines

import "fmt"

// variables that come from the upper scope have to passed as arguments to the goroutine, otherwise we cannot determine the value is ok.
func main() {
	go func(name string) {
		fmt.Println("ivan")
	}("IVAN")
}
