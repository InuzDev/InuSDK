package manifest

import (
	"fmt"
)

type PlatformBuild struct {
	URL      string `json:"url"`
	Checksum string `json:"checksum"`
	Bin      string `json:"bin"`
}

type Version struct {
	Windows map[string]PlatformBuild `json:"windows"`
	Linux   map[string]PlatformBuild `json:"linux"`
	Darwin  map[string]PlatformBuild `json:"darwin"`
}

type Manifest struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Homepage    string             `json:"homepage"`
	Versions    map[string]Version `json:"versions"`
}

func (m *Manifest) Resolve(version, goos, arch string) (*PlatformBuild, error) {
	ver, ok := m.Versions[version]
	if !ok {
		return nil, fmt.Errorf("version %s not found", version)
	}

	var builds map[string]PlatformBuild
	switch goos {
	case "windows":
		builds = ver.Windows
	case "linux":
		builds = ver.Linux
	case "darwin":
		builds = ver.Darwin
	default:
		return nil, fmt.Errorf("unsupported OS: %s", os)
	}

	build, ok := builds[arch]
	if !ok {
		return nil, fmt.Errorf("no build for %s/%s", os, arch)
	}

	return &build, nil
}
