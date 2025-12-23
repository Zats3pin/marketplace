package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("invalid DB_PORT: %v", err)
	}

	return &Config{
		ServerPort: os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
