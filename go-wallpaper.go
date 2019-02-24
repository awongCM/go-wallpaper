package main

import (
	"fmt"
	"os"
	// "os/user"
	"path/filepath"
	// "strconv"
	// "strings",
	"runtime"
)

// TODO - something magical is going to happen here...

func CheckOSEnviroment() {
	if runtime.GOOS == "windows" {
		fmt.Println("Running under Windows OS... ")
	} else if runtime.GOOS == "linux" {
		fmt.Println("Running under Unix/Linux OS... ")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("Running under Mac OS... ")
	}

}

func GetDefaultLocation() string {

	const desktopWallpaperFolder = "/Library/Desktop Pictures"
	root := desktopWallpaperFolder

	return root
}

func GetListOfPictures(rootFolder string) {
	var files []string
	fmt.Println("rootFolder: ", rootFolder)

	err := filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
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
}
