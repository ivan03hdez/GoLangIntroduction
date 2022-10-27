package string

import "strings"

func getSpaceChar() string { // Shared in "string" package, but not exported
	return " "
}

func TrimSpaces(s string) string { //Shared in "string" package, and also exported
	sc := getSpaceChar()
	return strings.Trim(s, sc)
}
