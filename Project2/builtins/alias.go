package builtins

import (
	"errors"
	"fmt"
)

var (
	ErrAliasFormat   = errors.New("invalid alias format")
	ErrAliasNotFound = errors.New("alias not found")
)

// AliasMap stores user-defined aliases.
var AliasMap = make(map[string]string)

func Alias(args ...string) error {
	switch len(args) {
	case 0:
		// Display all aliases
		for key, value := range AliasMap {
			fmt.Printf("%s='%s'\n", key, value)
		}
		return nil
	case 1:
		// Display the value of a specific alias
		if val, ok := AliasMap[args[0]]; ok {
			fmt.Printf("%s='%s'\n", args[0], val)
			return nil
		}
		return fmt.Errorf("%w: alias '%s' not found", ErrAliasNotFound, args[0])
	case 2:
		// Set an alias
		AliasMap[args[0]] = args[1]
		return nil
	default:
		return fmt.Errorf("%w: expected zero, one, or two arguments", ErrInvalidArgCount)
	}
}
