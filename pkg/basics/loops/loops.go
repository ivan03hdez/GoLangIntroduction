package loops

import "fmt"

func SimpleLoop() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
}

func InfiniteLoop() {
	for i := 0; ; i++ {
		fmt.Println("I WILL NEVER END: ", i)
	}
}

func LoopThroughSliceAndStructs() {
	slice := []string{"ivan", "sergio"}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("cadena:%s", slice[i])
	}

	for _, cadena := range slice {
		fmt.Printf("cadena:%s", cadena)
	}
}

func LoopThroughMap() {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	for i, key := range m {
		fmt.Printf("key:%s", key)
		fmt.Printf("i:%s", i)
	}
}
