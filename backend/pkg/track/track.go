package track

import "path/filepath"

// IsTrackFile checks provided file is audio file.
func IsTrackFile(path string) bool {
	extension := filepath.Ext(path)
	_, ok := map[string]bool{".mp3": true, ".m4a": true}[extension]
	return ok
}
