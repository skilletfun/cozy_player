package main

import (
	"gcozy_player/config"
	"gcozy_player/internal/container"
	"gcozy_player/internal/database"
	"gcozy_player/internal/router"
)

func main() {
	conf := config.LoadConfig()
	db := database.GetConnection(conf)

	container := container.NewContainer(conf, db)

	database.MigrateAll(db)

	router := router.GetRouter(container)
	router.Run("0.0.0.0:8000")
}
