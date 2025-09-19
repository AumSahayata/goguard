package parser

import (
	"os"

	"golang.org/x/mod/modfile"
)

type Module struct {
	Path    string
	Version string
}

func ParseGoMod(path string) ([]Module, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, err
	}

	var modules []Module
	for _, r := range modFile.Require {
		modules = append(modules, Module{
			Path:    r.Mod.Path,
			Version: r.Mod.Version,
		})
	}

	return modules, nil
}
