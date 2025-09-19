package cmd

import (
	"github.com/AumSahayata/goguard/internal/parser"
	"github.com/AumSahayata/goguard/internal/reporter"
	"github.com/AumSahayata/goguard/internal/scanner"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the current Go project",
	RunE: func(cmd *cobra.Command, args []string) error {
		mods, err := parser.ParseGoMod("go.mod")
		if err != nil {
			return err
		}

		results, err := scanner.ScanModules(mods)
		if err != nil {
			return err
		}

		reporter.PrintTable(results)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
