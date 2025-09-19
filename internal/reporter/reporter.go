package reporter

import (
	"fmt"
	"os"

	"github.com/AumSahayata/goguard/internal/scanner"
	"github.com/olekukonko/tablewriter"
)

func PrintTable(result []scanner.ModuleResult) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Module", "Version", "Latest", "Vulnerable", "CVES"})
	var err error
	for _, r := range result {
		err = table.Append([]string{
			r.Name,
			r.Version,
			r.Latest,
            fmt.Sprintf("%v", r.Vulnerable),
            fmt.Sprintf("%v", r.CVEs),
		})
		if err != nil {
			continue
		}
	}

	err = table.Render()
	if err != nil {
		fmt.Printf("Failed to render table: %v\n", err)
	}
}