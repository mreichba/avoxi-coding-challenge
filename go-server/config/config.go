package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application-wide configuration values loaded from the environment
// It can be extended later with DB URLs, log levels, API keys, etc.
type Config struct {
	GeoIPDBPath string
	Port        string
}

// LoadConfig loads environment variables into a Config struct
func LoadConfig() Config {
	_ = godotenv.Load()

	geoIPPath := os.Getenv("GEOIP_DB_PATH")
	if geoIPPath == "" {
		log.Fatal("GEOIP_DB_PATH is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default server port
	}

	return Config{
		GeoIPDBPath: geoIPPath,
		Port:        port,
	}
}
