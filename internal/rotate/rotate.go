package rotate

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/awongCM/go-wallpaper/internal/wallpaper"
)

// Run cycles through paths, setting each as the wallpaper until interrupted.
// On SIGINT or SIGTERM it restores originalPath and returns nil.
func Run(backend wallpaper.Backend, paths []string, interval time.Duration, originalPath string) error {
	if len(paths) == 0 {
		return fmt.Errorf("no wallpaper images found")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	restore := func() {
		if originalPath == "" {
			return
		}
		if err := backend.Set(originalPath); err != nil {
			log.Printf("restore wallpaper: %v", err)
			return
		}
		log.Printf("restored wallpaper: %s", originalPath)
	}

	go func() {
		<-stop
		log.Println("interrupted, restoring wallpaper...")
		restore()
		os.Exit(0)
	}()

	for {
		for _, imagePath := range paths {
			if err := backend.Set(imagePath); err != nil {
				log.Printf("set wallpaper %q: %v", imagePath, err)
			}
			time.Sleep(interval)
		}
	}
}
