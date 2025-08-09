package model

type Artist struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Cover string `json:"cover"`
}

type ArtistTrackInfo struct {
	TracksCount int `json:"tracksCount"`
	Duration    int `json:"duration"`
}

type ArtistInfo struct {
	Artist
	ArtistTrackInfo
}
