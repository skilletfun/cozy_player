package config

import "os"

type Config struct {
	MusicFolder string
	Database    string
}

func LoadConfig() *Config {
	config := Config{
		MusicFolder: "/music",
		Database:    "/db/database.db",
	}

	if musicFolder, exists := os.LookupEnv("COZY_PLAYER_MUSIC_FOLDER"); exists {
		config.MusicFolder = musicFolder
	}

	if database, exists := os.LookupEnv("COZY_PLAYER_DATABASE"); exists {
		config.Database = database
	}

	return &config
}
