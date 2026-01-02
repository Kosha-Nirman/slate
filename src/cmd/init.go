package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const samplePresentation = `---
title: My Presentation
author: Your Name
date: 2026-01-01
---

# Welcome to slate

A terminal-based presentation tool for markdown files.

---

## Features

- **Markdown Support**: Full markdown rendering with syntax highlighting
- **Themes**: Customizable themes (dark, light, dracula, etc.)
- **Navigation**: Easy keyboard navigation
- **Progress**: Visual progress indicators

---

## Getting Started

1. Write your slides in markdown
2. Separate slides with horizontal rules (---)
3. Run: slate present slides.md

---

## Code Examples

Syntax highlighting works great:

` + "```go" + `
func main() {
    fmt.Println("Hello, World!")
}
` + "```" + `

---

## Navigation

- **Next**: →, Space, L
- **Previous**: ←, H
- **First**: Home, G
- **Last**: End, Shift+G
- **Help**: ?
- **Quit**: Q, Esc

---

# Thank You!

Press ? for help anytime.

Enjoy your presentation!
`

var initCmd = &cobra.Command{
	Use:   "init [filename]",
	Short: "Create a sample presentation",
	Long: `Create a sample markdown presentation file.

If no filename is provided, it will create "presentation.md" in the current directory.

Example:
  slate init
  slate init my-slides.md`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// * Determine filename
		filename := "presentation.md"
		if len(args) > 0 {
			filename = args[0]
		}

		// * Add .md extension if not present
		if filepath.Ext(filename) != ".md" {
			filename += ".md"
		}

		// ? Check if file exists
		if _, err := os.Stat(filename); err == nil {
			fmt.Fprintf(os.Stderr, "Error: File %s already exists\n", filename)
			os.Exit(1)
		}

		// * Write sample presentation
		if err := os.WriteFile(filename, []byte(samplePresentation), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to create file %s\n", filename)
			os.Exit(1)
		}

		fmt.Printf("Created sample presentation: %s\n", filename)
		fmt.Printf("Run: slate present %s\n", filename)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
