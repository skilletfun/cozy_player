package controller

import (
	"errors"
	"gcozy_player/internal/container"
	"gcozy_player/internal/service"
	"gcozy_player/pkg/response"

	"github.com/gin-gonic/gin"
)

type ArtistController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetCover(c *gin.Context)
}

type artistController struct {
	service service.ArtistService
}

func NewArtistController(c container.Container) ArtistController {
	return &artistController{service: service.NewArtistService(c)}
}

// GetAll returns all artists.
func (a *artistController) GetAll(c *gin.Context) {
	result, err := a.service.GetAll()
	response.Response(c, result, err)
}

// GetByID returns artist by provided id.
func (a *artistController) GetByID(c *gin.Context) {
	result, err := a.service.GetInfoByID(c.GetInt("id"))
	response.Response(c, result, err)
}

// GetCover returns artist cover.
func (a *artistController) GetCover(c *gin.Context) {
	if artist, err := a.service.GetByID(c.GetInt("id")); err != nil {
		response.BadRequest(c, err)
	} else if artist.Cover == "" {
		response.NotFound(c, errors.New("cover no found"))
	} else {
		response.CachedFile(c, artist.Cover)
	}
}
