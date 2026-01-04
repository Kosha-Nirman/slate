package cmd

import (
	"github.com/spf13/cobra"
)

var presentCmd = &cobra.Command{
	Use:   "present [file]",
	Short: "Present a markdown file",
	Long: `Present a markdown file as a slide presentation.

The markdown file should use horizontal rules (---) to separate slides.
You can also include YAML frontmatter for presentation metadata.

Example:
  slate present slides.md`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(presentCmd)
}
