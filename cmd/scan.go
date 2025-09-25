package cmd

import (
	"fmt"

	"github.com/AumSahayata/goguard/internal/codes"
	"github.com/AumSahayata/goguard/internal/parser"
	"github.com/AumSahayata/goguard/internal/reporter"
	"github.com/AumSahayata/goguard/internal/scanner"
	"github.com/spf13/cobra"
)

var outputFile string
var verbose bool

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the current Go project for vulnerabilities, outdated dependencies, repo and license issues",
	Long: `Scan the current Go project in-depth. 
Exit codes:
  0 -> Scan completed successfully, no issues
  1 -> Warnings found (outdated or stale dependencies)
  2 -> Vulnerabilities detected (CVEs found)
You can use --json or --json-file to export results.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		cmd.SilenceUsage = true // Do not print usage for this specific error

		jsonOutput, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}

		// parse go.mod
		mods, err := parser.ParseGoMod("go.mod")
		if err != nil {
			return err
		}

		// scan results
		results, err := scanner.ScanModules(mods)
		if err != nil {
			return err
		}

		// output
		if jsonOutput {
			reporter.PrintJSON(results, "")
		} else if outputFile != "" {
			reporter.PrintJSON(results, outputFile)
		} else {
			reporter.PrintTable(results)
		}

		exitCode, reasons := codes.EvaluateExit(results)

		if verbose && len(reasons) > 0 {
			fmt.Println("Exit code reasons:")
			for _, reason := range reasons {
				fmt.Printf(" - %s\n", reason)
			}
		}

		if exitCode != codes.ExitCodeOK {
			return &codes.ExitCodeError{
				Code: exitCode,
				Msg:  fmt.Sprintf("scan completed with exit code %d", exitCode),
			}
		}
		return nil
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output for exit codes.")
	scanCmd.Flags().BoolP("json", "j", false, "Print output as json on the console.")
	scanCmd.Flags().StringVarP(&outputFile, "json-file", "f", "", "Get output in a json file.")

	rootCmd.AddCommand(scanCmd)
}
