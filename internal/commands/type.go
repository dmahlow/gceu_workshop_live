package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// NewTypeCommand creates the type command
func NewTypeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "type <text>",
		Short: "Type text at current cursor position",
		Long: `Type text at the current cursor position.

This command simulates keyboard input by typing the specified text at the current
cursor location. The text will be typed as if you were physically typing on the keyboard.`,
		Example: `  # Type a simple message
  desktop-automation type "Hello, World!"

  # Type a sentence with spaces
  desktop-automation type "This is a test message"

  # Type special characters
  desktop-automation type "user@example.com"

  # Type numbers and symbols
  desktop-automation type "Password123!"`,
		Args: cobra.ExactArgs(1),
		RunE: runTypeCommand,
	}

	return cmd
}

// runTypeCommand handles the type command execution
func runTypeCommand(cmd *cobra.Command, args []string) error {
	text := args[0]

	// Validate that text is not empty
	if strings.TrimSpace(text) == "" {
		return fmt.Errorf("text cannot be empty or contain only whitespace")
	}

	// TODO: Implement actual typing functionality
	// This is a placeholder for the actual implementation
	fmt.Printf("Typing text: %q\n", text)

	// Placeholder for actual automation implementation
	// return automation.TypeText(text)

	return nil
}
