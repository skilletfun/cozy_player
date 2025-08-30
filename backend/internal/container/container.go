package container

import (
	"gcozy_player/config"

	"gorm.io/gorm"
)

// Container store data (app config, DB connection)
// that sharing in overall application.
type Container interface {
	GetConfig() *config.Config
	GetConnection() *gorm.DB
}

type container struct {
	config *config.Config
	db     *gorm.DB
}

func NewContainer(conf *config.Config, db *gorm.DB) Container {
	return &container{config: conf, db: db}
}

// GetConfig returns the application config.
func (c *container) GetConfig() *config.Config {
	return c.config
}

// GetConnection returns database connection.
func (c *container) GetConnection() *gorm.DB {
	return c.db
}
