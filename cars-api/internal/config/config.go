package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	Port      string
	DBURL     string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
}

func Load() *Config {
	// load .env only in development
	if os.Getenv("GO_ENV") == "development" {
		_ = godotenv.Load(".env") // ignore error, env vars may already exist
	}

	cfg := &Config{
		AppName:   Get("APP_NAME", "MyApp"),
		Port:      Get("PORT", "3000"),
		DBURL:     Get("DB_URL", ""),
		DBHost:    Get("DB_HOST", "localhost"),
		DBPort:    Get("DB_PORT", "5432"),
		DBUser:    Get("DB_USER", "postgres"),
		DBPass:    Get("DB_PASSWORD", ""),
		DBName:    Get("DB_NAME", "cars_db"),
		DBSSLMode: Get("DB_SSLMODE", "disable"),
	}

	// enforce required in production
	if os.Getenv("GO_ENV") == "production" && cfg.DBURL == "" {
		log.Fatal("DB_URL required in production")
	}

	return cfg
}

func Get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
