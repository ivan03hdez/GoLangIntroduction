package functions

import "fmt"

func newFunction(parameter string) (string, error) {
	if parameter == "" {
		return "", fmt.Errorf("parameter cannot be blank")
	}

	return parameter, nil
}

var Fn = newFunction
