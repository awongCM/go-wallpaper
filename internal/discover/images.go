package discover

import (
	"os"
	"path/filepath"
	"strings"
)

var imageExtensions = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
	".webp": {},
}

var skipDirNames = map[string]struct{}{
	".localizations": {},
	".thumbnails":    {},
}

// ListImages walks root and returns paths to image files.
func ListImages(rootFolder string) ([]string, error) {
	if _, err := os.Stat(rootFolder); err != nil {
		return nil, err
	}

	var files []string
	err := filepath.Walk(rootFolder, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if fileInfo.IsDir() {
			if _, skip := skipDirNames[fileInfo.Name()]; skip {
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if _, ok := imageExtensions[ext]; !ok {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
