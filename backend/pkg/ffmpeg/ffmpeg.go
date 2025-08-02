package ffmpeg

import (
	"bytes"
	"encoding/json"
	"io"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type Tags struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

type FileInfo struct {
	FileName   string  `json:"filename"`
	FormatName string  `json:"format_name"`
	StartTime  float64 `json:"start_time,string"`
	Duration   float64 `json:"duration,string"`
	Size       uint    `json:"size,string"`
	BitRate    uint    `json:"bit_rate,string"`
	Tags       Tags    `json:"tags"`
}

type Format struct {
	Format FileInfo `json:"format"`
}

// GetFileInfo returns track file's FileInfo.
func GetFileInfo(filename string) (FileInfo, error) {
	if info, err := ffmpeg.Probe(filename); err == nil {
		var result Format
		if err := json.Unmarshal([]byte(info), &result); err != nil {
			return FileInfo{}, err
		}
		return result.Format, nil
	} else {
		return FileInfo{}, err
	}
}

// GetCover returns image from track file.
func GetCover(filename string) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)

	err := ffmpeg.Input(filename).
		Output("pipe:", ffmpeg.KwArgs{
			"map":      "0:v",
			"c:v":      "copy",
			"frames:v": "1",
			"f":        "mjpeg",
		}).
		WithOutput(buffer, io.Discard).
		Run()

	return buffer.Bytes(), err
}
