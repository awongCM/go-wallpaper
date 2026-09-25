package wallpaper

import (
	"fmt"
	"runtime"
)

// Backend sets and reads the desktop wallpaper for the current platform.
type Backend interface {
	Name() string
	Get() (string, error)
	Set(imagePath string) error
	DefaultWallpaperDir() string
}

// New returns a wallpaper backend for the current operating system.
func New() (Backend, error) {
	return newBackendFor(runtime.GOOS)
}

func newBackendFor(goos string) (Backend, error) {
	switch goos {
	case "darwin":
		return &darwinBackend{}, nil
	case "linux":
		return &gnomeBackend{}, nil
	default:
		return nil, fmt.Errorf("unsupported OS: %s", goos)
	}
}
