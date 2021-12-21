package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate sample config file",
	Run: func(cmd *cobra.Command, args []string) {
		sample := &InputData{
			URL:    "https://www.google.com",
			Method: "POST",
			Headers: []string{
				"Content-Type: application/json",
				"X-Requested-With: XMLHttpRequest",
			},
			Body: "action=transfer&amount=15",
		}

		f, err := os.Create("cfg.yaml")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %+v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		if err := yaml.NewEncoder(f).Encode(sample); err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding yaml: %+v\n", err)
			os.Exit(1)
		}
		fmt.Println("File cfg.yaml created successfully")
	},
}

func init() {
	RootCmd.AddCommand(genCmd)
}
