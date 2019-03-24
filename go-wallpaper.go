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
	TIMER_IN_MILLISECONDS time.Duration = 20 * time.Millisecond
	TIMER_IN_SECONDS      time.Duration = 10 * time.Second
	TIMER_IN_MINUTES      time.Duration = 5 * time.Minute
	TIMER_IN_HOURS        time.Duration = 1 * time.Hour
)

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

func SetCurrentWallpaper(imageFileLocation string) (string, error) {
	stdout, err := exec.Command("osascript", "-e", `tell application "System Events" to set picture of (reference to current desktop) to "`+imageFileLocation+`"`).Output()
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
	CheckOSEnviroment()

	wallpaperLocation := GetDefaultLocation()

	allWallpaperFiles := GetListOfPictures(wallpaperLocation)

	AlternateWallPapers(allWallpaperFiles)

}
