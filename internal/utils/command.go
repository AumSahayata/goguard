package utils

import (
	"os"
	"os/exec"
	"runtime"
)

// RunCommand executes a shell command cross-platform.
func RunCommand(cmdStr string) error {
	var c *exec.Cmd

	if runtime.GOOS == "windows" {
		// Use cmd.exe on Windows
		c = exec.Command("cmd", "/C", cmdStr)
	} else {
		// Use sh on Unix-like systems
		c = exec.Command("sh", "-c", cmdStr)
	}

	// forward stdout/stderr so user sees the output live
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
