package track

import "testing"

func TestIsTrackFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"wrong empty string", "", false},
		{"wrong string with spaces", "    \t  \n\n ", false},
		{"wrong string with digits", "123123", false},
		{"wrong string", "filepath", false},
		{"wrong extension mp4", "filepath.mp4", false},
		{"wrong extension avi", "filepath.avi", false},
		{"wrong string with slashes", "/home/user/Music/filepath.avi", false},
		{"wrong string with mp3", "mp3", false},
		{"correct string mp3", "filepath.mp3", true},
		{"correct string m4a", "filepath.m4a", true},
		{"correct string mp3 with slashes", "/home/user/Music/filepath.mp3", true},
		{"correct string m4a with slashes", "/home/user/Music/filepath.m4a", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsTrackFile(test.path)
			if result != test.expected {
				t.Errorf("IsTrackFile(%q) = %t, want %t", test.path, result, test.expected)
			}
		})
	}
}
