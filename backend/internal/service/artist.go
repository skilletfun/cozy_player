package service

import (
	"gcozy_player/config"
	"gcozy_player/internal/container"
	"gcozy_player/internal/model"
	"gcozy_player/pkg/cover"
	"os"

	"gorm.io/gorm"
)

type ArtistService interface {
	GetAll() (*[]model.Artist, error)
	GetByID(id int) (*model.Artist, error)
	GetInfoByID(id int) (*model.ArtistInfo, error)
	UpdateCovers() error
	UpdateCover(artist *model.Artist) error
	Create(artist *model.Artist) error
	BulkCreate(artists *[]model.Artist) error
	Delete(artist *model.Artist) error
	BulkDelete(artist *[]model.Artist) error
	DeleteWithoutTracks() error
}

type artistService struct {
	conf *config.Config
	db   *gorm.DB
}

func NewArtistService(c container.Container) ArtistService {
	return &artistService{conf: c.GetConfig(), db: c.GetConnection()}
}

// GetAll returns all artists sorted by name in ascending order.
func (a *artistService) GetAll() (*[]model.Artist, error) {
	artists := []model.Artist{}
	err := a.db.Order("name ASC").Find(&artists).Error
	return &artists, err
}

// GetByID returns artist by id.
func (a *artistService) GetByID(id int) (*model.Artist, error) {
	artist := model.Artist{}
	err := a.db.First(&artist, id).Error
	return &artist, err
}

// GetInfoByID returns artist by id with extra tracks info.
func (a *artistService) GetInfoByID(id int) (*model.ArtistInfo, error) {
	artist, err := a.GetByID(id)
	if err != nil {
		return nil, err
	}

	var artistTrackInfo model.ArtistTrackInfo
	err = a.db.Table("tracks").
		Select("COALESCE(COUNT(*), 0) AS tracks_count, COALESCE(SUM(duration), 0) AS duration").
		Where("artist_id = ?", artist.ID).
		Group("artist_id").
		Scan(&artistTrackInfo).Error

	if err != nil {
		return nil, err
	}

	return &model.ArtistInfo{Artist: *artist, ArtistTrackInfo: artistTrackInfo}, nil
}

// UpdateCovers is the method to update artist's covers.
// It runs UpdateCover method for all artists.
func (a *artistService) UpdateCovers() error {
	artists, err := a.GetAll()
	if err != nil {
		return err
	}

	for _, artist := range *artists {
		if err := a.UpdateCover(&artist); err != nil {
			return err
		}
	}

	return nil
}

// UpdateCover is the method to update artist's cover.
// Method checks that artist's cover exists.
// If cover does not exists, artist's cover property will be cleared.
// If artist has empty cover property, search for cover in artist folder.
func (a *artistService) UpdateCover(artist *model.Artist) error {
	if artist.Cover != "" {
		if _, err := os.Stat(artist.Cover); err != nil {
			if err := a.db.Model(artist).Update("cover", "").Error; err != nil {
				return err
			}
		}
	}

	if artist.Cover == "" {
		path := a.conf.MusicFolder + "/" + artist.Name
		if coverPath := cover.FindCover(path); coverPath != "" {
			coverPath = path + "/" + coverPath
			if err := a.db.Model(artist).Update("cover", coverPath).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *artistService) Create(artist *model.Artist) error {
	return a.db.Create(artist).Error
}

func (a *artistService) BulkCreate(artists *[]model.Artist) error {
	return a.db.Create(artists).Error
}

func (a *artistService) Delete(artist *model.Artist) error {
	return a.db.Delete(artist).Error
}

func (a *artistService) BulkDelete(artists *[]model.Artist) error {
	return a.db.Delete(artists).Error
}

// DeleteWithoutTracks is the method to delete all artists that have no tracks.
func (a *artistService) DeleteWithoutTracks() error {
	artists := []model.Artist{}

	a.db.Model(&artists).
		Joins("LEFT JOIN tracks t ON t.artist_id = artists.id").
		Where("t.artist_id IS NULL").
		Scan(&artists)

	if len(artists) > 0 {
		return a.db.Delete(&artists).Error
	}
	return nil
}
