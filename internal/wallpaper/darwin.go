package wallpaper

import (
	"fmt"
	"os/exec"
	"strings"
)

type darwinBackend struct{}

func (darwinBackend) Name() string { return "darwin" }

func (darwinBackend) DefaultWallpaperDir() string {
	return "/Library/Desktop Pictures/"
}

func (darwinBackend) Get() (string, error) {
	stdout, err := exec.Command(
		"osascript",
		"-e",
		`tell application "Finder" to get POSIX path of (get desktop picture as alias)`,
	).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}

func (b darwinBackend) Set(imagePath string) error {
	script := fmt.Sprintf(
		`tell application "System Events" to set picture of every desktop to POSIX file "%s"`,
		escapeAppleScriptString(imagePath),
	)
	_, err := exec.Command("osascript", "-e", script).CombinedOutput()
	return err
}

func escapeAppleScriptString(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return s
}
