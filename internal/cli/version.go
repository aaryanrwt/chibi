package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Chibi",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Chibi v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
