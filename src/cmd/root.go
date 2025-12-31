package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	noColor bool
)

var rootCmd = &cobra.Command{
	Use:   "slate",
	Short: "A terminal-based presentation tool for markdown files",
	Long: `slate is a powerful terminal-based presentation tool that renders
markdown files as beautiful slide presentations.

Features:
- Markdown rendering with syntax highlighting
- Customizable themes and keybindings
- Navigation history
- Progress indicators
- And more!`,
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func initConfig() {
	if noColor {
		color.NoColor = true
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVar(&noColor, "no-color", false, "Disable colored output")
}
