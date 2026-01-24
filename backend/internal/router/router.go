package router

import (
	"gcozy_player/internal/container"
	"gcozy_player/internal/controller"
	"gcozy_player/internal/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter(container container.Container) *gin.Engine {
	defaultRouter := gin.Default()
	defaultRouter.Use(middleware.CORS())

	router := defaultRouter.Group("/api")
	router.Use(middleware.ValidateID())

	SetUpArtistRouter(router, container)
	SetUpTrackRouter(router, container)
	SetUpQueueRouter(router, container)
	SetUpLibraryRouter(router, container)

	return defaultRouter
}

func SetUpArtistRouter(r *gin.RouterGroup, c container.Container) {
	artistController := controller.NewArtistController(c)

	r.GET("/artists", artistController.GetAll)
	r.GET("/artist/:id", artistController.GetByID)
	r.GET("/artist/cover/:id", artistController.GetCover)
}

func SetUpTrackRouter(r *gin.RouterGroup, c container.Container) {
	trackController := controller.NewTrackController(c)

	r.GET("/tracks", trackController.GetAll)
	r.GET("/track/:id", trackController.GetByID)
	r.GET("/track/info/:id", trackController.GetInfoByID)
	r.GET("/track/cover/:id", trackController.GetCover)
}

func SetUpQueueRouter(r *gin.RouterGroup, c container.Container) {
	queueController := controller.NewQueueController(c)

	r.GET("/queue/next", queueController.Next)
	r.GET("/queue/prev", queueController.Prev)
	r.POST("/queue", queueController.Generate)
}

func SetUpLibraryRouter(r *gin.RouterGroup, c container.Container) {
	libraryController := controller.NewLibraryController(c)

	r.POST("/library/rescan", libraryController.Rescan)
}
