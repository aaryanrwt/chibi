package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Stream intelligent logs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Streaming logs...")
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
