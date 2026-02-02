package tags

import (
	"go.senan.xyz/taglib"
)

type Tags struct {
	Title    string
	Artist   string
	Album    string
	Year     string
	Duration int
}

// GetTags returns track file's tags.
func GetTags(filename string) (Tags, error) {
	properties, err := taglib.ReadProperties(filename)
	if err != nil {
		return Tags{}, err
	}
	tags, err := taglib.ReadTags(filename)
	if err != nil {
		return Tags{}, err
	}

	year := ""
	if len(tags[taglib.Date]) > 0 {
		year = tags[taglib.Date][0]
	}

	return Tags{
		Title:    tags[taglib.Title][0],
		Artist:   tags[taglib.Artist][0],
		Album:    tags[taglib.Album][0],
		Year:     year,
		Duration: int(properties.Length.Seconds()),
	}, nil
}
