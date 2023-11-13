package builtins

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidArgCountTime = errors.New("invalid argument count for time")
)

func Time(args ...string) error {
	switch len(args) {
	case 0:
		// Display the current time
		currentTime := time.Now()
		fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
		return nil
	case 1:
		// Execute a command and measure its execution time
		startTime := time.Now()
		// Replace this with the actual code to execute the command
		// For demonstration purposes, let's just print the command
		fmt.Println("Executing command:", args[0])
		// Simulate command execution time
		time.Sleep(2 * time.Second)
		endTime := time.Now()
		executionTime := endTime.Sub(startTime)
		fmt.Printf("Execution time: %s\n", executionTime)
		return nil
	default:
		return fmt.Errorf("%w: expected zero or one argument", ErrInvalidArgCountTime)
	}
}
