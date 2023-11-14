package builtins_test

import (
	"errors"
	"testing"

	"github.com/edwinsoto88/CSCE4600/Project2/builtins"
)

func TestAliasCommand(t *testing.T) {
	builtins.AliasMap = make(map[string]string) // Reset AliasMap before each test case

	tests := []struct {
		name      string
		args      []string
		wantErr   error
		wantAlias string
	}{
		{
			name:      "display all aliases when no arguments provided",
			args:      []string{},
			wantErr:   nil, // No error expected
			wantAlias: "",  // Adjust if expecting specific output when no aliases set
		},
		{
			name:      "display specific alias value",
			args:      []string{"alias_name"},
			wantErr:   builtins.ErrAliasNotFound, // Expecting alias not found error
			wantAlias: "",                        // Adjust if expecting specific output for this case
		},
		{
			name:      "set a new alias",
			args:      []string{"new_alias", "some_command"},
			wantErr:   nil, // No error expected
			wantAlias: "",  // Adjust if expecting specific output for this case
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.Alias(tt.args...)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Alias() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Check alias value after setting
			if tt.wantErr == nil && tt.args[0] != "" {
				val, exists := builtins.AliasMap[tt.args[0]]
				if exists && val != tt.args[1] {
					t.Errorf("Alias value mismatch, got %s, want %s", val, tt.args[1])
				}
			}
		})
	}
}
