package config

type Config struct {
	MusicFolder string
	Database    string
}

func LoadConfig() *Config {
	config := Config{
		MusicFolder: "/music",
		Database:    "/db/database.db",
	}

	return &config
}
