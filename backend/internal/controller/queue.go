package controller

import (
	"fmt"
	"gcozy_player/internal/container"
	"gcozy_player/internal/service"
	"gcozy_player/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueueController interface {
	Generate(c *gin.Context)
	Next(c *gin.Context)
	Prev(c *gin.Context)
}

type GenerateQueueBy struct {
	ArtistId int
	TrackId  int
}

type queueController struct {
	service service.QueueService
}

func NewQueueController(c container.Container) QueueController {
	return &queueController{service: service.NewQueueService(c)}
}

// Generate is the method to generate new track's playing queue.
func (q *queueController) Generate(c *gin.Context) {
	generateBy := GenerateQueueBy{}
	if c.Bind(&generateBy) != nil {
		return
	}

	var err error
	if artistId := generateBy.ArtistId; artistId != 0 {
		err = q.service.GenerateByArtist(artistId)
	} else if trackId := generateBy.TrackId; trackId != 0 {
		err = q.service.GenerateByTrack(trackId)
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
