package reporter

import (
	"fmt"
	"html"
	"os"
	"strconv"
	"strings"

	"github.com/AumSahayata/goguard/internal/scanner"
)

func PrintHTML(results []scanner.ModuleResult, filename string) error {
	// Strip ANSI codes from statuses first
	for i := range results {
		results[i].Status = StripANSI(results[i].Status)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %q: %w", filename, err)
	}
	defer file.Close()

	// Write header and style
	file.WriteString(`<html>
<head>
<style>
table { border-collapse: collapse; width: 100%; }
th, td { border: 1px solid #ddd; padding: 8px; word-break: break-all; }
th { background-color: #f2f2f2; }
</style>
</head>
<body>
<table>
<tr>
<th>Package</th><th>Version</th><th>Latest</th><th>Status</th><th>Issues</th><th>Risk Score</th><th>Risk Level</th>
</tr>
`)

	for _, r := range results {
		name := html.EscapeString(r.Name)
		version := html.EscapeString(r.Version)
		latest := html.EscapeString(r.Latest)
		status := html.EscapeString(r.Status)
		issues := html.EscapeString(r.Issues)
		score := html.EscapeString(strconv.Itoa(r.RiskScore))
		risk := html.EscapeString(r.RiskLevel)

		// Decide color for status cell
		color := "green"
		if strings.Contains(r.Status, "FAIL") {
			color = "red"
		} else if strings.Contains(r.Status, "WARN") {
			color = "orange"
		}

		file.WriteString("<tr>")
		file.WriteString("<td>" + name + "</td>")
		file.WriteString("<td>" + version + "</td>")
		file.WriteString("<td>" + latest + "</td>")
		file.WriteString("<td style='color:" + color + "'>" + status + "</td>")
		file.WriteString("<td>" + issues + "</td>")
		file.WriteString("<td>" + score + "</td>")
		file.WriteString("<td>" + risk + "</td>")
		file.WriteString("</tr>\n")
	}

	file.WriteString("</table>\n</body>\n</html>")
	return nil
}
