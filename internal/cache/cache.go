package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type LicenseInfo struct {
	Name  string `json:"name"`
	Risky bool   `json:"risky"`
}

type ModuleMetadata struct {
	Module  string      `json:"module"`
	Version string      `json:"version"`
	License LicenseInfo `json:"license"`
	CVEs    []string    `json:"cves"`
}

// GetCacheDir returns default root directory to use cached data
func GetCacheDir() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("failed to fetch cache dir: %v", err)
	}

	path := filepath.Join(dir, "goguard")
	if err := os.MkdirAll(path, 0o755); err != nil {
		return "", fmt.Errorf("failed to create cache dir: %w", err)
	}

	return path, nil
}

// SaveJSON writes v to a cache file with given name.
func SaveJSON(name string, v any) error {
	dir, err := GetCacheDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, name)
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}

// LoadJSON reads JSON from cache file into v.
func LoadJSON(name string, v any) error {
	dir, err := GetCacheDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, name)
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("cache miss")
	} else if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// IsExpired checks if cache file is older than ttl.
func IsExpired(name string, ttl time.Duration) bool {
	dir, _ := GetCacheDir()
	path := filepath.Join(dir, name)
	info, err := os.Stat(path)
	if err != nil {
		return true // missing file = expired
	}

	return time.Since(info.ModTime()) > ttl
}
