package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goguard",
	Short: "Scan Go projects for vulnerable, outdated, or risky dependencies",
	Long: `GoGuard is a security and maintenance scanner for Go projects.

It inspects your project's go.mod file and:
  • Detects outdated dependencies
  • Flags vulnerabilities (CVEs)
  • Warns about unmaintained or archived repos
  • Checks license risks
  • Scores each module with a risk level

GoGuard can output results as a table, JSON, or HTML, and its exit codes
are designed for CI pipelines.

Examples:
  goguard scan            # Scan current project
  goguard scan --json     # Scan and print JSON
  goguard upgrade         # Show upgrade commands for outdated packages
  goguard upgrade --execute  # Actually upgrade outdated packages`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `goguard scan` to start scanning.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
