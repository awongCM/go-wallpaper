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
// On SIGINT or SIGTERM it restores originalPath when non-empty and returns nil.
func Run(backend wallpaper.Backend, paths []string, interval time.Duration, originalPath string) error {
	if len(paths) == 0 {
		return fmt.Errorf("no wallpaper images found")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(stop)

	interrupt := make(chan struct{})
	go func() {
		<-stop
		close(interrupt)
	}()

	return runUntilStop(backend, paths, interval, originalPath, interrupt)
}

func runUntilStop(
	backend wallpaper.Backend,
	paths []string,
	interval time.Duration,
	originalPath string,
	stop <-chan struct{},
) error {
	restore := func() {
		if originalPath == "" {
			log.Println("no original wallpaper path saved; skipping restore")
			return
		}
		if err := backend.Set(originalPath); err != nil {
			log.Printf("restore wallpaper: %v", err)
			return
		}
		log.Printf("restored wallpaper: %s", originalPath)
	}

	for {
		for _, imagePath := range paths {
			select {
			case <-stop:
				log.Println("interrupted, restoring wallpaper...")
				restore()
				return nil
			default:
			}

			if err := backend.Set(imagePath); err != nil {
				log.Printf("set wallpaper %q: %v", imagePath, err)
			}

			select {
			case <-stop:
				log.Println("interrupted, restoring wallpaper...")
				restore()
				return nil
			case <-time.After(interval):
			}
		}
	}
}
