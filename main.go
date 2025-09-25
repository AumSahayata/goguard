package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/AumSahayata/goguard/cmd"
	"github.com/AumSahayata/goguard/internal/codes"
)

func main() {
	if err := cmd.Execute(); err != nil {
		var exitErr *codes.ExitCodeError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.Code)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
