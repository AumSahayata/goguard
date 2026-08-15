package cmd

import (
	"fmt"
	"os"

	"github.com/AumSahayata/goguard/internal/cache"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage goguard cache",
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear cached data",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cache.GetCacheDir()
		if err != nil {
			fmt.Println("Failed to clear cache due to error:", err)
		}
		return os.RemoveAll(path)
	},
}

func init() {
	cacheCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(cacheCmd)
}
