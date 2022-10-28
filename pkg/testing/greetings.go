package greetings

import "fmt"

func Hello(s string) string {
	if s == "" {
		return "Hello anonymous!"
	}
	return fmt.Sprintf("Hi, %s", s)
}
