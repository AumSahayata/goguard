package scanner

import (
	"io"
	"net/http"
	"strings"
)

func checkLicense(modulePath string) (string, bool, error) {
	repoPath := strings.TrimPrefix(modulePath, "github.com/")

	urls := []string{
		"https://raw.githubusercontent.com/" + repoPath + "/master/LICENSE",
		"https://raw.githubusercontent.com/" + repoPath + "/main/LICENSE",
	}

	var content string
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		if resp.StatusCode != http.StatusOK {
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		content = string(body)
		resp.Body.Close()
		break
	}

	if content == "" {
		return "Unknown", false, nil
	}

	if strings.Contains(content, "MIT License") {
		return "MIT", false, nil
	}
	if strings.Contains(content, "Apache License") {
		return "Apache-2.0", false, nil
	}
	if strings.Contains(content, "GNU General Public License") {
		return "GPL", true, nil
	}
	if strings.Contains(content, "GNU Affero General Public License") {
		return "AGPL", true, nil
	}

	return "Other", false, nil
}
