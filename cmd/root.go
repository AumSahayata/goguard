package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goguard",
	Short: "Go Supply Chain Guard - scan Go projects for vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `gosupplyguard scan` to start scanning.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
