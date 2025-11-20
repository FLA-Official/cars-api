package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
	DBURL   string
}

func Load() *Config {
	// load .env only in development
	if os.Getenv("GO_ENV") == "development" {
		_ = godotenv.Load(".env") // ignore error, env vars may already exist
	}

	cfg := &Config{
		AppName: get("APP_NAME", "MyApp"),
		Port:    get("PORT", "3000"),
		DBURL:   get("DB_URL", ""),
	}

	// enforce required in production
	if os.Getenv("GO_ENV") == "production" && cfg.DBURL == "" {
		log.Fatal("DB_URL required in production")
	}

	return cfg
}

func get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
