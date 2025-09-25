package codes

import (
	"fmt"
	"strings"

	"github.com/AumSahayata/goguard/internal/scanner"
)

const (
	ExitCodeOK   = 0 // no issues
	ExitCodeWarn = 1 // warnings (outdated/stale/license risk)
	ExitCodeVuln = 2 // vulnerabilities
)

func EvaluateExit(result []scanner.ModuleResult) (int, []string) {
	reasons := []string{}
	exitcode := ExitCodeOK

	for _, mod := range result {
		if len(mod.CVEs) > 0 {
			reasons = append(reasons, fmt.Sprintf("%s has CVEs: %v", mod.Name, mod.CVEs))
			exitcode = ExitCodeVuln
			break
		} else if strings.Contains(mod.Status, "WARN") || mod.Issues != "-" {
			reasons = append(reasons, fmt.Sprintf("%s is outdated or has issues: %s", mod.Name, mod.Issues))
			exitcode = max(exitcode, ExitCodeWarn)
		}
	}

	return exitcode, reasons
}
