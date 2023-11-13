package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgRmdir = errors.New("invalid argument count for rmdir")
)

func RmDir(args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("%w: expected one argument (directory)", ErrInvalidArgRmdir)
	case 1:
		// Remove the specified directory.
		dir := args[0]
		return os.Remove(dir)
	default:
		return fmt.Errorf("%w: expected one argument (directory)", ErrInvalidArgRmdir)
	}
}
