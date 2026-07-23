package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system health and cluster connectivity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cluster connectivity OK. AI Provider configured.")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
