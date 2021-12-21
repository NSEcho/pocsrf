package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "pocsrf",
	Short: "generate html template",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}
