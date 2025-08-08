package config

import "os"

type Config struct {
	MusicFolder string
	Database    string
}

func LoadConfig() *Config {
	config := Config{
		MusicFolder: "/music",
		Database:    "database.db",
	}

	if musicFolder := os.Getenv("MUSIC_FOLDER"); musicFolder != "" {
		config.MusicFolder = musicFolder
	}

	if database := os.Getenv("DATABASE"); database != "" {
		config.Database = database
	}

	return &config
}
