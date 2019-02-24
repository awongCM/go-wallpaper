package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// TODOS ....

func CheckOSEnviroment() {
	if runtime.GOOS == "windows" {
		fmt.Println("Running under Windows OS... ")
	} else if runtime.GOOS == "linux" {
		fmt.Println("Running under Unix/Linux OS... ")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("Running under Mac OS... ")
	}
}

func GetCurrentWallpaper() (string, error) {
	stdout, err := exec.Command("osascript", "-e", `tell application "Finder" to get POSIX path of (get desktop picture as alias)`).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

func GetDefaultLocation() string {
	const desktopWallpaperFolder = "/Library/Desktop Pictures"
	root := desktopWallpaperFolder

	return root
}

func GetListOfPictures(rootFolder string) {
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

	for _, file := range files {
		fmt.Println("file: ", file)
	}
}

func main() {
	CheckOSEnviroment()

	wallpaperLocation := GetDefaultLocation()

	GetListOfPictures(wallpaperLocation)

	currentWallPaper, err := GetCurrentWallpaper()

	if err != nil {
		panic(err)
	}

	fmt.Println("Current Desktop Wallpaper: ", currentWallPaper)
}
