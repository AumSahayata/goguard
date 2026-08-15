package cmd

import (
	"fmt"
	"os"

	"github.com/AumSahayata/goguard/internal/cache"
	"github.com/AumSahayata/goguard/internal/parser"
	"github.com/AumSahayata/goguard/internal/scanner"
	"github.com/AumSahayata/goguard/internal/utils"
	"github.com/spf13/cobra"
)

var execFlag bool

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Suggest and optionally upgrade outdated packages.",
	Long:  "Scan for outdated packages and show 'go get' commands. Use --execute to actually run them.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var results []scanner.ModuleResult
		var usingCache = false

		// compute project cache file
		projectName, err := utils.ProjectCacheName()
		if err != nil {
			return fmt.Errorf("failed to compute project cache name: %v", err)
		}

		// try load project cache
		if !cache.IsExpired(projectName, ProjectCacheTTL) {
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

			results, _, err = scanner.ScanModules(mods, false)
			if err != nil {
				return err
			}

			if err := cache.SaveJSON(projectName, results); err != nil {
				fmt.Fprintf(os.Stderr, "warning: failed to cache scan results: %v\n", err)
			}

		}

		cmds := scanner.CollectUpgradeCommand(results)

		if len(cmds) == 0 {
			fmt.Println("All modules are up-to-date 🎉")
			return nil
		}

		if execFlag {
			var failed []string
			for i, cmdStr := range cmds {
				fmt.Printf("Upgrading %s with: %s\n", results[i].Name, cmdStr)
				if err := utils.RunCommand(cmdStr); err != nil {
					fmt.Fprintf(os.Stderr, "Failed to execute: %s (%v)\n", cmdStr, err)
					failed = append(failed, results[i].Name)
				}
			}
			if len(failed) > 0 {
				return fmt.Errorf("failed to upgrade %d modules: %v", len(failed), failed)
			}
			return nil
		}

		fmt.Println("Suggested commands to update:")
		for i, c := range cmds {
			fmt.Printf("%s: %s\n", results[i].Name, c)
		}

		fmt.Println("\n(Use --execute to actually run these commands)")

		return nil
	},
}

func init() {
	upgradeCmd.Flags().BoolVarP(&execFlag, "execute", "e", false, "Execute the suggested commands.")

	rootCmd.AddCommand(upgradeCmd)
}
