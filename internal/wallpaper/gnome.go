package wallpaper

import (
	"fmt"
	"net/url"
	"os/exec"
	"path/filepath"
	"strings"
)

type gnomeBackend struct{}

func (gnomeBackend) Name() string { return "gnome" }

func (gnomeBackend) DefaultWallpaperDir() string {
	return "/usr/share/backgrounds"
}

func (gnomeBackend) Get() (string, error) {
	stdout, err := exec.Command(
		"gsettings",
		"get",
		"org.gnome.desktop.background",
		"picture-uri",
	).Output()
	if err != nil {
		return "", err
	}
	raw := strings.TrimSpace(string(stdout))
	return FileURIToPath(raw)
}

func (gnomeBackend) Set(imagePath string) error {
	uri, err := PathToFileURI(imagePath)
	if err != nil {
		return err
	}
	_, err = exec.Command(
		"gsettings",
		"set",
		"org.gnome.desktop.background",
		"picture-uri",
		uri,
	).CombinedOutput()
	return err
}

// PathToFileURI converts a filesystem path to a file:// URI for gsettings.
func PathToFileURI(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	u := url.URL{Scheme: "file", Path: filepath.ToSlash(abs)}
	return u.String(), nil
}

// FileURIToPath converts a gsettings picture-uri value to a filesystem path.
func FileURIToPath(uri string) (string, error) {
	uri = strings.TrimSpace(uri)
	uri = strings.Trim(uri, "'\"")
	if uri == "" {
		return "", nil
	}
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", fmt.Errorf("parse wallpaper URI: %w", err)
	}
	if parsed.Scheme != "file" {
		return "", fmt.Errorf("expected file URI, got scheme %q", parsed.Scheme)
	}
	return filepath.FromSlash(parsed.Path), nil
}
