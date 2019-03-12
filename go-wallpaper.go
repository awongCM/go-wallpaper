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

	fmt.Println("ready to set wallpaper...")
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

	// TODO - need to place in another module for this
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
		for _, imageFile := range files {
			fmt.Println("imageFile: ", imageFile)

			SetCurrentWallpaper(imageFile)

			// test timer to set different wallpaper
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func endItHere() {
	fmt.Println("just cancelled...")
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
