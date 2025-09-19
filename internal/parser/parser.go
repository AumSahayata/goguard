package parser

import (
	"os"

	"golang.org/x/mod/modfile"
)

func ParseGoMod(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, err
	}

	var modules []string
	for _, r := range modFile.Require {
		modules = append(modules, r.Mod.Path+"@"+r.Mod.Version)
	}

	return modules, nil
}
