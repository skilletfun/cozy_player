package database

import (
	"gcozy_player/config"
	"gcozy_player/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// GetConnection returns the connection to database.
func GetConnection(config *config.Config) *gorm.DB {
	database := config.Database
	if db, err := gorm.Open(sqlite.Open(database)); err != nil {
		panic(err)
	} else {
		return db
	}
}

// MigrateAll is a method to migrate all models.
func MigrateAll(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Artist{}, &model.Track{}); err != nil {
		panic(err)
	}
}
