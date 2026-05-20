//go:build windows

package env

import "golang.org/x/sys/windows"

func documentsDir() (string, error) {
	path, err := windows.KnownFolderPath(windows.FOLDERID_Documents, 0)

	if err != nil {
		return "", err
	}

	return path, nil
}
