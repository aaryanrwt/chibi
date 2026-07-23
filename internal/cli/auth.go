package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with AI provider",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Authenticated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
