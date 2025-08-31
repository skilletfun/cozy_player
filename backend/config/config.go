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

	if music_folder, exists := os.LookupEnv("MUSIC_FOLDER"); exists {
		config.MusicFolder = music_folder
	} 

	if database, exists := os.LookupEnv("DATABASE"); exists {
		config.Database = database
	}

	return &config
}
