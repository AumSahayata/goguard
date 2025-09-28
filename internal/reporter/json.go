package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/AumSahayata/goguard/internal/scanner"
)

func PrintJSON(result []scanner.ModuleResult, outputJSON string) {
	for i := range result {
		result[i].Status = StripANSI(result[i].Status)
	}

	jsonData, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert scan results to JSON: %v\n", err)
		return
	}

	// Write to stdout
	if outputJSON == "" {
		fmt.Println(string(jsonData))
		return
	}

	// Write to file
	err = os.WriteFile(outputJSON, jsonData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write JSON to file '%s': %v\n", outputJSON, err)
		return
	}

	fmt.Println("Successfully written to", outputJSON)
}

func StripANSI(s string) string {
	re := regexp.MustCompile("\x1b\\[[0-9;]*m")
	return re.ReplaceAllString(s, "")
}
