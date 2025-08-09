package controller

import (
	"gcozy_player/internal/container"
	"gcozy_player/internal/service"
	"gcozy_player/pkg/response"

	"github.com/gin-gonic/gin"
)

type LibraryController interface {
	Rescan(c *gin.Context)
}

type libraryController struct {
	service service.LibraryService
}

func NewLibraryController(c container.Container) LibraryController {
	return &libraryController{service: service.NewLibraryService(c)}
}

// Rescan is the method to start rescan and update library.
func (l *libraryController) Rescan(c *gin.Context) {
	err := l.service.Rescan()
	response.ErrorOrNoContent(c, err)
}
