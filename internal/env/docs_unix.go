//go:build !windows

package env

import (
	"os"
	"path/filepath"
)

func documentsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "Documents"), nil
}
