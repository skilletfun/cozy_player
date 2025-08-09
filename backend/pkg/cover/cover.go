package cover

import (
	"os"
	"path/filepath"
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
