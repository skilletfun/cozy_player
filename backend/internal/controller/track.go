package controller

import (
	"errors"
	"gcozy_player/internal/container"
	"gcozy_player/internal/service"
	"gcozy_player/pkg/response"

	"github.com/gin-gonic/gin"
)

type TrackController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetInfoByID(c *gin.Context)
	GetCover(c *gin.Context)
}

type TrackFilter struct {
	ArtistId int `form:"artistId"`
}

type trackController struct {
	service service.TrackService
}

func NewTrackController(c container.Container) TrackController {
	return &trackController{service: service.NewTrackService(c)}
}

// GetAll returns all tracks.
func (t *trackController) GetAll(c *gin.Context) {
	filter := TrackFilter{}
	c.BindQuery(&filter)

	if artistId := filter.ArtistId; artistId != 0 {
		result, err := t.service.GetAllByArtist(artistId)
		response.Response(c, result, err)
	} else {
		result, err := t.service.GetAll()
		response.Response(c, result, err)
	}
}

// GetByID returns track binary data (stream) by id.
func (t *trackController) GetByID(c *gin.Context) {
	if track, err := t.service.GetByID(c.GetInt("id")); err != nil {
		response.BadRequest(c, err)
	} else {
		response.File(c, track.Path)
	}
}

// GetByID returns track info by id.
func (t *trackController) GetInfoByID(c *gin.Context) {
	track, err := t.service.GetByID(c.GetInt("id"))
	response.Response(c, track, err)
}

// GetCover returns track's cover.
func (t *trackController) GetCover(c *gin.Context) {
	mimeType, picture := t.service.GetCover(c.GetInt("id"))
	if picture == nil {
		response.NotFound(c, errors.New("cover no found"))
	} else {
		response.Data(c, mimeType, picture)
	}
}
