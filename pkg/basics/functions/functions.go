package functions

import "fmt"

func newFunctionWithError(parameter string) (string, error) {
	if parameter == "" {
		return "", fmt.Errorf("parameter cannot be blank")
	}

	return parameter, nil
}

var Fn = newFunctionWithError
