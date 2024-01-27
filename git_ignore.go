package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var defaultGitIgnoreFields = []string{
	".gitx",
	".idea",
	".vscode",
	"node_modules",
	"vendor",
	"*.log",
	"*.out",
	"*.pid",
	"*.seed",
	"*.key",
	"*.pem",
	"*.pub",
	"*.crt",
	"*.csr",
	".DS_Store",
	".env",
}

func AppendGitignoreFields(path string, fields ...string) error {
	f, err := os.OpenFile(filepath.Join(path, gitIgnoreFileName), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore: %w", err)
	}
	defer f.Close()

	origin := make(map[string]struct{})
	for _, field := range fields {
		origin[field] = struct{}{}
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		delete(origin, scanner.Text())
	}

	for field := range origin {
		if _, err := f.WriteString(field + "\n"); err != nil {
			return fmt.Errorf("failed to write .gitignore: %w", err)
		}
	}

	return nil
}
