package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Stream intelligent logs",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Streaming logs...")
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
