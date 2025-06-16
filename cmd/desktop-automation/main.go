package main

import (
	"fmt"
	"os"

	_ "github.com/charmbracelet/lipgloss"
	"github.com/dmahlow/desktop-automation/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "desktop-automation",
	Short:   "Beautiful Desktop Automation CLI",
	Long:    "Beautiful Desktop Automation CLI - A powerful command-line tool for automating desktop interactions including mouse clicks, cursor movements, and text input.",
	Version: "v0.1.0",
}

func main() {
	// Initialize cobra
	commands.AddCommands(rootCmd)

	// Execute root command and handle errors gracefully
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
