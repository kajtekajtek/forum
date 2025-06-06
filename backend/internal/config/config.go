package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	// server variables
	APIPort	string

	// database variables
	PostgresHost		string
	PostgresPort		string
	PostgresUser		string
	PostgresPassword	string
	PostgresDB			string

	// keycloak variables
	KeycloakHost		string
	KeycloakPort		string
	KeycloakRealm		string
	KeycloakClientID	string
}

func Load() (*Config, error) {
	var c Config

	err := godotenv.Load("../.env")
	if err != nil {
		return &Config{}, fmt.Errorf("load .env: %w", err)
	}

	c.APIPort	= os.Getenv("API_PORT")

	c.PostgresHost 		= os.Getenv("POSTGRES_HOST")
	c.PostgresPort 		= os.Getenv("POSTGRES_PORT")
	c.PostgresUser 		= os.Getenv("POSTGRES_USER")
	c.PostgresPassword 	= os.Getenv("POSTGRES_PASSWORD")
	c.PostgresDB 		= os.Getenv("POSTGRES_DB")

	c.KeycloakHost		= os.Getenv("KEYCLOAK_HOST")
	c.KeycloakPort		= os.Getenv("KEYCLOAK_PORT")
	c.KeycloakRealm		= os.Getenv("KEYCLOAK_REALM")
	c.KeycloakClientID	= os.Getenv("KEYCLOAK_CLIENT_ID")

	return &c, nil
}
