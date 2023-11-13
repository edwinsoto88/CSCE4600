package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgTouch = errors.New("invalid argument count for touch")
)

func Touch(args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("%w: expected one argument (file)", ErrInvalidArgTouch)
	case 1:
		// Create an empty file with the specified name.
		fileName := args[0]
		file, err := os.Create(fileName)

		// Check for errors.
		if err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidArgTouch, err)
		}

		// Close the file after creating it.
		return file.Close()
	default:
		return fmt.Errorf("%w: expected one argument (file)", ErrInvalidArgTouch)
	}
}
