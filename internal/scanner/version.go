package scanner

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type ProxyResponse struct {
	Version string `json:"Version"`
	Time    string `json:"Time"`
}

func fetchLatestVersion(moduleName string) string {
	escaped := url.PathEscape(moduleName)
	url := "https://proxy.golang.org/" + escaped + "/@latest"

	resp, err := http.Get(url)
	if err != nil {
		return "unknown"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "unknown"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "unknown"
	}

	var result ProxyResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "unknown"
	}

	return result.Version
}
