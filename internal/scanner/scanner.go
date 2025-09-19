package scanner

import (
	"fmt"

	"github.com/AumSahayata/goguard/internal/parser"
	"github.com/fatih/color"
	"golang.org/x/mod/semver"
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
)

type ModuleResult struct {
	Name       string
	Version    string
	Latest     string
	Vulnerable bool
	CVEs       []string
	Status     string
	Issues     string
}

func ScanModules(modules []parser.Module) ([]ModuleResult, error) {
	var results []ModuleResult

	for _, mod := range modules {
		latest := fetchLatestVersion(mod.Path)
		isVulnerable := false
		issues := "-"

		archived, stale, err := checkRepoStatus(mod.Path)
		if err != nil {
			// log but don’t stop
			fmt.Printf("warning: repo check failed for %s@%s: %v\n", mod.Path, mod.Version, err)
		}

		license, risky, err := checkLicense(mod.Path)

		cves, err := checkVulnerabilities(mod.Path, mod.Version)
		if err != nil {
			// log but don’t stop
			fmt.Printf("warning: vuln check failed for %s@%s: %v\n", mod.Path, mod.Version, err)
			cves = []string{}
		}

		var coloredStatus string

		switch {
		case archived:
			coloredStatus = red("[FAIL] Unmaintained")
			if issues == "-" {
				issues = "Repo archived"
			} else {
				issues += "; Repo archived"
			}

		case stale:
			coloredStatus = red("[FAIL] Unmaintained")
			if issues == "-" {
				issues = "Repo stale"
			} else {
				issues += "; Repo stale"
			}

		case len(cves) > 0:
			coloredStatus = yellow("[WARN] Vulnerable")
			issues = fmt.Sprintf("%d CVEs", len(cves))

		case risky:
			coloredStatus = yellow("[WARN] License")
			if issues == "-" {
				issues = fmt.Sprintf("License: %s", license)
			} else {
				issues += fmt.Sprintf("; License: %s", license)
			}

		case semver.Compare(mod.Version, latest) < 0:
			coloredStatus = yellow("[WARN] Outdated")

		default:
			coloredStatus = green("[OK] Up-to-date")
		}

		results = append(results, ModuleResult{
			Name:       mod.Path,
			Version:    mod.Version,
			Latest:     latest,
			Vulnerable: isVulnerable,
			CVEs:       cves,
			Status:     coloredStatus,
			Issues:     issues,
		})
	}

	return results, nil
}
