package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  "Manage slate configuration files.",
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a default configuration file",
	Long: `Create a default configuration file at ~/.config/slate/slate.yaml
	
This will create a configuration file with sensible defaults that you can customize.`,
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Create a default configuration file",
	Long:  "Create a default configuration file at ~/.config/slate/slate.yaml",
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show configuration file path",
	Long:  "Show the path to the configuration file being used",
}

var configExampleCmd = &cobra.Command{
	Use:   "example",
	Short: "Show an example configuration",
	Long:  "Display an example configuration file with all available options.",
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configPathCmd)
	configCmd.AddCommand(configExampleCmd)
}
