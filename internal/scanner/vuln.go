package scanner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Vulns struct {
	ID      string   `json:"id"`
	Aliases []string `json:"aliases"`
}

type VulnResponse struct {
	Vulns []Vulns `json:"vulns"`
}

func checkVulnerabilities(modulePath string, version string) ([]string, error) {
	payload := map[string]interface{}{
		"package": map[string]string{
			"name":      modulePath,
			"ecosystem": "Go",
		},
		"version": version,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://api.osv.dev/v1/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status at vuln: %s", resp.Status)
	}

	var results VulnResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	var cves []string
	for _, v := range results.Vulns {
		if len(v.Aliases) > 0 {
			cves = append(cves, v.Aliases...)
		} else {
			cves = append(cves, v.ID)
		}
	}

	return cves, nil
}
