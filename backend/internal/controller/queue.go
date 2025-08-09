package controller

import (
	"fmt"
	"gcozy_player/internal/container"
	"gcozy_player/internal/service"
	"gcozy_player/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueueController interface {
	Generate(c *gin.Context)
	Next(c *gin.Context)
	Prev(c *gin.Context)
}

type queueController struct {
	service service.QueueService
}

func NewQueueController(c container.Container) QueueController {
	return &queueController{service: service.NewQueueService(c)}
}

// Generate is the method to generate new track's playing queue.
func (q *queueController) Generate(c *gin.Context) {
	var err error

	if artistId := c.Query("artistId"); artistId != "" {
		if artistId, err := strconv.Atoi(artistId); err != nil {
			response.BadRequest(c, err)
		} else {
			err = q.service.GenerateByArtist(artistId)
		}
	} else if trackId := c.Query("trackId"); trackId != "" {
		if trackId, err := strconv.Atoi(trackId); err != nil {
			response.BadRequest(c, err)
		} else {
			err = q.service.GenerateByTrack(trackId)
		}
	} else {
		err = q.service.GenerateByAll()
	}

	response.ErrorOrNoContent(c, err)
}

// Next returns the next track in the playing queue.
func (q *queueController) Next(c *gin.Context) {
	if trackId, err := q.service.Next(); err != nil {
		response.BadRequest(c, err)
	} else {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/track/info/%d", trackId))
	}
}

// Prev returns the track played before current.
func (q *queueController) Prev(c *gin.Context) {
	if trackId, err := q.service.Prev(); err != nil {
		response.BadRequest(c, err)
	} else {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/track/info/%d", trackId))
	}
}
