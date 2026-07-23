package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Chibi configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuration loaded.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
