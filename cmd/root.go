package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goguard",
	Short: "Go Guard - scan Go projects for vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `goguard scan` to start scanning.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
