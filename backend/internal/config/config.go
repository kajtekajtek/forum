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
	KeycloakHost		string
	KeycloakPort		string
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

	c.KeycloakHost		= os.Getenv("KEYCLOAK_HOST")
	c.KeycloakPort		= os.Getenv("KEYCLOAK_PORT")
	c.KeycloakRealm		= os.Getenv("KEYCLOAK_REALM")
	c.KeycloakClientID	= os.Getenv("KEYCLOAK_CLIENT_ID")

	return &c, nil
}
