package scanner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func checkRepoStatus(modulePath string) (bool, bool, error) {
	mod := strings.TrimPrefix(modulePath, "github.com/")
	url := "https://api.github.com/repos/" + mod

	resp, err := http.Get(url)
	if err != nil {
		return false, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, false, fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var repoInfo struct {
		Archived bool
		PushedAt string
	}

	if err := json.NewDecoder(resp.Body).Decode(&repoInfo); err != nil {
		return false, false, err
	}

	stale := false
	t, err := time.Parse(time.RFC3339, repoInfo.PushedAt)
	if err == nil && time.Since(t) > 2*365*24*time.Hour {
		stale = true
	}

	return repoInfo.Archived, stale, nil
}
