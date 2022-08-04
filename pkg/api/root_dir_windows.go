//go:build windows
// +build windows

package api

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

func GetRootDir() string {
	if rootDir := os.Getenv("AQUA_ROOT_DIR"); rootDir != "" {
		return rootDir
	}
	xdgDataHome := xdg.DataHome
	if xdgDataHome == "" {
		xdgDataHome = filepath.Join(os.Getenv("HOME"), ".local", "share")
	}
	return filepath.Join(xdg.DataHome, "aquaproj-aqua")
}
