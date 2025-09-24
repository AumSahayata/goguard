package cmd

import (
	"github.com/AumSahayata/goguard/internal/parser"
	"github.com/AumSahayata/goguard/internal/reporter"
	"github.com/AumSahayata/goguard/internal/scanner"
	"github.com/spf13/cobra"
)

var outputFile string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the current Go project",
	RunE: func(cmd *cobra.Command, args []string) error {

		jsonOutput, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}

		mods, err := parser.ParseGoMod("go.mod")
		if err != nil {
			return err
		}

		results, err := scanner.ScanModules(mods)
		if err != nil {
			return err
		}

		if jsonOutput {
			reporter.PrintJSON(results, "")
		} else if outputFile != "" {
			reporter.PrintJSON(results, outputFile)
		} else {
			reporter.PrintTable(results)
		}

		return nil
	},
}

func init() {
	scanCmd.Flags().BoolP("json", "j", false, "Print output as json on the console.")
	scanCmd.Flags().StringVarP(&outputFile, "json-file", "f", "", "Get output in a json file")

	rootCmd.AddCommand(scanCmd)
}
