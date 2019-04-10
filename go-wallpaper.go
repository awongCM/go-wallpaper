package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// We have few time options here
const (
	TIMER_IN_MILLISECONDS time.Duration = 200 * time.Millisecond
	TIMER_IN_SECONDS      time.Duration = 10 * time.Second
	TIMER_IN_MINUTES      time.Duration = 5 * time.Minute
	TIMER_IN_HOURS        time.Duration = 1 * time.Hour
)

// OperatingSystem ...
type OperatingSystem struct {
	osRuntime, desktopWallPaperLocation    string
	executableName, getCommand, setCommand string
}

// OSMap ...
var OSMap = map[string]OperatingSystem{
	"linux": OperatingSystem{
		"linux", "/usr/share/backgrounds",
		"gsettings", "get org.gnome.desktop.background picture-uri",
		"set org.gnome.desktop.background picture-uri",
	},
	"darwin": OperatingSystem{
		"darwin", "/Library/Desktop Pictures/",
		"osascript", `-e tell application "Finder" to get POSIX path of (get desktop picture as alias)`,
		`-e tell application "System Events" to set picture of (reference to current desktop) to`,
	},
}

// CheckOSEnviroment
func CheckOSEnviroment() string {
	return OSMap[runtime.GOOS].osRuntime
}

// GetCurrentWallpaper
func GetCurrentWallpaper() (string, error) {
	stdout, err := exec.Command(OSMap[runtime.GOOS].executableName, OSMap[runtime.GOOS].getCommand).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

// SetCurrentWallpaper
func SetCurrentWallpaper(imageFileLocation string) (string, error) {
	var setImageFileCommand = OSMap[runtime.GOOS].setCommand + `"` + imageFileLocation + `"`

	stdout, err := exec.Command(OSMap[runtime.GOOS].executableName, setImageFileCommand).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

// GetDefaultLocation
func GetDefaultLocation() string {

	return OSMap[runtime.GOOS].desktopWallPaperLocation
}

// GetListOfPictures
func GetListOfPictures(rootFolder string) []string {
	var files []string
	fmt.Println("rootFolder: ", rootFolder)

	var skipFolders [2]string

	skipFolders[0] = ".localizations"
	skipFolders[1] = ".thumbnails"

	err := filepath.Walk(rootFolder, func(path string, fileInfo os.FileInfo, err error) error {
		for _, unwantedFolder := range skipFolders {
			if fileInfo.IsDir() && fileInfo.Name() == unwantedFolder {
				return filepath.SkipDir
			}
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}

// AlternateWallPapers
func AlternateWallPapers(wallpapper_files []string) {

	// Keyboard interrupt here
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		endItHere()
		os.Exit(1)
	}()

	// loop forever until keyboard interrupt signal kicks in
	for {
		for _, wallpaperFile := range wallpapper_files {
			fmt.Println("wallpaperFile: ", wallpaperFile)

			SetCurrentWallpaper(wallpaperFile)

			// test timer to set different wallpaper
			time.Sleep(TIMER_IN_MILLISECONDS)
		}
	}
}

func endItHere() {
	fmt.Println("got keyboard interrupt...")

	currentWallPaper, err := GetCurrentWallpaper()

	if err != nil {
		panic(err)
	}

	fmt.Println("Current Desktop Wallpaper: ", currentWallPaper)
}

func main() {

	fmt.Println("Current running OS: ", CheckOSEnviroment())

	wallpaperLocation := GetDefaultLocation()

	allWallpaperFiles := GetListOfPictures(wallpaperLocation)

	AlternateWallPapers(allWallpaperFiles)

}
