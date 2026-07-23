package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Chibi",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chibi v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
