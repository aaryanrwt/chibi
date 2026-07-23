package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Chibi configuration",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Configuration loaded.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
