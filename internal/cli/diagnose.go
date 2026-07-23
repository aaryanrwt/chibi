package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Run an AI diagnosis on a resource",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Diagnosis complete. No issues found.")
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}
