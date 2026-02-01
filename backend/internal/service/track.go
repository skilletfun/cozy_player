package service

import (
	"gcozy_player/config"
	"gcozy_player/internal/container"
	"gcozy_player/internal/model"
	"gcozy_player/pkg/cover"
	"os"
	"strings"

	"gorm.io/gorm"
)

type TrackService interface {
	GetAll() (*[]model.Track, error)
	GetAllByArtist(artistId int) (*[]model.Track, error)
	GetByID(id int) (*model.Track, error)
	GetCover(id int) (string, []byte)
	IncrementPlayCount(id int) error
	Create(track *model.Track) error
	BulkCreate(tracks *[]model.Track) error
	Delete(track *model.Track) error
	BulkDelete(tracks *[]model.Track) error
}

type trackService struct {
	conf *config.Config
	db   *gorm.DB
}

func NewTrackService(c container.Container) TrackService {
	return &trackService{conf: c.GetConfig(), db: c.GetConnection()}
}

// GetAll returns all tracks sorted by album and title in ascending order.
func (t *trackService) GetAll() (*[]model.Track, error) {
	tracks := []model.Track{}
	err := t.db.Order("album ASC, title ASC").Find(&tracks).Error
	return &tracks, err
}

// GetAllByArtist returns all artist's tracks sorted by album and title in ascending order.
func (t *trackService) GetAllByArtist(artistId int) (*[]model.Track, error) {
	tracks := []model.Track{}
	err := t.db.
		Where("artist_id = ?", artistId).
		Order("album ASC, title ASC").
		Find(&tracks).Error
	return &tracks, err
}

// GetByID returns track by id.
func (t *trackService) GetByID(id int) (*model.Track, error) {
	track := model.Track{}
	err := t.db.Joins("Artist").First(&track, id).Error
	return &track, err
}

// IncrementPlayCount is the method to increment track's play_count property.
func (t *trackService) IncrementPlayCount(id int) error {
	if track, err := t.GetByID(id); err != nil {
		return err
	} else {
		return t.db.Model(track).Update("play_count", track.PlayCount+1).Error
	}
}

// GetCover returns track cover.
// If track has no cover, cover from artist will returns.
func (t *trackService) GetCover(id int) (string, []byte) {
	if track, err := t.GetByID(id); err != nil {
		return "", nil
	} else if picture, _ := cover.GetCover(track.Path); len(picture) > 0 {
		return "image/jpeg", picture
	} else {
		return t.GetCoverFromArtist(track)
	}
}

// GetCoverFromArtist returns cover from track's artist.
func (t *trackService) GetCoverFromArtist(track *model.Track) (string, []byte) {
	if track.Artist.Cover == "" {
		return "", nil
	}

	path := track.Artist.Cover
	if picture, err := os.ReadFile(path); err != nil {
		return "", nil
	} else {
		splitted := strings.Split(path, ".")
		return "image/" + splitted[len(splitted)-1], picture
	}
}

func (t *trackService) Create(track *model.Track) error {
	return t.db.Create(track).Error
}

func (t *trackService) BulkCreate(tracks *[]model.Track) error {
	return t.db.Create(tracks).Error
}

func (t *trackService) Delete(track *model.Track) error {
	return t.db.Delete(track).Error
}

func (t *trackService) BulkDelete(tracks *[]model.Track) error {
	return t.db.Delete(tracks).Error
}
