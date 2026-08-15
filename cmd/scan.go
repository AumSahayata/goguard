package cmd

import (
	"fmt"
	"time"

	"github.com/AumSahayata/goguard/internal/cache"
	"github.com/AumSahayata/goguard/internal/codes"
	"github.com/AumSahayata/goguard/internal/parser"
	"github.com/AumSahayata/goguard/internal/reporter"
	"github.com/AumSahayata/goguard/internal/scanner"
	"github.com/AumSahayata/goguard/internal/utils"
	"github.com/spf13/cobra"
)

const (
	ProjectCacheTTL = 10 * time.Minute // cache TTL for whole-project results
)

var jsonFile, htmlFile string
var json, verbose, strict, no_cache bool

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the current Go project for vulnerabilities, outdated dependencies, repo and license issues",
	Long: `Scan the current Go project in-depth.
Exit codes:
  0 -> Scan completed successfully, no issues
  1 -> Warnings found (outdated or stale dependencies)
  2 -> Archived, Vulnerabilities detected (CVEs found)
You can use --json-file or --html-file to export results.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true // suppress usage on error

		// compute project cache file
		projectName, err := utils.ProjectCacheName()
		if err != nil {
			return fmt.Errorf("failed to compute project cache name: %v", err)
		}

		var results []scanner.ModuleResult
		usingCache := false

		// try load project cache
		if !no_cache && !cache.IsExpired(projectName, ProjectCacheTTL) {
			if err := cache.LoadJSON(projectName, &results); err == nil {
				usingCache = true
			}
		}

		// if no cache, run a fresh scan
		if !usingCache {
			mods, err := parser.ParseGoMod("go.mod")
			if err != nil {
				return err
			}

			results, _, err = scanner.ScanModules(mods, no_cache)
			if err != nil {
				return err
			}
			// save results for project cache
			if err := cache.SaveJSON(projectName, results); err != nil {
				fmt.Printf("warning: failed to save project cache: %v\n", err)
			}
		}

		// handle output format
		if htmlFile != "" {
			if err := reporter.PrintHTML(results, htmlFile); err != nil {
				return err
			}
		} else if json {
			reporter.PrintJSON(results, "")
		} else if jsonFile != "" {
			reporter.PrintJSON(results, jsonFile)
		} else {
			reporter.PrintTable(results)
		}

		// evaluate exit code
		exitCode, reasons := codes.EvaluateExit(results, strict)

		if verbose && len(reasons) > 0 {
			fmt.Println("Exit code reasons:")
			for _, reason := range reasons {
				fmt.Printf(" - %s\n", reason)
			}
		}

		if usingCache {
			fmt.Println("Loaded results from cache")
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
	scanCmd.Flags().BoolVarP(&no_cache, "no-cache", "c", false, "Do not use cache.")
	scanCmd.Flags().BoolVarP(&strict, "strict", "s", false, "Exit with non-zero code on warnings")
	scanCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output for exit codes.")
	scanCmd.Flags().BoolVarP(&json, "json", "j", false, "Print output as JSON on the console.")
	scanCmd.Flags().StringVarP(&jsonFile, "json-file", "f", "", "Write results to a JSON file.")
	scanCmd.Flags().StringVarP(&htmlFile, "html-file", "t", "", "Write results to an HTML file.")

	rootCmd.AddCommand(scanCmd)
}
