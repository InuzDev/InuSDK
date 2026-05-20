package env

import (
	"os"
	"path/filepath"
	"runtime"
)

// This returns the path of the documents folder no matter what in windows.
func BaseDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	base := filepath.Join(home, ".inusdk")

	if runtime.GOOS == "windows" {
		if err := os.MkdirAll(base, 0755); err != nil {
			docs, docErr := documentsDir()
			if docErr != nil {
				return "", err
			}
			fallback := filepath.Join(docs, "InuSDK")

			if err2 := os.MkdirAll(fallback, 0755); err2 != nil {
				return "", err2
			}
			return fallback, nil
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
