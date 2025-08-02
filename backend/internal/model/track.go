package model

type Track struct {
	ID        uint    `gorm:"primarykey" json:"id"`
	Title     string  `gorm:"not null" json:"title"`
	ArtistID  uint    `gorm:"not null" json:"artistId"`
	Artist    *Artist `json:"-"`
	Album     string  `json:"album"`
	Path      string  `gorm:"not null" json:"-"`
	Duration  uint16  `gorm:"not null;default:0" json:"duration"`
	PlayCount uint16  `gorm:"not null;default:0" json:"-"`
}
