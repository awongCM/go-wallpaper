package main

import (
	"log"
	"time"

	"github.com/awongCM/go-wallpaper/internal/discover"
	"github.com/awongCM/go-wallpaper/internal/rotate"
	"github.com/awongCM/go-wallpaper/internal/wallpaper"
)

const defaultInterval = 200 * time.Millisecond

func main() {
	backend, err := wallpaper.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("wallpaper backend: %s", backend.Name())

	originalPath, err := backend.Get()
	if err != nil {
		log.Printf("read current wallpaper: %v (restore on exit will be skipped)", err)
	} else if originalPath == "" {
		log.Println("current wallpaper is not a local file path; restore on exit will be skipped")
	} else {
		log.Printf("saved original wallpaper: %s", originalPath)
	}

	wallpaperDir := backend.DefaultWallpaperDir()
	images, err := discover.ListImages(wallpaperDir)
	if err != nil {
		log.Fatalf("list images in %q: %v", wallpaperDir, err)
	}

	log.Printf("found %d images in %q", len(images), wallpaperDir)

	if err := rotate.Run(backend, images, defaultInterval, originalPath); err != nil {
		log.Fatal(err)
	}
}
