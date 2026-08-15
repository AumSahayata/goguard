package utils

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

func SimpleHash(data string) string {
	h := fnv.New32a()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func ProjectCacheName() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	projectName := modfile.ModulePath(data)
	projectName = strings.ReplaceAll(projectName, "/", "_")

	hash := SimpleHash(projectName)

	return fmt.Sprintf("scan-%s_%s.json", projectName, hash), nil
}
