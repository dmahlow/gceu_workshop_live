package commands

import (
	"fmt"

	"github.com/dmahlow/desktop-automation/internal/automation"
	"github.com/spf13/cobra"
)

// NewScreenshotCommand creates the screenshot command
func NewScreenshotCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "screenshot",
		Short: "Capture a screenshot of the entire screen",
		Long: `Capture a screenshot of the entire screen and save it to a temporary location.

This command captures the full screen and saves it as a PNG file in the system's
temporary directory. The path to the saved screenshot is printed to stdout,
making it easy to use in scripts and automation workflows.`,
		Example: `  # Take a screenshot
  desktop-automation screenshot

  # Use the screenshot path in a script
  SCREENSHOT_PATH=$(desktop-automation screenshot)
  echo "Screenshot saved to: $SCREENSHOT_PATH"`,
		Args: cobra.NoArgs,
		RunE: runScreenshotCommand,
	}

	return cmd
}

// runScreenshotCommand handles the screenshot command execution
func runScreenshotCommand(cmd *cobra.Command, args []string) error {
	// Capture the screenshot
	filepath, err := automation.CaptureScreenshot()
	if err != nil {
		return fmt.Errorf("failed to capture screenshot: %w", err)
	}

	// Print the file path to stdout (for scripting)
	fmt.Println(filepath)

	return nil
}
