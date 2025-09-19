package reporter

import (
	"log"
	"os"

	"github.com/AumSahayata/goguard/internal/scanner"
	"github.com/olekukonko/tablewriter"
)

func PrintTable(result []scanner.ModuleResult) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Module", "Version", "Latest", "Status", "Issues"})
	var err error

	for _, r := range result {
		err = table.Append([]string{
			r.Name,
			r.Version,
			r.Latest,
			r.Status,
			r.Issues,
		})
		if err != nil {
			continue
		}
	}

	err = table.Render()
	if err != nil {
		log.Fatalf("Failed to render table: %v\n", err)
	}
}
