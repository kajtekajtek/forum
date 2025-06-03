package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	// server variables
	Port	string

	// database variables
	DBHost		string
	DBPort		string
	DBUser		string
	DBPassword	string
	DBName		string

	// keycloak variables
	KeycloakURL			string
	KeycloakRealm		string
	KeycloakClientID	string
}

func Load() (*Config, error) {
	var c Config

	err := godotenv.Load(".env")
	if err != nil {
		return &Config{}, fmt.Errorf("load .env: %w", err)
	}

	c.Port	= os.Getenv("PORT")

	c.DBHost 		= os.Getenv("DB_HOST")
	c.DBPort 		= os.Getenv("DB_PORT")
	c.DBUser 		= os.Getenv("DB_USER")
	c.DBPassword 	= os.Getenv("DB_PASSWORD")
	c.DBName 		= os.Getenv("DB_NAME")

	c.KeycloakURL		= os.Getenv("KEYCLOAK_URL")
	c.KeycloakRealm		= os.Getenv("KEYCLOAK_REALM")
	c.KeycloakClientID	= os.Getenv("KEYCLOAK_REALM")

	return &c, nil
}
