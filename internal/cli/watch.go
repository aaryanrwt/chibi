package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Start the predictive daemon",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Watching cluster for anomalies...")
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
