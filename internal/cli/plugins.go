package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Manage WASM plugins",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No WASM plugins installed.")
	},
}

func init() {
	rootCmd.AddCommand(pluginsCmd)
}
