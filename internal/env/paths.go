package env

import (
	"os"
	"path/filepath"
	"runtime"
)

func BaseDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	base := filepath.Join(home, ".inusdk")

	if runtime.GOOS == "windows" {
		if err := os.MkdirAll(base, 0755); err != nil {
			docs := filepath.Join(home, "Documents", "InuSDK")
			if err2 := os.MkdirAll(docs, 0755); err2 != nil {
				return "", err2
			}
			return docs, nil
		}
	} else {
		if err := os.MkdirAll(base, 0755); err != nil {
			return "", err
		}
	}

	return base, nil
}

func CandidatesDir() (string, error) {
	base, err := BaseDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "candidates"), nil
}

func CandidateVersionDir(sdk, version string) (string, error) {
	candidates, err := CandidatesDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(candidates, sdk, version), nil
}
