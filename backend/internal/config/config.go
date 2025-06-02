package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost		string
	DBPort		string
	DBUser		string
	DBPassword	string
	DBName		string
}

func Load() (*Config, error) {
	var c Config

	err := godotenv.Load(".env")
	if err != nil {
		return &Config{}, fmt.Errorf("load .env: %w", err)
	}

	c.DBHost = os.Getenv("DB_HOST")
	c.DBPort = os.Getenv("DB_PORT")
	c.DBUser = os.Getenv("DB_USER")
	c.DBPassword = os.Getenv("DB_PASSWORD")
	c.DBName = os.Getenv("DB_NAME")

	return &c, nil
}
