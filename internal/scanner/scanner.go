package scanner

import (
	"fmt"
	"strings"
	"time"

	"github.com/AumSahayata/goguard/internal/cache"
	"github.com/AumSahayata/goguard/internal/parser"
	"github.com/fatih/color"
	"golang.org/x/mod/semver"
)

const (
	RiskHighThreshold   = 7
	RiskMediumThreshold = 3

	moduleCacheTTL = 7 * 24 * time.Hour // cache TTL for each module metadata
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
)

type ModuleResult struct {
	Name              string
	Version           string
	Latest            string
	Vulnerable        bool
	CVEs              []string
	Status            string
	Issues            string
	RiskScore         int
	RiskLevel         string
	UpgradeSuggestion string
}

func ScanModules(modules []parser.Module, no_cache bool) ([]ModuleResult, bool, error) {
	var results []ModuleResult
	usingCache := false

	for _, mod := range modules {
		upgradeSuggestion := ""
		statusText := "[OK] Up-to-date"
		latest := fetchLatestVersion(mod.Path)
		cacheFileName := strings.ReplaceAll(mod.Path, "/", "_") + "@" + mod.Version + ".json"

		archived, stale, err := checkRepoStatus(mod.Path)
		if err != nil {
			fmt.Printf("warning: repo check failed for %s@%s: %v\n", mod.Path, mod.Version, err)
		}

		// License and CVEs
		var license string
		var risky bool
		var cves []string
		var metadata cache.ModuleMetadata

		if no_cache || cache.IsExpired(cacheFileName, moduleCacheTTL) {
			cves, err = checkVulnerabilities(mod.Path, mod.Version)
			if err != nil {
				fmt.Printf("warning: vuln check failed for %s@%s: %v\n", mod.Path, mod.Version, err)
				cves = []string{}
			}

			license, risky, err = checkLicense(mod.Path)
			if err != nil {
				fmt.Printf("warning: license check failed for %s@%s: %v\n", mod.Path, mod.Version, err)
			}

			metadata = cache.ModuleMetadata{
				Module:  mod.Path,
				Version: mod.Version,
				License: cache.LicenseInfo{Name: license, Risky: risky},
				CVEs:    cves,
			}

			err = cache.SaveJSON(cacheFileName, metadata)
			if err != nil {
				fmt.Printf("failed to cache for %s@%s: %v\n", mod.Path, mod.Version, err)
			}
		} else {
			err = cache.LoadJSON(cacheFileName, &metadata)
			if err != nil {
				cves, err = checkVulnerabilities(mod.Path, mod.Version)
				if err != nil {
					fmt.Printf("warning: vuln check failed for %s@%s: %v\n", mod.Path, mod.Version, err)
					cves = []string{}
				}

				license, risky, err = checkLicense(mod.Path)
				if err != nil {
					fmt.Printf("warning: license check failed for %s@%s: %v\n", mod.Path, mod.Version, err)
				}
			} else {
				cves = metadata.CVEs
				license = metadata.License.Name
				risky = metadata.License.Risky

				usingCache = true
			}
		}
		isVulnerable := len(cves) > 0

		// risk scoring
		score := 0
		issuesList := []string{}

		if semver.Compare(mod.Version, latest) < 0 {
			score += 2
			issuesList = append(issuesList, "Outdated")
			statusText = "[WARN] Outdated"
			upgradeSuggestion = "go get " + mod.Path + "@latest"
		}
		if risky {
			if license == "Unknown" {
				score += 2
			} else {
				score += 3
			}
			issuesList = append(issuesList, fmt.Sprintf("License: %s", license))
			statusText = "[WARN] License"
		}
		if stale {
			score += 3
			issuesList = append(issuesList, "Repo stale")
			statusText = "[WARN] Unmaintained"
		}
		if isVulnerable {
			score += 7
			issuesList = append(issuesList, fmt.Sprintf("%d CVEs", len(cves)))
			statusText = "[FAIL] Vulnerable"
		}
		if archived {
			score += 10
			issuesList = append(issuesList, "Repo archived")
			statusText = "[FAIL] Archived"
		}

		issues := "-"
		if len(issuesList) > 0 {
			issues = strings.Join(issuesList, "; ")
		}

		// map score to risk level
		level := "Low"
		if score >= RiskHighThreshold {
			level = "High"
		} else if score >= RiskMediumThreshold {
			level = "Medium"
		}

		colorFn := green
		if strings.Contains(statusText, "FAIL") {
			colorFn = red
		} else if strings.Contains(statusText, "WARN") {
			colorFn = yellow
		}

		coloredStatus := colorFn(statusText)

		results = append(results, ModuleResult{
			Name:              mod.Path,
			Version:           mod.Version,
			Latest:            latest,
			Vulnerable:        isVulnerable,
			CVEs:              cves,
			Status:            coloredStatus,
			Issues:            issues,
			RiskScore:         score,
			RiskLevel:         level,
			UpgradeSuggestion: upgradeSuggestion,
		})
	}

	return results, usingCache, nil
}
