package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// RunCommand runs a command with arguments and prints its stdout/stderr in real-time.
func RunCommand(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("→ Running: %s %v\n", name, args)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("command failed: %s %v: %w", name, args, err)
	}
	return nil
}

// EnsureDir creates the directory (and parent dirs) if it doesn't exist.
func EnsureDir(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("could not resolve absolute path: %w", err)
	}

	if _, err := os.Stat(abs); os.IsNotExist(err) {
		fmt.Printf("→ Creating directory: %s\n", abs)
		if err := os.MkdirAll(abs, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", abs, err)
		}
	}
	return nil
}
