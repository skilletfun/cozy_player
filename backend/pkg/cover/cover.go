package cover

import (
	"fmt"
	"os"
	"path/filepath"

	"go.senan.xyz/taglib"
)

// FindCover checks all files in provided dir and returns first image.
// Returns empty string if no image found.
func FindCover(path string) string {
	entries, _ := os.ReadDir(path)
	for _, entry := range entries {
		extension := filepath.Ext(path + "/" + entry.Name())
		extensions := map[string]bool{".jpg": true, ".png": true, ".jpeg": true}
		if _, ok := extensions[extension]; ok {
			return entry.Name()
		}
	}

	return ""
}

// GetCover returns image from track file.
func GetCover(filename string) ([]byte, error) {
	imageBytes, err := taglib.ReadImage(filename)
	if err != nil {
		return nil, err
	}
	if imageBytes == nil {
		return nil, fmt.Errorf("File %s contains no image", filename)
	}
	return imageBytes, nil
}
