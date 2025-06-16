package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// NewMoveCommand creates the move command
func NewMoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "move <x> <y>",
		Short: "Move the mouse cursor to coordinates",
		Long: `Move the mouse cursor to specific screen coordinates.

This command moves the mouse cursor to the specified x and y coordinates on the screen
without clicking. The coordinates are measured in pixels from the top-left corner of
the screen (0,0).`,
		Example: `  # Move cursor to coordinates (500, 300)
  desktop-automation move 500 300

  # Move cursor to the center of a 1920x1080 screen
  desktop-automation move 960 540

  # Move cursor to the top-left corner
  desktop-automation move 0 0

  # Move cursor to bottom-right of a typical screen
  desktop-automation move 1919 1079`,
		Args: cobra.ExactArgs(2),
		RunE: runMoveCommand,
	}

	return cmd
}

// runMoveCommand handles the move command execution
func runMoveCommand(cmd *cobra.Command, args []string) error {
	// Parse X coordinate
	x, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid x coordinate '%s': must be a valid integer", args[0])
	}

	// Parse Y coordinate
	y, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid y coordinate '%s': must be a valid integer", args[1])
	}

	// Validate coordinates are not negative
	if x < 0 {
		return fmt.Errorf("x coordinate cannot be negative: %d", x)
	}
	if y < 0 {
		return fmt.Errorf("y coordinate cannot be negative: %d", y)
	}

	// TODO: Implement actual mouse movement functionality
	// This is a placeholder for the actual implementation
	fmt.Printf("Moving mouse cursor to coordinates (%d, %d)\n", x, y)

	// Placeholder for actual automation implementation
	// return automation.MoveMouse(x, y)

	return nil
}
