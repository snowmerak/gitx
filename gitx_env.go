package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

const (
	gitxFolder         = ".gitx"
	gitxConfigFileName = "gitx.toml"
	gitxBranchDir      = "branches"

	gitIgnoreFileName = ".gitignore"
)

type Config struct {
	SSHKeyFile string
}

func CheckGitiIgnore(path string) error {
	f, err := os.Open(filepath.Join(path, gitIgnoreFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to open .gitignore: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == gitxFolder {
			return nil
		}
	}

	return fmt.Errorf("gitx is not ignored")
}

func AppendGitIgnore(path string) error {
	f, err := os.OpenFile(filepath.Join(path, gitIgnoreFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .gitignore: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(gitxFolder + "\n"); err != nil {
		return fmt.Errorf("failed to write .gitignore: %w", err)
	}

	return nil
}

func WriteGitxConfig(path string, config Config) error {
	if err := os.MkdirAll(filepath.Join(path, gitxFolder), 0755); err != nil {
		return fmt.Errorf("failed to create .gitx directory: %w", err)
	}

	f, err := os.Create(filepath.Join(path, gitxFolder, gitxConfigFileName))
	if err != nil {
		return fmt.Errorf("failed to create gitx.toml: %w", err)
	}
	defer f.Close()

	enc := toml.NewEncoder(f)
	if err := enc.Encode(config); err != nil {
		return fmt.Errorf("failed to write gitx.toml: %w", err)
	}

	return nil
}

func ReadGitxConfig(path string) (Config, error) {
	cfg := Config{}

	f, err := os.Open(filepath.Join(path, gitxFolder, gitxConfigFileName))
	if err != nil {
		return cfg, fmt.Errorf("failed to open gitx.toml: %w", err)
	}

	if err := toml.NewDecoder(f).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to decode gitx.toml: %w", err)
	}

	return cfg, nil
}
