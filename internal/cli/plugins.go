package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Manage WASM plugins",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("No WASM plugins installed.")
	},
}

func init() {
	rootCmd.AddCommand(pluginsCmd)
}
