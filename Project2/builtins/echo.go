package builtins

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidArgEcho = errors.New("invalid arg count for echo")
)

func Echo(args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("%w: expected at least one arg", ErrInvalidArgEcho)
	default:
		// Convert []string to []interface{} before passing to fmt.Println
		interfaceArgs := make([]interface{}, len(args))
		for i, v := range args {
			interfaceArgs[i] = v
		}

		// Print the arguments to the console.
		fmt.Println(interfaceArgs...)
		return nil
	}
}
