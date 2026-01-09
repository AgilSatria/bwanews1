package cmd

import (
	"bwanews/internal/app"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Bwanews server",
	Long:  `Start`,
	Run: func(cmd *cobra.Command, args []string) {
		app.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
